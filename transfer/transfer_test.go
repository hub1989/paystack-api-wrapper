package transfer

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
)

var c *client.Client
var service *DefaultTransferService

func init() {
	apiKey := client.MustGetTestKey()
	c = configuration.NewClient(apiKey, nil)
	service = &DefaultTransferService{Client: c}
}

func TestInitiateTransfer(t *testing.T) {
	_, err := service.EnableOTP()
	if err != nil {
		t.Error(err)
	}

	recipient := &Recipient{
		Type:          "Nuban",
		Name:          "Customer 1",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

	recipient1, err := service.CreateRecipient(recipient)

	req := &Request{
		Source:    "balance",
		Reason:    "Delivery pickup",
		Amount:    300,
		Recipient: recipient1.RecipientCode,
	}

	transfer, err := service.Initiate(req)

	if err != nil {
		t.Error(err)
	}

	if transfer.TransferCode == "" {
		t.Errorf("Expected transfer code, got %+v", transfer.TransferCode)
	}

	// fetch transfer
	trf, err := service.Get(transfer.TransferCode)
	if err != nil {
		t.Error(err)
	}

	if trf.TransferCode == "" {
		t.Errorf("Expected transfer code, got %+v", trf.TransferCode)
	}
}

/* FAILS: Error message: Invalid amount passed
func TestBulkTransfer(t *testing.T) {
	// You need to disable the Transfers OTP requirement to use this endpoint
	service.DisableOTP()

	// retrieve the transfer recipient list
	recipients, err := createDemoRecipients()

	if err != nil {
		t.Error(err)
	}

	transfer := &BulkTransfer{
		Source:   "balance",
		Currency: "NGN",
		Transfers: []map[string]interface{}{
			{
				"amount":    50000,
				"recipient": recipients[0].RecipientCode,
			},
			{
				"amount":    50000,
				"recipient": recipients[1].RecipientCode,
			},
		},
	}

	_, err = service.MakeBulkTransfer(transfer)

	if err != nil {
		t.Error(err)
	}
}
*/

func TestTransferList(t *testing.T) {
	// retrieve the transfer list
	transfers, err := service.List()
	if err != nil {
		t.Errorf("Expected Transfer list, got %d, returned error %v", len(transfers.Values), err)
	}
}

func TestTransferRecipientList(t *testing.T) {
	//fmt.Println("createDemoRecipients <<<<<<<")
	//_, err := createDemoRecipients()

	//if err != nil {
	//	t.Error(err)
	//}

	//fmt.Println("ListRecipients <<<<<<<")
	// retrieve the transfer recipient list
	recipients, err := service.ListRecipients()

	if err != nil || !(len(recipients.Values) > 0) || !(recipients.Meta.Total > 0) {
		t.Errorf("Expected Recipients list, got %d, returned error %v", len(recipients.Values), err)
	}
}

func createDemoRecipients() ([]*Recipient, error) {
	recipient1 := &Recipient{
		Type:          "Nuban",
		Name:          "Customer 1",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Carpenter"},
	}

	recipient2 := &Recipient{
		Type:          "Nuban",
		Name:          "Customer 2",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Chef"},
	}

	recipient3 := &Recipient{
		Type:          "Nuban",
		Name:          "Customer 2",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

	_, err := service.CreateRecipient(recipient1)
	_, err = service.CreateRecipient(recipient2)
	_, err = service.CreateRecipient(recipient3)

	return []*Recipient{recipient1, recipient2, recipient3}, err
}
