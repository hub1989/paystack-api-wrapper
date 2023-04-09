package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/response"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Service interface {
	Call(ctx context.Context, method, path string, body, v interface{}) error
	ResolveCardBIN(ctx context.Context, bin int) (response.Response, error)
	CheckBalance(ctx context.Context) (response.Response, error)
	GetSessionTimeout(ctx context.Context) (response.Response, error)
	UpdateSessionTimeout(ctx context.Context, timeout int) (response.Response, error)
	decodeResponse(httpResp *http.Response, v interface{}) error
}

type DefaultPaystackService struct {
	Client *Client
}

// Logger interface for custom loggers
type Logger interface {
	Printf(format string, v ...interface{})
}

// Metadata is an key-value pairs added to Paystack API requests
type Metadata map[string]interface{}

// Client manages communication with the Paystack API
type Client struct {
	Client *http.Client // HTTP Client used to communicate with the API.
	// the API Key used to authenticate all Paystack API requests
	Key            string
	BaseURL        *url.URL
	LoggingEnabled bool
}

// Call actually does the HTTP request to Paystack API
func (c *Client) Call(ctx context.Context, method, path string, body, v interface{}) error {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, _ := c.BaseURL.Parse(path)
	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)

	if err != nil {
		if c.LoggingEnabled {
			log.WithError(err).Error("Cannot create Paystack request")
		}
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.Key)
	req.Header.Set("User-Agent", userAgent)

	if c.LoggingEnabled {
		log.WithFields(log.Fields{
			"method": req.Method,
			"host":   req.URL.Host,
			"path":   req.URL.Path,
		}).Infoln("Requesting")

		log.WithFields(log.Fields{
			"body": buf,
		}).Infoln("POST request data", buf)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return c.decodeResponse(resp, v)
}

// ResolveCardBIN docs https://developers.paystack.co/v1.0/reference#resolve-card-bin
func (c *Client) ResolveCardBIN(ctx context.Context, bin int) (response.Response, error) {
	u := fmt.Sprintf("/decision/bin/%d", bin)
	resp := response.Response{}
	err := c.Call(ctx, http.MethodGet, u, nil, &resp)

	return resp, err
}

// CheckBalance docs https://developers.paystack.co/v1.0/reference#resolve-card-bin
func (c *Client) CheckBalance(ctx context.Context) (response.Response, error) {
	resp := response.Response{}
	err := c.Call(ctx, http.MethodGet, "balance", nil, &resp)
	// check balance 'data' node is an array
	resp2 := resp["data"].([]interface{})[0].(map[string]interface{})
	return resp2, err
}

// GetSessionTimeout fetches payment session timeout
func (c *Client) GetSessionTimeout(ctx context.Context) (response.Response, error) {
	resp := response.Response{}
	err := c.Call(ctx, http.MethodGet, "/integration/payment_session_timeout", nil, &resp)
	return resp, err
}

// UpdateSessionTimeout updates payment session timeout
func (c *Client) UpdateSessionTimeout(ctx context.Context, timeout int) (response.Response, error) {
	data := url.Values{}
	data.Add("timeout", strconv.Itoa(timeout))
	resp := response.Response{}
	u := "/integration/payment_session_timeout"
	err := c.Call(ctx, http.MethodPut, u, data, &resp)
	return resp, err
}

// decodeResponse decodes the JSON response from the Twitter API.
// The actual response will be written to the `v` parameter
func (c *Client) decodeResponse(httpResp *http.Response, v interface{}) error {
	var resp response.Response
	respBody, err := io.ReadAll(httpResp.Body)
	json.Unmarshal(respBody, &resp)

	if status, _ := resp["status"].(bool); !status || httpResp.StatusCode >= 400 {
		if c.LoggingEnabled {
			log.WithError(err).Error("Paystack error")
			log.WithFields(log.Fields{
				"resp": resp,
			}).Error("HTTP response")
		}
		return response.NewAPIError(httpResp)
	}

	if c.LoggingEnabled {
		log.WithFields(log.Fields{
			"resp": resp,
		}).Infoln("Paystack response")
	}

	if data, ok := resp["data"]; ok {
		switch t := resp["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(resp, v)
		}
	}
	// if response data does not contain data key, map entire response to v
	return mapstruct(resp, v)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}
