package charge

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/response"
	"net/http"
	"net/url"
)

type Service interface {
	Create(ctx context.Context, req *ChargeRequest) (response.Response, error)
	Tokenize(ctx context.Context, req *ChargeRequest) (response.Response, error)
	SubmitPIN(ctx context.Context, pin, reference string) (response.Response, error)
	SubmitOTP(ctx context.Context, otp, reference string) (response.Response, error)
	SubmitPhone(ctx context.Context, phone, reference string) (response.Response, error)
	SubmitBirthday(ctx context.Context, birthday, reference string) (response.Response, error)
	CheckPending(ctx context.Context, reference string) (response.Response, error)
}

// DefaultChargeService handles operations related to bulk charges
// For more details see https://developers.paystack.co/v1.0/reference#charge-tokenize
type DefaultChargeService struct {
	*client.Client
}

// Create submits a charge request using card details or bank details or authorization code
// For more details see https://developers.paystack.co/v1.0/reference#charge
func (s *DefaultChargeService) Create(ctx context.Context, req *ChargeRequest) (response.Response, error) {
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/charge", req, &resp)
	return resp, err
}

// Tokenize tokenizes payment instrument before a charge
// For more details see https://developers.paystack.co/v1.0/reference#charge-tokenize
func (s *DefaultChargeService) Tokenize(ctx context.Context, req *ChargeRequest) (response.Response, error) {
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/charge/tokenize", req, &resp)
	return resp, err
}

// SubmitPIN submits PIN to continue a charge
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *DefaultChargeService) SubmitPIN(ctx context.Context, pin, reference string) (response.Response, error) {
	data := url.Values{}
	data.Add("pin", pin)
	data.Add("reference", reference)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/charge/submit_pin", data, &resp)
	return resp, err
}

// SubmitOTP submits OTP to continue a charge
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *DefaultChargeService) SubmitOTP(ctx context.Context, otp, reference string) (response.Response, error) {
	data := url.Values{}
	data.Add("pin", otp)
	data.Add("reference", reference)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/charge/submit_otp", data, &resp)
	return resp, err
}

// SubmitPhone submits Phone when requested
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *DefaultChargeService) SubmitPhone(ctx context.Context, phone, reference string) (response.Response, error) {
	data := url.Values{}
	data.Add("pin", phone)
	data.Add("reference", reference)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/charge/submit_phone", data, &resp)
	return resp, err
}

// SubmitBirthday submits Birthday when requested
// For more details see https://developers.paystack.co/v1.0/reference#submit-pin
func (s *DefaultChargeService) SubmitBirthday(ctx context.Context, birthday, reference string) (response.Response, error) {
	data := url.Values{}
	data.Add("pin", birthday)
	data.Add("reference", reference)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodPost, "/charge/submit_birthday", data, &resp)
	return resp, err
}

// CheckPending returns pending charges
// When you get "pending" as a charge status, wait 30 seconds or more,
// then make a check to see if its status has changed. Don't call too early as you may get a lot more pending than you should.
// For more details see https://developers.paystack.co/v1.0/reference#check-pending-charge
func (s *DefaultChargeService) CheckPending(ctx context.Context, reference string) (response.Response, error) {
	u := fmt.Sprintf("/charge/%s", reference)
	resp := response.Response{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, &resp)
	return resp, err
}
