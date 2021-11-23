//////////////////////////////////////////////////////////////////////
// task_post_client.go
//////////////////////////////////////////////////////////////////////
package standard

import (
    "encoding/json"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/dataforseo/result/v3/serp/google/organic/standard"
)

type TaskPostClient struct {
    *myAuthentication.Authentication
    EndpointUrl string
}

type TaskPostData struct {
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
// New TaskPostClient object.
//////////////////////////////////////////////////////////////////////
func NewTaskPostClient(login, password string, isSandbox bool) *TaskPostClient {
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3_SANDBOX
    }
    return &TaskPostClient{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// New TaskPostData.
//////////////////////////////////////////////////////////////////////
func NewTaskPostData(keyword, postbackData string, langCode string, locationCode int) *TaskPostData {
    return &TaskPostData{
        Keyword: keyword,
        PostbackData: postbackData,
        LanguageCode: langCode,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *TaskPostClient) Run(data []*TaskPostData) (int, *myResult.TaskPostResults, error) {
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
    var result *myResult.TaskPostResults
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}


//////////////////////////////////////////////////////////////////////
// Set "url".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetUrl(url string) *TaskPostData {
    o.Url = url
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "priority".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetPriority(priority int) *TaskPostData {
    o.Priority = priority
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "depth".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetDepth(depth int) *TaskPostData {
    o.Depth = depth
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "location_name".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetLocationName(locationName string) *TaskPostData {
    o.LocationName = locationName
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "location_coordinate".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetLocationCoordinate(locationCoordinate string) *TaskPostData {
    o.LocationCoordinate = locationCoordinate
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "language_name".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetLanguageName(languageName string) *TaskPostData {
    o.LanguageName = languageName
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "se_domain".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetSeDomain(seDomain string) *TaskPostData {
    o.SeDomain = seDomain
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "device".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetDevice(device string) *TaskPostData {
    o.Device = device
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "os".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetOs(os string) *TaskPostData {
    o.Os = os
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "calculate_rectangles".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetCalculateRectangles(calculateRectangles bool) *TaskPostData {
    o.CalculateRectangles = calculateRectangles
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_width".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetBrowserScreenWidth(browserScreenWidth int) *TaskPostData {
    o.BrowserScreenWidth = browserScreenWidth
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_height".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetBrowserScreenHeight(browserScreenHeight int) *TaskPostData {
    o.BrowserScreenHeight = browserScreenHeight
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_resolution_ratio".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetBrowserScreenResolutionRatio(browserScreenResolutionRatio int) *TaskPostData {
    o.BrowserScreenResolutionRatio = browserScreenResolutionRatio
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "search_param".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetSearchParam(searchParam string) *TaskPostData {
    o.SearchParam = searchParam
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetTag(tag string) *TaskPostData {
    o.Tag = tag
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "postback_url".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetPostbackUrl(postbackUrl string) *TaskPostData {
    o.PostbackUrl = postbackUrl
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "pingback_url".
//////////////////////////////////////////////////////////////////////
func (o *TaskPostData) SetPingbackUrl(pingbackUrl string) *TaskPostData {
    o.PingbackUrl = pingbackUrl
    return o
}
