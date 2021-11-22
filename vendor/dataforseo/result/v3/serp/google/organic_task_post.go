//////////////////////////////////////////////////////////////////////
// organic_task_post.go
//////////////////////////////////////////////////////////////////////
package google

import (
    myResult "github.com/noknow-hub/pkg-go/vendor/dataforseo/result/v3"
)

type Result struct {
    *myResult.General
    Tasks []*Task  `json:"tasks"`
    Raw string
}
type Task struct {
    *myResult.Task
    Result []*Result  `json:"result"`
}
