package settlement

import "github.com/hub1989/paystack-api-wrapper/response"

// List is a list object for settlements.
type List struct {
	Meta   response.ListMeta
	Values []response.Response `json:"data,omitempty"`
}
