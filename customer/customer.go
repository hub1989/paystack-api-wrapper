package customer

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/url"
)

type Service interface {
	Create(customer *Customer) (*Customer, error)
	Update(customer *Customer) (*Customer, error)
	Get(customerCode string) (*Customer, error)
	List() (*List, error)
	ListN(count, offset int) (*List, error)
	SetRiskAction(customerCode, riskAction string) (*Customer, error)
	DeactivateAuthorization(authorizationCode string) (*response.Response, error)
}

// DefaultCustomerService handles operations related to the customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
type DefaultCustomerService struct {
	*client.Client
}

// Create creates a new customer
// For more details see https://developers.paystack.co/v1.0/reference#create-customer
func (s *DefaultCustomerService) Create(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("/customer")
	cust := &Customer{}
	err := s.Client.Call("POST", u, customer, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-customer
func (s *DefaultCustomerService) Update(customer *Customer) (*Customer, error) {
	u := fmt.Sprintf("customer/%d", customer.ID)
	cust := &Customer{}
	err := s.Client.Call("PUT", u, customer, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details see https://paystack.com/docs/api/#customer-fetch
func (s *DefaultCustomerService) Get(customerCode string) (*Customer, error) {
	u := fmt.Sprintf("/customer/%s", customerCode)
	cust := &Customer{}
	err := s.Client.Call("GET", u, nil, cust)

	return cust, err
}

// List returns a list of customers.
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *DefaultCustomerService) List() (*List, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of customers
// For more details see https://developers.paystack.co/v1.0/reference#list-customers
func (s *DefaultCustomerService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/customer", count, offset)
	cust := &List{}
	err := s.Client.Call("GET", u, nil, cust)
	return cust, err
}

// SetRiskAction can be used to either whitelist or blacklist a customer
// For more details see https://developers.paystack.co/v1.0/reference#whiteblacklist-customer
func (s *DefaultCustomerService) SetRiskAction(customerCode, riskAction string) (*Customer, error) {
	reqBody := struct {
		Customer    string `json:"customer"`
		Risk_action string `json:"risk_action"`
	}{
		Customer:    customerCode,
		Risk_action: riskAction,
	}
	cust := &Customer{}
	err := s.Client.Call("POST", "/customer/set_risk_action", reqBody, cust)

	return cust, err
}

// DeactivateAuthorization deactivates an authorization
// For more details see https://developers.paystack.co/v1.0/reference#deactivate-authorization
func (s *DefaultCustomerService) DeactivateAuthorization(authorizationCode string) (*response.Response, error) {
	params := url.Values{}
	params.Add("authorization_code", authorizationCode)

	resp := &response.Response{}
	err := s.Client.Call("POST", "/customer/deactivate_authorization", params, resp)

	return resp, err
}
