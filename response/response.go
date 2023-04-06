package response

import (
	"net/url"
)

// RequestValues aliased to url.Values as a workaround
type RequestValues url.Values

// ListMeta is pagination metadata for paginated responses from the Paystack API
type ListMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}
