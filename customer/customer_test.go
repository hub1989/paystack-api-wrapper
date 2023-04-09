package customer

import (
	"context"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
)

var service *DefaultCustomerService

func init() {
	apiKey := client.MustGetTestKey()
	c := configuration.NewClient(apiKey, nil, true)
	service = &DefaultCustomerService{Client: c}
}

func TestCustomerCRUD(t *testing.T) {
	cust := &Customer{
		FirstName: "User123",
		LastName:  "AdminUser",
		Email:     "user123@gmail.com",
		Phone:     "+23400000000000000",
	}
	// create the customer
	customer, err := service.Create(context.TODO(), cust)
	if err != nil {
		t.Errorf("CREATE Customer returned error: %v", err)
	}

	// retrieve the customer
	customer, err = service.Get(context.TODO(), customer.CustomerCode)
	if err != nil {
		t.Errorf("GET Customer returned error: %v", err)
	}

	if customer.Email != cust.Email {
		t.Errorf("Expected Customer email %v, got %v", cust.Email, customer.Email)
	}

	if customer.FirstName != cust.FirstName {
		t.Errorf("Expected Customer first name %v, got %v", cust.FirstName, customer.FirstName)
	}

	if customer.LastName != cust.LastName {
		t.Errorf("Expected Customer last name %v, got %v", cust.FirstName, customer.LastName)
	}

	if customer.Phone != cust.Phone {
		t.Errorf("Expected Customer phone %v, got %v", cust.Phone, customer.Phone)
	}

	// retrieve the customer list
	customers, err := service.List(context.TODO())
	if err != nil || !(len(customers.Values) > 0) || !(customers.Meta.Total > 0) {
		t.Errorf("Expected Customer list, got %d, returned error %v", len(customers.Values), err)
	}

}

func TestCustomerRiskAction(t *testing.T) {
	cust := &Customer{
		FirstName: "User123",
		LastName:  "AdminUser",
		Email:     "user1-deny@gmail.com",
		Phone:     "+2341000000000000",
	}
	customer1, _ := service.Create(context.TODO(), cust)

	//TODO: investigate why 'allow' returns: 403 You cannot whitelist customers on this integration
	customer, err := service.SetRiskAction(context.TODO(), customer1.CustomerCode, "deny")
	if err != nil {
		t.Errorf("Customer risk action returned error %v", err)
	}

	if customer.Email != customer1.Email {
		t.Errorf("Expected Customer email %v, got %v", cust.Email, customer.Email)
	}
}

func TestListCustomer(t *testing.T) {
	customers, err := service.List(context.TODO())
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(customers.Values) == 0 {
		t.Errorf("Expected more than 0 zero customers, got %v", len(customers.Values))
	}
}
