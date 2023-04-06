package subscription

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/url"
)

type Service interface {
	Create(subscription *Request) (*Subscription, error)
	Update(subscription *Subscription) (*Subscription, error)
	Get(id int) (*Subscription, error)
	List() (*List, error)
	ListN(count, offset int) (*List, error)
	Enable(subscriptionCode, emailToken string) (response.Response, error)
	Disable(subscriptionCode, emailToken string) (response.Response, error)
}

// DefaultSubscriptionService handles operations related to the subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
type DefaultSubscriptionService struct {
	*client.Client
}

// Create creates a new subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
func (s *DefaultSubscriptionService) Create(subscription *Request) (*Subscription, error) {
	u := fmt.Sprintf("/subscription")
	sub := &Subscription{}
	err := s.Client.Call("POST", u, subscription, sub)
	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-subscription
func (s *DefaultSubscriptionService) Update(subscription *Subscription) (*Subscription, error) {
	u := fmt.Sprintf("subscription/%d", subscription.ID)
	sub := &Subscription{}
	err := s.Client.Call("PUT", u, subscription, sub)
	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-subscription
func (s *DefaultSubscriptionService) Get(id int) (*Subscription, error) {
	u := fmt.Sprintf("/subscription/%d", id)
	sub := &Subscription{}
	err := s.Client.Call("GET", u, nil, sub)
	return sub, err
}

// List returns a list of subscriptions.
// For more details see https://developers.paystack.co/v1.0/reference#list-subscriptions
func (s *DefaultSubscriptionService) List() (*List, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of subscriptions
// For more details see https://developers.paystack.co/v1.0/reference#list-subscriptions
func (s *DefaultSubscriptionService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/subscription", count, offset)
	sub := &List{}
	err := s.Client.Call("GET", u, nil, sub)
	return sub, err
}

// Enable enables a subscription
// For more details see https://developers.paystack.co/v1.0/reference#enable-subscription
func (s *DefaultSubscriptionService) Enable(subscriptionCode, emailToken string) (response.Response, error) {
	params := url.Values{}
	params.Add("code", subscriptionCode)
	params.Add("token", emailToken)
	resp := response.Response{}
	err := s.Client.Call("POST", "/subscription/enable", params, &resp)
	return resp, err
}

// Disable disables a subscription
// For more details see https://developers.paystack.co/v1.0/reference#disable-subscription
func (s *DefaultSubscriptionService) Disable(subscriptionCode, emailToken string) (response.Response, error) {
	params := url.Values{}
	params.Add("code", subscriptionCode)
	params.Add("token", emailToken)
	resp := response.Response{}
	err := s.Client.Call("POST", "/subscription/disable", params, &resp)
	return resp, err
}
