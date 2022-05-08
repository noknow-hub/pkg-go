//////////////////////////////////////////////////////////////////////
// tasks_ready.go
//////////////////////////////////////////////////////////////////////
package tasks_ready

import (
    myResult "github.com/noknow-hub/pkg-go/dataforseo/response/v3"
)

type Response struct {
    *myResult.General
    Tasks []*Task  `json:"tasks"`
    Raw string
}

type Task struct {
    *myResult.Task
    Result []*Result  `json:"result"`
}

type Result struct {
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
// Get Results of Response.
//////////////////////////////////////////////////////////////////////
func (r *Response) GetResults() []*Result {
    if len(r.Tasks) == 0 {
        return nil
    }
    var result []*Result
    for _, o := range r.Tasks {
        for _, oo := range o.Result {
            result = append(result, oo)
        }
    }
    return result
}
