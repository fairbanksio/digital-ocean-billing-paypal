# digital-ocean-billing-paypal

Automated DigitalOcean billing via PayPal


### Prerequisites

- [Setup a PayPal Developer account](https://developer.paypal.com/developer/accounts/)
    - [Create a Sandbox seller account](https://developer.paypal.com/developer/accounts/)
    - [Create a REST API app attached to the seller account](https://developer.paypal.com/developer/applications/create)
    - [Create a Sandbox buyer account with a balance to test](https://developer.paypal.com/developer/accounts/)
- [Get a DigitalOcean API key](https://cloud.digitalocean.com/account/api/tokens)
- Update `.env-sample` with your values and rename to `.env`


### Usage

```
go run main.go
```
