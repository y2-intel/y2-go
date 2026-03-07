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

// Subscription delivery management
//
// SubscriptionService contains methods and other services that help with
// interacting with the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
	r.Options = opts
	return
}

// Changes the delivery method for a subscription. When setting to `webhook`, a
// valid `webhookConfigId` must be provided. The webhook must be active.
func (r *SubscriptionService) UpdateDelivery(ctx context.Context, subscriptionID string, body SubscriptionUpdateDeliveryParams, opts ...option.RequestOption) (res *SubscriptionUpdateDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/delivery", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SubscriptionUpdateDeliveryResponse struct {
	Data SubscriptionUpdateDeliveryResponseData `json:"data" api:"required"`
	Meta SubscriptionUpdateDeliveryResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionUpdateDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionUpdateDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdateDeliveryResponseData struct {
	// Subscription delivery method
	//
	// Any of "email", "sms", "webhook", "both_email_sms".
	DeliveryMethod string `json:"deliveryMethod" api:"required"`
	SubscriptionID string `json:"subscriptionId" api:"required"`
	Success        bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryMethod respjson.Field
		SubscriptionID respjson.Field
		Success        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionUpdateDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionUpdateDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdateDeliveryResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionUpdateDeliveryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionUpdateDeliveryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdateDeliveryParams struct {
	// Subscription delivery method
	//
	// Any of "email", "sms", "webhook", "both_email_sms".
	DeliveryMethod SubscriptionUpdateDeliveryParamsDeliveryMethod `json:"deliveryMethod,omitzero" api:"required"`
	// Required when deliveryMethod is "webhook"
	WebhookConfigID param.Opt[string] `json:"webhookConfigId,omitzero"`
	paramObj
}

func (r SubscriptionUpdateDeliveryParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdateDeliveryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdateDeliveryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription delivery method
type SubscriptionUpdateDeliveryParamsDeliveryMethod string

const (
	SubscriptionUpdateDeliveryParamsDeliveryMethodEmail        SubscriptionUpdateDeliveryParamsDeliveryMethod = "email"
	SubscriptionUpdateDeliveryParamsDeliveryMethodSMS          SubscriptionUpdateDeliveryParamsDeliveryMethod = "sms"
	SubscriptionUpdateDeliveryParamsDeliveryMethodWebhook      SubscriptionUpdateDeliveryParamsDeliveryMethod = "webhook"
	SubscriptionUpdateDeliveryParamsDeliveryMethodBothEmailSMS SubscriptionUpdateDeliveryParamsDeliveryMethod = "both_email_sms"
)
