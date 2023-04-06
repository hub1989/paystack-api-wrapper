package page

import (
	"github.com/hub1989/paystack-api-wrapper/client"
	"github.com/hub1989/paystack-api-wrapper/configuration"
	"testing"
)

var c *client.Client
var service *DefaultPageService

func init() {
	apiKey := client.MustGetTestKey()
	c = configuration.NewClient(apiKey, nil)
	service = &DefaultPageService{Client: c}
}

func TestPageCRUD(t *testing.T) {
	page1 := &Page{
		Name:        "Demo page",
		Description: "Paystack Go client test page",
	}

	// create the page
	page, err := service.Create(page1)
	if err != nil {
		t.Errorf("CREATE Page returned error: %v", err)
	}

	// retrieve the page
	page, err = service.Get(page.ID)
	if err != nil {
		t.Errorf("GET Page returned error: %v", err)
	}

	if page.Name != page1.Name {
		t.Errorf("Expected Page Name %v, got %v", page.Name, page1.Name)
	}

	// retrieve the page list
	pages, err := service.List()
	if err != nil || !(len(pages.Values) > 0) || !(pages.Meta.Total > 0) {
		t.Errorf("Expected Page list, got %d, returned error %v", len(pages.Values), err)
	}
}
