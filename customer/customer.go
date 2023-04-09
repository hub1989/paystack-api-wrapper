package customer

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
	"net/url"
)

type Service interface {
	Create(ctx context.Context, customer *Customer) (*Customer, error)
	Update(ctx context.Context, customer *Customer) (*Customer, error)
	Get(ctx context.Context, customerCode string) (*Customer, error)
	List(ctx context.Context) (*List, error)
	ListN(ctx context.Context, count, offset int) (*List, error)
	SetRiskAction(ctx context.Context, customerCode, riskAction string) (*Customer, error)
	DeactivateAuthorization(ctx context.Context, authorizationCode string) (*response.Response, error)
	ValidateCustomer(ctx context.Context, customerId string, request *ValidateCustomerRequest) (bool, error)
}

// DefaultCustomerService handles operations related to the customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
type DefaultCustomerService struct {
	*client.Client
}

// Create creates a new customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
func (s *DefaultCustomerService) Create(ctx context.Context, customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customer")
	cust := &Customer{}
	err := s.Client.Call(ctx, http.MethodPost, u, customer, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-customer
func (s *DefaultCustomerService) Update(ctx context.Context, customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("customer/%d", customer.ID)
	cust := &Customer{}
	err := s.Client.Call(ctx, http.MethodPut, u, customer, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details see https://paystack.com/docs/api/#customer-fetch
func (s *DefaultCustomerService) Get(ctx context.Context, customerCode string) (*Customer, error) {
	u := fmt.Sprintf("/customer/%s", customerCode)
	cust := &Customer{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, cust)

	return cust, err
}

// List returns a list of customers.
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *DefaultCustomerService) List(ctx context.Context) (*List, error) {
	return s.ListN(ctx, 10, 0)
}

// ListN returns a list of customers
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *DefaultCustomerService) ListN(ctx context.Context, count, offset int) (*List, error) {
	u := client.PaginateURL("/customer", count, offset)
	cust := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, cust)
	return cust, err
}

// SetRiskAction can be used to either whitelist or blacklist a customer
// For more details see https://developers.paystack.co/v1.0/reference#whiteblacklist-customer
func (s *DefaultCustomerService) SetRiskAction(ctx context.Context, customerCode, riskAction string) (*Customer, error) {
	reqBody := struct {
		Customer    string `json:"customer"`
		Risk_action string `json:"risk_action"`
	}{
		Customer:    customerCode,
		Risk_action: riskAction,
	}
	cust := &Customer{}
	err := s.Client.Call(ctx, http.MethodPost, "/customer/set_risk_action", reqBody, cust)

	return cust, err
}

// DeactivateAuthorization deactivates an authorization
// For more details see https://developers.paystack.co/v1.0/reference#deactivate-authorization
func (s *DefaultCustomerService) DeactivateAuthorization(ctx context.Context, authorizationCode string) (*response.Response, error) {
	params := url.Values{}
	params.Add("authorization_code", authorizationCode)

	resp := &response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/customer/deactivate_authorization", params, resp)

	return resp, err
}

func (s *DefaultCustomerService) ValidateCustomer(ctx context.Context, customerId string, request *ValidateCustomerRequest) (bool, error) {
	endpoint := fmt.Sprintf("/customer/%s/identification", customerId)
	resp := &response.Response{}

	err := s.Client.Call(ctx, http.MethodPost, endpoint, request, &resp)
	if err != nil {
		return false, err
	}

	return true, nil
}
