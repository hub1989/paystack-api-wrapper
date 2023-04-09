package configuration

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

// NewClient creates a new Paystack API Client with the given API key
// and HTTP Client, allowing overriding of the HTTP Client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func NewClient(key string, httpClient *http.Client, loggingEnabled bool) *client.Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: client.DefaultHTTPTimeout}
	}

	u, _ := url.Parse(client.BaseURL)
	c := &client.Client{
		Client:         httpClient,
		Key:            key,
		BaseURL:        u,
		LoggingEnabled: loggingEnabled,
	}

	if loggingEnabled {
		log.Info("logging is enabled..all requests and errors to paystack will be logged")
	}
	return c
}
