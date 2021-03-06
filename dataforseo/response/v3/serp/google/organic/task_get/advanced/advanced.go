//////////////////////////////////////////////////////////////////////
// advanced.go
//////////////////////////////////////////////////////////////////////
package advanced

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
    Keyword string       `json:"keyword"`
    Type string          `json:"type"`
    SeDomain string      `json:"se_domain"`
    LocationCode int     `json:"location_code"`
    LanguageCode string  `json:"language_code"`
    CheckUrl string      `json:"check_url"`
    Datetime string      `json:"datetime"`
    Spell *Spell         `json:"spell"`
    ItemTypes []string   `json:"item_types"`
    SeResultsCount int   `json:"se_results_count"`
    ItemsCount int       `json:"items_count"`
    Items []*Item        `json:"items"`
}

type Spell struct {
    Keyword string  `json:"keyword"`
    Type string     `json:"type"`
}

type Item struct {
    Type string                        `json:"type"`
    RankGroup int                      `json:"rank_group"`
    RankAbsolute int                   `json:"rank_absolute"`
    Position string                    `json:"position"`
    Xpath string                       `json:"xpath"`
    Domain string                      `json:"domain"`
    Title string                       `json:"title"`
    Url string                         `json:"url"`
    CacheUrl string                    `json:"cache_url"`
    RelatedSearchUrl string            `json:"related_search_url"`
    Breadcrumb string                  `json:"breadcrumb"`
    IsImage bool                       `json:"is_image"`
    IsVideo bool                       `json:"is_video"`
    IsFeaturedSnippet bool             `json:"is_featured_snippet"`
    IsMalicious bool                   `json:"is_malicious"`
    IsWebStory bool                    `json:"is_web_story"`
    Description string                 `json:"description"`
    PreSnippet string                  `json:"pre_snippet"`
    ExtendedSnippet string             `json:"extended_snippet"`
    Images []*Image                    `json:"images"`
    AmpVersion bool                    `json:"amp_version"`
    Rating *Rating                     `json:"rating"`
    Highlighted []string               `json:"highlighted"`
    Links []*Link                      `json:"links"`
    Faq *Faq                           `json:"faq"`
    ExtendedPeopleAlsoSearch []string  `json:"extended_people_also_search"`
    Timestamp string                   `json:"timestamp"`
    Price *Price                       `json:"price"`
    Rectangle *Rectangle               `json:"rectangle"`
}

type Image struct {
    Type string      `json:"type"`
    Alt string       `json:"alt"`
    Url string       `json:"url"`
    ImageUrl string  `json:"image_url"`
}

type Rating struct {
    RatingType string  `json:"rating_type"`
    Value float64      `json:"value"`
    VotesCount int     `json:"votes_count"`
    RatingMax float64  `json:"rating_max"`
}

type Link struct {
    Type string         `json:"type"`
    Title string        `json:"title"`
    Description string  `json:"description"`
    Url string          `json:"url"`
}

type Faq struct {
    Type string       `json:"type"`
    Items []*FaqItem  `json:"items"`
}

type FaqItem struct {
    Type string         `json:"type"`
    Title string        `json:"title"`
    Description string  `json:"description"`
    Links []*Link       `json:"links"`
}

type Price struct {
    Current float64        `json:"current"`
    Regular float64        `json:"regular"`
    MaxValue float64       `json:"max_value"`
    Currency string        `json:"currency"`
    IsPriceRange bool      `json:"is_price_range"`
    DisplayedPrice string  `json:"displayed_price"`
}

type Rectangle struct {
    X int       `json:"x"`
    Y int       `json:"y"`
    Width int   `json:"width"`
    Height int  `json:"height"`
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
