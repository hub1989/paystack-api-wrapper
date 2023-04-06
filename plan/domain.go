package plan

import "github.com/hub1989/paystack-api-wrapper/response"

// Plan represents a
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
type Plan struct {
	ID                int     `json:"id,omitempty"`
	CreatedAt         string  `json:"createdAt,omitempty"`
	UpdatedAt         string  `json:"updatedAt,omitempty"`
	Domain            string  `json:"domain,omitempty"`
	Integration       int     `json:"integration,omitempty"`
	Name              string  `json:"name,omitempty"`
	Description       string  `json:"description,omitempty"`
	PlanCode          string  `json:"plan_code,omitempty"`
	Amount            float32 `json:"amount,omitempty"`
	Interval          string  `json:"interval,omitempty"`
	SendInvoices      bool    `json:"send_invoices,omitempty"`
	SendSMS           bool    `json:"send_sms,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	InvoiceLimit      float32 `json:"invoice_limit,omitempty"`
	HostedPage        string  `json:"hosted_page,omitempty"`
	HostedPageURL     string  `json:"hosted_page_url,omitempty"`
	HostedPageSummary string  `json:"hosted_page_summary,omitempty"`
}

// List is a list object for Plans.
type List struct {
	Meta   response.ListMeta
	Values []Plan `json:"data"`
}
