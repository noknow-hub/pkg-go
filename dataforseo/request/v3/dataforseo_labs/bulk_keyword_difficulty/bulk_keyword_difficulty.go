//////////////////////////////////////////////////////////////////////
// bulk_keyword_difficulty.go
//////////////////////////////////////////////////////////////////////
package bulk_keyword_difficulty

import (
    "encoding/json"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/dataforseo/response/v3/dataforseo_labs/bulk_keyword_difficulty"
)

type Client struct {
    *myAuthentication.Authentication
    EndpointUrl string
}

type Data struct {
    Keywords []string      `json:"keywords,omitempty"`
    LanguageCode string    `json:"language_code,omitempty"`
    LanguageName string    `json:"language_name,omitempty"`
    LocationCode int       `json:"location_code,omitempty"`
    LocationName string    `json:"location_name,omitempty"`
    Tag string             `json:"tag,omitempty"`
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
// New Data with "language_code" and "location_code".
//////////////////////////////////////////////////////////////////////
func NewDataWithCode(keywords []string, languageCode string, locationCode int) *Data {
    return &Data{
        Keywords: keywords,
        LanguageCode: languageCode,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "language_name" and "location_name".
//////////////////////////////////////////////////////////////////////
func NewDataWithName(keywords []string, languageName, locationName string) *Data {
    return &Data{
        Keywords: keywords,
        LanguageName: languageName,
        LocationName: locationName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run(data []*Data) (int, *myResult.Response, error) {
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
    var result *myResult.Response
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetTag(tag string) *Data {
    d.Tag = tag
    return d
}
