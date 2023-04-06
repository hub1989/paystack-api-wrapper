package transfer

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
)

type Request struct {
	Source    string  `json:"source,omitempty"`
	Amount    float32 `json:"amount,omitempty"`
	Currency  string  `json:"currency,omitempty"`
	Reason    string  `json:"reason,omitempty"`
	Recipient string  `json:"recipient,omitempty"`
}

// Transfer is the resource representing your Paystack transfer.
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
type Transfer struct {
	ID           int     `json:"id,omitempty"`
	CreatedAt    string  `json:"createdAt,omitempty"`
	UpdatedAt    string  `json:"updatedAt,omitempty"`
	Domain       string  `json:"domain,omitempty"`
	Integration  int     `json:"integration,omitempty"`
	Source       string  `json:"source,omitempty"`
	Amount       float32 `json:"amount,omitempty"`
	Currency     string  `json:"currency,omitempty"`
	Reason       string  `json:"reason,omitempty"`
	TransferCode string  `json:"transfer_code,omitempty"`
	// Initiate returns recipient ID as recipient value, Fetch returns recipient object
	Recipient interface{} `json:"recipient,omitempty"`
	Status    string      `json:"status,omitempty"`
	// confirm types for source_details and failures
	SourceDetails interface{} `json:"source_details,omitempty"`
	Failures      interface{} `json:"failures,omitempty"`
	TransferredAt string      `json:"transferred_at,omitempty"`
	TitanCode     string      `json:"titan_code,omitempty"`
}

// Recipient represents a Paystack transfer recipient
// For more details see https://developers.paystack.co/v1.0/reference#create-transfer-recipient
type Recipient struct {
	ID            int                    `json:"id,omitempty"`
	CreatedAt     string                 `json:"createdAt,omitempty"`
	UpdatedAt     string                 `json:"updatedAt,omitempty"`
	Type          string                 `json:",omitempty"`
	Name          string                 `json:"name,omitempty"`
	Metadata      client.Metadata        `json:"metadata,omitempty"`
	AccountNumber string                 `json:"account_number,omitempty"`
	BankCode      string                 `json:"bank_code,omitempty"`
	Currency      string                 `json:"currency,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Active        bool                   `json:"active,omitempty"`
	Details       map[string]interface{} `json:"details,omitempty"`
	Domain        string                 `json:"domain,omitempty"`
	RecipientCode string                 `json:"recipient_code,omitempty"`
}

// BulkTransfer represents a Paystack bulk transfer
// You need to disable the Transfers OTP requirement to use this endpoint
type BulkTransfer struct {
	Currency  string                   `json:"currency,omitempty"`
	Source    string                   `json:"source,omitempty"`
	Transfers []map[string]interface{} `json:"transfers,omitempty"`
}

// List is a list object for transfers.
type List struct {
	Meta   response.ListMeta
	Values []Transfer `json:"data,omitempty"`
}

// RecipientList is a list object for transfer recipient.
type RecipientList struct {
	Meta   response.ListMeta
	Values []Recipient `json:"data,omitempty"`
}
