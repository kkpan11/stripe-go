//
//
// File generated from our OpenAPI spec
//
//

package stripe

// You can also delete webhook endpoints via the [webhook endpoint management](https://dashboard.stripe.com/account/webhooks) page of the Stripe dashboard.
type WebhookEndpointParams struct {
	Params `form:"*"`
	// Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.
	Connect *bool `form:"connect"`
	// An optional description of what the webhook is used for.
	Description *string `form:"description"`
	// Disable the webhook endpoint if set to true.
	Disabled *bool `form:"disabled"`
	// The list of events to enable for this endpoint. You may specify `['*']` to enable all events, except those that require explicit selection.
	EnabledEvents []*string `form:"enabled_events"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The URL of the webhook endpoint.
	URL *string `form:"url"`
	// This parameter is only available on creation.
	// We recommend setting the API version that the library is pinned to. See apiversion in stripe.go
	// Events sent to this endpoint will be generated with this Stripe Version instead of your account's default Stripe Version.
	APIVersion *string `form:"api_version"`
}

// AddExpand appends a new field to expand.
func (p *WebhookEndpointParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *WebhookEndpointParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Returns a list of your webhook endpoints.
type WebhookEndpointListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *WebhookEndpointListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// You can also delete webhook endpoints via the [webhook endpoint management](https://dashboard.stripe.com/account/webhooks) page of the Stripe dashboard.
type WebhookEndpointDeleteParams struct {
	Params `form:"*"`
}

// Retrieves the webhook endpoint with the given ID.
type WebhookEndpointRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *WebhookEndpointRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the webhook endpoint. You may edit the url, the list of enabled_events, and the status of your endpoint.
type WebhookEndpointUpdateParams struct {
	Params `form:"*"`
	// An optional description of what the webhook is used for.
	Description *string `form:"description"`
	// Disable the webhook endpoint if set to true.
	Disabled *bool `form:"disabled"`
	// The list of events to enable for this endpoint. You may specify `['*']` to enable all events, except those that require explicit selection.
	EnabledEvents []*string `form:"enabled_events"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The URL of the webhook endpoint.
	URL *string `form:"url"`
}

// AddExpand appends a new field to expand.
func (p *WebhookEndpointUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *WebhookEndpointUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A webhook endpoint must have a url and a list of enabled_events. You may optionally specify the Boolean connect parameter. If set to true, then a Connect webhook endpoint that notifies the specified url about events from all connected accounts is created; otherwise an account webhook endpoint that notifies the specified url only about events from your account is created. You can also create webhook endpoints in the [webhooks settings](https://dashboard.stripe.com/account/webhooks) section of the Dashboard.
type WebhookEndpointCreateParams struct {
	Params `form:"*"`
	// Events sent to this endpoint will be generated with this Stripe Version instead of your account's default Stripe Version.
	APIVersion *string `form:"api_version"`
	// Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.
	Connect *bool `form:"connect"`
	// An optional description of what the webhook is used for.
	Description *string `form:"description"`
	// The list of events to enable for this endpoint. You may specify `['*']` to enable all events, except those that require explicit selection.
	EnabledEvents []*string `form:"enabled_events"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The URL of the webhook endpoint.
	URL *string `form:"url"`
}

// AddExpand appends a new field to expand.
func (p *WebhookEndpointCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *WebhookEndpointCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// You can configure [webhook endpoints](https://docs.stripe.com/webhooks/) via the API to be
// notified about events that happen in your Stripe account or connected
// accounts.
//
// Most users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.
//
// Related guide: [Setting up webhooks](https://docs.stripe.com/webhooks/configure)
type WebhookEndpoint struct {
	APIResource
	// The API version events are rendered as for this webhook endpoint.
	APIVersion string `json:"api_version"`
	// The ID of the associated Connect application.
	Application string `json:"application"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	Deleted bool  `json:"deleted"`
	// An optional description of what the webhook is used for.
	Description string `json:"description"`
	// The list of events to enable for this endpoint. `['*']` indicates that all events are enabled, except those that require explicit selection.
	EnabledEvents []string `json:"enabled_events"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The endpoint's secret, used to generate [webhook signatures](https://docs.stripe.com/webhooks/signatures). Only returned at creation.
	Secret string `json:"secret"`
	// The status of the webhook. It can be `enabled` or `disabled`.
	Status string `json:"status"`
	// The URL of the webhook endpoint.
	URL string `json:"url"`
}

// WebhookEndpointList is a list of WebhookEndpoints as retrieved from a list endpoint.
type WebhookEndpointList struct {
	APIResource
	ListMeta
	Data []*WebhookEndpoint `json:"data"`
}
