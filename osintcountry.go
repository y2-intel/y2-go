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

// Situation Room OSINT intelligence operations
//
// OsintCountryService contains methods and other services that help with
// interacting with the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOsintCountryService] method instead.
type OsintCountryService struct {
	Options []option.RequestOption
}

// NewOsintCountryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOsintCountryService(opts ...option.RequestOption) (r OsintCountryService) {
	r = OsintCountryService{}
	r.Options = opts
	return
}

// Returns the per-country Conflict Indicators Index (CII) score, including
// baseline, delta, and component breakdown.
func (r *OsintCountryService) GetCountryInstabilityIndex(ctx context.Context, countryCode string, opts ...option.RequestOption) (res *OsintCountryGetCountryInstabilityIndexResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if countryCode == "" {
		err = errors.New("missing required countryCode parameter")
		return
	}
	path := fmt.Sprintf("osint/countries/%s/cii", countryCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns recent news items specific to a given country, sourced from the OSINT
// event pipeline.
func (r *OsintCountryService) GetCountryNews(ctx context.Context, countryCode string, query OsintCountryGetCountryNewsParams, opts ...option.RequestOption) (res *OsintCountryGetCountryNewsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if countryCode == "" {
		err = errors.New("missing required countryCode parameter")
		return
	}
	path := fmt.Sprintf("osint/countries/%s/news", countryCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns an AI-generated intelligence brief for a specific country. Briefs are
// generated periodically and cached.
func (r *OsintCountryService) GetIntelligenceBrief(ctx context.Context, countryCode string, opts ...option.RequestOption) (res *OsintCountryGetIntelligenceBriefResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if countryCode == "" {
		err = errors.New("missing required countryCode parameter")
		return
	}
	path := fmt.Sprintf("osint/countries/%s/brief", countryCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns prediction market data for a specific country, including probabilities
// and trading volumes.
func (r *OsintCountryService) GetPredictionMarkets(ctx context.Context, countryCode string, query OsintCountryGetPredictionMarketsParams, opts ...option.RequestOption) (res *OsintCountryGetPredictionMarketsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if countryCode == "" {
		err = errors.New("missing required countryCode parameter")
		return
	}
	path := fmt.Sprintf("osint/countries/%s/predictions", countryCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns the primary stock market index data for a specific country, including
// weekly change and currency.
func (r *OsintCountryService) GetStockMarketIndex(ctx context.Context, countryCode string, opts ...option.RequestOption) (res *OsintCountryGetStockMarketIndexResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if countryCode == "" {
		err = errors.New("missing required countryCode parameter")
		return
	}
	path := fmt.Sprintf("osint/countries/%s/markets", countryCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OsintCountryGetCountryInstabilityIndexResponse struct {
	Data OsintCountryGetCountryInstabilityIndexResponseData `json:"data" api:"required"`
	Meta OsintCountryGetCountryInstabilityIndexResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryInstabilityIndexResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetCountryInstabilityIndexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetCountryInstabilityIndexResponseData struct {
	// Baseline instability score for this country
	Baseline float64 `json:"baseline" api:"required"`
	// Computation time as Unix timestamp (milliseconds)
	ComputedAt int64 `json:"computedAt" api:"required"`
	// ISO alpha-2 country code
	CountryCode string `json:"countryCode" api:"required"`
	// Country display name
	CountryName string `json:"countryName" api:"required"`
	// Change from baseline
	Delta float64 `json:"delta" api:"required"`
	// Current instability index (0-100)
	Value float64 `json:"value" api:"required"`
	// Breakdown of instability components
	Components OsintCountryGetCountryInstabilityIndexResponseDataComponents `json:"components"`
	// Computation time as ISO 8601 string
	ComputedAtISO time.Time `json:"computedAtISO" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Baseline      respjson.Field
		ComputedAt    respjson.Field
		CountryCode   respjson.Field
		CountryName   respjson.Field
		Delta         respjson.Field
		Value         respjson.Field
		Components    respjson.Field
		ComputedAtISO respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryInstabilityIndexResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetCountryInstabilityIndexResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of instability components
type OsintCountryGetCountryInstabilityIndexResponseDataComponents struct {
	EventCounts any     `json:"eventCounts"`
	Info        float64 `json:"info"`
	Multiplier  float64 `json:"multiplier"`
	Security    float64 `json:"security"`
	Unrest      float64 `json:"unrest"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventCounts respjson.Field
		Info        respjson.Field
		Multiplier  respjson.Field
		Security    respjson.Field
		Unrest      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryInstabilityIndexResponseDataComponents) RawJSON() string {
	return r.JSON.raw
}
func (r *OsintCountryGetCountryInstabilityIndexResponseDataComponents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetCountryInstabilityIndexResponseMeta struct {
	CountryCode string `json:"countryCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryInstabilityIndexResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetCountryInstabilityIndexResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetCountryNewsResponse struct {
	Data []OsintCountryGetCountryNewsResponseData `json:"data" api:"required"`
	Meta OsintCountryGetCountryNewsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetCountryNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetCountryNewsResponseData struct {
	// News item ID
	ID string `json:"id" api:"required"`
	// OSINT event category classification
	//
	// Any of "seismic", "conflict", "political", "economic", "weather", "health",
	// "cyber", "maritime", "fire", "aviation", "other".
	Category string `json:"category" api:"required"`
	// Event time as Unix timestamp (milliseconds)
	EventTime int64 `json:"eventTime" api:"required"`
	// Event severity level
	//
	// Any of "low", "medium", "high", "critical".
	Severity string `json:"severity" api:"required"`
	// News headline
	Title string `json:"title" api:"required"`
	// News description/summary
	Description string `json:"description"`
	// Data source type
	SourceType string `json:"sourceType"`
	// Source URL
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		EventTime   respjson.Field
		Severity    respjson.Field
		Title       respjson.Field
		Description respjson.Field
		SourceType  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryNewsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetCountryNewsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetCountryNewsResponseMeta struct {
	Count       int64  `json:"count"`
	CountryCode string `json:"countryCode"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		CountryCode respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetCountryNewsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetCountryNewsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetIntelligenceBriefResponse struct {
	Data OsintCountryGetIntelligenceBriefResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetIntelligenceBriefResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetIntelligenceBriefResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetIntelligenceBriefResponseData struct {
	// AI-generated intelligence brief text
	BriefText string `json:"briefText" api:"required"`
	// Generation time as Unix timestamp (milliseconds)
	GeneratedAt int64 `json:"generatedAt" api:"required"`
	// LLM model used for generation
	Model string `json:"model" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BriefText   respjson.Field
		GeneratedAt respjson.Field
		Model       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetIntelligenceBriefResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetIntelligenceBriefResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetPredictionMarketsResponse struct {
	Data []OsintCountryGetPredictionMarketsResponseData `json:"data" api:"required"`
	Meta OsintCountryGetPredictionMarketsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetPredictionMarketsResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetPredictionMarketsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetPredictionMarketsResponseData struct {
	// Prediction market identifier
	MarketID string `json:"marketId" api:"required"`
	// Current probability (0-1)
	Probability float64 `json:"probability" api:"required"`
	// Market question/title
	Title string `json:"title" api:"required"`
	// Market resolution date
	EndDate time.Time `json:"endDate" format:"date-time"`
	// Market liquidity (null if unavailable)
	Liquidity float64 `json:"liquidity" api:"nullable"`
	// Outcome prices corresponding to each outcome (null if unavailable)
	OutcomePrices []string `json:"outcomePrices" api:"nullable"`
	// Possible market outcomes
	Outcomes []string `json:"outcomes"`
	// Polymarket URL for this market
	PolymarketURL string `json:"polymarketUrl" format:"uri"`
	// URL-friendly market slug (null if unavailable)
	Slug string `json:"slug" api:"nullable"`
	// Trading volume
	Volume float64 `json:"volume"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MarketID      respjson.Field
		Probability   respjson.Field
		Title         respjson.Field
		EndDate       respjson.Field
		Liquidity     respjson.Field
		OutcomePrices respjson.Field
		Outcomes      respjson.Field
		PolymarketURL respjson.Field
		Slug          respjson.Field
		Volume        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetPredictionMarketsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetPredictionMarketsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetPredictionMarketsResponseMeta struct {
	Count       int64  `json:"count"`
	CountryCode string `json:"countryCode"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		CountryCode respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetPredictionMarketsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetPredictionMarketsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetStockMarketIndexResponse struct {
	Data OsintCountryGetStockMarketIndexResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetStockMarketIndexResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetStockMarketIndexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetStockMarketIndexResponseData struct {
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Data fetch time as Unix timestamp (milliseconds)
	FetchedAt int64 `json:"fetchedAt" api:"required"`
	// Human-readable index name
	IndexName string `json:"indexName" api:"required"`
	// Stock index ticker symbol
	Ticker string `json:"ticker" api:"required"`
	// Weekly change percentage
	WeeklyChange float64 `json:"weeklyChange" api:"required"`
	// Current index price/value
	CurrentPrice float64 `json:"currentPrice"`
	// Data fetch time as ISO 8601 string
	FetchedAtISO time.Time `json:"fetchedAtISO" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency     respjson.Field
		FetchedAt    respjson.Field
		IndexName    respjson.Field
		Ticker       respjson.Field
		WeeklyChange respjson.Field
		CurrentPrice respjson.Field
		FetchedAtISO respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintCountryGetStockMarketIndexResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintCountryGetStockMarketIndexResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintCountryGetCountryNewsParams struct {
	// Maximum number of news items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintCountryGetCountryNewsParams]'s query parameters as
// `url.Values`.
func (r OsintCountryGetCountryNewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OsintCountryGetPredictionMarketsParams struct {
	// Maximum number of predictions to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintCountryGetPredictionMarketsParams]'s query parameters
// as `url.Values`.
func (r OsintCountryGetPredictionMarketsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
