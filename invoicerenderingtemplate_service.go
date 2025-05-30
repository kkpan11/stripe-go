//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1InvoiceRenderingTemplateService is used to invoke /v1/invoice_rendering_templates APIs.
type v1InvoiceRenderingTemplateService struct {
	B   Backend
	Key string
}

// Retrieves an invoice rendering template with the given ID. It by default returns the latest version of the template. Optionally, specify a version to see previous versions.
func (c v1InvoiceRenderingTemplateService) Retrieve(ctx context.Context, id string, params *InvoiceRenderingTemplateRetrieveParams) (*InvoiceRenderingTemplate, error) {
	if params == nil {
		params = &InvoiceRenderingTemplateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoice_rendering_templates/%s", id)
	invoicerenderingtemplate := &InvoiceRenderingTemplate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoicerenderingtemplate)
	return invoicerenderingtemplate, err
}

// Updates the status of an invoice rendering template to ‘archived' so no new Stripe objects (customers, invoices, etc.) can reference it. The template can also no longer be updated. However, if the template is already set on a Stripe object, it will continue to be applied on invoices generated by it.
func (c v1InvoiceRenderingTemplateService) Archive(ctx context.Context, id string, params *InvoiceRenderingTemplateArchiveParams) (*InvoiceRenderingTemplate, error) {
	if params == nil {
		params = &InvoiceRenderingTemplateArchiveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoice_rendering_templates/%s/archive", id)
	invoicerenderingtemplate := &InvoiceRenderingTemplate{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, invoicerenderingtemplate)
	return invoicerenderingtemplate, err
}

// Unarchive an invoice rendering template so it can be used on new Stripe objects again.
func (c v1InvoiceRenderingTemplateService) Unarchive(ctx context.Context, id string, params *InvoiceRenderingTemplateUnarchiveParams) (*InvoiceRenderingTemplate, error) {
	if params == nil {
		params = &InvoiceRenderingTemplateUnarchiveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoice_rendering_templates/%s/unarchive", id)
	invoicerenderingtemplate := &InvoiceRenderingTemplate{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, invoicerenderingtemplate)
	return invoicerenderingtemplate, err
}

// List all templates, ordered by creation date, with the most recently created template appearing first.
func (c v1InvoiceRenderingTemplateService) List(ctx context.Context, listParams *InvoiceRenderingTemplateListParams) Seq2[*InvoiceRenderingTemplate, error] {
	if listParams == nil {
		listParams = &InvoiceRenderingTemplateListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*InvoiceRenderingTemplate, ListContainer, error) {
		list := &InvoiceRenderingTemplateList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/invoice_rendering_templates", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
