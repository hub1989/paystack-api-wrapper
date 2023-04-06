package bank

import "github.com/hub1989/paystack-api-wrapper/response"

// Bank represents a Paystack bank
type Bank struct {
	ID        int    `json:"id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Name      string `json:"name,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Code      string `json:"code,omitempty"`
	LongCode  string `json:"long_code,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Active    bool   `json:"active,omitempty"`
	IsDeleted bool   `json:"is_deleted,omitempty"`
}

// List is a list object for banks.
type List struct {
	Meta   response.ListMeta
	Values []Bank `json:"data,omitempty"`
}

// BVNResponse represents response from resolve_bvn endpoint
type BVNResponse struct {
	Meta struct {
		CallsThisMonth int `json:"calls_this_month,omitempty"`
		FreeCallsLeft  int `json:"free_calls_left,omitempty"`
	}
	BVN string
}
