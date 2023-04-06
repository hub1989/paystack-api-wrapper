package transaction

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/customer"
	"github.com/hub1989/paystack-api-wrapper/plan"
	"github.com/hub1989/paystack-api-wrapper/response"
	"github.com/hub1989/paystack-api-wrapper/subaccount"
)

// List is a list object for transactions.
type List struct {
	Meta   response.ListMeta
	Values []Transaction `json:"data"`
}

// Request represents a request to start a transaction.
type Request struct {
	CallbackURL       string          `json:"callback_url,omitempty"`
	Reference         string          `json:"reference,omitempty"`
	AuthorizationCode string          `json:"authorization_code,omitempty"`
	Currency          string          `json:"currency,omitempty"`
	Amount            float32         `json:"amount,omitempty"`
	Email             string          `json:"email,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	InvoiceLimit      int             `json:"invoice_limit,omitempty"`
	Metadata          client.Metadata `json:"metadata,omitempty"`
	SubAccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            string          `json:"bearer,omitempty"`
	Channels          []string        `json:"channels,omitempty"`
}

// AuthorizationRequest represents a request to enable/revoke an authorization
type AuthorizationRequest struct {
	Reference         string          `json:"reference,omitempty"`
	AuthorizationCode string          `json:"authorization_code,omitempty"`
	Amount            int             `json:"amount,omitempty"`
	Currency          string          `json:"currency,omitempty"`
	Email             string          `json:"email,omitempty"`
	Metadata          client.Metadata `json:"metadata,omitempty"`
}

// Transaction is the resource representing your Paystack transaction.
// For more details see https://developers.paystack.co/v1.0/reference#initialize-a-transaction
type Transaction struct {
	ID              int                    `json:"id,omitempty"`
	CreatedAt       string                 `json:"createdAt,omitempty"`
	Domain          string                 `json:"domain,omitempty"`
	Metadata        string                 `json:"metadata,omitempty"` //TODO: why is transaction metadata a string?
	Status          string                 `json:"status,omitempty"`
	Reference       string                 `json:"reference,omitempty"`
	Amount          float32                `json:"amount,omitempty"`
	Message         string                 `json:"message,omitempty"`
	GatewayResponse string                 `json:"gateway_response,omitempty"`
	PaidAt          string                 `json:"piad_at,omitempty"`
	Channel         string                 `json:"channel,omitempty"`
	Currency        string                 `json:"currency,omitempty"`
	IPAddress       string                 `json:"ip_address,omitempty"`
	Log             map[string]interface{} `json:"log,omitempty"` // TODO: same as timeline?
	Fees            int                    `json:"int,omitempty"`
	FeesSplit       string                 `json:"fees_split,omitempty"` // TODO: confirm data type
	Customer        customer.Customer      `json:"customer,omitempty"`
	Authorization   Authorization          `json:"authorization,omitempty"`
	Plan            plan.Plan              `json:"plan,omitempty"`
	SubAccount      subaccount.SubAccount  `json:"sub_account,omitempty"`
}

// Authorization represents Paystack authorization object
type Authorization struct {
	AuthorizationCode string `json:"authorization_code,omitempty"`
	Bin               string `json:"bin,omitempty"`
	Last4             string `json:"last4,omitempty"`
	ExpMonth          string `json:"exp_month,omitempty"`
	ExpYear           string `json:"exp_year,omitempty"`
	Channel           string `json:"channel,omitempty"`
	CardType          string `json:"card_type,omitempty"`
	Bank              string `json:"bank,omitempty"`
	CountryCode       string `json:"country_code,omitempty"`
	Brand             string `json:"brand,omitempty"`
	Resusable         bool   `json:"reusable,omitempty"`
	Signature         string `json:"signature,omitempty"`
}

// Timeline represents a timeline of events in a transaction session
type Timeline struct {
	TimeSpent      int                      `json:"time_spent,omitempty"`
	Attempts       int                      `json:"attempts,omitempty"`
	Authentication string                   `json:"authentication,omitempty"` // TODO: confirm type
	Errors         int                      `json:"errors,omitempty"`
	Success        bool                     `json:"success,omitempty"`
	Mobile         bool                     `json:"mobile,omitempty"`
	Input          []string                 `json:"input,omitempty"` // TODO: confirm type
	Channel        string                   `json:"channel,omitempty"`
	History        []map[string]interface{} `json:"history,omitempty"`
}
