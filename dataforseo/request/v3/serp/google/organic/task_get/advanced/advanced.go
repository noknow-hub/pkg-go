//////////////////////////////////////////////////////////////////////
// advanced.go
//////////////////////////////////////////////////////////////////////
package advanced

import (
    "encoding/json"
    "strings"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResponse "github.com/noknow-hub/pkg-go/dataforseo/response/v3/serp/google/organic/task_get/advanced"
)

type Client struct {
    *myAuthentication.Authentication
    EndpointUrl string
}


//////////////////////////////////////////////////////////////////////
// New Client
//////////////////////////////////////////////////////////////////////
func NewClient(login, password string, isSandbox bool, id string) *Client {
    endpointUrl := strings.Replace(myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_GET_ADVANCED_V3, "$id", id, 1)
    if isSandbox {
        endpointUrl = strings.Replace(myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_GET_ADVANCED_V3_SANDBOX, "$id", id, 1)
    }
    return &Client{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run() (int, *myResponse.Response, error) {
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.
        AddHeaderContentType(myHttpClient.CONTENT_TYPE_JSON).
        AddHeaderAuthorization("Basic " + c.Cred)
    resp, err := httpClient.Get()
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
