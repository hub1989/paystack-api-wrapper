package page

import (
	"fmt"
	"github.com/hub1989/paystack-api-wrapper/client"
)

type Service interface {
	Create(page *Page) (*Page, error)
	Update(page *Page) (*Page, error)
	Get(id int) (*Page, error)
	List() (*List, error)
	ListN(count, offset int) (*List, error)
}

// DefaultPageService handles operations related to the page
// For more details see https://developers.paystack.co/v1.0/reference#create-page
type DefaultPageService struct {
	*client.Client
}

// Create creates a new page
// For more details see https://developers.paystack.co/v1.0/reference#create-page
func (s *DefaultPageService) Create(page *Page) (*Page, error) {
	u := fmt.Sprintf("/page")
	pg := &Page{}
	err := s.Client.Call("POST", u, page, pg)

	return pg, err
}

// Update updates a page's properties.
// For more details see https://developers.paystack.co/v1.0/reference#update-page
func (s *DefaultPageService) Update(page *Page) (*Page, error) {
	u := fmt.Sprintf("page/%d", page.ID)
	pg := &Page{}
	err := s.Client.Call("PUT", u, page, pg)

	return pg, err
}

// Get returns the details of a page.
// For more details see https://developers.paystack.co/v1.0/reference#fetch-page
func (s *DefaultPageService) Get(id int) (*Page, error) {
	u := fmt.Sprintf("/page/%d", id)
	pg := &Page{}
	err := s.Client.Call("GET", u, nil, pg)

	return pg, err
}

// List returns a list of pages.
// For more details see https://developers.paystack.co/v1.0/reference#list-pages
func (s *DefaultPageService) List() (*List, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of pages
// For more details see https://developers.paystack.co/v1.0/reference#list-pages
func (s *DefaultPageService) ListN(count, offset int) (*List, error) {
	u := client.PaginateURL("/page", count, offset)
	pg := &List{}
	err := s.Client.Call("GET", u, nil, pg)
	return pg, err
}
