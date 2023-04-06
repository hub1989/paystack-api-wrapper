package subscription

//import (
//	"github.com/hub1989/paystack-api-wrapper/client"
//	"github.com/hub1989/paystack-api-wrapper/configuration"
//	"github.com/hub1989/paystack-api-wrapper/customer"
//	plan2 "github.com/hub1989/paystack-api-wrapper/plan"
//	"testing"
//)
//
//var c *client.Client
//var service DefaultSubscriptionService
//var customerService customer.DefaultCustomerService
//var planService plan2.DefaultPlanService
//
//func init() {
//	apiKey := client.MustGetTestKey()
//	c = configuration.NewClient(apiKey, nil)
//	service = DefaultSubscriptionService{Client: c}
//	customerService = customer.DefaultCustomerService{Client: c}
//	planService = plan2.DefaultPlanService{Client: c}
//}
//
//func TestSubscriptionCRUD(t *testing.T) {
//	cust := &customer.Customer{
//		FirstName: "User123",
//		LastName:  "AdminUser",
//		Email:     "user123-subscription@gmail.com",
//		Phone:     "+23400000000000000",
//	}
//	// create the customer
//	create, err := customerService.Create(cust)
//	if err != nil {
//		t.Errorf("CREATE Subscription Customer returned error: %v", err)
//	}
//
//	plan1 := &plan2.Plan{
//		Name:     "Monthly subscription retainer",
//		Interval: "monthly",
//		Amount:   250000,
//	}
//
//	// create the plan
//	plan, err := planService.Create(plan1)
//	if err != nil {
//		t.Errorf("CREATE Plan returned error: %v", err)
//	}
//
//	subscription1 := &Request{
//		Customer: create.CustomerCode,
//		Plan:     plan.PlanCode,
//	}
//
//	// create the subscription
//	_, err = service.Create(subscription1)
//	if err == nil {
//		t.Errorf("Expected CREATE Subscription to fail with aunthorized create, got %+v", err)
//	}
//
//	// retrieve the subscription list
//	subscriptions, err := service.List()
//	if err != nil {
//		t.Errorf("Expected Subscription list, got %d, returned error %v", len(subscriptions.Values), err)
//	}
//}
