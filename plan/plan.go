package plan

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
)

type Service interface {
	Create(ctx context.Context, plan *Plan) (*Plan, error)
	Update(ctx context.Context, plan *Plan) (response.Response, error)
	Get(ctx context.Context, id int) (*Plan, error)
	List(ctx context.Context) (*List, error)
	ListN(ctx context.Context, count, offset int) (*List, error)
}

// DefaultPlanService handles operations related to the plan
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
type DefaultPlanService struct {
	*client.Client
}

// Create creates a new plan
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
func (s *DefaultPlanService) Create(ctx context.Context, plan *Plan) (*Plan, error) {
	u := fmt.Sprintf("/plan")
	plan2 := &Plan{}
	err := s.Client.Call(ctx, http.MethodPost, u, plan, plan2)
	return plan2, err
}

// Update updates a plan's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-plan
func (s *DefaultPlanService) Update(ctx context.Context, plan *Plan) (response.Response, error) {
	u := fmt.Sprintf("plan/%d", plan.ID)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPut, u, plan, &resp)
	return resp, err
}

// Get returns the details of a plan.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-plan
func (s *DefaultPlanService) Get(ctx context.Context, id int) (*Plan, error) {
	u := fmt.Sprintf("/plan/%d", id)
	plan2 := &Plan{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, plan2)
	return plan2, err
}

// List returns a list of plans.
// For more details see https://developers.paystack.co/v1.0/reference#list-plans
func (s *DefaultPlanService) List(ctx context.Context) (*List, error) {
	return s.ListN(ctx, 10, 0)
}

// ListN returns a list of plans
// For more details see https://developers.paystack.co/v1.0/reference#list-plans
func (s *DefaultPlanService) ListN(ctx context.Context, count, offset int) (*List, error) {
	u := client.PaginateURL("/plan", count, offset)
	plan2 := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, plan2)
	return plan2, err
}
