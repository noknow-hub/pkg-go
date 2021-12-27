//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package article

import (
    "time"
)

const (
    COL_CREATED_AT = "created_at"
    COL_EXCERPT = "excerpt"
    COL_ID = "id"
    COL_LANG_CODE = "lang_code"
    COL_PASSWORD = "password"
    COL_STATUS = "status"
    COL_TEXT = "text"
    COL_THUMBNAIL_URL = "thumbnail_url"
    COL_TITLE = "title"
    COL_TYPE = "type"
    COL_UPDATED_AT = "updated_at"
    COL_URL = "url"
    NUM_COLS = 12
    TABLE_NAME = "articles"
    VAL_TYPE_DEFAULT = "default"
    VAL_TYPE_ADVERTISEMENT = "ads"
    VAL_TYPE_PAID = "paid"
    VAL_STATUS_PUBLIC = "public"
    VAL_STATUS_PRIVATE = "private"
    VAL_STATUS_DELETED = "deleted"
    VAL_STATUS_REVIEW = "review"
)

type Article struct {
    Id string
    Status string
    Title string
    Url string
    Text string
    LangCode string
    Excerpt string
    ThumbnailUrl string
    Password string
    Type string
    CreatedAt time.Time
    UpdatedAt time.Time
}
