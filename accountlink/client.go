//
//
// File generated from our OpenAPI spec
//
//

// Package accountlink provides the /v1/account_links APIs
package accountlink

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/account_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.
func New(params *stripe.AccountLinkParams) (*stripe.AccountLink, error) {
	return getC().New(params)
}

// Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.AccountLinkParams) (*stripe.AccountLink, error) {
	accountlink := &stripe.AccountLink{}
	err := c.B.Call(
		http.MethodPost, "/v1/account_links", c.Key, params, accountlink)
	return accountlink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
