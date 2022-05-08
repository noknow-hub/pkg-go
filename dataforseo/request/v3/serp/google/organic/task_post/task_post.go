//////////////////////////////////////////////////////////////////////
// task_post.go
//////////////////////////////////////////////////////////////////////
package task_post

import (
    "encoding/json"
    "time"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResponse "github.com/noknow-hub/pkg-go/dataforseo/response/v3/serp/google/organic/task_post"
)

var (
    NumOfApiCalls = 0
    LastCalledAt time.Time
    ApiCallQueues []time.Time
)

type Client struct {
    *myAuthentication.Authentication
    EndpointUrl string
}

type Data struct {
    Keyword string                    `json:"keyword,omitempty"`
    Url string                        `json:"url,omitempty"`
    Priority int                      `json:"priority,omitempty"`
    Depth int                         `json:"depth,omitempty"`
    LocationName string               `json:"location_name,omitempty"`
    LocationCode int                  `json:"location_code,omitempty"`
    LocationCoordinate string         `json:"location_coordinate,omitempty"`
    LanguageName string               `json:"language_name,omitempty"`
    LanguageCode string               `json:"language_code,omitempty"`
    SeDomain string                   `json:"se_domain,omitempty"`
    Device string                     `json:"device,omitempty"`
    Os string                         `json:"os,omitempty"`
    CalculateRectangles bool          `json:"calculate_rectangles"`
    BrowserScreenWidth int            `json:"browser_screen_width,omitempty"`
    BrowserScreenHeight int           `json:"browser_screen_height,omitempty"`
    BrowserScreenResolutionRatio int  `json:"browser_screen_resolution_ratio,omitempty"`
    SearchParam string                `json:"search_param,omitempty"`
    Tag string                        `json:"tag,omitempty"`
    PostbackUrl string                `json:"postback_url,omitempty"`
    PostbackData string               `json:"postback_data,omitempty"`
    PingbackUrl string                `json:"pingback_url,omitempty"`
}


//////////////////////////////////////////////////////////////////////
// Clear NumOfApiCalls
//////////////////////////////////////////////////////////////////////
func ClearNumOfApiCalls() {
    NumOfApiCalls = 0
}


//////////////////////////////////////////////////////////////////////
// Is over the number of tasks
//////////////////////////////////////////////////////////////////////
func IsOverNumOfTasks(datas []*Data) bool {
    return len(datas) > myConstant.LIMIT_NUM_OF_TASKS_PER_REQ_FOR_SERP_GOOGLE_ORGANIC_TASK_POST_V3
}

//////////////////////////////////////////////////////////////////////
// New Client
//////////////////////////////////////////////////////////////////////
func NewClient(login, password string, isSandbox bool) *Client {
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3_SANDBOX
    }
    return &Client{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data.
//////////////////////////////////////////////////////////////////////
func NewData(keyword, postbackData string, langCode string, locationCode int) *Data {
    return &Data{
        Keyword: keyword,
        PostbackData: postbackData,
        LanguageCode: langCode,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run(data []*Data) (int, *myResponse.Response, error) {
    minAgo := time.Now().Add(time.Minute * -1)
    for i, ApiCallQueue := range ApiCallQueues {
        if ApiCallQueue.After(minAgo) {
            ApiCallQueues = ApiCallQueues[i:]
            break
        }
    }

    if len(ApiCallQueues) > myConstant.LIMIT_NUM_OF_CALLS_PER_MIN_FOR_SERP_GOOGLE_ORGANIC_TASK_POST_V3 {
        time.Sleep(60 * time.Second)
        ApiCallQueues = nil
    }
    ApiCallQueues = append(ApiCallQueues, time.Now())

    jsonData, err := json.Marshal(data)
    if err != nil {
        return 0, nil, err
    }

    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.
        AddHeaderContentType(myHttpClient.CONTENT_TYPE_JSON).
        AddHeaderAuthorization("Basic " + c.Cred).
        SetJsonData(jsonData)
    resp, err := httpClient.Post()
    if err != nil {
        return 0, nil, err
    }
    var result *myResponse.Response
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}


//////////////////////////////////////////////////////////////////////
// Set "url".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetUrl(url string) *Data {
    o.Url = url
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "priority".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetPriority(priority int) *Data {
    o.Priority = priority
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "depth".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetDepth(depth int) *Data {
    o.Depth = depth
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "location_name".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetLocationName(locationName string) *Data {
    o.LocationName = locationName
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "location_coordinate".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetLocationCoordinate(locationCoordinate string) *Data {
    o.LocationCoordinate = locationCoordinate
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "language_name".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetLanguageName(languageName string) *Data {
    o.LanguageName = languageName
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "se_domain".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetSeDomain(seDomain string) *Data {
    o.SeDomain = seDomain
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "device".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetDevice(device string) *Data {
    o.Device = device
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "os".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetOs(os string) *Data {
    o.Os = os
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "calculate_rectangles".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetCalculateRectangles(calculateRectangles bool) *Data {
    o.CalculateRectangles = calculateRectangles
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_width".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetBrowserScreenWidth(browserScreenWidth int) *Data {
    o.BrowserScreenWidth = browserScreenWidth
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_height".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetBrowserScreenHeight(browserScreenHeight int) *Data {
    o.BrowserScreenHeight = browserScreenHeight
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_resolution_ratio".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetBrowserScreenResolutionRatio(browserScreenResolutionRatio int) *Data {
    o.BrowserScreenResolutionRatio = browserScreenResolutionRatio
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "search_param".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetSearchParam(searchParam string) *Data {
    o.SearchParam = searchParam
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetTag(tag string) *Data {
    o.Tag = tag
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "postback_url".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetPostbackUrl(postbackUrl string) *Data {
    o.PostbackUrl = postbackUrl
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "pingback_url".
//////////////////////////////////////////////////////////////////////
func (o *Data) SetPingbackUrl(pingbackUrl string) *Data {
    o.PingbackUrl = pingbackUrl
    return o
}
