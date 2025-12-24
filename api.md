# Reports

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#AudioMetadata">AudioMetadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportGetResponse">ReportGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportListResponse">ReportListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportGetAudioResponse">ReportGetAudioResponse</a>

Methods:

- <code title="get /reports/{reportId}">client.Reports.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportGetResponse">ReportGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /reports">client.Reports.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportListParams">ReportListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportListResponse">ReportListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /reports/{reportId}/audio">client.Reports.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportService.GetAudio">GetAudio</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportGetAudioParams">ReportGetAudioParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ReportGetAudioResponse">ReportGetAudioResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ProfileListResponse">ProfileListResponse</a>

Methods:

- <code title="get /profiles">client.Profiles.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#ProfileListResponse">ProfileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# News

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#TimeframeEnum">TimeframeEnum</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#TimeframeEnum">TimeframeEnum</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#TopicEnum">TopicEnum</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsListResponse">NewsListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsGetRecapsResponse">NewsGetRecapsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsListFeedsResponse">NewsListFeedsResponse</a>

Methods:

- <code title="get /news">client.News.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsListParams">NewsListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsListResponse">NewsListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /news/recaps">client.News.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsService.GetRecaps">GetRecaps</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsGetRecapsParams">NewsGetRecapsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsGetRecapsResponse">NewsGetRecapsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /news/feeds">client.News.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsService.ListFeeds">ListFeeds</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/y2-go#NewsListFeedsResponse">NewsListFeedsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
