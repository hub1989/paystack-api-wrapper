package settlement

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
)

var c *client.Client
var service DefaultSettlementService

func init() {
	apiKey := client.MustGetTestKey()
	c = configuration.NewClient(apiKey, nil, true)
	service = DefaultSettlementService{Client: c}
}

func TestSettlementList(t *testing.T) {
	// retrieve the settlement list
	settlements, err := service.List(context.TODO())

	if err != nil {
		t.Error(err)
	}

	if err == nil {
		fmt.Printf("Settlements total: %d", len(settlements.Values))
	}
}
