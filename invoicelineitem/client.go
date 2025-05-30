//
//
// File generated from our OpenAPI spec
//
//

// Package invoicelineitem provides the /v1/invoices/{invoice}/lines APIs
package invoicelineitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/invoices/{invoice}/lines APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates an invoice's line item. Some fields, such as tax_amounts, only live on the invoice line item,
// so they can only be updated through this endpoint. Other fields, such as amount, live on both the invoice
// item and the invoice line item, so updates on this endpoint will propagate to the invoice item as well.
// Updating an invoice's line item is only possible before the invoice is finalized.
func Update(id string, params *stripe.InvoiceLineItemParams) (*stripe.InvoiceLineItem, error) {
	return getC().Update(id, params)
}

// Updates an invoice's line item. Some fields, such as tax_amounts, only live on the invoice line item,
// so they can only be updated through this endpoint. Other fields, such as amount, live on both the invoice
// item and the invoice line item, so updates on this endpoint will propagate to the invoice item as well.
// Updating an invoice's line item is only possible before the invoice is finalized.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.InvoiceLineItemParams) (*stripe.InvoiceLineItem, error) {
	path := stripe.FormatURLPath(
		"/v1/invoices/%s/lines/%s", stripe.StringValue(params.Invoice), id)
	invoicelineitem := &stripe.InvoiceLineItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoicelineitem)
	return invoicelineitem, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
