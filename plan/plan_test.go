package plan

import (
	"context"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
)

var c *client.Client
var service *DefaultPlanService

func init() {
	apiKey := client.MustGetTestKey()
	c = configuration.NewClient(apiKey, nil, true)
	service = &DefaultPlanService{Client: c}
}

func TestPlanCRUD(t *testing.T) {
	plan1 := &Plan{
		Name:     "Monthly retainer",
		Interval: "monthly",
		Amount:   500000,
	}

	// create the plan
	plan, err := service.Create(context.TODO(), plan1)
	if err != nil {
		t.Errorf("CREATE Plan returned error: %v", err)
	}

	if plan.PlanCode == "" {
		t.Errorf("Expected Plan code to be set")
	}

	// retrieve the plan
	plan, err = service.Get(context.TODO(), plan.ID)
	if err != nil {
		t.Errorf("GET Plan returned error: %v", err)
	}

	if plan.Name != plan1.Name {
		t.Errorf("Expected Plan Name %v, got %v", plan.Name, plan1.Name)
	}

	// retrieve the plan list
	plans, err := service.List(context.TODO())
	if err != nil || !(len(plans.Values) > 0) || !(plans.Meta.Total > 0) {
		t.Errorf("Expected Plan list, got %d, returned error %v", len(plans.Values), err)
	}
}
