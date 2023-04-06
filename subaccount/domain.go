package subaccount

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
)

// SubAccount is the resource representing your Paystack subaccount.
// For more details see https://developers.paystack.co/v1.0/reference#create-subaccount
type SubAccount struct {
	ID                  int             `json:"id,omitempty"`
	CreatedAt           string          `json:"createdAt,omitempty"`
	UpdatedAt           string          `json:"updatedAt,omitempty"`
	Domain              string          `json:"domain,omitempty"`
	Integration         int             `json:"integration,omitempty"`
	BusinessName        string          `json:"business_name,omitempty"`
	SubAccountCode      string          `json:"subaccount_code,omitempty"`
	Description         string          `json:"description,omitempty"`
	PrimaryContactName  string          `json:"primary_contact_name,omitempty"`
	PrimaryContactEmail string          `json:"primary_contact_email,omitempty"`
	PrimaryContactPhone string          `json:"primary_contact_phone,omitempty"`
	Metadata            client.Metadata `json:"metadata,omitempty"`
	PercentageCharge    float32         `json:"percentage_charge,omitempty"`
	IsVerified          bool            `json:"is_verified,omitempty"`
	SettlementBank      string          `json:"settlement_bank,omitempty"`
	AccountNumber       string          `json:"account_number,omitempty"`
	SettlementSchedule  string          `json:"settlement_schedule,omitempty"`
	Active              bool            `json:"active,omitempty"`
	Migrate             bool            `json:"migrate,omitempty"`
}

// SubAccountList is a list object for subaccounts.
type SubAccountList struct {
	Meta   response.ListMeta
	Values []SubAccount `json:"data"`
}
