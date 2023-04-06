package customer

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"github.com/hub1989/paystack-api-wrapper/subscription"
)

// Customer is the resource representing your Paystack customer.
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
type Customer struct {
	ID             int                         `json:"id,omitempty"`
	CreatedAt      string                      `json:"createdAt,omitempty"`
	UpdatedAt      string                      `json:"updatedAt,omitempty"`
	Domain         string                      `json:"domain,omitempty"`
	Integration    int                         `json:"integration,omitempty"`
	FirstName      string                      `json:"first_name,omitempty"`
	LastName       string                      `json:"last_name,omitempty"`
	Email          string                      `json:"email,omitempty"`
	Phone          string                      `json:"phone,omitempty"`
	Metadata       client.Metadata             `json:"metadata,omitempty"`
	CustomerCode   string                      `json:"customer_code,omitempty"`
	Subscriptions  []subscription.Subscription `json:"subscriptions,omitempty"`
	Authorizations []interface{}               `json:"authorizations,omitempty"`
	RiskAction     string                      `json:"risk_action"`
}

// List is a list object for customers.
type List struct {
	Meta   response.ListMeta
	Values []Customer `json:"data"`
}
