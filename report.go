// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/y2-intel/y2-go/internal/apijson"
	"github.com/y2-intel/y2-go/internal/apiquery"
	"github.com/y2-intel/y2-go/internal/requestconfig"
	"github.com/y2-intel/y2-go/option"
	"github.com/y2-intel/y2-go/packages/param"
	"github.com/y2-intel/y2-go/packages/respjson"
)

// Intelligence report operations
//
// ReportService contains methods and other services that help with interacting
// with the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportService] method instead.
type ReportService struct {
	Options []option.RequestOption
}

// NewReportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReportService(opts ...option.RequestOption) (r ReportService) {
	r = ReportService{}
	r.Options = opts
	return
}

// Returns the full content of a specific intelligence report, including HTML
// content, sources, and audio metadata.
func (r *ReportService) Get(ctx context.Context, reportID string, opts ...option.RequestOption) (res *ReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return
	}
	path := fmt.Sprintf("reports/%s", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of reports for the user's subscribed profiles. Results are sorted
// by generation date (newest first).
func (r *ReportService) List(ctx context.Context, query ReportListParams, opts ...option.RequestOption) (res *ReportListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns audio file metadata or redirects to the CDN URL. Requires the
// `reports:audio` scope.
func (r *ReportService) GetAudio(ctx context.Context, reportID string, query ReportGetAudioParams, opts ...option.RequestOption) (res *ReportGetAudioResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return
	}
	path := fmt.Sprintf("reports/%s/audio", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AudioMetadata struct {
	// Duration in seconds
	Duration float64 `json:"duration"`
	// Duration as HH:MM:SS
	DurationFormatted string `json:"durationFormatted"`
	// File size in bytes
	FileSize int64 `json:"fileSize"`
	// Any of "mp3".
	Format AudioMetadataFormat `json:"format"`
	// Any of "audio/mpeg".
	MimeType AudioMetadataMimeType `json:"mimeType"`
	// Convex storage ID for internal reference
	StorageID string `json:"storageId"`
	// CDN URL for audio file
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration          respjson.Field
		DurationFormatted respjson.Field
		FileSize          respjson.Field
		Format            respjson.Field
		MimeType          respjson.Field
		StorageID         respjson.Field
		URL               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioMetadata) RawJSON() string { return r.JSON.raw }
func (r *AudioMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioMetadataFormat string

const (
	AudioMetadataFormatMP3 AudioMetadataFormat = "mp3"
)

type AudioMetadataMimeType string

const (
	AudioMetadataMimeTypeAudioMpeg AudioMetadataMimeType = "audio/mpeg"
)

type ReportGetResponse struct {
	Data ReportGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportGetResponseData struct {
	ID             string                        `json:"id" api:"required"`
	Content        ReportGetResponseDataContent  `json:"content" api:"required"`
	GeneratedAt    int64                         `json:"generatedAt" api:"required"`
	GeneratedAtISO time.Time                     `json:"generatedAtISO" api:"required" format:"date-time"`
	ProfileID      string                        `json:"profileId" api:"required"`
	Audio          AudioMetadata                 `json:"audio" api:"nullable"`
	Metadata       ReportGetResponseDataMetadata `json:"metadata"`
	ProfileName    string                        `json:"profileName"`
	ProfileTopic   string                        `json:"profileTopic"`
	Sources        []string                      `json:"sources" format:"uri"`
	Topic          string                        `json:"topic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Content        respjson.Field
		GeneratedAt    respjson.Field
		GeneratedAtISO respjson.Field
		ProfileID      respjson.Field
		Audio          respjson.Field
		Metadata       respjson.Field
		ProfileName    respjson.Field
		ProfileTopic   respjson.Field
		Sources        respjson.Field
		Topic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReportGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportGetResponseDataContent struct {
	// Full HTML content
	HTML string `json:"html"`
	// SMS-friendly summary
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTML        respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetResponseDataContent) RawJSON() string { return r.JSON.raw }
func (r *ReportGetResponseDataContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportGetResponseDataMetadata struct {
	// Source freshness validation results
	FreshnessMetadata ReportGetResponseDataMetadataFreshnessMetadata `json:"freshnessMetadata"`
	Model             string                                         `json:"model"`
	// Metadata about recursive research execution
	RecursionMetadata ReportGetResponseDataMetadataRecursionMetadata `json:"recursionMetadata"`
	TotalCost         float64                                        `json:"totalCost"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FreshnessMetadata respjson.Field
		Model             respjson.Field
		RecursionMetadata respjson.Field
		TotalCost         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetResponseDataMetadata) RawJSON() string { return r.JSON.raw }
func (r *ReportGetResponseDataMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source freshness validation results
type ReportGetResponseDataMetadataFreshnessMetadata struct {
	AccessibleLinks int64 `json:"accessibleLinks"`
	// Average source age in milliseconds
	AverageAgeMs int64 `json:"averageAgeMs"`
	// Overall freshness score (higher = fresher)
	FreshnessScore    float64 `json:"freshnessScore"`
	StaleSourcesCount int64   `json:"staleSourcesCount"`
	TotalLinks        int64   `json:"totalLinks"`
	ValidatedAt       int64   `json:"validatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessibleLinks   respjson.Field
		AverageAgeMs      respjson.Field
		FreshnessScore    respjson.Field
		StaleSourcesCount respjson.Field
		TotalLinks        respjson.Field
		ValidatedAt       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetResponseDataMetadataFreshnessMetadata) RawJSON() string { return r.JSON.raw }
func (r *ReportGetResponseDataMetadataFreshnessMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about recursive research execution
type ReportGetResponseDataMetadataRecursionMetadata struct {
	// Recursion depth achieved (0 = standard report)
	Depth int64 `json:"depth"`
	// Reason if fallback to standard generation occurred
	FallbackReason  string `json:"fallbackReason"`
	LayersProcessed int64  `json:"layersProcessed"`
	// Any of "breadth-first", "depth-first", "hybrid".
	Strategy                string   `json:"strategy"`
	SubtopicsGenerated      []string `json:"subtopicsGenerated"`
	TotalSourcesCollected   int64    `json:"totalSourcesCollected"`
	TotalTimeMs             int64    `json:"totalTimeMs"`
	UniqueSourcesAggregated int64    `json:"uniqueSourcesAggregated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Depth                   respjson.Field
		FallbackReason          respjson.Field
		LayersProcessed         respjson.Field
		Strategy                respjson.Field
		SubtopicsGenerated      respjson.Field
		TotalSourcesCollected   respjson.Field
		TotalTimeMs             respjson.Field
		UniqueSourcesAggregated respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetResponseDataMetadataRecursionMetadata) RawJSON() string { return r.JSON.raw }
func (r *ReportGetResponseDataMetadataRecursionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListResponse struct {
	Data []ReportListResponseData `json:"data" api:"required"`
	Meta ReportListResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListResponseData struct {
	// Report ID (Convex document ID)
	ID string `json:"id" api:"required"`
	// Unix timestamp (milliseconds)
	GeneratedAt int64 `json:"generatedAt" api:"required"`
	// ISO 8601 timestamp
	GeneratedAtISO time.Time `json:"generatedAtISO" api:"required" format:"date-time"`
	// Profile ID this report belongs to
	ProfileID string `json:"profileId" api:"required"`
	// Whether audio narration is available
	HasAudio bool `json:"hasAudio"`
	// LLM model used for generation
	Model string `json:"model"`
	// Brief SMS-friendly summary
	Summary string `json:"summary"`
	// Report topic
	Topic string `json:"topic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		GeneratedAt    respjson.Field
		GeneratedAtISO respjson.Field
		ProfileID      respjson.Field
		HasAudio       respjson.Field
		Model          respjson.Field
		Summary        respjson.Field
		Topic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReportListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListResponseMeta struct {
	Count int64 `json:"count"`
	Limit int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ReportListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportGetAudioResponse struct {
	Data AudioMetadata `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportGetAudioResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportGetAudioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListParams struct {
	// Maximum number of reports to return (hard-capped at 5)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter reports by profile ID
	ProfileID param.Opt[string] `query:"profileId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportListParams]'s query parameters as `url.Values`.
func (r ReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportGetAudioParams struct {
	// If true, returns 302 redirect to audio CDN URL
	Redirect param.Opt[bool] `query:"redirect,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportGetAudioParams]'s query parameters as `url.Values`.
func (r ReportGetAudioParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
