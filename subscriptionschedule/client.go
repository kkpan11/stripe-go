//
//
// File generated from our OpenAPI spec
//
//

// Package subscriptionschedule provides the /v1/subscription_schedules APIs
package subscriptionschedule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/subscription_schedules APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
func New(params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	return getC().New(params)
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	subscriptionschedule := &stripe.SubscriptionSchedule{}
	err := c.B.Call(
		http.MethodPost, "/v1/subscription_schedules", c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Retrieves the details of an existing subscription schedule. You only need to supply the unique subscription schedule identifier that was returned upon subscription schedule creation.
func Get(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing subscription schedule. You only need to supply the unique subscription schedule identifier that was returned upon subscription schedule creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s", id)
	subscriptionschedule := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Updates an existing subscription schedule.
func Update(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Update(id, params)
}

// Updates an existing subscription schedule.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s", id)
	subscriptionschedule := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
func Cancel(id string, params *stripe.SubscriptionScheduleCancelParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Cancel(id, params)
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.SubscriptionScheduleCancelParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s/cancel", id)
	subscriptionschedule := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
func Release(id string, params *stripe.SubscriptionScheduleReleaseParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Release(id, params)
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Release(id string, params *stripe.SubscriptionScheduleReleaseParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s/release", id)
	subscriptionschedule := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Retrieves the list of your subscription schedules.
func List(params *stripe.SubscriptionScheduleListParams) *Iter {
	return getC().List(params)
}

// Retrieves the list of your subscription schedules.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SubscriptionScheduleListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SubscriptionScheduleList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/subscription_schedules", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for subscription schedules.
type Iter struct {
	*stripe.Iter
}

// SubscriptionSchedule returns the subscription schedule which the iterator is currently pointing to.
func (i *Iter) SubscriptionSchedule() *stripe.SubscriptionSchedule {
	return i.Current().(*stripe.SubscriptionSchedule)
}

// SubscriptionScheduleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SubscriptionScheduleList() *stripe.SubscriptionScheduleList {
	return i.List().(*stripe.SubscriptionScheduleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
