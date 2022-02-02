//////////////////////////////////////////////////////////////////////
// bulk_keyword_difficulty.go
//////////////////////////////////////////////////////////////////////
package bulk_keyword_difficulty

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

type Result struct {
    LocationCode int        `json:"location_code"`
    LanguageCode string     `json:"language_code"`
    TotalCount int          `json:"total_count"`
    ItemsCount int          `json:"items_count"`
    Items []*Item           `json:"items"`
}

type Item struct {
    Keyword string           `json:"keyword"`
    keywordDifficulty int    `json:"keyword_difficulty"`
}
