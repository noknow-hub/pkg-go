//////////////////////////////////////////////////////////////////////
// task_post_client.go
//////////////////////////////////////////////////////////////////////
package search_volume

import (
    "encoding/json"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/dataforseo/result/v3/keywords_data/google_ads/search_volume"
)

type TaskPostClient struct {
    *myAuthentication.Authentication
    EndpointUrl string
}

type TaskPostData struct {
    DateFrom string              `json:"date_from,omitempty"`
    DateTo string                `json:"date_to,omitempty"`
    Keywords []string            `json:"keywords,omitempty"`
    LanguageName string          `json:"language_name,omitempty"`
    LanguageCode string          `json:"language_code,omitempty"`
    LocationName string          `json:"location_name,omitempty"`
    LocationCode int             `json:"location_code,omitempty"`
    LocationCoordinate string    `json:"location_coordinate,omitempty"`
    PingbackUrl string           `json:"pingback_url,omitempty"`
    PostbackUrl string           `json:"postback_url,omitempty"`
    SearchPartners bool          `json:"search_partners,omitempty"`
    SortBy string                `json:"sort_by,omitempty"`
    Tag string                   `json:"tag,omitempty"`
}


//////////////////////////////////////////////////////////////////////
// New TaskPostClient object.
//////////////////////////////////////////////////////////////////////
func NewTaskPostClient(login, password string, isSandbox bool) *TaskPostClient {
    endpointUrl := myConstant.ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3_SANDBOX
    }
    return &TaskPostClient{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// New TaskPostData with "location_name".
//////////////////////////////////////////////////////////////////////
func NewTaskPostDataWithLocationName(keywords []string, locationName string) *TaskPostData {
    return &TaskPostData{
        Keywords: keywords,
        LocationName: locationName,
    }
}


//////////////////////////////////////////////////////////////////////
// New TaskPostData with "location_code".
//////////////////////////////////////////////////////////////////////
func NewTaskPostDataWithLocationCode(keywords []string, locationCode int) *TaskPostData {
    return &TaskPostData{
        Keywords: keywords,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// New TaskPostData with "location_coordinate".
//////////////////////////////////////////////////////////////////////
func NewTaskPostDataWithLocationCoordinate(keywords []string, locationCoordinate string) *TaskPostData {
    return &TaskPostData{
        Keywords: keywords,
        LocationCoordinate: locationCoordinate,
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
// Set "date_from".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetDateFrom(dateFrom string) *TaskPostData {
    d.DateFrom = dateFrom
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "date_to".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetDateTo(dateTo string) *TaskPostData {
    d.DateTo = dateTo
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "language_code".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetLanguageCode(languageCode string) *TaskPostData {
    d.LanguageCode = languageCode
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "language_name".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetLanguageName(languageName string) *TaskPostData {
    d.LanguageName = languageName
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "pingback_url".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetPingbackUrl(pingbackUrl string) *TaskPostData {
    d.PingbackUrl = pingbackUrl
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "postback_url".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetPostbackUrl(postbackUrl string) *TaskPostData {
    d.PostbackUrl = postbackUrl
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "search_partners".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetSearchPartners(searchPartners bool) *TaskPostData {
    d.SearchPartners = searchPartners
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "sort_by".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetSortBy(sortBy string) *TaskPostData {
    d.SortBy = sortBy
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (d *TaskPostData) SetTag(tag string) *TaskPostData {
    d.Tag = tag
    return d
}


