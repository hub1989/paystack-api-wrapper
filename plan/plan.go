package plan

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
)

type Service interface {
	Create(plan *Plan) (*Plan, error)
	Update(plan *Plan) (response.Response, error)
	Get(id int) (*Plan, error)
	List() (*List, error)
	ListN(count, offset int) (*List, error)
}

// DefaultPlanService handles operations related to the plan
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
type DefaultPlanService struct {
	*client.Client
}

// Create creates a new plan
// For more details see https://developers.paystack.co/v1.0/reference#create-plan
func (s *DefaultPlanService) Create(plan *Plan) (*Plan, error) {
	u := fmt.Sprintf("/plan")
	plan2 := &Plan{}
	err := s.Client.Call("POST", u, plan, plan2)
	return plan2, err
}

// Update updates a plan's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-plan
func (s *DefaultPlanService) Update(plan *Plan) (response.Response, error) {
	u := fmt.Sprintf("plan/%d", plan.ID)
	resp := response.Response{}
	err := s.Client.Call("PUT", u, plan, &resp)
	return resp, err
}

// Get returns the details of a plan.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-plan
func (s *DefaultPlanService) Get(id int) (*Plan, error) {
	u := fmt.Sprintf("/plan/%d", id)
	plan2 := &Plan{}
	err := s.Client.Call("GET", u, nil, plan2)
	return plan2, err
}

// List returns a list of plans.
// For more details see https://developers.paystack.co/v1.0/reference#list-plans
func (s *DefaultPlanService) List() (*List, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of plans
// For more details see https://developers.paystack.co/v1.0/reference#list-plans
func (s *DefaultPlanService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/plan", count, offset)
	plan2 := &List{}
	err := s.Client.Call("GET", u, nil, plan2)
	return plan2, err
}
