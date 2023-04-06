package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/response"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Service interface {
	Call(method, path string, body, v interface{}) error
	ResolveCardBIN(bin int) (response.Response, error)
	CheckBalance() (response.Response, error)
	GetSessionTimeout() (response.Response, error)
	UpdateSessionTimeout(timeout int) (response.Response, error)
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
	Key     string
	BaseURL *url.URL

	logger         Logger
	LoggingEnabled bool
	Log            Logger
}

// Call actually does the HTTP request to Paystack API
func (c *Client) Call(method, path string, body, v interface{}) error {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, _ := c.BaseURL.Parse(path)
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		if c.LoggingEnabled {
			c.Log.Printf("Cannot create Paystack request: %v\n", err)
		}
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.Key)
	req.Header.Set("User-Agent", userAgent)

	if c.LoggingEnabled {
		c.Log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		c.Log.Printf("POST request data %v\n", buf)
	}

	start := time.Now()

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if c.LoggingEnabled {
		c.Log.Printf("Completed in %v\n", time.Since(start))
	}

	defer resp.Body.Close()
	return c.decodeResponse(resp, v)
}

// ResolveCardBIN docs https://developers.paystack.co/v1.0/reference#resolve-card-bin
func (c *Client) ResolveCardBIN(bin int) (response.Response, error) {
	u := fmt.Sprintf("/decision/bin/%d", bin)
	resp := response.Response{}
	err := c.Call("GET", u, nil, &resp)

	return resp, err
}

// CheckBalance docs https://developers.paystack.co/v1.0/reference#resolve-card-bin
func (c *Client) CheckBalance() (response.Response, error) {
	resp := response.Response{}
	err := c.Call("GET", "balance", nil, &resp)
	// check balance 'data' node is an array
	resp2 := resp["data"].([]interface{})[0].(map[string]interface{})
	return resp2, err
}

// GetSessionTimeout fetches payment session timeout
func (c *Client) GetSessionTimeout() (response.Response, error) {
	resp := response.Response{}
	err := c.Call("GET", "/integration/payment_session_timeout", nil, &resp)
	return resp, err
}

// UpdateSessionTimeout updates payment session timeout
func (c *Client) UpdateSessionTimeout(timeout int) (response.Response, error) {
	data := url.Values{}
	data.Add("timeout", strconv.Itoa(timeout))
	resp := response.Response{}
	u := "/integration/payment_session_timeout"
	err := c.Call("PUT", u, data, &resp)
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
			c.Log.Printf("Paystack error: %+v", err)
			c.Log.Printf("HTTP Response: %+v", resp)
		}
		return response.NewAPIError(httpResp)
	}

	if c.LoggingEnabled {
		c.Log.Printf("Paystack response: %v\n", resp)
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
