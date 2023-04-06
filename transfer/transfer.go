package transfer

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/url"
)

type Service interface {
	Initiate(req *Request) (*Transfer, error)
	Finalize(code, otp string) (response.Response, error)
	MakeBulkTransfer(req *BulkTransfer) (response.Response, error)
	Get(idCode string) (*Transfer, error)
	List() (*List, error)
	ListN(count, offset int) (*List, error)
	ResendOTP(transferCode, reason string) (response.Response, error)
	EnableOTP() (response.Response, error)
	FinalizeOTPDisable(otp string) (response.Response, error)
	CreateRecipient(recipient *Recipient) (*Recipient, error)
	ListRecipients() (*RecipientList, error)
	ListRecipientsN(count, offset int) (*RecipientList, error)
	DisableOTP() (response.Response, error)
}

// DefaultTransferService handles operations related to the transfer
// For more details see https://developers.paystack.co/v1.0/reference#create-transfer
type DefaultTransferService struct {
	*client.Client
}

// TransferRequest represents a request to create a transfer.

// Initiate initiates a new transfer
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
func (s *DefaultTransferService) Initiate(req *Request) (*Transfer, error) {
	transfer := &Transfer{}
	err := s.Client.Call("POST", "/transfer", req, transfer)
	return transfer, err
}

// Finalize completes a transfer request
// For more details see https://developers.paystack.co/v1.0/reference#finalize-transfer
func (s *DefaultTransferService) Finalize(code, otp string) (response.Response, error) {
	u := fmt.Sprintf("/transfer/finalize_transfer")
	req := url.Values{}
	req.Add("transfer_code", code)
	req.Add("otp", otp)
	resp := response.Response{}
	err := s.Client.Call("POST", u, req, &resp)
	return resp, err
}

// MakeBulkTransfer initiates a new bulk transfer request
// You need to disable the Transfers OTP requirement to use this endpoint
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-transfer
func (s *DefaultTransferService) MakeBulkTransfer(req *BulkTransfer) (response.Response, error) {
	u := fmt.Sprintf("/transfer")
	resp := response.Response{}
	err := s.Client.Call("POST", u, req, &resp)
	return resp, err
}

// Get returns the details of a transfer.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-transfer
func (s *DefaultTransferService) Get(idCode string) (*Transfer, error) {
	u := fmt.Sprintf("/transfer/%s", idCode)
	transfer := &Transfer{}
	err := s.Client.Call("GET", u, nil, transfer)
	return transfer, err
}

// List returns a list of transfers.
// For more details see https://developers.paystack.co/v1.0/reference#list-transfers
func (s *DefaultTransferService) List() (*List, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of transfers
// For more details see https://developers.paystack.co/v1.0/reference#list-transfers
func (s *DefaultTransferService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/transfer", count, offset)
	transfers := &List{}
	err := s.Client.Call("GET", u, nil, transfers)
	return transfers, err
}

// ResendOTP generates a new OTP and sends to customer in the event they are having trouble receiving one.
// For more details see https://developers.paystack.co/v1.0/reference#resend-otp-for-transfer
func (s *DefaultTransferService) ResendOTP(transferCode, reason string) (response.Response, error) {
	data := url.Values{}
	data.Add("transfer_code", transferCode)
	data.Add("reason", reason)
	resp := response.Response{}
	err := s.Client.Call("POST", "/transfer/resend_otp", data, &resp)
	return resp, err
}

// EnableOTP enables OTP requirement for Transfers
// In the event that a customer wants to stop being able to complete
// transfers programmatically, this endpoint helps turn OTP requirement back on.
// No arguments required.
func (s *DefaultTransferService) EnableOTP() (response.Response, error) {
	resp := response.Response{}
	err := s.Client.Call("POST", "/transfer/enable_otp", nil, &resp)
	return resp, err
}

// DisableOTP disables OTP requirement for Transfers
// In the event that you want to be able to complete transfers
// programmatically without use of OTPs, this endpoint helps disable that….
// with an OTP. No arguments required. You will get an OTP.
func (s *DefaultTransferService) DisableOTP() (response.Response, error) {
	resp := response.Response{}
	err := s.Client.Call("POST", "/transfer/disable_otp", nil, &resp)
	return resp, err
}

// FinalizeOTPDisable finalizes disabling of OTP requirement for Transfers
// For more details see https://developers.paystack.co/v1.0/reference#finalize-disabling-of-otp-requirement-for-transfers
func (s *DefaultTransferService) FinalizeOTPDisable(otp string) (response.Response, error) {
	data := url.Values{}
	data.Add("otp", otp)
	resp := response.Response{}
	err := s.Client.Call("POST", "/transfer/disable_otp_finalize", data, &resp)
	return resp, err
}

// CreateRecipient creates a new transfer recipient
// For more details see https://developers.paystack.co/v1.0/reference#create-transferrecipient
func (s *DefaultTransferService) CreateRecipient(recipient *Recipient) (*Recipient, error) {
	recipient1 := &Recipient{}
	err := s.Client.Call("POST", "/transferrecipient", recipient, recipient1)
	return recipient1, err
}

// ListRecipients returns a list of transfer recipients.
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *DefaultTransferService) ListRecipients() (*RecipientList, error) {
	return s.ListRecipientsN(10, 1)
}

// ListRecipientsN returns a list of transfer recipients
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *DefaultTransferService) ListRecipientsN(count, offset int) (*RecipientList, error) {
	u := client.PaginateURL("/transferrecipient", count, offset)
	resp := &RecipientList{}
	err := s.Client.Call("GET", u, nil, &resp)
	return resp, err
}
