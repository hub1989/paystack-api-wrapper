package bank

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
)

type Service interface {
	List(ctx context.Context) (*List, error)
	ResolveBVN(ctx context.Context, bvn int) (*BVNResponse, error)
	ResolveAccountNumber(ctx context.Context, accountNumber, bankCode string) (response.Response, error)
}

// DefaultBankService handles operations related to the bank
// For more details see https://developers.paystack.co/v1.0/reference#bank
type DefaultBankService struct {
	*client.Client
}

// List returns a list of all the banks.
// For more details see https://developers.paystack.co/v1.0/reference#list-banks
func (s *DefaultBankService) List(ctx context.Context) (*List, error) {
	banks := &List{}
	err := s.Client.Call(ctx, http.MethodGet, "/bank", nil, banks)
	return banks, err
}

// ResolveBVN docs https://developers.paystack.co/v1.0/reference#resolve-bvn
func (s *DefaultBankService) ResolveBVN(ctx context.Context, bvn int) (*BVNResponse, error) {
	u := fmt.Sprintf("/bank/resolve_bvn/%d", bvn)
	resp := &BVNResponse{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, resp)
	return resp, err
}

// ResolveAccountNumber docs https://developers.paystack.co/v1.0/reference#resolve-account-number
func (s *DefaultBankService) ResolveAccountNumber(ctx context.Context, accountNumber, bankCode string) (response.Response, error) {
	u := fmt.Sprintf("/bank/resolve?account_number=%s&bank_code=%s", accountNumber, bankCode)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)
	return resp, err
}
