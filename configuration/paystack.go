package configuration

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"log"
	"net/http"
	"net/url"
	"os"
)

// NewClient creates a new Paystack API Client with the given API key
// and HTTP Client, allowing overriding of the HTTP Client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func NewClient(key string, httpClient *http.Client) *client.Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: client.DefaultHTTPTimeout}
	}

	u, _ := url.Parse(client.BaseURL)
	c := &client.Client{
		Client:         httpClient,
		Key:            key,
		BaseURL:        u,
		LoggingEnabled: true,
		Log:            log.New(os.Stderr, "", log.LstdFlags),
	}
	return c
}
