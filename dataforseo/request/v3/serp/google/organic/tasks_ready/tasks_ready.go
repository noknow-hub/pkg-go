//////////////////////////////////////////////////////////////////////
// tasks_ready.go
//////////////////////////////////////////////////////////////////////
package tasks_ready

import (
    "encoding/json"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    Response "github.com/noknow-hub/pkg-go/dataforseo/response/v3/serp/google/organic/tasks_ready"
)

type Client struct {
    *myAuthentication.Authentication
    EndpointUrl string
}


//////////////////////////////////////////////////////////////////////
// New Client
//////////////////////////////////////////////////////////////////////
func NewClient(login, password string, isSandbox bool) *Client {
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASKS_READY_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASKS_READY_V3_SANDBOX
    }
    return &Client{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// Run
//////////////////////////////////////////////////////////////////////
func (c *Client) Run() (int, *Response.Response, error) {
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.
        AddHeaderContentType(myHttpClient.CONTENT_TYPE_JSON).
        AddHeaderAuthorization("Basic " + c.Cred)
    resp, err := httpClient.Get()
    if err != nil {
        return 0, nil, err
    }
    var result *Response.Response
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}
