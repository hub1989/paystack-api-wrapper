package refund

import (
	"context"
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
	"net/http"
)

type Service interface {
	RefundById(ctx context.Context, id int64) (*Response, error)
	RefundByReference(ctx context.Context, reference string) (*Response, error)
}

type DefaultRefundService struct {
	*client.Client
}

func (d DefaultRefundService) RefundById(ctx context.Context, id int64) (*Response, error) {
	endpoint := fmt.Sprintf("/refund")
	request := struct {
		Transaction int64 `json:"transaction"`
	}{
		Transaction: id,
	}

	response := &Response{}
	err := d.Client.Call(ctx, http.MethodPost, endpoint, &request, response)
	return response, err
}

func (d DefaultRefundService) RefundByReference(ctx context.Context, reference string) (*Response, error) {
	endpoint := fmt.Sprintf("/refund")
	request := struct {
		Transaction string `json:"transaction"`
	}{
		Transaction: reference,
	}

	response := &Response{}
	err := d.Client.Call(ctx, http.MethodPost, endpoint, &request, response)
	return response, err
}
