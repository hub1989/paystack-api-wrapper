package settlement

import (
	"github.com/hub1989/paystack-api-wrapper/client"
)

type Service interface {
	List() (*List, error)
	ListN(count, offset int) (*List, error)
}

// DefaultSettlementService handles operations related to the settlement
// For more details see https://developers.paystack.co/v1.0/reference#create-settlement
type DefaultSettlementService struct {
	*client.Client
}

// List returns a list of settlements.
// For more details see https://developers.paystack.co/v1.0/reference#settlements
func (s *DefaultSettlementService) List() (*List, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of settlements
// For more details see https://developers.paystack.co/v1.0/reference#settlements
func (s *DefaultSettlementService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/settlement", count, offset)
	pg := &List{}
	err := s.Client.Call("GET", u, nil, pg)
	return pg, err
}
