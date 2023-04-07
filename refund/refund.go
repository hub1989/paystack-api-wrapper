package refund

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
)

type Service interface {
	RefundById(id int64) (*Response, error)
	RefundByReference(reference *string) (*Response, error)
}

type DefaultRefundService struct {
	*client.Client
}

func (d DefaultRefundService) RefundById(id int64) (*Response, error) {
	endpoint := fmt.Sprintf("/refund")
	request := struct {
		Transaction int64 `json:"transaction"`
	}{
		Transaction: id,
	}

	response := &Response{}
	err := d.Client.Call("POST", endpoint, &request, response)
	return response, err
}

func (d DefaultRefundService) RefundByReference(reference string) (*Response, error) {
	endpoint := fmt.Sprintf("/refund")
	request := struct {
		Transaction string `json:"transaction"`
	}{
		Transaction: reference,
	}

	response := &Response{}
	err := d.Client.Call("POST", endpoint, &request, response)
	return response, err
}
