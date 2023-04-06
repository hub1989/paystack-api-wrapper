package page

import "github.com/hub1989/paystack-api-wrapper/response"

// Page represents a Paystack page
// For more details see https://developers.paystack.co/v1.0/reference#create-page
type Page struct {
	ID           int                 `json:"id,omitempty"`
	CreatedAt    string              `json:"createdAt,omitempty"`
	UpdatedAt    string              `json:"updatedAt,omitempty"`
	Domain       string              `json:"domain,omitempty"`
	Integration  int                 `json:"integration,omitempty"`
	Name         string              `json:"name,omitempty"`
	Slug         string              `json:"slug,omitempty"`
	Description  string              `json:"description,omitempty"`
	Amount       float32             `json:"amount,omitempty"`
	Currency     string              `json:"currency,omitempty"`
	Active       bool                `json:"active,omitempty"`
	RedirectURL  string              `json:"redirect_url,omitempty"`
	CustomFields []map[string]string `json:"custom_fields,omitempty"`
}

// List is a list object for pages.
type List struct {
	Meta   response.ListMeta
	Values []Page `json:"data,omitempty"`
}
