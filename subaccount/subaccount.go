package subaccount

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"net/http"
)

type Service interface {
	Create(ctx context.Context, subaccount *SubAccount) (*SubAccount, error)
	Update(ctx context.Context, subaccount *SubAccount) (*SubAccount, error)
	Get(ctx context.Context, id int) (*SubAccount, error)
	List(ctx context.Context) (*SubAccountList, error)
	ListN(ctx context.Context, count, offset int) (*SubAccountList, error)
}

// DefaultSubAccountService handles operations related to subaccounts
// For more details see https://developers.paystack.co/v1.0/reference#create-subaccount
type DefaultSubAccountService struct {
	*client.Client
}

// Create creates a new subaccount
// For more details see https://paystack.com/docs/api/#subaccount-create
func (s *DefaultSubAccountService) Create(ctx context.Context, subaccount *SubAccount) (*SubAccount, error) {
	u := fmt.Sprintf("/subaccount")
	acc := &SubAccount{}
	err := s.Client.Call(ctx, http.MethodPost, u, subaccount, acc)
	return acc, err
}

// Update updates a subaccount's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-subaccount
// TODO: use ID or slug
func (s *DefaultSubAccountService) Update(ctx context.Context, subaccount *SubAccount) (*SubAccount, error) {
	u := fmt.Sprintf("subaccount/%d", subaccount.ID)
	acc := &SubAccount{}
	err := s.Client.Call(ctx, http.MethodPut, u, subaccount, acc)

	return acc, err
}

// Get returns the details of a subaccount.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-subaccount
// TODO: use ID or slug
func (s *DefaultSubAccountService) Get(ctx context.Context, id int) (*SubAccount, error) {
	u := fmt.Sprintf("/subaccount/%d", id)
	acc := &SubAccount{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, acc)

	return acc, err
}

// List returns a list of subaccounts.
// For more details see https://developers.paystack.co/v1.0/reference#list-subaccounts
func (s *DefaultSubAccountService) List(ctx context.Context) (*SubAccountList, error) {
	return s.ListN(ctx, 10, 1)
}

// ListN returns a list of subaccounts
// For more details see https://paystack.com/docs/api/#subaccount-list
func (s *DefaultSubAccountService) ListN(ctx context.Context, count, offset int) (*SubAccountList, error) {
	u := client.PaginateURL("/subaccount", count, offset)
	acc := &SubAccountList{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, acc)
	return acc, err
}
