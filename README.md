# Go library for the Paystack API.
[![hub1989 - paystack-api-wrapper](https://img.shields.io/static/v1?label=hub1989&message=paystack-api-wrapper&color=blue&logo=github)](https://github.com/hub1989/paystack-api-wrapper "Go to GitHub repo")
[![stars - paystack-api-wrapper](https://img.shields.io/github/stars/hub1989/paystack-api-wrapper?style=social)](https://github.com/hub1989/paystack-api-wrapper)
[![forks - paystack-api-wrapper](https://img.shields.io/github/forks/hub1989/paystack-api-wrapper?style=social)](https://github.com/hub1989/paystack-api-wrapper)

[![CI Pipeline](https://github.com/hub1989/paystack-api-wrapper/workflows/CI%20Pipeline/badge.svg)](https://github.com/hub1989/paystack-api-wrapper/actions?query=workflow:"CI+Pipeline")
[![GitHub tag](https://img.shields.io/github/tag/hub1989/paystack-api-wrapper?include_prereleases=&sort=semver&color=blue)](https://github.com/hub1989/paystack-api-wrapper/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)
[![issues - paystack-api-wrapper](https://img.shields.io/github/issues/hub1989/paystack-api-wrapper)](https://github.com/hub1989/paystack-api-wrapper/issues)

###### this library is based off the original 00"github.com/rpip/paystack-go"
paystack-api-wrapper is a Go client library for accessing the Paystack API.

Where possible, the services available on the client groups the API into logical chunks and correspond to the structure of the Paystack API documentation at https://developers.paystack.co/v1.0/reference.

## Usage

``` go
import "github.com/hub1989/paystack-api-wrapper"

apiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"

// second param is an optional http client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
client := paystack.NewClient(apiKey)

recipient := &TransferRecipient{
    Type:          "Nuban",
    Name:          "Customer 1",
    Description:   "Demo customer",
    AccountNumber: "0100000010",
    BankCode:      "044",
    Currency:      "NGN",
    Metadata:      map[string]interface{}{"job": "Plumber"},
}

recipient1, err := client.Transfer.CreateRecipient(recipient)

req := &TransferRequest{
    Source:    "balance",
    Reason:    "Delivery pickup",
    Amount:    30,
    Recipient: recipient1.RecipientCode,
}

transfer, err := client.Transfer.Initiate(req)
if err != nil {
    // do something with error
}

// retrieve list of plans
plans, err := client.Plan.List()

for i, plan := range plans.Values {
  fmt.Printf("%+v", plan)
}

cust := &Customer{
    FirstName: "User123",
    LastName:  "AdminUser",
    Email:     "user123@gmail.com",
    Phone:     "+23400000000000000",
}
// create the customer
customer, err := client.Customer.Create(cust)
if err != nil {
    // do something with error
}

// Get customer by ID
customer, err := client.Customers.Get(customer.ID)
```

See the test files for more examples.

## Docker

Test this library in a docker container:

```bash
# PAYSTACK_KEY is an environment variable that should be added to your rc file. i.e .bashrc
$ make docker && docker run -e PAYSTACK_KEY -i -t paystack:latest
```

## TODO
- [ ] Maybe support request context?
- [ ] Test on App Engine

## CONTRIBUTING
Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. Please make sure the build succeeds and the test suite passes.

### test data
https://paystack.com/docs/payments/test-payments

## License

Released under [MIT](/LICENSE) by [@hub1989](https://github.com/hub1989).