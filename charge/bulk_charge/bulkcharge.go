package bulk_charge

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
)

type Service interface {
	Initiate(ctx context.Context, req *BulkChargeRequest) (*BulkChargeBatch, error)
	List(ctx context.Context) (*BulkChargeBatchList, error)
	ListN(ctx context.Context, count, offset int) (*BulkChargeBatchList, error)
	Get(ctx context.Context, idCode string) (*BulkChargeBatch, error)
	GetBatchCharges(ctx context.Context, idCode string) (response.Response, error)
	PauseBulkCharge(ctx context.Context, batchCode string) (response.Response, error)
	ResumeBulkCharge(ctx context.Context, batchCode string) (response.Response, error)
}

// DefaultBulkChargeService handles operations related to the bulkcharge
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
type DefaultBulkChargeService struct {
	client.Client
}

// Initiate initiates a new bulkcharge
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-charge
func (s *DefaultBulkChargeService) Initiate(ctx context.Context, req *BulkChargeRequest) (*BulkChargeBatch, error) {
	bulkcharge := &BulkChargeBatch{}
	err := s.Client.Call(ctx, http.MethodPost, "/bulkcharge", req.Items, bulkcharge)
	return bulkcharge, err
}

// List returns a list of bulkcharges.
// For more details see https://developers.paystack.co/v1.0/reference#list-bulkcharges
func (s *DefaultBulkChargeService) List(ctx context.Context) (*BulkChargeBatchList, error) {
	return s.ListN(ctx, 10, 0)
}

// ListN returns a list of bulkcharges
// For more details see https://developers.paystack.co/v1.0/reference#list-bulkcharges
func (s *DefaultBulkChargeService) ListN(ctx context.Context, count, offset int) (*BulkChargeBatchList, error) {
	u := client.PaginateURL("/bulkcharge", count, offset)
	bulkcharges := &BulkChargeBatchList{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, bulkcharges)
	return bulkcharges, err
}

// Get returns a bulk charge batch
// This endpoint retrieves a specific batch code.
// It also returns useful information on its progress by way of
// the total_charges and pending_charges attributes.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-bulk-charge-batch
func (s *DefaultBulkChargeService) Get(ctx context.Context, idCode string) (*BulkChargeBatch, error) {
	u := fmt.Sprintf("/bulkcharge/%s", idCode)
	bulkcharge := &BulkChargeBatch{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, bulkcharge)
	return bulkcharge, err
}

// GetBatchCharges returns charges in a batch
// This endpoint retrieves the charges associated with a specified batch code.
// Pagination parameters are available. You can also filter by status.
// Charge statuses can be pending, success or failed.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-charges-in-a-batch
func (s *DefaultBulkChargeService) GetBatchCharges(ctx context.Context, idCode string) (response.Response, error) {
	u := fmt.Sprintf("/bulkcharge/%s/charges", idCode)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)
	return resp, err
}

// PauseBulkCharge stops processing a batch
// For more details see https://developers.paystack.co/v1.0/reference#pause-bulk-charge-batch
func (s *DefaultBulkChargeService) PauseBulkCharge(ctx context.Context, batchCode string) (response.Response, error) {
	u := fmt.Sprintf("/bulkcharge/pause/%s", batchCode)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)

	return resp, err
}

// ResumeBulkCharge stops processing a batch
// For more details see https://developers.paystack.co/v1.0/reference#resume-bulk-charge-batch
func (s *DefaultBulkChargeService) ResumeBulkCharge(ctx context.Context, batchCode string) (response.Response, error) {
	u := fmt.Sprintf("/bulkcharge/resume/%s", batchCode)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)

	return resp, err
}
