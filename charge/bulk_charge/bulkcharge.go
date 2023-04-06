package bulk_charge

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
)

type Service interface {
	Initiate(req *BulkChargeRequest) (*BulkChargeBatch, error)
	List() (*BulkChargeBatchList, error)
	ListN(count, offset int) (*BulkChargeBatchList, error)
	Get(idCode string) (*BulkChargeBatch, error)
	GetBatchCharges(idCode string) (response.Response, error)
	PauseBulkCharge(batchCode string) (response.Response, error)
	ResumeBulkCharge(batchCode string) (response.Response, error)
}

// DefaultBulkChargeService handles operations related to the bulkcharge
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
type DefaultBulkChargeService struct {
	client.Client
}

// Initiate initiates a new bulkcharge
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
func (s *DefaultBulkChargeService) Initiate(req *BulkChargeRequest) (*BulkChargeBatch, error) {
	bulkcharge := &BulkChargeBatch{}
	err := s.Client.Call("POST", "/bulkcharge", req.Items, bulkcharge)
	return bulkcharge, err
}

// List returns a list of bulkcharges.
// For more details see https://developers.paystack.co/v1.0/reference#list-bulkcharges
func (s *DefaultBulkChargeService) List() (*BulkChargeBatchList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of bulkcharges
// For more details see https://developers.paystack.co/v1.0/reference#list-bulkcharges
func (s *DefaultBulkChargeService) ListN(count, offset int) (*BulkChargeBatchList, error) {
	u := client.PaginateURL("/bulkcharge", count, offset)
	bulkcharges := &BulkChargeBatchList{}
	err := s.Client.Call("GET", u, nil, bulkcharges)
	return bulkcharges, err
}

// Get returns a bulk charge batch
// This endpoint retrieves a specific batch code.
// It also returns useful information on its progress by way of
// the total_charges and pending_charges attributes.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-bulk-charge-batch
func (s *DefaultBulkChargeService) Get(idCode string) (*BulkChargeBatch, error) {
	u := fmt.Sprintf("/bulkcharge/%s", idCode)
	bulkcharge := &BulkChargeBatch{}
	err := s.Client.Call("GET", u, nil, bulkcharge)
	return bulkcharge, err
}

// GetBatchCharges returns charges in a batch
// This endpoint retrieves the charges associated with a specified batch code.
// Pagination parameters are available. You can also filter by status.
// Charge statuses can be pending, success or failed.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-charges-in-a-batch
func (s *DefaultBulkChargeService) GetBatchCharges(idCode string) (response.Response, error) {
	u := fmt.Sprintf("/bulkcharge/%s/charges", idCode)
	resp := response.Response{}
	err := s.Client.Call("GET", u, nil, &resp)
	return resp, err
}

// PauseBulkCharge stops processing a batch
// For more details see https://developers.paystack.co/v1.0/reference#pause-bulk-charge-batch
func (s *DefaultBulkChargeService) PauseBulkCharge(batchCode string) (response.Response, error) {
	u := fmt.Sprintf("/bulkcharge/pause/%s", batchCode)
	resp := response.Response{}
	err := s.Client.Call("GET", u, nil, &resp)

	return resp, err
}

// ResumeBulkCharge stops processing a batch
// For more details see https://developers.paystack.co/v1.0/reference#resume-bulk-charge-batch
func (s *DefaultBulkChargeService) ResumeBulkCharge(batchCode string) (response.Response, error) {
	u := fmt.Sprintf("/bulkcharge/resume/%s", batchCode)
	resp := response.Response{}
	err := s.Client.Call("GET", u, nil, &resp)

	return resp, err
}
