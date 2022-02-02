//////////////////////////////////////////////////////////////////////
// task_post.go
//////////////////////////////////////////////////////////////////////
package task_post

import (
    myResult "github.com/noknow-hub/pkg-go/dataforseo/response/v3"
)

type Response struct {
    *myResult.General
    Tasks []*Task    `json:"tasks"`
    Raw string
}

type Task struct {
    *myResult.Task
    Result []*Result    `json:"result"`
}

type Result struct {}
