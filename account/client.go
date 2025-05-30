//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the /v1/accounts APIs
package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/accounts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// With [Connect](https://docs.stripe.com/docs/connect), you can create Stripe accounts for your users.
// To do this, you'll first need to [register your platform](https://dashboard.stripe.com/account/applications/settings).
//
// If you've already collected information for your connected accounts, you [can prefill that information](https://docs.stripe.com/docs/connect/best-practices#onboarding) when
// creating the account. Connect Onboarding won't ask for the prefilled information during account onboarding.
// You can prefill any information on the account.
func New(params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().New(params)
}

// With [Connect](https://docs.stripe.com/docs/connect), you can create Stripe accounts for your users.
// To do this, you'll first need to [register your platform](https://dashboard.stripe.com/account/applications/settings).
//
// If you've already collected information for your connected accounts, you [can prefill that information](https://docs.stripe.com/docs/connect/best-practices#onboarding) when
// creating the account. Connect Onboarding won't ask for the prefilled information during account onboarding.
// You can prefill any information on the account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.AccountParams) (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call(http.MethodPost, "/v1/accounts", c.Key, params, account)
	return account, err
}

// Get retrieves the authenticating account.
func Get() (*stripe.Account, error) {
	return getC().Get()
}

// Get retrieves the authenticating account.
func (c Client) Get() (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call(http.MethodGet, "/v1/account", c.Key, nil, account)
	return account, err
}

// Retrieves the details of an account.
func GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().GetByID(id, params)
}

// Retrieves the details of an account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Updates a [connected account](https://docs.stripe.com/connect/accounts) by setting the values of the parameters passed. Any parameters not provided are
// left unchanged.
//
// For accounts where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection)
// is application, which includes Custom accounts, you can update any information on the account.
//
// For accounts where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection)
// is stripe, which includes Standard and Express accounts, you can update all information until you create
// an [Account Link or <a href="/api/account_sessions">Account Session](https://docs.stripe.com/api/account_links) to start Connect onboarding,
// after which some properties can no longer be updated.
//
// To update your own account, use the [Dashboard](https://dashboard.stripe.com/settings/account). Refer to our
// [Connect](https://docs.stripe.com/docs/connect/updating-accounts) documentation to learn more about updating accounts.
func Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Update(id, params)
}

// Updates a [connected account](https://docs.stripe.com/connect/accounts) by setting the values of the parameters passed. Any parameters not provided are
// left unchanged.
//
// For accounts where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection)
// is application, which includes Custom accounts, you can update any information on the account.
//
// For accounts where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection)
// is stripe, which includes Standard and Express accounts, you can update all information until you create
// an [Account Link or <a href="/api/account_sessions">Account Session](https://docs.stripe.com/api/account_links) to start Connect onboarding,
// after which some properties can no longer be updated.
//
// To update your own account, use the [Dashboard](https://dashboard.stripe.com/settings/account). Refer to our
// [Connect](https://docs.stripe.com/docs/connect/updating-accounts) documentation to learn more about updating accounts.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// With [Connect](https://docs.stripe.com/connect), you can delete accounts you manage.
//
// Test-mode accounts can be deleted at any time.
//
// Live-mode accounts where Stripe is responsible for negative account balances cannot be deleted, which includes Standard accounts. Live-mode accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be deleted when all [balances](https://docs.stripe.com/api/balance/balance_object) are zero.
//
// If you want to delete your own account, use the [account information tab in your account settings](https://dashboard.stripe.com/settings/account) instead.
func Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Del(id, params)
}

// With [Connect](https://docs.stripe.com/connect), you can delete accounts you manage.
//
// Test-mode accounts can be deleted at any time.
//
// Live-mode accounts where Stripe is responsible for negative account balances cannot be deleted, which includes Standard accounts. Live-mode accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be deleted when all [balances](https://docs.stripe.com/api/balance/balance_object) are zero.
//
// If you want to delete your own account, use the [account information tab in your account settings](https://dashboard.stripe.com/settings/account) instead.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, account)
	return account, err
}

// With [Connect](https://docs.stripe.com/connect), you can reject accounts that you have flagged as suspicious.
//
// Only accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be rejected. Test-mode accounts can be rejected at any time. Live-mode accounts can only be rejected after all balances are zero.
func Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	return getC().Reject(id, params)
}

// With [Connect](https://docs.stripe.com/connect), you can reject accounts that you have flagged as suspicious.
//
// Only accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be rejected. Test-mode accounts can be rejected at any time. Live-mode accounts can only be rejected after all balances are zero.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s/reject", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Returns a list of accounts connected to your platform via [Connect](https://docs.stripe.com/docs/connect). If you're not a platform, the list is empty.
func List(params *stripe.AccountListParams) *Iter {
	return getC().List(params)
}

// Returns a list of accounts connected to your platform via [Connect](https://docs.stripe.com/docs/connect). If you're not a platform, the list is empty.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.AccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.AccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/accounts", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for accounts.
type Iter struct {
	*stripe.Iter
}

// Account returns the account which the iterator is currently pointing to.
func (i *Iter) Account() *stripe.Account {
	return i.Current().(*stripe.Account)
}

// AccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) AccountList() *stripe.AccountList {
	return i.List().(*stripe.AccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
