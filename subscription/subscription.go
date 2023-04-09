package subscription

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
	"net/url"
)

type Service interface {
	Create(ctx context.Context, subscription *Request) (*Subscription, error)
	Update(ctx context.Context, subscription *Subscription) (*Subscription, error)
	Get(ctx context.Context, id int) (*Subscription, error)
	List(ctx context.Context) (*List, error)
	ListN(ctx context.Context, count, offset int) (*List, error)
	Enable(ctx context.Context, subscriptionCode, emailToken string) (response.Response, error)
	Disable(ctx context.Context, subscriptionCode, emailToken string) (response.Response, error)
}

// DefaultSubscriptionService handles operations related to the subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
type DefaultSubscriptionService struct {
	*client.Client
}

// Create creates a new subscription
// For more details see https://developers.paystack.co/v1.0/reference#create-subscription
func (s *DefaultSubscriptionService) Create(ctx context.Context, subscription *Request) (*Subscription, error) {
	u := fmt.Sprintf("/subscription")
	sub := &Subscription{}
	err := s.Client.Call(ctx, http.MethodPost, u, subscription, sub)
	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-subscription
func (s *DefaultSubscriptionService) Update(ctx context.Context, subscription *Subscription) (*Subscription, error) {
	u := fmt.Sprintf("subscription/%d", subscription.ID)
	sub := &Subscription{}
	err := s.Client.Call(ctx, http.MethodPut, u, subscription, sub)
	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-subscription
func (s *DefaultSubscriptionService) Get(ctx context.Context, id int) (*Subscription, error) {
	u := fmt.Sprintf("/subscription/%d", id)
	sub := &Subscription{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, sub)
	return sub, err
}

// List returns a list of subscriptions.
// For more details see https://developers.paystack.co/v1.0/reference#list-subscriptions
func (s *DefaultSubscriptionService) List(ctx context.Context) (*List, error) {
	return s.ListN(ctx, 10, 0)
}

// ListN returns a list of subscriptions
// For more details see https://developers.paystack.co/v1.0/reference#list-subscriptions
func (s *DefaultSubscriptionService) ListN(ctx context.Context, count, offset int) (*List, error) {
	u := client.PaginateURL("/subscription", count, offset)
	sub := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, sub)
	return sub, err
}

// Enable enables a subscription
// For more details see https://developers.paystack.co/v1.0/reference#enable-subscription
func (s *DefaultSubscriptionService) Enable(ctx context.Context, subscriptionCode, emailToken string) (response.Response, error) {
	params := url.Values{}
	params.Add("code", subscriptionCode)
	params.Add("token", emailToken)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/subscription/enable", params, &resp)
	return resp, err
}

// Disable disables a subscription
// For more details see https://developers.paystack.co/v1.0/reference#disable-subscription
func (s *DefaultSubscriptionService) Disable(ctx context.Context, subscriptionCode, emailToken string) (response.Response, error) {
	params := url.Values{}
	params.Add("code", subscriptionCode)
	params.Add("token", emailToken)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/subscription/disable", params, &resp)
	return resp, err
}
