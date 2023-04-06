package bulk_charge

import (
	"github.com/hub1989/paystack-api-wrapper/response"
)

// BulkChargeBatch represents a bulk charge batch object
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
type BulkChargeBatch struct {
	ID            int    `json:"id,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
	BatchCode     string `json:"batch_code,omitempty"`
	Status        string `json:"status,omitempty"`
	Integration   int    `json:"integration,omitempty"`
	Domain        string `json:"domain,omitempty"`
	TotalCharges  string `json:"total_charges,omitempty"`
	PendingCharge string `json:"pending_charge,omitempty"`
}

// BulkChargeRequest is an array of objects with authorization codes and amount
type BulkChargeRequest struct {
	Items []BulkItem
}

// BulkItem represents a single bulk charge request item
type BulkItem struct {
	Authorization string  `json:"authorization,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
}

// BulkChargeBatchList is a list object for bulkcharges.
type BulkChargeBatchList struct {
	Meta   response.ListMeta
	Values []BulkChargeBatch `json:"data,omitempty"`
}
