//////////////////////////////////////////////////////////////////////
// standard.go
//////////////////////////////////////////////////////////////////////
package standard

import (
    myResult "github.com/noknow-hub/pkg-go/dataforseo/result/v3"
)

type TasksReadyResults struct {
    *myResult.General
    Tasks []*TasksReadyTask  `json:"tasks"`
    Raw string
}

type TasksReadyTask struct {
    *myResult.Task
    Result []*TasksReadyResult  `json:"result"`
}

type TasksReadyResult struct {
    Id string                `json:"id"`
    Se string                `json:"se"`
    SeType string            `json:"se_type"`
    DatePosted string        `json:"date_posted"`
    Tag string               `json:"tag"`
    EndpointRegular string   `json:"endpoint_regular"`
    EndpointAdvanced string  `json:"endpoint_advanced"`
    EndpointHtml string      `json:"endpoint_html"`
}


//////////////////////////////////////////////////////////////////////
// Get Results of TasksReadyResults.
//////////////////////////////////////////////////////////////////////
func (r *TasksReadyResults) GetResults() []*TasksReadyResult {
    if len(r.Tasks) == 0 {
        return nil
    }
    var result []*TasksReadyResult
    for _, o := range r.Tasks {
        for _, oo := range o.Result {
            result = append(result, oo)
        }
    }
    return result
}
