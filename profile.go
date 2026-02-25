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

// Creates a new intelligence profile with the specified configuration. The profile
// will be owned by the authenticated user and start with `active` status.
func (r *ProfileService) New(ctx context.Context, body ProfileNewParams, opts ...option.RequestOption) (res *ProfileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Replaces all mutable fields of an existing intelligence profile. Only profiles
// owned by the authenticated user can be updated.
func (r *ProfileService) Update(ctx context.Context, profileID string, body ProfileUpdateParams, opts ...option.RequestOption) (res *ProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
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

// Permanently deletes an intelligence profile and all associated subscriptions.
// Only profiles owned by the authenticated user can be deleted. This action cannot
// be undone.
func (r *ProfileService) Delete(ctx context.Context, profileID string, opts ...option.RequestOption) (res *ProfileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Partially updates an existing intelligence profile. Only the fields included in
// the request body will be modified; all other fields remain unchanged. Only
// profiles owned by the authenticated user can be updated.
func (r *ProfileService) PartialUpdate(ctx context.Context, profileID string, body ProfilePartialUpdateParams, opts ...option.RequestOption) (res *ProfilePartialUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ProfileNewResponse struct {
	Data ProfileNewResponseData `json:"data" api:"required"`
	Meta ProfileNewResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewResponseData struct {
	// The ID of the newly created profile
	ProfileID string `json:"profileId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProfileID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateResponse struct {
	Data ProfileUpdateResponseData `json:"data" api:"required"`
	Meta ProfileUpdateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateResponseData struct {
	ProfileID string `json:"profileId" api:"required"`
	Success   bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProfileID   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListResponse struct {
	Data []ProfileListResponseData `json:"data" api:"required"`
	Meta ProfileListResponseMeta   `json:"meta" api:"required"`
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
	DeliveryMethod string                         `json:"deliveryMethod" api:"required"`
	IsActive       bool                           `json:"isActive" api:"required"`
	ProfileID      string                         `json:"profileId" api:"required"`
	SubscribedAt   int64                          `json:"subscribedAt" api:"required"`
	SubscriptionID string                         `json:"subscriptionId" api:"required"`
	Profile        ProfileListResponseDataProfile `json:"profile" api:"nullable"`
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

type ProfileDeleteResponse struct {
	Data ProfileDeleteResponseData `json:"data" api:"required"`
	Meta ProfileDeleteResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileDeleteResponseData struct {
	Deleted   bool   `json:"deleted" api:"required"`
	ProfileID string `json:"profileId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ProfileID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileDeleteResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDeleteResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileDeleteResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePartialUpdateResponse struct {
	Data ProfilePartialUpdateResponseData `json:"data" api:"required"`
	Meta ProfilePartialUpdateResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfilePartialUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfilePartialUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePartialUpdateResponseData struct {
	ProfileID string `json:"profileId" api:"required"`
	Success   bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProfileID   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfilePartialUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfilePartialUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePartialUpdateResponseMeta struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfilePartialUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfilePartialUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewParams struct {
	// Report generation frequency
	//
	// Any of "daily", "weekly", "biweekly", "monthly".
	Frequency ProfileNewParamsFrequency `json:"frequency,omitzero" api:"required"`
	// Profile display name
	Name string `json:"name" api:"required"`
	// Time of day for report generation (HH:mm, UTC)
	ScheduleTimeOfDay string `json:"scheduleTimeOfDay" api:"required"`
	// Topic description for research
	Topic string `json:"topic" api:"required"`
	// Custom BLUF report structure template
	BlufStructure param.Opt[string] `json:"blufStructure,omitzero"`
	// Custom system prompt for the AI analyst
	CustomPrompt param.Opt[string] `json:"customPrompt,omitzero"`
	// Whether this is a community (public) profile
	IsCommunity param.Opt[bool] `json:"isCommunity,omitzero"`
	// Day of month for monthly profiles
	ScheduleDayOfMonth param.Opt[string] `json:"scheduleDayOfMonth,omitzero"`
	// Day of week for weekly/biweekly profiles
	ScheduleDayOfWeek param.Opt[string]               `json:"scheduleDayOfWeek,omitzero"`
	RecursionConfig   ProfileNewParamsRecursionConfig `json:"recursionConfig,omitzero"`
	// Tags for categorization
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Report generation frequency
type ProfileNewParamsFrequency string

const (
	ProfileNewParamsFrequencyDaily    ProfileNewParamsFrequency = "daily"
	ProfileNewParamsFrequencyWeekly   ProfileNewParamsFrequency = "weekly"
	ProfileNewParamsFrequencyBiweekly ProfileNewParamsFrequency = "biweekly"
	ProfileNewParamsFrequencyMonthly  ProfileNewParamsFrequency = "monthly"
)

// The properties Enabled, MaxDepth, Strategy are required.
type ProfileNewParamsRecursionConfig struct {
	Enabled  bool  `json:"enabled" api:"required"`
	MaxDepth int64 `json:"maxDepth" api:"required"`
	// Any of "breadth-first", "depth-first", "hybrid".
	Strategy string `json:"strategy,omitzero" api:"required"`
	paramObj
}

func (r ProfileNewParamsRecursionConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsRecursionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsRecursionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileNewParamsRecursionConfig](
		"strategy", "breadth-first", "depth-first", "hybrid",
	)
}

type ProfileUpdateParams struct {
	BlufStructure param.Opt[string] `json:"blufStructure,omitzero"`
	// Branding template ID (Pro feature)
	BrandingTemplateID param.Opt[string]               `json:"brandingTemplateId,omitzero"`
	CustomPrompt       param.Opt[string]               `json:"customPrompt,omitzero"`
	IsCommunity        param.Opt[bool]                 `json:"isCommunity,omitzero"`
	Name               param.Opt[string]               `json:"name,omitzero"`
	ScheduleDayOfMonth param.Opt[string]               `json:"scheduleDayOfMonth,omitzero"`
	ScheduleDayOfWeek  param.Opt[string]               `json:"scheduleDayOfWeek,omitzero"`
	ScheduleTimeOfDay  param.Opt[string]               `json:"scheduleTimeOfDay,omitzero"`
	Topic              param.Opt[string]               `json:"topic,omitzero"`
	AudioConfig        ProfileUpdateParamsAudioConfig  `json:"audioConfig,omitzero"`
	BudgetConfig       ProfileUpdateParamsBudgetConfig `json:"budgetConfig,omitzero"`
	// Report generation frequency
	//
	// Any of "daily", "weekly", "biweekly", "monthly".
	Frequency       ProfileUpdateParamsFrequency       `json:"frequency,omitzero"`
	FreshnessConfig ProfileUpdateParamsFreshnessConfig `json:"freshnessConfig,omitzero"`
	ModelConfig     ProfileUpdateParamsModelConfig     `json:"modelConfig,omitzero"`
	RecursionConfig ProfileUpdateParamsRecursionConfig `json:"recursionConfig,omitzero"`
	SearchConfig    ProfileUpdateParamsSearchConfig    `json:"searchConfig,omitzero"`
	// Profile status
	//
	// Any of "active", "paused", "cancelled".
	Status ProfileUpdateParamsStatus `json:"status,omitzero"`
	Tags   []string                  `json:"tags,omitzero"`
	paramObj
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParamsAudioConfig struct {
	Enabled param.Opt[bool]    `json:"enabled,omitzero"`
	Speed   param.Opt[float64] `json:"speed,omitzero"`
	VoiceID param.Opt[string]  `json:"voiceId,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsAudioConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsAudioConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsAudioConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParamsBudgetConfig struct {
	AlertThreshold   param.Opt[float64] `json:"alertThreshold,omitzero"`
	MaxCostPerReport param.Opt[float64] `json:"maxCostPerReport,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBudgetConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBudgetConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBudgetConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Report generation frequency
type ProfileUpdateParamsFrequency string

const (
	ProfileUpdateParamsFrequencyDaily    ProfileUpdateParamsFrequency = "daily"
	ProfileUpdateParamsFrequencyWeekly   ProfileUpdateParamsFrequency = "weekly"
	ProfileUpdateParamsFrequencyBiweekly ProfileUpdateParamsFrequency = "biweekly"
	ProfileUpdateParamsFrequencyMonthly  ProfileUpdateParamsFrequency = "monthly"
)

type ProfileUpdateParamsFreshnessConfig struct {
	Enabled             param.Opt[bool]    `json:"enabled,omitzero"`
	MaxAgeMs            param.Opt[int64]   `json:"maxAgeMs,omitzero"`
	PreferRecentSources param.Opt[bool]    `json:"preferRecentSources,omitzero"`
	RecencyWeight       param.Opt[float64] `json:"recencyWeight,omitzero"`
	ValidateLinks       param.Opt[bool]    `json:"validateLinks,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsFreshnessConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsFreshnessConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsFreshnessConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParamsModelConfig struct {
	MaxOutputTokens param.Opt[int64]   `json:"maxOutputTokens,omitzero"`
	ModelID         param.Opt[string]  `json:"modelId,omitzero"`
	Temperature     param.Opt[float64] `json:"temperature,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsModelConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsModelConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsModelConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Enabled, MaxDepth, Strategy are required.
type ProfileUpdateParamsRecursionConfig struct {
	Enabled  bool  `json:"enabled" api:"required"`
	MaxDepth int64 `json:"maxDepth" api:"required"`
	// Any of "breadth-first", "depth-first", "hybrid".
	Strategy string `json:"strategy,omitzero" api:"required"`
	paramObj
}

func (r ProfileUpdateParamsRecursionConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsRecursionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsRecursionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileUpdateParamsRecursionConfig](
		"strategy", "breadth-first", "depth-first", "hybrid",
	)
}

type ProfileUpdateParamsSearchConfig struct {
	MaxResults     param.Opt[int64]  `json:"maxResults,omitzero"`
	TimeRange      param.Opt[string] `json:"timeRange,omitzero"`
	Topic          param.Opt[string] `json:"topic,omitzero"`
	ExcludeDomains []string          `json:"excludeDomains,omitzero"`
	IncludeDomains []string          `json:"includeDomains,omitzero"`
	// Any of "basic", "advanced".
	SearchDepth string `json:"searchDepth,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsSearchConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsSearchConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsSearchConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileUpdateParamsSearchConfig](
		"searchDepth", "basic", "advanced",
	)
}

// Profile status
type ProfileUpdateParamsStatus string

const (
	ProfileUpdateParamsStatusActive    ProfileUpdateParamsStatus = "active"
	ProfileUpdateParamsStatusPaused    ProfileUpdateParamsStatus = "paused"
	ProfileUpdateParamsStatusCancelled ProfileUpdateParamsStatus = "cancelled"
)

type ProfilePartialUpdateParams struct {
	BlufStructure param.Opt[string] `json:"blufStructure,omitzero"`
	// Branding template ID (Pro feature)
	BrandingTemplateID param.Opt[string]                      `json:"brandingTemplateId,omitzero"`
	CustomPrompt       param.Opt[string]                      `json:"customPrompt,omitzero"`
	IsCommunity        param.Opt[bool]                        `json:"isCommunity,omitzero"`
	Name               param.Opt[string]                      `json:"name,omitzero"`
	ScheduleDayOfMonth param.Opt[string]                      `json:"scheduleDayOfMonth,omitzero"`
	ScheduleDayOfWeek  param.Opt[string]                      `json:"scheduleDayOfWeek,omitzero"`
	ScheduleTimeOfDay  param.Opt[string]                      `json:"scheduleTimeOfDay,omitzero"`
	Topic              param.Opt[string]                      `json:"topic,omitzero"`
	AudioConfig        ProfilePartialUpdateParamsAudioConfig  `json:"audioConfig,omitzero"`
	BudgetConfig       ProfilePartialUpdateParamsBudgetConfig `json:"budgetConfig,omitzero"`
	// Report generation frequency
	//
	// Any of "daily", "weekly", "biweekly", "monthly".
	Frequency       ProfilePartialUpdateParamsFrequency       `json:"frequency,omitzero"`
	FreshnessConfig ProfilePartialUpdateParamsFreshnessConfig `json:"freshnessConfig,omitzero"`
	ModelConfig     ProfilePartialUpdateParamsModelConfig     `json:"modelConfig,omitzero"`
	RecursionConfig ProfilePartialUpdateParamsRecursionConfig `json:"recursionConfig,omitzero"`
	SearchConfig    ProfilePartialUpdateParamsSearchConfig    `json:"searchConfig,omitzero"`
	// Profile status
	//
	// Any of "active", "paused", "cancelled".
	Status ProfilePartialUpdateParamsStatus `json:"status,omitzero"`
	Tags   []string                         `json:"tags,omitzero"`
	paramObj
}

func (r ProfilePartialUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePartialUpdateParamsAudioConfig struct {
	Enabled param.Opt[bool]    `json:"enabled,omitzero"`
	Speed   param.Opt[float64] `json:"speed,omitzero"`
	VoiceID param.Opt[string]  `json:"voiceId,omitzero"`
	paramObj
}

func (r ProfilePartialUpdateParamsAudioConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParamsAudioConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParamsAudioConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePartialUpdateParamsBudgetConfig struct {
	AlertThreshold   param.Opt[float64] `json:"alertThreshold,omitzero"`
	MaxCostPerReport param.Opt[float64] `json:"maxCostPerReport,omitzero"`
	paramObj
}

func (r ProfilePartialUpdateParamsBudgetConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParamsBudgetConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParamsBudgetConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Report generation frequency
type ProfilePartialUpdateParamsFrequency string

const (
	ProfilePartialUpdateParamsFrequencyDaily    ProfilePartialUpdateParamsFrequency = "daily"
	ProfilePartialUpdateParamsFrequencyWeekly   ProfilePartialUpdateParamsFrequency = "weekly"
	ProfilePartialUpdateParamsFrequencyBiweekly ProfilePartialUpdateParamsFrequency = "biweekly"
	ProfilePartialUpdateParamsFrequencyMonthly  ProfilePartialUpdateParamsFrequency = "monthly"
)

type ProfilePartialUpdateParamsFreshnessConfig struct {
	Enabled             param.Opt[bool]    `json:"enabled,omitzero"`
	MaxAgeMs            param.Opt[int64]   `json:"maxAgeMs,omitzero"`
	PreferRecentSources param.Opt[bool]    `json:"preferRecentSources,omitzero"`
	RecencyWeight       param.Opt[float64] `json:"recencyWeight,omitzero"`
	ValidateLinks       param.Opt[bool]    `json:"validateLinks,omitzero"`
	paramObj
}

func (r ProfilePartialUpdateParamsFreshnessConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParamsFreshnessConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParamsFreshnessConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePartialUpdateParamsModelConfig struct {
	MaxOutputTokens param.Opt[int64]   `json:"maxOutputTokens,omitzero"`
	ModelID         param.Opt[string]  `json:"modelId,omitzero"`
	Temperature     param.Opt[float64] `json:"temperature,omitzero"`
	paramObj
}

func (r ProfilePartialUpdateParamsModelConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParamsModelConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParamsModelConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Enabled, MaxDepth, Strategy are required.
type ProfilePartialUpdateParamsRecursionConfig struct {
	Enabled  bool  `json:"enabled" api:"required"`
	MaxDepth int64 `json:"maxDepth" api:"required"`
	// Any of "breadth-first", "depth-first", "hybrid".
	Strategy string `json:"strategy,omitzero" api:"required"`
	paramObj
}

func (r ProfilePartialUpdateParamsRecursionConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParamsRecursionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParamsRecursionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfilePartialUpdateParamsRecursionConfig](
		"strategy", "breadth-first", "depth-first", "hybrid",
	)
}

type ProfilePartialUpdateParamsSearchConfig struct {
	MaxResults     param.Opt[int64]  `json:"maxResults,omitzero"`
	TimeRange      param.Opt[string] `json:"timeRange,omitzero"`
	Topic          param.Opt[string] `json:"topic,omitzero"`
	ExcludeDomains []string          `json:"excludeDomains,omitzero"`
	IncludeDomains []string          `json:"includeDomains,omitzero"`
	// Any of "basic", "advanced".
	SearchDepth string `json:"searchDepth,omitzero"`
	paramObj
}

func (r ProfilePartialUpdateParamsSearchConfig) MarshalJSON() (data []byte, err error) {
	type shadow ProfilePartialUpdateParamsSearchConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfilePartialUpdateParamsSearchConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfilePartialUpdateParamsSearchConfig](
		"searchDepth", "basic", "advanced",
	)
}

// Profile status
type ProfilePartialUpdateParamsStatus string

const (
	ProfilePartialUpdateParamsStatusActive    ProfilePartialUpdateParamsStatus = "active"
	ProfilePartialUpdateParamsStatusPaused    ProfilePartialUpdateParamsStatus = "paused"
	ProfilePartialUpdateParamsStatusCancelled ProfilePartialUpdateParamsStatus = "cancelled"
)
