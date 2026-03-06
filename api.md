# Reports

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#AudioMetadata">AudioMetadata</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportGetResponse">ReportGetResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportListResponse">ReportListResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportGetAudioResponse">ReportGetAudioResponse</a>

Methods:

- <code title="get /reports/{reportId}">client.Reports.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportGetResponse">ReportGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /reports">client.Reports.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportListParams">ReportListParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportListResponse">ReportListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /reports/{reportId}/audio">client.Reports.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportService.GetAudio">GetAudio</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportGetAudioParams">ReportGetAudioParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ReportGetAudioResponse">ReportGetAudioResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileNewResponse">ProfileNewResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileUpdateResponse">ProfileUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileListResponse">ProfileListResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileDeleteResponse">ProfileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfilePartialUpdateResponse">ProfilePartialUpdateResponse</a>

Methods:

- <code title="post /profiles">client.Profiles.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileNewParams">ProfileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileNewResponse">ProfileNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /profiles/{profileId}">client.Profiles.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileUpdateParams">ProfileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileUpdateResponse">ProfileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /profiles">client.Profiles.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileListResponse">ProfileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /profiles/{profileId}">client.Profiles.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileDeleteResponse">ProfileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /profiles/{profileId}">client.Profiles.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfileService.PartialUpdate">PartialUpdate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfilePartialUpdateParams">ProfilePartialUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#ProfilePartialUpdateResponse">ProfilePartialUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# News

Params Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#TimeframeEnum">TimeframeEnum</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#TimeframeEnum">TimeframeEnum</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#TopicEnum">TopicEnum</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsListResponse">NewsListResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsGetRecapsResponse">NewsGetRecapsResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsListFeedsResponse">NewsListFeedsResponse</a>

Methods:

- <code title="get /news">client.News.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsListParams">NewsListParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsListResponse">NewsListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /news/recaps">client.News.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsService.GetRecaps">GetRecaps</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsGetRecapsParams">NewsGetRecapsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsGetRecapsResponse">NewsGetRecapsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /news/feeds">client.News.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsService.ListFeeds">ListFeeds</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#NewsListFeedsResponse">NewsListFeedsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookUpdateResponse">WebhookUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookDeleteResponse">WebhookDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookTestResponse">WebhookTestResponse</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookUpdateResponse">WebhookUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookDeleteResponse">WebhookDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/{webhookId}/test">client.Webhooks.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookService.Test">Test</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#WebhookTestResponse">WebhookTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#SubscriptionUpdateDeliveryResponse">SubscriptionUpdateDeliveryResponse</a>

Methods:

- <code title="patch /subscriptions/{subscriptionId}/delivery">client.Subscriptions.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#SubscriptionService.UpdateDelivery">UpdateDelivery</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#SubscriptionUpdateDeliveryParams">SubscriptionUpdateDeliveryParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#SubscriptionUpdateDeliveryResponse">SubscriptionUpdateDeliveryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Osint

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetConflictIndicatorsResponse">OsintGetConflictIndicatorsResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetGpsJammingZonesResponse">OsintGetGpsJammingZonesResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetMilitaryPostureResponse">OsintGetMilitaryPostureResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListAircraftResponse">OsintListAircraftResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListEventsResponse">OsintListEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListVesselsResponse">OsintListVesselsResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintMapEventsResponse">OsintMapEventsResponse</a>

Methods:

- <code title="get /osint/cii">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.GetConflictIndicators">GetConflictIndicators</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetConflictIndicatorsParams">OsintGetConflictIndicatorsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetConflictIndicatorsResponse">OsintGetConflictIndicatorsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/gps-jamming">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.GetGpsJammingZones">GetGpsJammingZones</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetGpsJammingZonesParams">OsintGetGpsJammingZonesParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetGpsJammingZonesResponse">OsintGetGpsJammingZonesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/military-posture">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.GetMilitaryPosture">GetMilitaryPosture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetMilitaryPostureParams">OsintGetMilitaryPostureParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintGetMilitaryPostureResponse">OsintGetMilitaryPostureResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/aircraft">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.ListAircraft">ListAircraft</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListAircraftParams">OsintListAircraftParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListAircraftResponse">OsintListAircraftResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/events">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.ListEvents">ListEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListEventsParams">OsintListEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListEventsResponse">OsintListEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/vessels">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.ListVessels">ListVessels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListVesselsParams">OsintListVesselsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintListVesselsResponse">OsintListVesselsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/map">client.Osint.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintService.MapEvents">MapEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintMapEventsParams">OsintMapEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintMapEventsResponse">OsintMapEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Countries

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetCountryInstabilityIndexResponse">OsintCountryGetCountryInstabilityIndexResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetCountryNewsResponse">OsintCountryGetCountryNewsResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetIntelligenceBriefResponse">OsintCountryGetIntelligenceBriefResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetPredictionMarketsResponse">OsintCountryGetPredictionMarketsResponse</a>
- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetStockMarketIndexResponse">OsintCountryGetStockMarketIndexResponse</a>

Methods:

- <code title="get /osint/countries/{countryCode}/cii">client.Osint.Countries.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryService.GetCountryInstabilityIndex">GetCountryInstabilityIndex</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, countryCode <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetCountryInstabilityIndexResponse">OsintCountryGetCountryInstabilityIndexResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/countries/{countryCode}/news">client.Osint.Countries.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryService.GetCountryNews">GetCountryNews</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, countryCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetCountryNewsParams">OsintCountryGetCountryNewsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetCountryNewsResponse">OsintCountryGetCountryNewsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/countries/{countryCode}/brief">client.Osint.Countries.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryService.GetIntelligenceBrief">GetIntelligenceBrief</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, countryCode <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetIntelligenceBriefResponse">OsintCountryGetIntelligenceBriefResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/countries/{countryCode}/predictions">client.Osint.Countries.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryService.GetPredictionMarkets">GetPredictionMarkets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, countryCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetPredictionMarketsParams">OsintCountryGetPredictionMarketsParams</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetPredictionMarketsResponse">OsintCountryGetPredictionMarketsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /osint/countries/{countryCode}/markets">client.Osint.Countries.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryService.GetStockMarketIndex">GetStockMarketIndex</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, countryCode <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintCountryGetStockMarketIndexResponse">OsintCountryGetStockMarketIndexResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sources

Response Types:

- <a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintSourceGetDataSourceHealthResponse">OsintSourceGetDataSourceHealthResponse</a>

Methods:

- <code title="get /osint/sources/status">client.Osint.Sources.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintSourceService.GetDataSourceHealth">GetDataSourceHealth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/y2-intel/y2-go">y2</a>.<a href="https://pkg.go.dev/github.com/y2-intel/y2-go#OsintSourceGetDataSourceHealthResponse">OsintSourceGetDataSourceHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
