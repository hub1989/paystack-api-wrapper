package subscription

import "github.com/hub1989/paystack-api-wrapper/response"

// Subscription represents a Paystack subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
type Subscription struct {
	ID          int    `json:"id,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Integration int    `json:"integration,omitempty"`
	// inconsistent API response. Create returns Customer code, Fetch returns an object
	Customer  interface{} `json:"customer,omitempty"`
	Plan      string      `json:"plan,omitempty"`
	StartDate string      `json:"start,omitempty"`
	// inconsistent API response. Fetch returns string, List returns an object
	Authorization    interface{}   `json:"authorization,omitempty"`
	Invoices         []interface{} `json:"invoices,omitempty"`
	Status           string        `json:"status,omitempty"`
	Quantity         int           `json:"quantity,omitempty"`
	Amount           int           `json:"amount,omitempty"`
	SubscriptionCode string        `json:"subscription_code,omitempty"`
	EmailToken       string        `json:"email_token,omitempty"`
	EasyCronID       string        `json:"easy_cron_id,omitempty"`
	CronExpression   string        `json:"cron_expression,omitempty"`
	NextPaymentDate  string        `json:"next_payment_date,omitempty"`
	OpenInvoice      string        `json:"open_invoice,omitempty"`
}

// Request represents a Paystack subscription request
type Request struct {
	// customer code or email address
	Customer string `json:"customer,omitempty"`
	// plan code
	Plan          string `json:"plan,omitempty"`
	Authorization string `json:"authorization,omitempty"`
	StartDate     string `json:"start,omitempty"`
}

// List is a list object for subscriptions.
type List struct {
	Meta   response.ListMeta
	Values []Subscription `json:"data"`
}
