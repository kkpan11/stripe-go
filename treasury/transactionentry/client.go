//
//
// File generated from our OpenAPI spec
//
//

// Package transactionentry provides the /v1/treasury/transaction_entries APIs
package transactionentry

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/treasury/transaction_entries APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a TransactionEntry object.
func Get(id string, params *stripe.TreasuryTransactionEntryParams) (*stripe.TreasuryTransactionEntry, error) {
	return getC().Get(id, params)
}

// Retrieves a TransactionEntry object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TreasuryTransactionEntryParams) (*stripe.TreasuryTransactionEntry, error) {
	path := stripe.FormatURLPath("/v1/treasury/transaction_entries/%s", id)
	transactionentry := &stripe.TreasuryTransactionEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transactionentry)
	return transactionentry, err
}

// Retrieves a list of TransactionEntry objects.
func List(params *stripe.TreasuryTransactionEntryListParams) *Iter {
	return getC().List(params)
}

// Retrieves a list of TransactionEntry objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TreasuryTransactionEntryListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryTransactionEntryList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/transaction_entries", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury transaction entries.
type Iter struct {
	*stripe.Iter
}

// TreasuryTransactionEntry returns the treasury transaction entry which the iterator is currently pointing to.
func (i *Iter) TreasuryTransactionEntry() *stripe.TreasuryTransactionEntry {
	return i.Current().(*stripe.TreasuryTransactionEntry)
}

// TreasuryTransactionEntryList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryTransactionEntryList() *stripe.TreasuryTransactionEntryList {
	return i.List().(*stripe.TreasuryTransactionEntryList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
