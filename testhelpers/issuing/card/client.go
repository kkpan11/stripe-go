//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the /v1/issuing/cards APIs
package card

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/issuing/cards APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates the shipping status of the specified Issuing Card object to delivered.
func DeliverCard(id string, params *stripe.TestHelpersIssuingCardDeliverCardParams) (*stripe.IssuingCard, error) {
	return getC().DeliverCard(id, params)
}

// Updates the shipping status of the specified Issuing Card object to delivered.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeliverCard(id string, params *stripe.TestHelpersIssuingCardDeliverCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/deliver", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to failure.
func FailCard(id string, params *stripe.TestHelpersIssuingCardFailCardParams) (*stripe.IssuingCard, error) {
	return getC().FailCard(id, params)
}

// Updates the shipping status of the specified Issuing Card object to failure.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FailCard(id string, params *stripe.TestHelpersIssuingCardFailCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/fail", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to returned.
func ReturnCard(id string, params *stripe.TestHelpersIssuingCardReturnCardParams) (*stripe.IssuingCard, error) {
	return getC().ReturnCard(id, params)
}

// Updates the shipping status of the specified Issuing Card object to returned.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnCard(id string, params *stripe.TestHelpersIssuingCardReturnCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/return", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to shipped.
func ShipCard(id string, params *stripe.TestHelpersIssuingCardShipCardParams) (*stripe.IssuingCard, error) {
	return getC().ShipCard(id, params)
}

// Updates the shipping status of the specified Issuing Card object to shipped.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ShipCard(id string, params *stripe.TestHelpersIssuingCardShipCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/ship", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Updates the shipping status of the specified Issuing Card object to submitted. This method requires Stripe Version ‘2024-09-30.acacia' or later.
func SubmitCard(id string, params *stripe.TestHelpersIssuingCardSubmitCardParams) (*stripe.IssuingCard, error) {
	return getC().SubmitCard(id, params)
}

// Updates the shipping status of the specified Issuing Card object to submitted. This method requires Stripe Version ‘2024-09-30.acacia' or later.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SubmitCard(id string, params *stripe.TestHelpersIssuingCardSubmitCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/cards/%s/shipping/submit", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
