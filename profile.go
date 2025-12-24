// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2

import (
	"context"
	"net/http"
	"slices"

	"github.com/y2-intel/y2-go/internal/apijson"
	"github.com/y2-intel/y2-go/internal/requestconfig"
	"github.com/y2-intel/y2-go/option"
	"github.com/y2-intel/y2-go/packages/respjson"
)

// ProfileService contains methods and other services that help with interacting
// with the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	return
}

// Returns a list of intelligence profiles the user is subscribed to, including
// subscription status and delivery preferences.
func (r *ProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *ProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProfileListResponse struct {
	Data []ProfileListResponseData `json:"data,required"`
	Meta ProfileListResponseMeta   `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListResponseData struct {
	// Any of "email", "sms", "webhook", "both_email_sms".
	DeliveryMethod string                         `json:"deliveryMethod,required"`
	IsActive       bool                           `json:"isActive,required"`
	ProfileID      string                         `json:"profileId,required"`
	SubscribedAt   int64                          `json:"subscribedAt,required"`
	SubscriptionID string                         `json:"subscriptionId,required"`
	Profile        ProfileListResponseDataProfile `json:"profile,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryMethod respjson.Field
		IsActive       respjson.Field
		ProfileID      respjson.Field
		SubscribedAt   respjson.Field
		SubscriptionID respjson.Field
		Profile        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListResponseDataProfile struct {
	AudioEnabled bool   `json:"audioEnabled"`
	Frequency    string `json:"frequency"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Topic        string `json:"topic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioEnabled respjson.Field
		Frequency    respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		Topic        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseDataProfile) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListResponseMeta struct {
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
