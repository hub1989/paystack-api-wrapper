package charge

import "github.com/hub1989/paystack-api-wrapper/client"

// Card represents a Card object
type Card struct {
	Number            string `json:"card_number,omitempty"`
	CVV               string `json:"card_cvc,omitempty"`
	ExpirtyMonth      string `json:"expiry_month,omitempty"`
	ExpiryYear        string `json:"expiry_year,omitempty"`
	AddressLine1      string `json:"address_line1,omitempty"`
	AddressLine2      string `json:"address_line2,omitempty"`
	AddressLine3      string `json:"address_line3,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	AddressPostalCode string `json:"address_postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
}

// BankAccount is used as bank in a charge request
type BankAccount struct {
	Code          string `json:"code,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
}

// ChargeRequest represents a Paystack charge request
type ChargeRequest struct {
	Email             string           `json:"email,omitempty"`
	Amount            float32          `json:"amount,omitempty"`
	Birthday          string           `json:"birthday,omitempty"`
	Card              *Card            `json:"card,omitempty"`
	Bank              *BankAccount     `json:"bank,omitempty"`
	AuthorizationCode string           `json:"authorization_code,omitempty"`
	Pin               string           `json:"pin,omitempty"`
	Metadata          *client.Metadata `json:"metadata,omitempty"`
	Reference         string           `json:"reference,omitempty"`
}
