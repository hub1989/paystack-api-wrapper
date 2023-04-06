package client

import (
	"fmt"
	"os"
	"time"
)

const (
	// library version
	version = "0.1.0"

	// DefaultHTTPTimeout defaultHTTPTimeout is the default timeout on the http Client
	DefaultHTTPTimeout = 60 * time.Second

	// base URL for all Paystack API requests
	BaseURL = "https://api.paystack.co"

	// User agent used when communicating with the Paystack API.
	// userAgent = "paystack-go/" + version
	userAgent = "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1"
)

// PaginateURL INTERNALS
func PaginateURL(path string, count, offset int) string {
	return fmt.Sprintf("%s?perPage=%d&page=%d", path, count, offset)
}

func MustGetTestKey() string {
	key := os.Getenv("PAYSTACK_KEY")

	if len(key) == 0 {
		panic("PAYSTACK_KEY environment variable is not set\n")
	}

	return key
}
