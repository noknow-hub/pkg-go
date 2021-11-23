//////////////////////////////////////////////////////////////////////
// tasks_ready_client.go
//////////////////////////////////////////////////////////////////////
package standard

import (
    "encoding/json"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/dataforseo/result/v3/serp/google/organic/standard"
)

type TasksReadyClient struct {
    *myAuthentication.Authentication
    EndpointUrl string
}


//////////////////////////////////////////////////////////////////////
// New TasksReadyClient object.
//////////////////////////////////////////////////////////////////////
func NewTasksReadyClient(login, password string, isSandbox bool) *TasksReadyClient {
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASKS_READY_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASKS_READY_V3_SANDBOX
    }
    return &TasksReadyClient{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *TasksReadyClient) Run() (int, *myResult.TasksReadyResults, error) {
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.
        AddHeaderContentType(myHttpClient.CONTENT_TYPE_JSON).
        AddHeaderAuthorization("Basic " + c.Cred)
    resp, err := httpClient.Get()
    if err != nil {
        return 0, nil, err
    }
    var result *myResult.TasksReadyResults
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}
