//
//
// File generated from our OpenAPI spec
//
//

// Package review provides the /v1/reviews APIs
package review

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/reviews APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Review object.
func Get(id string, params *stripe.ReviewParams) (*stripe.Review, error) {
	return getC().Get(id, params)
}

// Retrieves a Review object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ReviewParams) (*stripe.Review, error) {
	path := stripe.FormatURLPath("/v1/reviews/%s", id)
	review := &stripe.Review{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, review)
	return review, err
}

// Approves a Review object, closing it and removing it from the list of reviews.
func Approve(id string, params *stripe.ReviewApproveParams) (*stripe.Review, error) {
	return getC().Approve(id, params)
}

// Approves a Review object, closing it and removing it from the list of reviews.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Approve(id string, params *stripe.ReviewApproveParams) (*stripe.Review, error) {
	path := stripe.FormatURLPath("/v1/reviews/%s/approve", id)
	review := &stripe.Review{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, review)
	return review, err
}

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.ReviewListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ReviewListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ReviewList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/reviews", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for reviews.
type Iter struct {
	*stripe.Iter
}

// Review returns the review which the iterator is currently pointing to.
func (i *Iter) Review() *stripe.Review {
	return i.Current().(*stripe.Review)
}

// ReviewList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReviewList() *stripe.ReviewList {
	return i.List().(*stripe.ReviewList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
