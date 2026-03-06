// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2

import (
	"context"
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
// OsintService contains methods and other services that help with interacting with
// the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOsintService] method instead.
type OsintService struct {
	Options []option.RequestOption
	// Situation Room OSINT intelligence operations
	Countries OsintCountryService
	// Situation Room OSINT intelligence operations
	Sources OsintSourceService
}

// NewOsintService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOsintService(opts ...option.RequestOption) (r OsintService) {
	r = OsintService{}
	r.Options = opts
	r.Countries = NewOsintCountryService(opts...)
	r.Sources = NewOsintSourceService(opts...)
	return
}

// Returns the Conflict Indicators Index (CII) values. Each item represents a
// conflict indicator with a score from 0-100 and a delta showing recent change.
// Supports filtering by region and category.
func (r *OsintService) GetConflictIndicators(ctx context.Context, query OsintGetConflictIndicatorsParams, opts ...option.RequestOption) (res *OsintGetConflictIndicatorsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/cii"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns GPS interference zones detected via ADS-B navigation accuracy
// degradation analysis, aggregated into H3 hex cells.
func (r *OsintService) GetGpsJammingZones(ctx context.Context, query OsintGetGpsJammingZonesParams, opts ...option.RequestOption) (res *OsintGetGpsJammingZonesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/gps-jamming"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns military posture assessments for monitored theaters, based on detected
// military aircraft activity from the OpenSky Network. Each theater has a posture
// level (normal, elevated, critical) and aircraft breakdown by type.
func (r *OsintService) GetMilitaryPosture(ctx context.Context, query OsintGetMilitaryPostureParams, opts ...option.RequestOption) (res *OsintGetMilitaryPostureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/military-posture"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns tracked military aircraft positions from the OpenSky Network, filtered
// and classified by type (tanker, AWACS, fighter, etc.).
func (r *OsintService) ListAircraft(ctx context.Context, query OsintListAircraftParams, opts ...option.RequestOption) (res *OsintListAircraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/aircraft"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns OSINT threat events from the Situation Room. Supports filtering by
// category, severity, region, and country.
func (r *OsintService) ListEvents(ctx context.Context, query OsintListEventsParams, opts ...option.RequestOption) (res *OsintListEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns naval vessel positions sourced from USNI fleet tracker data, including
// carrier strike groups and individual warships.
func (r *OsintService) ListVessels(ctx context.Context, query OsintListVesselsParams, opts ...option.RequestOption) (res *OsintListVesselsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/vessels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns OSINT events with geographic coordinates for map display. Events without
// coordinates are excluded.
func (r *OsintService) MapEvents(ctx context.Context, query OsintMapEventsParams, opts ...option.RequestOption) (res *OsintMapEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "osint/map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OsintGetConflictIndicatorsResponse struct {
	Data []OsintGetConflictIndicatorsResponseData `json:"data" api:"required"`
	Meta OsintGetConflictIndicatorsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetConflictIndicatorsResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintGetConflictIndicatorsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetConflictIndicatorsResponseData struct {
	// CII item ID
	ID string `json:"id" api:"required"`
	// Conflict indicator category
	Category string `json:"category" api:"required"`
	// Computation time as Unix timestamp (milliseconds)
	ComputedAt int64 `json:"computedAt" api:"required"`
	// Recent change in value
	Delta float64 `json:"delta" api:"required"`
	// Conflict indicator value (0-100)
	Value float64 `json:"value" api:"required"`
	// Breakdown of indicator components
	Components map[string]any `json:"components"`
	// Computation time as ISO 8601 string
	ComputedAtISO time.Time `json:"computedAtISO" format:"date-time"`
	// Geographic region identifier
	//
	// Any of "mena", "africa", "latam", "asiapac", "europe", "namerica".
	Region string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Category      respjson.Field
		ComputedAt    respjson.Field
		Delta         respjson.Field
		Value         respjson.Field
		Components    respjson.Field
		ComputedAtISO respjson.Field
		Region        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetConflictIndicatorsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintGetConflictIndicatorsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetConflictIndicatorsResponseMeta struct {
	Count   int64                                         `json:"count"`
	Filters OsintGetConflictIndicatorsResponseMetaFilters `json:"filters"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Filters     respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetConflictIndicatorsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintGetConflictIndicatorsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetConflictIndicatorsResponseMetaFilters struct {
	Category string `json:"category" api:"nullable"`
	Region   string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetConflictIndicatorsResponseMetaFilters) RawJSON() string { return r.JSON.raw }
func (r *OsintGetConflictIndicatorsResponseMetaFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetGpsJammingZonesResponse struct {
	Data []OsintGetGpsJammingZonesResponseData `json:"data" api:"required"`
	Meta OsintGetGpsJammingZonesResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetGpsJammingZonesResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintGetGpsJammingZonesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetGpsJammingZonesResponseData struct {
	// Zone ID
	ID          string                                         `json:"id" api:"required"`
	Coordinates OsintGetGpsJammingZonesResponseDataCoordinates `json:"coordinates" api:"required"`
	// Fetch time as Unix timestamp (milliseconds)
	FetchedAt int64 `json:"fetchedAt" api:"required"`
	// H3 hexagonal cell index
	H3Index string `json:"h3Index" api:"required"`
	// Severity classification of interference
	Severity string `json:"severity" api:"required"`
	// Number of distinct aircraft reporting
	AircraftCount int64 `json:"aircraftCount"`
	// Fetch time as ISO 8601 string
	FetchedAtISO time.Time `json:"fetchedAtISO" format:"date-time"`
	// Average navigation accuracy category
	NavAccuracyAvg float64 `json:"navAccuracyAvg"`
	// Number of ADS-B samples in this cell
	SampleCount int64 `json:"sampleCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Coordinates    respjson.Field
		FetchedAt      respjson.Field
		H3Index        respjson.Field
		Severity       respjson.Field
		AircraftCount  respjson.Field
		FetchedAtISO   respjson.Field
		NavAccuracyAvg respjson.Field
		SampleCount    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetGpsJammingZonesResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintGetGpsJammingZonesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetGpsJammingZonesResponseDataCoordinates struct {
	Lat float64 `json:"lat" api:"required"`
	Lon float64 `json:"lon" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetGpsJammingZonesResponseDataCoordinates) RawJSON() string { return r.JSON.raw }
func (r *OsintGetGpsJammingZonesResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetGpsJammingZonesResponseMeta struct {
	Count   int64                                      `json:"count"`
	Filters OsintGetGpsJammingZonesResponseMetaFilters `json:"filters"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Filters     respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetGpsJammingZonesResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintGetGpsJammingZonesResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetGpsJammingZonesResponseMetaFilters struct {
	Severity string `json:"severity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetGpsJammingZonesResponseMetaFilters) RawJSON() string { return r.JSON.raw }
func (r *OsintGetGpsJammingZonesResponseMetaFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetMilitaryPostureResponse struct {
	Data []OsintGetMilitaryPostureResponseData `json:"data" api:"required"`
	Meta OsintGetMilitaryPostureResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetMilitaryPostureResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintGetMilitaryPostureResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetMilitaryPostureResponseData struct {
	// Number of detected military aircraft
	AircraftCount int64 `json:"aircraftCount" api:"required"`
	// Assessment time as Unix timestamp (milliseconds)
	ComputedAt int64 `json:"computedAt" api:"required"`
	// Current military posture assessment
	//
	// Any of "normal", "elevated", "critical".
	Posture string `json:"posture" api:"required"`
	// Theater name (e.g. "iran", "taiwan", "baltic")
	Theater string `json:"theater" api:"required"`
	// Aircraft count by type
	Breakdown OsintGetMilitaryPostureResponseDataBreakdown `json:"breakdown"`
	// Assessment time as ISO 8601 string
	ComputedAtISO time.Time `json:"computedAtISO" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AircraftCount respjson.Field
		ComputedAt    respjson.Field
		Posture       respjson.Field
		Theater       respjson.Field
		Breakdown     respjson.Field
		ComputedAtISO respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetMilitaryPostureResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintGetMilitaryPostureResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aircraft count by type
type OsintGetMilitaryPostureResponseDataBreakdown struct {
	Awacs           int64 `json:"awacs"`
	Bomber          int64 `json:"bomber"`
	Drone           int64 `json:"drone"`
	Fighter         int64 `json:"fighter"`
	MilitaryGeneric int64 `json:"military_generic"`
	Recon           int64 `json:"recon"`
	Tanker          int64 `json:"tanker"`
	Transport       int64 `json:"transport"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Awacs           respjson.Field
		Bomber          respjson.Field
		Drone           respjson.Field
		Fighter         respjson.Field
		MilitaryGeneric respjson.Field
		Recon           respjson.Field
		Tanker          respjson.Field
		Transport       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetMilitaryPostureResponseDataBreakdown) RawJSON() string { return r.JSON.raw }
func (r *OsintGetMilitaryPostureResponseDataBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetMilitaryPostureResponseMeta struct {
	Count int64 `json:"count"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintGetMilitaryPostureResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintGetMilitaryPostureResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListAircraftResponse struct {
	Data []OsintListAircraftResponseData `json:"data" api:"required"`
	Meta OsintListAircraftResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListAircraftResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintListAircraftResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListAircraftResponseData struct {
	// Track ID
	ID string `json:"id" api:"required"`
	// Aircraft callsign
	Callsign    string                                   `json:"callsign" api:"required"`
	Coordinates OsintListAircraftResponseDataCoordinates `json:"coordinates" api:"required"`
	// ICAO 24-bit transponder address
	Icao24 string `json:"icao24" api:"required"`
	// Last seen time as Unix timestamp (milliseconds)
	LastSeenAt int64 `json:"lastSeenAt" api:"required"`
	// Theater assignment
	Theater string `json:"theater" api:"required"`
	// Classified aircraft type (tanker, awacs, fighter, recon, etc.)
	AircraftType string `json:"aircraftType"`
	// Altitude in meters
	Altitude float64 `json:"altitude"`
	// Classification confidence score
	Confidence float64 `json:"confidence"`
	// Heading in degrees
	Heading float64 `json:"heading"`
	// Whether flagged as operationally interesting
	IsInteresting bool `json:"isInteresting"`
	// Last seen time as ISO 8601 string
	LastSeenAtISO time.Time `json:"lastSeenAtISO" format:"date-time"`
	// Operating entity
	Operator string `json:"operator"`
	// Ground speed in m/s
	Speed float64 `json:"speed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Callsign      respjson.Field
		Coordinates   respjson.Field
		Icao24        respjson.Field
		LastSeenAt    respjson.Field
		Theater       respjson.Field
		AircraftType  respjson.Field
		Altitude      respjson.Field
		Confidence    respjson.Field
		Heading       respjson.Field
		IsInteresting respjson.Field
		LastSeenAtISO respjson.Field
		Operator      respjson.Field
		Speed         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListAircraftResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintListAircraftResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListAircraftResponseDataCoordinates struct {
	Lat float64 `json:"lat" api:"required"`
	Lon float64 `json:"lon" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListAircraftResponseDataCoordinates) RawJSON() string { return r.JSON.raw }
func (r *OsintListAircraftResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListAircraftResponseMeta struct {
	Count   int64                                `json:"count"`
	Filters OsintListAircraftResponseMetaFilters `json:"filters"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Filters     respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListAircraftResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintListAircraftResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListAircraftResponseMetaFilters struct {
	Theater string `json:"theater" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Theater     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListAircraftResponseMetaFilters) RawJSON() string { return r.JSON.raw }
func (r *OsintListAircraftResponseMetaFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListEventsResponse struct {
	Data []OsintListEventsResponseData `json:"data" api:"required"`
	Meta OsintListEventsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListEventsResponseData struct {
	// Event ID
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
	// Source-specific event identifier
	SourceID string `json:"sourceId" api:"required"`
	// Data source type
	SourceType string `json:"sourceType" api:"required"`
	// Event title/headline
	Title string `json:"title" api:"required"`
	// Geographic coordinates (null if unavailable)
	Coordinates OsintListEventsResponseDataCoordinates `json:"coordinates" api:"nullable"`
	// ISO alpha-2 country code
	CountryCode string `json:"countryCode"`
	// Detailed event description
	Description string `json:"description"`
	// Event time as ISO 8601 string
	EventTimeISO time.Time `json:"eventTimeISO" format:"date-time"`
	// Data fetch time as Unix timestamp (milliseconds), null if unavailable
	FetchedAt int64 `json:"fetchedAt" api:"nullable"`
	// Data fetch time as ISO 8601 string, null if unavailable
	FetchedAtISO time.Time `json:"fetchedAtISO" api:"nullable" format:"date-time"`
	// Human-readable location name
	LocationName string `json:"locationName"`
	// Geographic region identifier
	//
	// Any of "mena", "africa", "latam", "asiapac", "europe", "namerica".
	Region string `json:"region"`
	// Source URL for the event
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Category     respjson.Field
		EventTime    respjson.Field
		Severity     respjson.Field
		SourceID     respjson.Field
		SourceType   respjson.Field
		Title        respjson.Field
		Coordinates  respjson.Field
		CountryCode  respjson.Field
		Description  respjson.Field
		EventTimeISO respjson.Field
		FetchedAt    respjson.Field
		FetchedAtISO respjson.Field
		LocationName respjson.Field
		Region       respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintListEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Geographic coordinates (null if unavailable)
type OsintListEventsResponseDataCoordinates struct {
	// Latitude
	Lat float64 `json:"lat"`
	// Longitude
	Lon float64 `json:"lon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListEventsResponseDataCoordinates) RawJSON() string { return r.JSON.raw }
func (r *OsintListEventsResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListEventsResponseMeta struct {
	Count   int64                              `json:"count"`
	Filters OsintListEventsResponseMetaFilters `json:"filters"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Filters     respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListEventsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintListEventsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListEventsResponseMetaFilters struct {
	Category string `json:"category" api:"nullable"`
	Severity string `json:"severity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListEventsResponseMetaFilters) RawJSON() string { return r.JSON.raw }
func (r *OsintListEventsResponseMetaFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListVesselsResponse struct {
	Data []OsintListVesselsResponseData `json:"data" api:"required"`
	Meta OsintListVesselsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListVesselsResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintListVesselsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListVesselsResponseData struct {
	// Vessel ID
	ID          string                                  `json:"id" api:"required"`
	Coordinates OsintListVesselsResponseDataCoordinates `json:"coordinates" api:"required"`
	// Vessel name
	Name string `json:"name" api:"required"`
	// Vessel type classification
	VesselType string `json:"vesselType" api:"required"`
	// Position confidence
	Confidence float64 `json:"confidence"`
	// Fetch time as Unix timestamp (milliseconds)
	FetchedAt int64 `json:"fetchedAt"`
	// Fetch time as ISO 8601 string
	FetchedAtISO time.Time `json:"fetchedAtISO" format:"date-time"`
	// Hull number designation
	HullNumber string `json:"hullNumber"`
	// Operating navy
	Operator string `json:"operator"`
	// Deployment region
	Region string `json:"region"`
	// Data source
	Source string `json:"source"`
	// Current operational status
	Status string `json:"status"`
	// Strike group assignment
	StrikeGroup string `json:"strikeGroup"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Coordinates  respjson.Field
		Name         respjson.Field
		VesselType   respjson.Field
		Confidence   respjson.Field
		FetchedAt    respjson.Field
		FetchedAtISO respjson.Field
		HullNumber   respjson.Field
		Operator     respjson.Field
		Region       respjson.Field
		Source       respjson.Field
		Status       respjson.Field
		StrikeGroup  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListVesselsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintListVesselsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListVesselsResponseDataCoordinates struct {
	Lat float64 `json:"lat" api:"required"`
	Lon float64 `json:"lon" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListVesselsResponseDataCoordinates) RawJSON() string { return r.JSON.raw }
func (r *OsintListVesselsResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListVesselsResponseMeta struct {
	Count   int64                               `json:"count"`
	Filters OsintListVesselsResponseMetaFilters `json:"filters"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Filters     respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListVesselsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintListVesselsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintListVesselsResponseMetaFilters struct {
	Region string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintListVesselsResponseMetaFilters) RawJSON() string { return r.JSON.raw }
func (r *OsintListVesselsResponseMetaFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintMapEventsResponse struct {
	Data []OsintMapEventsResponseData `json:"data" api:"required"`
	Meta OsintMapEventsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintMapEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *OsintMapEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintMapEventsResponseData struct {
	// Event ID
	ID string `json:"id" api:"required"`
	// OSINT event category classification
	//
	// Any of "seismic", "conflict", "political", "economic", "weather", "health",
	// "cyber", "maritime", "fire", "aviation", "other".
	Category string `json:"category" api:"required"`
	// Geographic coordinates (always present — events without coordinates are
	// excluded)
	Coordinates OsintMapEventsResponseDataCoordinates `json:"coordinates" api:"required"`
	// Event time as Unix timestamp (milliseconds)
	EventTime int64 `json:"eventTime" api:"required"`
	// Event severity level
	//
	// Any of "low", "medium", "high", "critical".
	Severity string `json:"severity" api:"required"`
	// Event title/headline
	Title string `json:"title" api:"required"`
	// ISO alpha-2 country code
	CountryCode string `json:"countryCode"`
	// Detailed event description
	Description string `json:"description"`
	// Event time as ISO 8601 string
	EventTimeISO time.Time `json:"eventTimeISO" format:"date-time"`
	// Human-readable location name
	LocationName string `json:"locationName"`
	// Geographic region identifier
	//
	// Any of "mena", "africa", "latam", "asiapac", "europe", "namerica".
	Region string `json:"region"`
	// Data source type
	SourceType string `json:"sourceType"`
	// Source URL for the event
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Category     respjson.Field
		Coordinates  respjson.Field
		EventTime    respjson.Field
		Severity     respjson.Field
		Title        respjson.Field
		CountryCode  respjson.Field
		Description  respjson.Field
		EventTimeISO respjson.Field
		LocationName respjson.Field
		Region       respjson.Field
		SourceType   respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintMapEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *OsintMapEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Geographic coordinates (always present — events without coordinates are
// excluded)
type OsintMapEventsResponseDataCoordinates struct {
	// Latitude
	Lat float64 `json:"lat" api:"required"`
	// Longitude
	Lon float64 `json:"lon" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintMapEventsResponseDataCoordinates) RawJSON() string { return r.JSON.raw }
func (r *OsintMapEventsResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintMapEventsResponseMeta struct {
	Count   int64                             `json:"count"`
	Filters OsintMapEventsResponseMetaFilters `json:"filters"`
	// Whether more results are available beyond the current limit
	HasMore bool  `json:"hasMore"`
	Limit   int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Filters     respjson.Field
		HasMore     respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintMapEventsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OsintMapEventsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintMapEventsResponseMetaFilters struct {
	Region string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OsintMapEventsResponseMetaFilters) RawJSON() string { return r.JSON.raw }
func (r *OsintMapEventsResponseMetaFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OsintGetConflictIndicatorsParams struct {
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by event category
	//
	// Any of "seismic", "conflict", "political", "economic", "weather", "health",
	// "cyber", "maritime", "fire", "aviation", "other".
	Category OsintGetConflictIndicatorsParamsCategory `query:"category,omitzero" json:"-"`
	// Filter by geographic region
	//
	// Any of "mena", "africa", "latam", "asiapac", "europe", "namerica".
	Region OsintGetConflictIndicatorsParamsRegion `query:"region,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintGetConflictIndicatorsParams]'s query parameters as
// `url.Values`.
func (r OsintGetConflictIndicatorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by event category
type OsintGetConflictIndicatorsParamsCategory string

const (
	OsintGetConflictIndicatorsParamsCategorySeismic   OsintGetConflictIndicatorsParamsCategory = "seismic"
	OsintGetConflictIndicatorsParamsCategoryConflict  OsintGetConflictIndicatorsParamsCategory = "conflict"
	OsintGetConflictIndicatorsParamsCategoryPolitical OsintGetConflictIndicatorsParamsCategory = "political"
	OsintGetConflictIndicatorsParamsCategoryEconomic  OsintGetConflictIndicatorsParamsCategory = "economic"
	OsintGetConflictIndicatorsParamsCategoryWeather   OsintGetConflictIndicatorsParamsCategory = "weather"
	OsintGetConflictIndicatorsParamsCategoryHealth    OsintGetConflictIndicatorsParamsCategory = "health"
	OsintGetConflictIndicatorsParamsCategoryCyber     OsintGetConflictIndicatorsParamsCategory = "cyber"
	OsintGetConflictIndicatorsParamsCategoryMaritime  OsintGetConflictIndicatorsParamsCategory = "maritime"
	OsintGetConflictIndicatorsParamsCategoryFire      OsintGetConflictIndicatorsParamsCategory = "fire"
	OsintGetConflictIndicatorsParamsCategoryAviation  OsintGetConflictIndicatorsParamsCategory = "aviation"
	OsintGetConflictIndicatorsParamsCategoryOther     OsintGetConflictIndicatorsParamsCategory = "other"
)

// Filter by geographic region
type OsintGetConflictIndicatorsParamsRegion string

const (
	OsintGetConflictIndicatorsParamsRegionMena     OsintGetConflictIndicatorsParamsRegion = "mena"
	OsintGetConflictIndicatorsParamsRegionAfrica   OsintGetConflictIndicatorsParamsRegion = "africa"
	OsintGetConflictIndicatorsParamsRegionLatam    OsintGetConflictIndicatorsParamsRegion = "latam"
	OsintGetConflictIndicatorsParamsRegionAsiapac  OsintGetConflictIndicatorsParamsRegion = "asiapac"
	OsintGetConflictIndicatorsParamsRegionEurope   OsintGetConflictIndicatorsParamsRegion = "europe"
	OsintGetConflictIndicatorsParamsRegionNamerica OsintGetConflictIndicatorsParamsRegion = "namerica"
)

type OsintGetGpsJammingZonesParams struct {
	// Maximum number of zones to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by interference severity
	//
	// Any of "low", "moderate", "severe", "critical".
	Severity OsintGetGpsJammingZonesParamsSeverity `query:"severity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintGetGpsJammingZonesParams]'s query parameters as
// `url.Values`.
func (r OsintGetGpsJammingZonesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by interference severity
type OsintGetGpsJammingZonesParamsSeverity string

const (
	OsintGetGpsJammingZonesParamsSeverityLow      OsintGetGpsJammingZonesParamsSeverity = "low"
	OsintGetGpsJammingZonesParamsSeverityModerate OsintGetGpsJammingZonesParamsSeverity = "moderate"
	OsintGetGpsJammingZonesParamsSeveritySevere   OsintGetGpsJammingZonesParamsSeverity = "severe"
	OsintGetGpsJammingZonesParamsSeverityCritical OsintGetGpsJammingZonesParamsSeverity = "critical"
)

type OsintGetMilitaryPostureParams struct {
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintGetMilitaryPostureParams]'s query parameters as
// `url.Values`.
func (r OsintGetMilitaryPostureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OsintListAircraftParams struct {
	// Maximum number of aircraft to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by theater ID (e.g. "iran", "taiwan", "baltic")
	Theater param.Opt[string] `query:"theater,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintListAircraftParams]'s query parameters as
// `url.Values`.
func (r OsintListAircraftParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OsintListEventsParams struct {
	// Maximum number of events to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by event category
	//
	// Any of "seismic", "conflict", "political", "economic", "weather", "health",
	// "cyber", "maritime", "fire", "aviation", "other".
	Category OsintListEventsParamsCategory `query:"category,omitzero" json:"-"`
	// Filter by severity level
	//
	// Any of "low", "medium", "high", "critical".
	Severity OsintListEventsParamsSeverity `query:"severity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintListEventsParams]'s query parameters as `url.Values`.
func (r OsintListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by event category
type OsintListEventsParamsCategory string

const (
	OsintListEventsParamsCategorySeismic   OsintListEventsParamsCategory = "seismic"
	OsintListEventsParamsCategoryConflict  OsintListEventsParamsCategory = "conflict"
	OsintListEventsParamsCategoryPolitical OsintListEventsParamsCategory = "political"
	OsintListEventsParamsCategoryEconomic  OsintListEventsParamsCategory = "economic"
	OsintListEventsParamsCategoryWeather   OsintListEventsParamsCategory = "weather"
	OsintListEventsParamsCategoryHealth    OsintListEventsParamsCategory = "health"
	OsintListEventsParamsCategoryCyber     OsintListEventsParamsCategory = "cyber"
	OsintListEventsParamsCategoryMaritime  OsintListEventsParamsCategory = "maritime"
	OsintListEventsParamsCategoryFire      OsintListEventsParamsCategory = "fire"
	OsintListEventsParamsCategoryAviation  OsintListEventsParamsCategory = "aviation"
	OsintListEventsParamsCategoryOther     OsintListEventsParamsCategory = "other"
)

// Filter by severity level
type OsintListEventsParamsSeverity string

const (
	OsintListEventsParamsSeverityLow      OsintListEventsParamsSeverity = "low"
	OsintListEventsParamsSeverityMedium   OsintListEventsParamsSeverity = "medium"
	OsintListEventsParamsSeverityHigh     OsintListEventsParamsSeverity = "high"
	OsintListEventsParamsSeverityCritical OsintListEventsParamsSeverity = "critical"
)

type OsintListVesselsParams struct {
	// Maximum number of vessels to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by region name
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintListVesselsParams]'s query parameters as `url.Values`.
func (r OsintListVesselsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OsintMapEventsParams struct {
	// Maximum number of events to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by geographic region
	//
	// Any of "mena", "africa", "latam", "asiapac", "europe", "namerica".
	Region OsintMapEventsParamsRegion `query:"region,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OsintMapEventsParams]'s query parameters as `url.Values`.
func (r OsintMapEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by geographic region
type OsintMapEventsParamsRegion string

const (
	OsintMapEventsParamsRegionMena     OsintMapEventsParamsRegion = "mena"
	OsintMapEventsParamsRegionAfrica   OsintMapEventsParamsRegion = "africa"
	OsintMapEventsParamsRegionLatam    OsintMapEventsParamsRegion = "latam"
	OsintMapEventsParamsRegionAsiapac  OsintMapEventsParamsRegion = "asiapac"
	OsintMapEventsParamsRegionEurope   OsintMapEventsParamsRegion = "europe"
	OsintMapEventsParamsRegionNamerica OsintMapEventsParamsRegion = "namerica"
)
