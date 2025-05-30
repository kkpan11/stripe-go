//
//
// File generated from our OpenAPI spec
//
//

// Package supplier provides the /v1/climate/suppliers APIs
package supplier

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/climate/suppliers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Climate supplier object.
func Get(id string, params *stripe.ClimateSupplierParams) (*stripe.ClimateSupplier, error) {
	return getC().Get(id, params)
}

// Retrieves a Climate supplier object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ClimateSupplierParams) (*stripe.ClimateSupplier, error) {
	path := stripe.FormatURLPath("/v1/climate/suppliers/%s", id)
	supplier := &stripe.ClimateSupplier{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, supplier)
	return supplier, err
}

// Lists all available Climate supplier objects.
func List(params *stripe.ClimateSupplierListParams) *Iter {
	return getC().List(params)
}

// Lists all available Climate supplier objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ClimateSupplierListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ClimateSupplierList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/climate/suppliers", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for climate suppliers.
type Iter struct {
	*stripe.Iter
}

// ClimateSupplier returns the climate supplier which the iterator is currently pointing to.
func (i *Iter) ClimateSupplier() *stripe.ClimateSupplier {
	return i.Current().(*stripe.ClimateSupplier)
}

// ClimateSupplierList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ClimateSupplierList() *stripe.ClimateSupplierList {
	return i.List().(*stripe.ClimateSupplierList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
