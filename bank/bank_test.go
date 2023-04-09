package bank

import (
	"context"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
)

var c *client.Client
var service *DefaultBankService

func init() {
	apiKey := client.MustGetTestKey()
	c = configuration.NewClient(apiKey, nil, true)
	service = &DefaultBankService{Client: c}
}

func TestBankList(t *testing.T) {
	// retrieve the bank list
	banks, err := service.List(context.TODO())

	if err != nil || !(len(banks.Values) > 0) {
		t.Errorf("Expected Bank list, got %d, returned error %v", len(banks.Values), err)
	}
}

func TestResolveBVN(t *testing.T) {
	// Test invlaid BVN.
	// Err not nill. Resp status code is 400
	resp, err := service.ResolveBVN(context.Background(), 21212917)
	if err == nil {
		t.Errorf("Expected error for invalid BVN, got %+v'", resp)
	}

	// Test free calls limit
	// Error is nil
	// &{Meta:{CallsThisMonth:0 FreeCallsLeft:0} BVN:cZ+MKrsLAqJCUi+hxIdQqw==}’
	resp, err = service.ResolveBVN(context.TODO(), 21212917741)
	if resp.Meta.FreeCallsLeft != 0 {
		t.Errorf("Expected free calls limit exceeded, got %+v'", resp)
	}
	// TODO(yao): Reproduce error: Your balance is not enough to fulfill this request
}

func TestResolveAccountNumber(t *testing.T) {
	resp, err := service.ResolveAccountNumber(context.TODO(), "0022728151", "063")
	if err == nil {
		t.Errorf("Expected error, got %+v'", resp)
	}

	/*
		if _, ok := resp["account_number"]; !ok {
			t.Errorf("Expected response to contain 'account_number'")
		}

		if _, ok := resp["account_name"]; !ok {
			t.Errorf("Expected response to contain 'account_name'")
		}
	*/
}
