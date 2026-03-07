// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/y2-intel/y2-go/internal/apijson"
	"github.com/y2-intel/y2-go/internal/requestconfig"
	"github.com/y2-intel/y2-go/option"
	"github.com/y2-intel/y2-go/packages/respjson"
)

// Situation Room OSINT intelligence operations
//
// OsintSourceService contains methods and other services that help with
// interacting with the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOsintSourceService] method instead.
type OsintSourceService struct {
	Options []option.RequestOption
}

// NewOsintSourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOsintSourceService(opts ...option.RequestOption) (r OsintSourceService) {
	r = OsintSourceService{}
	r.Options = opts
	return
}

// Returns the health status of all OSINT data sources, including circuit breaker
// state and failure counts.
func (r *OsintSourceService) GetDataSourceHealth(ctx context.Context, opts ...option.RequestOption) (res *OsintSourceGetDataSourceHealthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/sources/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OsintSourceGetDataSourceHealthResponse struct {
	Data []OsintSourceGetDataSourceHealthResponseData `json:"data" api:"required"`
	Meta OsintSourceGetDataSourceHealthResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintSourceGetDataSourceHealthResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintSourceGetDataSourceHealthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintSourceGetDataSourceHealthResponseData struct {
	// Number of consecutive failures
	FailureCount int64 `json:"failureCount" api:"required"`
	// Data source identifier
	SourceType string `json:"sourceType" api:"required"`
	// Circuit breaker state
	//
	// Any of "closed", "open", "half_open".
	State string `json:"state" api:"required"`
	// Most recent error message
	LastError string `json:"lastError"`
	// Last failure time (milliseconds)
	LastFailureAt int64 `json:"lastFailureAt"`
	// Last failure time as ISO 8601 string
	LastFailureAtISO time.Time `json:"lastFailureAtISO" api:"nullable" format:"date-time"`
	// Last successful fetch time (milliseconds)
	LastSuccessAt int64 `json:"lastSuccessAt"`
	// Last successful fetch time as ISO 8601 string
	LastSuccessAtISO time.Time `json:"lastSuccessAtISO" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailureCount     respjson.Field
		SourceType       respjson.Field
		State            respjson.Field
		LastError        respjson.Field
		LastFailureAt    respjson.Field
		LastFailureAtISO respjson.Field
		LastSuccessAt    respjson.Field
		LastSuccessAtISO respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintSourceGetDataSourceHealthResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintSourceGetDataSourceHealthResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintSourceGetDataSourceHealthResponseMeta struct {
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintSourceGetDataSourceHealthResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintSourceGetDataSourceHealthResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
