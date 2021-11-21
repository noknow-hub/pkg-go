//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package organic_item

import (
    "time"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
)

const (
    COL_CACHE_URL = "cache_url"
    COL_CREATED_AT = "created_at"
    COL_DOMAIN = "domain"
    COL_GROUP_RANKING = "group_ranking"
    COL_ID = "id"
    COL_RANKING = "ranking"
    COL_RELATED_URL = "related_url"
    COL_SERP_ID = "serp_id"
    COL_SNIPPET =  "snippet"
    COL_TIMESTAMP = "timestamp"
    COL_TITLE = "title"
    COL_UPDATED_AT = "updated_at"
    COL_URL = "url"
    NUM_COLS = 13
    TABLE_NAME = "serp_organic_items"
)

type OrganicItem struct {
    Id string
    SerpId string
    Ranking string
    GroupRanking string
    Domain string
    Url string
    Title string
    Timestamp time.Time
    Snippet string
    RelatedUrl string
    CacheUrl string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SerpOrganicItem struct {
    *mySerp.Serp
    OrganicItems []*OrganicItem
}
