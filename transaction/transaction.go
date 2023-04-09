package transaction

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
)

type Service interface {
	Initialize(ctx context.Context, txn *Request) (response.Response, error)
	Verify(ctx context.Context, reference string) (*Transaction, error)
	List(ctx context.Context) (*List, error)
	ListForCustomer(ctx context.Context, customerId string) (*List, error)
	ListN(ctx context.Context, count, offset int) (*List, error)
	Get(ctx context.Context, id int) (*Transaction, error)
	ChargeAuthorization(ctx context.Context, req *Request) (*Transaction, error)
	Timeline(ctx context.Context, reference string) (*Timeline, error)
	Totals(ctx context.Context) (response.Response, error)
	Export(ctx context.Context, params response.RequestValues) (response.Response, error)
	ReAuthorize(ctx context.Context, req AuthorizationRequest) (response.Response, error)
	CheckAuthorization(ctx context.Context, req AuthorizationRequest) (response.Response, error)
}

// DefaultTransactionService handles operations related to transactions
// For more details see https://developers.paystack.co/v1.0/reference#create-transaction
type DefaultTransactionService struct {
	*client.Client
}

// Initialize initiates a transaction process
// For more details see https://developers.paystack.co/v1.0/reference#initialize-a-transaction
func (s *DefaultTransactionService) Initialize(ctx context.Context, txn *Request) (response.Response, error) {
	u := fmt.Sprintf("/transaction/initialize")
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, u, txn, &resp)
	return resp, err
}

// Verify checks that transaction with the given reference exists
// For more details see https://api.paystack.co/transaction/verify/reference
func (s *DefaultTransactionService) Verify(ctx context.Context, reference string) (*Transaction, error) {
	u := fmt.Sprintf("/transaction/verify/%s", reference)
	txn := &Transaction{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, txn)
	return txn, err
}

// List returns a list of transactions.
// For more details see https://paystack.com/docs/api/#transaction-list
func (s *DefaultTransactionService) List(ctx context.Context) (*List, error) {
	return s.ListN(ctx, 10, 1)
}

// ListN returns a list of transactions
// For more details see https://developers.paystack.co/v1.0/reference#list-transactions
func (s *DefaultTransactionService) ListN(ctx context.Context, count, offset int) (*List, error) {
	u := client.PaginateURL("/transaction", count, offset)
	txns := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, txns)
	return txns, err
}

// Get returns the details of a transaction.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-transaction
func (s *DefaultTransactionService) Get(ctx context.Context, id int) (*Transaction, error) {
	u := fmt.Sprintf("/transaction/%d", id)
	txn := &Transaction{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, txn)
	return txn, err
}

// ChargeAuthorization is for charging all  authorizations marked as reusable whenever you need to recieve payments.
// For more details see https://developers.paystack.co/v1.0/reference#charge-authorization
func (s *DefaultTransactionService) ChargeAuthorization(ctx context.Context, req *Request) (*Transaction, error) {
	txn := &Transaction{}
	err := s.Client.Call(ctx, http.MethodPost, "/transaction/charge_authorization", req, txn)
	return txn, err
}

// Timeline fetches the transaction timeline. Reference can be ID or transaction reference
// For more details see https://developers.paystack.co/v1.0/reference#view-transaction-timeline
func (s *DefaultTransactionService) Timeline(ctx context.Context, reference string) (*Timeline, error) {
	u := fmt.Sprintf("/transaction/timeline/%s", reference)
	timeline := &Timeline{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, timeline)
	return timeline, err
}

// Totals returns total amount received on your account
// For more details see https://developers.paystack.co/v1.0/reference#transaction-totals
func (s *DefaultTransactionService) Totals(ctx context.Context) (response.Response, error) {
	u := fmt.Sprintf("/transaction/totals")
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)
	return resp, err
}

// Export exports transactions to a downloadable file and returns a link to the file
// For more details see https://developers.paystack.co/v1.0/reference#export-transactions
func (s *DefaultTransactionService) Export(ctx context.Context, params response.RequestValues) (response.Response, error) {
	u := fmt.Sprintf("/transaction/export")
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)
	return resp, err
}

// ReAuthorize requests reauthorization
// For more details see https://developers.paystack.co/v1.0/reference#request-reauthorization
func (s *DefaultTransactionService) ReAuthorize(ctx context.Context, req AuthorizationRequest) (response.Response, error) {
	u := fmt.Sprintf("/transaction/request_reauthorization")
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, u, nil, &resp)
	return resp, err
}

// CheckAuthorization checks authorization
// For more details see https://developers.paystack.co/v1.0/reference#check-authorization
func (s *DefaultTransactionService) CheckAuthorization(ctx context.Context, req AuthorizationRequest) (response.Response, error) {
	u := fmt.Sprintf("/transaction/check_reauthorization")
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, u, nil, &resp)
	return resp, err
}

func (s *DefaultTransactionService) ListForCustomer(ctx context.Context, customerId string) (*List, error) {
	u := fmt.Sprintf("/transaction?customer=%s", customerId)
	txns := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, txns)
	return txns, err
}
