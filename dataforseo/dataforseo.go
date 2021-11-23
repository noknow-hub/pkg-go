//////////////////////////////////////////////////////////////////////
// dataforseo.go
//////////////////////////////////////////////////////////////////////
package dataforseo

import (
    myV3SerpGoogleOrganicStandard "github.com/noknow-hub/pkg-go/dataforseo/client/v3/serp/google/organic/standard"
)


//////////////////////////////////////////////////////////////////////
// New client for "v3: serp: google: organic: task: get: advanced".
//////////////////////////////////////////////////////////////////////
func NewClientV3SerpGoogleOrganicTaskGetAdvanced(login, password string, isSandbox bool, id string) *myV3SerpGoogleOrganicStandard.TaskGetAdvancedClient {
    return myV3SerpGoogleOrganicStandard.NewTaskGetAdvancedClient(login, password, isSandbox, id)
}


//////////////////////////////////////////////////////////////////////
// New client for "v3: serp: google: organic: task: post".
//////////////////////////////////////////////////////////////////////
func NewClientV3SerpGoogleOrganicTaskPost(login, password string, isSandbox bool) *myV3SerpGoogleOrganicStandard.TaskPostClient {
    return myV3SerpGoogleOrganicStandard.NewTaskPostClient(login, password, isSandbox)
}


//////////////////////////////////////////////////////////////////////
// New client for "v3: serp: google: organic: tasks: ready".
//////////////////////////////////////////////////////////////////////
func NewClientV3SerpGoogleOrganicTasksReady(login, password string, isSandbox bool) *myV3SerpGoogleOrganicStandard.TasksReadyClient {
    return myV3SerpGoogleOrganicStandard.NewTasksReadyClient(login, password, isSandbox)
}
