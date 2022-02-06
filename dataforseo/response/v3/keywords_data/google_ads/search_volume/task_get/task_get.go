//////////////////////////////////////////////////////////////////////
// task_get.go
//////////////////////////////////////////////////////////////////////
package task_get

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
    Results []*Result    `json:"result"`
}

type Result struct {
    Keyword string                      `json:"keyword"`
    Spell string                        `json:"spell"`
    LocationCode int                    `json:"location_code"`
    LanguageCode string                 `json:"language_code"`
    SearchPartners bool                 `json:"search_partners"`
    Competition string                  `json:"competition"`
    CompetitionIndex int                `json:"competition_index"`
    SearchVolume int                    `json:"search_volume"`
    LowTopOfPageBid float64             `json:"low_top_of_page_bid"`
    HighTopOfPageBid float64            `json:"high_top_of_page_bid"`
    MonthlySearches []*MonthlySearch    `json:"monthly_searches"`
}

type MonthlySearch struct {
    Year int            `json:"year"`
    Month int           `json:"month"`
    SearchVolume int    `json:"search_volume"`
}
