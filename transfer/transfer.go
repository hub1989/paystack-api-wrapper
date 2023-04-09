package transfer

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
	"net/url"
)

type Service interface {
	Initiate(ctx context.Context, req *Request) (*Transfer, error)
	Finalize(ctx context.Context, code, otp string) (response.Response, error)
	MakeBulkTransfer(ctx context.Context, req *BulkTransfer) (response.Response, error)
	Get(ctx context.Context, idCode string) (*Transfer, error)
	List(ctx context.Context) (*List, error)
	ListN(ctx context.Context, count, offset int) (*List, error)
	ResendOTP(ctx context.Context, transferCode, reason string) (response.Response, error)
	EnableOTP(ctx context.Context) (response.Response, error)
	FinalizeOTPDisable(ctx context.Context, otp string) (response.Response, error)
	CreateRecipient(ctx context.Context, recipient *Recipient) (*Recipient, error)
	ListRecipients(ctx context.Context) (*RecipientList, error)
	ListRecipientsN(ctx context.Context, count, offset int) (*RecipientList, error)
	DisableOTP(ctx context.Context) (response.Response, error)
}

// DefaultTransferService handles operations related to the transfer
// For more details see https://developers.paystack.co/v1.0/reference#create-transfer
type DefaultTransferService struct {
	*client.Client
}

// TransferRequest represents a request to create a transfer.

// Initiate initiates a new transfer
// For more details see https://developers.paystack.co/v1.0/reference#initiate-transfer
func (s *DefaultTransferService) Initiate(ctx context.Context, req *Request) (*Transfer, error) {
	transfer := &Transfer{}
	err := s.Client.Call(ctx, http.MethodPost, "/transfer", req, transfer)
	return transfer, err
}

// Finalize completes a transfer request
// For more details see https://developers.paystack.co/v1.0/reference#finalize-transfer
func (s *DefaultTransferService) Finalize(ctx context.Context, code, otp string) (response.Response, error) {
	u := fmt.Sprintf("/transfer/finalize_transfer")
	req := url.Values{}
	req.Add("transfer_code", code)
	req.Add("otp", otp)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, u, req, &resp)
	return resp, err
}

// MakeBulkTransfer initiates a new bulk transfer request
// You need to disable the Transfers OTP requirement to use this endpoint
// For more details see https://developers.paystack.co/v1.0/reference#initiate-bulk-transfer
func (s *DefaultTransferService) MakeBulkTransfer(ctx context.Context, req *BulkTransfer) (response.Response, error) {
	u := fmt.Sprintf("/transfer")
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, u, req, &resp)
	return resp, err
}

// Get returns the details of a transfer.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-transfer
func (s *DefaultTransferService) Get(ctx context.Context, idCode string) (*Transfer, error) {
	u := fmt.Sprintf("/transfer/%s", idCode)
	transfer := &Transfer{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, transfer)
	return transfer, err
}

// List returns a list of transfers.
// For more details see https://developers.paystack.co/v1.0/reference#list-transfers
func (s *DefaultTransferService) List(ctx context.Context) (*List, error) {
	return s.ListN(ctx, 10, 0)
}

// ListN returns a list of transfers
// For more details see https://developers.paystack.co/v1.0/reference#list-transfers
func (s *DefaultTransferService) ListN(ctx context.Context, count, offset int) (*List, error) {
	u := client.PaginateURL("/transfer", count, offset)
	transfers := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, transfers)
	return transfers, err
}

// ResendOTP generates a new OTP and sends to customer in the event they are having trouble receiving one.
// For more details see https://developers.paystack.co/v1.0/reference#resend-otp-for-transfer
func (s *DefaultTransferService) ResendOTP(ctx context.Context, transferCode, reason string) (response.Response, error) {
	data := url.Values{}
	data.Add("transfer_code", transferCode)
	data.Add("reason", reason)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/transfer/resend_otp", data, &resp)
	return resp, err
}

// EnableOTP enables OTP requirement for Transfers
// In the event that a customer wants to stop being able to complete
// transfers programmatically, this endpoint helps turn OTP requirement back on.
// No arguments required.
func (s *DefaultTransferService) EnableOTP(ctx context.Context) (response.Response, error) {
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/transfer/enable_otp", nil, &resp)
	return resp, err
}

// DisableOTP disables OTP requirement for Transfers
// In the event that you want to be able to complete transfers
// programmatically without use of OTPs, this endpoint helps disable that….
// with an OTP. No arguments required. You will get an OTP.
func (s *DefaultTransferService) DisableOTP(ctx context.Context) (response.Response, error) {
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/transfer/disable_otp", nil, &resp)
	return resp, err
}

// FinalizeOTPDisable finalizes disabling of OTP requirement for Transfers
// For more details see https://developers.paystack.co/v1.0/reference#finalize-disabling-of-otp-requirement-for-transfers
func (s *DefaultTransferService) FinalizeOTPDisable(ctx context.Context, otp string) (response.Response, error) {
	data := url.Values{}
	data.Add("otp", otp)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/transfer/disable_otp_finalize", data, &resp)
	return resp, err
}

// CreateRecipient creates a new transfer recipient
// For more details see https://developers.paystack.co/v1.0/reference#create-transferrecipient
func (s *DefaultTransferService) CreateRecipient(ctx context.Context, recipient *Recipient) (*Recipient, error) {
	recipient1 := &Recipient{}
	err := s.Client.Call(ctx, http.MethodPost, "/transferrecipient", recipient, recipient1)
	return recipient1, err
}

// ListRecipients returns a list of transfer recipients.
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *DefaultTransferService) ListRecipients(ctx context.Context) (*RecipientList, error) {
	return s.ListRecipientsN(ctx, 10, 1)
}

// ListRecipientsN returns a list of transfer recipients
// For more details see https://developers.paystack.co/v1.0/reference#list-transferrecipients
func (s *DefaultTransferService) ListRecipientsN(ctx context.Context, count, offset int) (*RecipientList, error) {
	u := client.PaginateURL("/transferrecipient", count, offset)
	resp := &RecipientList{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)
	return resp, err
}
