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

// NewsService contains methods and other services that help with interacting with
// the y2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNewsService] method instead.
type NewsService struct {
	Options []option.RequestOption
}

// NewNewsService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNewsService(opts ...option.RequestOption) (r NewsService) {
	r = NewsService{}
	r.Options = opts
	return
}

// Returns news items from the GloriaAI terminal cache. Supports filtering by
// topics and pagination.
func (r *NewsService) List(ctx context.Context, query NewsListParams, opts ...option.RequestOption) (res *NewsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "news"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns AI-generated recap summaries for specified topics within a given
// timeframe.
func (r *NewsService) GetRecaps(ctx context.Context, query NewsGetRecapsParams, opts ...option.RequestOption) (res *NewsGetRecapsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "news/recaps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns all available news feed topics with descriptions.
func (r *NewsService) ListFeeds(ctx context.Context, opts ...option.RequestOption) (res *NewsListFeedsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "news/feeds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Time period for recap data
type TimeframeEnum string

const (
	TimeframeEnum12h TimeframeEnum = "12h"
	TimeframeEnum24h TimeframeEnum = "24h"
	TimeframeEnum3d  TimeframeEnum = "3d"
	TimeframeEnum7d  TimeframeEnum = "7d"
)

// Available news feed topics from GloriaAI
type TopicEnum string

const (
	TopicEnumAI              TopicEnum = "ai"
	TopicEnumAIAgents        TopicEnum = "ai_agents"
	TopicEnumAptos           TopicEnum = "aptos"
	TopicEnumBase            TopicEnum = "base"
	TopicEnumBitcoin         TopicEnum = "bitcoin"
	TopicEnumCrypto          TopicEnum = "crypto"
	TopicEnumDats            TopicEnum = "dats"
	TopicEnumDefi            TopicEnum = "defi"
	TopicEnumEthereum        TopicEnum = "ethereum"
	TopicEnumHyperliquid     TopicEnum = "hyperliquid"
	TopicEnumMachineLearning TopicEnum = "machine_learning"
	TopicEnumMacro           TopicEnum = "macro"
	TopicEnumOndo            TopicEnum = "ondo"
	TopicEnumPerps           TopicEnum = "perps"
	TopicEnumRipple          TopicEnum = "ripple"
	TopicEnumRwa             TopicEnum = "rwa"
	TopicEnumSolana          TopicEnum = "solana"
	TopicEnumTech            TopicEnum = "tech"
	TopicEnumVirtuals        TopicEnum = "virtuals"
)

type NewsListResponse struct {
	Data []NewsListResponseData `json:"data,required"`
	Meta NewsListResponseMeta   `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsListResponse) RawJSON() string { return r.JSON.raw }
func (r *NewsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsListResponseData struct {
	ID string `json:"id,required"`
	// Primary signal/headline
	Signal string `json:"signal,required"`
	// Unix timestamp (seconds)
	Timestamp    int64     `json:"timestamp,required"`
	TimestampISO time.Time `json:"timestampISO,required" format:"date-time"`
	Author       string    `json:"author"`
	Categories   []string  `json:"categories"`
	// Full context
	Content     string `json:"content"`
	NarrativeID string `json:"narrativeId"`
	// Sentiment classification for news items
	//
	// Any of "bullish", "bearish", "neutral".
	Sentiment      string   `json:"sentiment,nullable"`
	SentimentValue float64  `json:"sentimentValue"`
	Sources        []string `json:"sources"`
	// Short context summary
	Summary string `json:"summary"`
	// Related tokens/assets
	Tokens   []string `json:"tokens"`
	TweetURL string   `json:"tweetUrl" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Signal         respjson.Field
		Timestamp      respjson.Field
		TimestampISO   respjson.Field
		Author         respjson.Field
		Categories     respjson.Field
		Content        respjson.Field
		NarrativeID    respjson.Field
		Sentiment      respjson.Field
		SentimentValue respjson.Field
		Sources        respjson.Field
		Summary        respjson.Field
		Tokens         respjson.Field
		TweetURL       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NewsListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsListResponseMeta struct {
	Count  int64       `json:"count"`
	Limit  int64       `json:"limit"`
	Topics []TopicEnum `json:"topics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		Topics      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *NewsListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsGetRecapsResponse struct {
	Data map[string]any            `json:"data,required"`
	Meta NewsGetRecapsResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsGetRecapsResponse) RawJSON() string { return r.JSON.raw }
func (r *NewsGetRecapsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsGetRecapsResponseMeta struct {
	// Time period for recap data
	//
	// Any of "12h", "24h", "3d", "7d".
	Timeframe TimeframeEnum `json:"timeframe"`
	Topics    []TopicEnum   `json:"topics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timeframe   respjson.Field
		Topics      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsGetRecapsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *NewsGetRecapsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsListFeedsResponse struct {
	Data []NewsListFeedsResponseData `json:"data,required"`
	Meta NewsListFeedsResponseMeta   `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsListFeedsResponse) RawJSON() string { return r.JSON.raw }
func (r *NewsListFeedsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsListFeedsResponseData struct {
	// Available news feed topics from GloriaAI
	//
	// Any of "ai", "ai_agents", "aptos", "base", "bitcoin", "crypto", "dats", "defi",
	// "ethereum", "hyperliquid", "machine_learning", "macro", "ondo", "perps",
	// "ripple", "rwa", "solana", "tech", "virtuals".
	ID TopicEnum `json:"id,required"`
	// Human-readable name
	Name string `json:"name,required"`
	// Feed description
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsListFeedsResponseData) RawJSON() string { return r.JSON.raw }
func (r *NewsListFeedsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsListFeedsResponseMeta struct {
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsListFeedsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *NewsListFeedsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsListParams struct {
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma-separated list of topics to filter by. Valid topics: ai, ai_agents, aptos,
	// base, bitcoin, crypto, dats, defi, ethereum, hyperliquid, machine_learning,
	// macro, ondo, perps, ripple, rwa, solana, tech, virtuals. Default: crypto,
	// ai_agents, macro, bitcoin, ethereum, tech
	Topics param.Opt[string] `query:"topics,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NewsListParams]'s query parameters as `url.Values`.
func (r NewsListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NewsGetRecapsParams struct {
	// Comma-separated list of topics. Valid topics: ai, ai_agents, aptos, base,
	// bitcoin, crypto, dats, defi, ethereum, hyperliquid, machine_learning, macro,
	// ondo, perps, ripple, rwa, solana, tech, virtuals
	Topics param.Opt[string] `query:"topics,omitzero" json:"-"`
	// Time period for recaps
	//
	// Any of "12h", "24h", "3d", "7d".
	Timeframe TimeframeEnum `query:"timeframe,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NewsGetRecapsParams]'s query parameters as `url.Values`.
func (r NewsGetRecapsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
