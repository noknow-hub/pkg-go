//////////////////////////////////////////////////////////////////////
// organic_task_post_client.go
//////////////////////////////////////////////////////////////////////
package google

import (
    "encoding/json"
    myConstant "github.com/noknow-hub/pkg-go/vendor/dataforseo/constant"
    myAuthentication "github.com/noknow-hub/pkg-go/vendor/dataforseo/client/authentication"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/vendor/dataforseo/result/v3/serp/google"
)

type OrganicTaskPostClient struct {
    *myAuthentication.Authentication,
    EndpointUrl string
}

type OrganicTaskPostData struct {
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
// New OrganicTaskPostClient object.
//////////////////////////////////////////////////////////////////////
func NewOrganicTaskPostClient(login, password string, isSandbox bool) *OrganicTaskPostClient {
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3_SANDBOX
    }
    return &OrganicTaskPostClient{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// New OrganicTaskPostData.
//////////////////////////////////////////////////////////////////////
func NewOrganicTaskPostData(keyword, postbackData string, langCode int, locationCode string) *OrganicTaskPostData {
    return &OrganicTaskPostData{
        Keyword: keyword,
        PostbackData: postbackData,
        LanguageCode: langCode,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// Do.
//////////////////////////////////////////////////////////////////////
func (c *OrganicTaskPostClient) Do(data []*OrganicTaskPostData) (int, *myResult.Result, error) {
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
    var result *myResult.Result
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(respBody)

    return statusCode, result, nil
}


//////////////////////////////////////////////////////////////////////
// Set "url".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetUrl(url string) *OrganicTaskPostData {
    o.Url = url
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "priority".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetPriority(priority int) *OrganicTaskPostData {
    o.Priority = priority
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "depth".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetDepth(depth int) *OrganicTaskPostData {
    o.Depth = depth
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "location_name".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetLocationName(locationName string) *OrganicTaskPostData {
    o.LocationName = locationName
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "location_coordinate".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetLocationCoordinate(locationCoordinate string) *OrganicTaskPostData {
    o.LocationCoordinate = locationCoordinate
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "language_name".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetLanguageName(languageName string) *OrganicTaskPostData {
    o.LanguageName = languageName
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "se_domain".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetSeDomain(seDomain string) *OrganicTaskPostData {
    o.SeDomain = seDomain
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "device".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetDevice(device string) *OrganicTaskPostData {
    o.Device = device
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "os".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetOs(os string) *OrganicTaskPostData {
    o.Os = os
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "calculate_rectangles".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetCalculateRectangles(calculateRectangles bool) *OrganicTaskPostData {
    o.CalculateRectangles = calculateRectangles
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_width".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetBrowserScreenWidth(browserScreenWidth int) *OrganicTaskPostData {
    o.BrowserScreenWidth = browserScreenWidth
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_height".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetBrowserScreenHeight(browserScreenHeight int) *OrganicTaskPostData {
    o.BrowserScreenHeight = browserScreenHeight
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "browser_screen_resolution_ratio".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetBrowserScreenResolutionRatio(browserScreenResolutionRatio int) *OrganicTaskPostData {
    o.BrowserScreenResolutionRatio = browserScreenResolutionRatio
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "search_param".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetSearchParam(searchParam string) *OrganicTaskPostData {
    o.SearchParam = searchParam
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetTag(tag string) *OrganicTaskPostData {
    o.Tag = tag
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "postback_url".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetPostbackUrl(postbackUrl string) *OrganicTaskPostData {
    o.PostbackUrl = postbackUrl
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "pingback_url".
//////////////////////////////////////////////////////////////////////
func (o *OrganicTaskPostData) SetPingbackUrl(pingbackUrl string) *OrganicTaskPostData {
    o.PingbackUrl = pingbackUrl
    return o
}
