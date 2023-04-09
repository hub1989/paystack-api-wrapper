package transaction

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
	"time"
)

var c *client.Client
var service DefaultTransactionService

func init() {
	apiKey := client.MustGetTestKey()
	c = configuration.NewClient(apiKey, nil, true)
	service = DefaultTransactionService{Client: c}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func TestInitializeTransaction(t *testing.T) {
	txn := &Request{
		Email:     "user123@gmail.com",
		Amount:    6000,
		Reference: "Txn-" + fmt.Sprintf("%d", makeTimestamp()),
	}
	resp, err := service.Initialize(context.TODO(), txn)
	if err != nil {
		t.Error(err)
	}

	if resp["authorization_code"] == "" {
		t.Error("Missing transaction authorization code")
	}

	if resp["access_code"] == "" {
		t.Error("Missing transaction access code")
	}

	if resp["reference"] == "" {
		t.Error("Missing transaction reference")
	}

	txn1, err := service.Verify(context.TODO(), resp["reference"].(string))

	if err != nil {
		t.Error(err)
	}

	if txn1.Amount != txn.Amount {
		t.Errorf("Expected transaction amount %f, got %+v", txn.Amount, txn1.Amount)
	}

	if txn1.Reference == "" {
		t.Errorf("Missing transaction reference")
	}

	_, err = service.Get(context.TODO(), txn1.ID)

	if err != nil {
		t.Error(err)
	}
}

func TestTransactionList(t *testing.T) {
	// retrieve the transaction list
	transactions, err := service.List(context.TODO())
	if err != nil {
		t.Errorf("Expected Transaction list, got %d, returned error %v", len(transactions.Values), err)
	}
}

func TestTransactionTotals(t *testing.T) {
	_, err := service.Totals(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestExportTransaction(t *testing.T) {
	resp, err := service.Export(context.TODO(), nil)
	if err != nil {
		t.Error(err)
	}

	if _, ok := resp["path"]; !ok {
		t.Error("Expected transactiion export path")
	}
}
