package transaction

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
)

type Service interface {
	Initialize(txn *Request) (response.Response, error)
	Verify(reference string) (*Transaction, error)
	List() (*List, error)
	ListForCustomer(customerId string) (*List, error)
	ListN(count, offset int) (*List, error)
	Get(id int) (*Transaction, error)
	ChargeAuthorization(req *Request) (*Transaction, error)
	Timeline(reference string) (*Timeline, error)
	Totals() (response.Response, error)
	Export(params response.RequestValues) (response.Response, error)
	ReAuthorize(req AuthorizationRequest) (response.Response, error)
	CheckAuthorization(req AuthorizationRequest) (response.Response, error)
}

// DefaultTransactionService handles operations related to transactions
// For more details see https://developers.paystack.co/v1.0/reference#create-transaction
type DefaultTransactionService struct {
	*client.Client
}

// Initialize initiates a transaction process
// For more details see https://developers.paystack.co/v1.0/reference#initialize-a-transaction
func (s *DefaultTransactionService) Initialize(txn *Request) (response.Response, error) {
	u := fmt.Sprintf("/transaction/initialize")
	resp := response.Response{}
	err := s.Client.Call("POST", u, txn, &resp)
	return resp, err
}

// Verify checks that transaction with the given reference exists
// For more details see https://api.paystack.co/transaction/verify/reference
func (s *DefaultTransactionService) Verify(reference string) (*Transaction, error) {
	u := fmt.Sprintf("/transaction/verify/%s", reference)
	txn := &Transaction{}
	err := s.Client.Call("GET", u, nil, txn)
	return txn, err
}

// List returns a list of transactions.
// For more details see https://paystack.com/docs/api/#transaction-list
func (s *DefaultTransactionService) List() (*List, error) {
	return s.ListN(10, 1)
}

// ListN returns a list of transactions
// For more details see https://developers.paystack.co/v1.0/reference#list-transactions
func (s *DefaultTransactionService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/transaction", count, offset)
	txns := &List{}
	err := s.Client.Call("GET", u, nil, txns)
	return txns, err
}

// Get returns the details of a transaction.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-transaction
func (s *DefaultTransactionService) Get(id int) (*Transaction, error) {
	u := fmt.Sprintf("/transaction/%d", id)
	txn := &Transaction{}
	err := s.Client.Call("GET", u, nil, txn)
	return txn, err
}

// ChargeAuthorization is for charging all  authorizations marked as reusable whenever you need to recieve payments.
// For more details see https://developers.paystack.co/v1.0/reference#charge-authorization
func (s *DefaultTransactionService) ChargeAuthorization(req *Request) (*Transaction, error) {
	txn := &Transaction{}
	err := s.Client.Call("POST", "/transaction/charge_authorization", req, txn)
	return txn, err
}

// Timeline fetches the transaction timeline. Reference can be ID or transaction reference
// For more details see https://developers.paystack.co/v1.0/reference#view-transaction-timeline
func (s *DefaultTransactionService) Timeline(reference string) (*Timeline, error) {
	u := fmt.Sprintf("/transaction/timeline/%s", reference)
	timeline := &Timeline{}
	err := s.Client.Call("GET", u, nil, timeline)
	return timeline, err
}

// Totals returns total amount received on your account
// For more details see https://developers.paystack.co/v1.0/reference#transaction-totals
func (s *DefaultTransactionService) Totals() (response.Response, error) {
	u := fmt.Sprintf("/transaction/totals")
	resp := response.Response{}
	err := s.Client.Call("GET", u, nil, &resp)
	return resp, err
}

// Export exports transactions to a downloadable file and returns a link to the file
// For more details see https://developers.paystack.co/v1.0/reference#export-transactions
func (s *DefaultTransactionService) Export(params response.RequestValues) (response.Response, error) {
	u := fmt.Sprintf("/transaction/export")
	resp := response.Response{}
	err := s.Client.Call("GET", u, nil, &resp)
	return resp, err
}

// ReAuthorize requests reauthorization
// For more details see https://developers.paystack.co/v1.0/reference#request-reauthorization
func (s *DefaultTransactionService) ReAuthorize(req AuthorizationRequest) (response.Response, error) {
	u := fmt.Sprintf("/transaction/request_reauthorization")
	resp := response.Response{}
	err := s.Client.Call("POST", u, nil, &resp)
	return resp, err
}

// CheckAuthorization checks authorization
// For more details see https://developers.paystack.co/v1.0/reference#check-authorization
func (s *DefaultTransactionService) CheckAuthorization(req AuthorizationRequest) (response.Response, error) {
	u := fmt.Sprintf("/transaction/check_reauthorization")
	resp := response.Response{}
	err := s.Client.Call("POST", u, nil, &resp)
	return resp, err
}

func (s *DefaultTransactionService) ListForCustomer(customerId string) (*List, error) {
	u := fmt.Sprintf("/transaction?customer=%s", customerId)
	txns := &List{}
	err := s.Client.Call("GET", u, nil, txns)
	return txns, err
}
