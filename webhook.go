// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/y2-intel/y2-go/internal/apijson"
	"github.com/y2-intel/y2-go/internal/requestconfig"
	"github.com/y2-intel/y2-go/option"
	"github.com/y2-intel/y2-go/packages/param"
	"github.com/y2-intel/y2-go/packages/respjson"
)

// Webhook configuration management (Pro feature)
//
// WebhookService contains methods and other services that help with interacting
// with the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook configuration. Requires an active Pro subscription. The
// webhook URL must be HTTPS and pass SSRF security validation.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing webhook configuration. All fields are optional. Only
// provided fields will be updated.
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all webhook configurations for the authenticated user. Secrets are
// masked in the response.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a webhook configuration. Fails with 409 if the webhook is currently in
// use by any subscriptions.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Sends a test payload to the webhook URL and returns the result. Returns 422 if
// the webhook endpoint responds with an error.
func (r *WebhookService) Test(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/test", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type WebhookNewResponse struct {
	Data WebhookNewResponseData `json:"data" api:"required"`
	Meta WebhookNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseData struct {
	// The ID of the newly created webhook configuration
	WebhookID string `json:"webhookId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WebhookID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponse struct {
	Data WebhookUpdateResponseData `json:"data" api:"required"`
	Meta WebhookUpdateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponseData struct {
	Success   bool   `json:"success" api:"required"`
	WebhookID string `json:"webhookId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		WebhookID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	Data []WebhookListResponseData `json:"data" api:"required"`
	Meta WebhookListResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseData struct {
	// Webhook configuration ID
	ID string `json:"id" api:"required"`
	// Creation timestamp (ms)
	CreatedAt int64 `json:"createdAt" api:"required"`
	// Consecutive failure count (auto-disabled at 5)
	FailureCount int64 `json:"failureCount" api:"required"`
	// Whether a secret is configured (actual secret is never exposed)
	HasSecret bool `json:"hasSecret" api:"required"`
	// Whether the webhook is active
	IsActive bool `json:"isActive" api:"required"`
	// Webhook display name
	Name string `json:"name" api:"required"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required" format:"uri"`
	// Last delivery timestamp (ms)
	LastUsedAt int64 `json:"lastUsedAt"`
	// Last update timestamp (ms)
	UpdatedAt int64 `json:"updatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		FailureCount respjson.Field
		HasSecret    respjson.Field
		IsActive     respjson.Field
		Name         respjson.Field
		URL          respjson.Field
		LastUsedAt   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseMeta struct {
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponse struct {
	Data WebhookDeleteResponseData `json:"data" api:"required"`
	Meta WebhookDeleteResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponseData struct {
	Deleted   bool   `json:"deleted" api:"required"`
	WebhookID string `json:"webhookId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		WebhookID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeleteResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponse struct {
	Data WebhookTestResponseData `json:"data" api:"required"`
	Meta WebhookTestResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponseData struct {
	Success bool `json:"success" api:"required"`
	// Error message if the test failed
	Error string `json:"error"`
	// Response time in milliseconds
	ResponseTime int64 `json:"responseTime"`
	// HTTP status code from the webhook endpoint
	StatusCode int64 `json:"statusCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success      respjson.Field
		Error        respjson.Field
		ResponseTime respjson.Field
		StatusCode   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// Webhook display name
	Name string `json:"name" api:"required"`
	// Webhook endpoint URL (must be HTTPS)
	URL string `json:"url" api:"required" format:"uri"`
	// Shared secret for signature verification
	Secret param.Opt[string] `json:"secret,omitzero"`
	// Custom headers to include in webhook deliveries
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	IsActive param.Opt[bool]   `json:"isActive,omitzero"`
	Name     param.Opt[string] `json:"name,omitzero"`
	Secret   param.Opt[string] `json:"secret,omitzero"`
	URL      param.Opt[string] `json:"url,omitzero" format:"uri"`
	Headers  map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
