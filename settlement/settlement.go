package settlement

import (
	"context"
	"github.com/hub1989/paystack-api-wrapper/client"
	"net/http"
)

type Service interface {
	List(ctx context.Context) (*List, error)
	ListN(ctx context.Context, count, offset int) (*List, error)
}

// DefaultSettlementService handles operations related to the settlement
// For more details see https://developers.paystack.co/v1.0/reference#create-settlement
type DefaultSettlementService struct {
	*client.Client
}

// List returns a list of settlements.
// For more details see https://developers.paystack.co/v1.0/reference#settlements
func (s *DefaultSettlementService) List(ctx context.Context) (*List, error) {
	return s.ListN(ctx, 10, 0)
}

// ListN returns a list of settlements
// For more details see https://developers.paystack.co/v1.0/reference#settlements
func (s *DefaultSettlementService) ListN(ctx context.Context, count, offset int) (*List, error) {
	u := client.PaginateURL("/settlement", count, offset)
	pg := &List{}
	err := s.Client.Call(ctx, http.MethodGet, u, nil, pg)
	return pg, err
}
