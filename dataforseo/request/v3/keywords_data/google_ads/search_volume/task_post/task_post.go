//////////////////////////////////////////////////////////////////////
// task_post.go
//////////////////////////////////////////////////////////////////////
package task_post

import (
    "encoding/json"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/dataforseo/response/v3/keywords_data/google_ads/search_volume/task_post"
)

type Client struct {
    *myAuthentication.Authentication
    EndpointUrl string
}

type Data struct {
    DateFrom string             `json:"date_from,omitempty"`
    DateTo string               `json:"date_to,omitempty"`
    Keywords []string           `json:"keywords,omitempty"`
    LanguageName string         `json:"language_name,omitempty"`
    LanguageCode string         `json:"language_code,omitempty"`
    LocationName string         `json:"location_name,omitempty"`
    LocationCode int            `json:"location_code,omitempty"`
    LocationCoordinate string   `json:"location_coordinate,omitempty"`
    PingbackUrl string          `json:"pingback_url,omitempty"`
    PostbackUrl string          `json:"postback_url,omitempty"`
    SearchPartners bool         `json:"search_partners,omitempty"`
    SortBy string               `json:"sort_by,omitempty"`
    Tag string                  `json:"tag,omitempty"`
}


//////////////////////////////////////////////////////////////////////
// New Client object.
//////////////////////////////////////////////////////////////////////
func NewClient(login, password string, isSandbox bool) *Client {
    endpointUrl := myConstant.ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3_SANDBOX
    }
    return &Client{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "location_name".
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationName(keywords []string, locationName string) *Data {
    return &Data{
        Keywords: keywords,
        LocationName: locationName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "location_code".
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationCode(keywords []string, locationCode int) *Data {
    return &Data{
        Keywords: keywords,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "location_coordinate".
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationCoordinate(keywords []string, locationCoordinate string) *Data {
    return &Data{
        Keywords: keywords,
        LocationCoordinate: locationCoordinate,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run(datas []*Data) (int, *myResult.Response, error) {
    jsonData, err := json.Marshal(datas)
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
    var result *myResult.Response
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}


//////////////////////////////////////////////////////////////////////
// Set "date_from".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetDateFrom(dateFrom string) *Data {
    d.DateFrom = dateFrom
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "date_to".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetDateTo(dateTo string) *Data {
    d.DateTo = dateTo
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "language_code".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetLanguageCode(languageCode string) *Data {
    d.LanguageCode = languageCode
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "language_name".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetLanguageName(languageName string) *Data {
    d.LanguageName = languageName
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "pingback_url".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetPingbackUrl(pingbackUrl string) *Data {
    d.PingbackUrl = pingbackUrl
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "postback_url".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetPostbackUrl(postbackUrl string) *Data {
    d.PostbackUrl = postbackUrl
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "search_partners".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetSearchPartners(searchPartners bool) *Data {
    d.SearchPartners = searchPartners
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "sort_by".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetSortBy(sortBy string) *Data {
    d.SortBy = sortBy
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetTag(tag string) *Data {
    d.Tag = tag
    return d
}


