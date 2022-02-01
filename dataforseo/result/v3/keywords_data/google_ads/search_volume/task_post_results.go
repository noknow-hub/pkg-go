//////////////////////////////////////////////////////////////////////
// task_post_results.go
//////////////////////////////////////////////////////////////////////
package search_volume

import (
    myResult "github.com/noknow-hub/pkg-go/dataforseo/result/v3"
)

type TaskPostResults struct {
    *myResult.General
    Tasks []*TaskPostTask  `json:"tasks"`
    Raw string
}

type TaskPostTask struct {
    *myResult.Task
    Result []*TaskPostResult  `json:"result"`
}

type TaskPostResult struct {}
