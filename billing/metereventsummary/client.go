//
//
// File generated from our OpenAPI spec
//
//

// Package metereventsummary provides the /v1/billing/meters/{id}/event_summaries APIs
package metereventsummary

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/billing/meters/{id}/event_summaries APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a list of billing meter event summaries.
func List(params *stripe.BillingMeterEventSummaryListParams) *Iter {
	return getC().List(params)
}

// Retrieve a list of billing meter event summaries.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingMeterEventSummaryListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/billing/meters/%s/event_summaries", stripe.StringValue(listParams.ID))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.BillingMeterEventSummaryList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for billing meter event summaries.
type Iter struct {
	*stripe.Iter
}

// BillingMeterEventSummary returns the billing meter event summary which the iterator is currently pointing to.
func (i *Iter) BillingMeterEventSummary() *stripe.BillingMeterEventSummary {
	return i.Current().(*stripe.BillingMeterEventSummary)
}

// BillingMeterEventSummaryList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingMeterEventSummaryList() *stripe.BillingMeterEventSummaryList {
	return i.List().(*stripe.BillingMeterEventSummaryList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
