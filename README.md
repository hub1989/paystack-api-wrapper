# Go library for the Paystack API.

[![hub1989 - paystack-api-wrapper](https://img.shields.io/static/v1?label=hub1989&message=paystack-api-wrapper&color=blue&logo=github)](https://github.com/hub1989/paystack-api-wrapper "Go to GitHub repo")
[![stars - paystack-api-wrapper](https://img.shields.io/github/stars/hub1989/paystack-api-wrapper?style=social)](https://github.com/hub1989/paystack-api-wrapper)
[![forks - paystack-api-wrapper](https://img.shields.io/github/forks/hub1989/paystack-api-wrapper?style=social)](https://github.com/hub1989/paystack-api-wrapper)

[![CI Pipeline](https://github.com/hub1989/paystack-api-wrapper/workflows/CI%20Pipeline/badge.svg)](https://github.com/hub1989/paystack-api-wrapper/actions?query=workflow:"CI+Pipeline")
[![GitHub tag](https://img.shields.io/github/tag/hub1989/paystack-api-wrapper?include_prereleases=&sort=semver&color=blue)](https://github.com/hub1989/paystack-api-wrapper/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)
[![issues - paystack-api-wrapper](https://img.shields.io/github/issues/hub1989/paystack-api-wrapper)](https://github.com/hub1989/paystack-api-wrapper/issues)

###### this library is based off the original "github.com/rpip/paystack-go"

paystack-api-wrapper is a Go client library for accessing the Paystack API.

Where possible, the services available on the client groups the API into logical chunks and correspond to the structure
of the Paystack API documentation at https://developers.paystack.co/v1.0/reference.

## logging

This library uses the `github.com/sirupsen/logrus` to enable better logging experience.

## structure
The code is structured to follow `paystack's` API structure.
Each domain has its own directory and corresponding service

- bank
- charge
- customer
- page
- plan
- refund
- response
- settlement
- subaccount
- subscription
- transaction
- transfer

You could customize the logging library to output in json format for example.
```go
package main

import log "github.com/sirupsen/logrus"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}
```
## context
All client requests take a context object. This can add value if used in an environment where for example `otel` is used.
You could pass a httpClient which supports `otel` and in that case, the context becomes valuable for every request.

## Usage

``` go
import "github.com/hub1989/paystack-api-wrapper"

apiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"

// second param is an optional http client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// This is also useful when you want to use a specific type of client.. e.g a client with otel capabilities
// where the http.DefaultClient is not available.
// The third param is a flag to enable verbose logging or not.
client := configuration.NewClient(apiKey, nil, true)

// There is a default implementation for every API domain. You can override this providing an implementation of the transfer.Service interface. 
transferService = transfer.DefaultTransferService{Client: c}

	recipient := &transfer.Recipient{
		Type:          "Nuban",
		Name:          "Customer 1",
		Description:   "Demo customer",
		AccountNumber: "0001234560",
		BankCode:      "058",
		Currency:      "NGN",
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

recipient1, err := client.Transfer.CreateRecipient(recipient)

req := &transfer.TransferRequest{
    Source:    "balance",
    Reason:    "Delivery pickup",
    Amount:    30,
    Recipient: recipient1.RecipientCode,
}

transfer, err := service.Initiate(context.TODO(), req)
if err != nil {
    // do something with error
}

// retrieve list of plans
planService = transfer.DefaultPlanService{Client: c}
plans, err := planService.List(context.TODO())

for i, plan := range plans.Values {
  fmt.Printf("%+v", plan)
}

cust := &customer.Customer{
    FirstName: "User123",
    LastName:  "AdminUser",
    Email:     "user123@gmail.com",
    Phone:     "+23400000000000000",
}
// create the customer
customerService = transfer.DefaultCustomerService{Client: c}
customer, err := customerService.Create(cust)
if err != nil {
    // do something with error
}

// Get customer by ID
customer, err := customerService.Get(customer.ID)
```

See the test files for more examples.

## Docker

Test this library in a docker container:

```bash
# PAYSTACK_KEY is an environment variable that should be added to your rc file. i.e .bashrc
$ make docker && docker run -e PAYSTACK_KEY -i -t paystack:latest
```

## TODO

- [ ] Test on App Engine

## CONTRIBUTING

Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is
relatively straightforward. Please make sure the build succeeds and the test suite passes.

### test data

https://paystack.com/docs/payments/test-payments

## License

Released under [MIT](/LICENSE) by [@hub1989](https://github.com/hub1989).