//
//
// File generated from our OpenAPI spec
//
//

// Package personalizationdesign provides the /v1/issuing/personalization_designs APIs
package personalizationdesign

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/issuing/personalization_designs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a personalization design object.
func New(params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().New(params)
}

// Creates a personalization design object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/personalization_designs", c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Retrieves a personalization design object.
func Get(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Get(id, params)
}

// Retrieves a personalization design object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Updates a card personalization object.
func Update(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Update(id, params)
}

// Updates a card personalization object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Returns a list of personalization design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingPersonalizationDesignListParams) *Iter {
	return getC().List(params)
}

// Returns a list of personalization design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingPersonalizationDesignListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingPersonalizationDesignList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/personalization_designs", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing personalization designs.
type Iter struct {
	*stripe.Iter
}

// IssuingPersonalizationDesign returns the issuing personalization design which the iterator is currently pointing to.
func (i *Iter) IssuingPersonalizationDesign() *stripe.IssuingPersonalizationDesign {
	return i.Current().(*stripe.IssuingPersonalizationDesign)
}

// IssuingPersonalizationDesignList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingPersonalizationDesignList() *stripe.IssuingPersonalizationDesignList {
	return i.List().(*stripe.IssuingPersonalizationDesignList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
