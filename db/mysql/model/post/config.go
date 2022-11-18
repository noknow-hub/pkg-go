//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package post

import (
    "time"
)

const (
    COL_COUNTRY_CODE = "country_code"
    COL_CREATED_AT = "created_at"
    COL_ID = "id"
    COL_LANG_CODE = "lang_code"
    COL_PARENT_ID = "parent_id"
    COL_STATUS = "status"
    COL_TEXT = "text"
    COL_TYPE = "type"
    COL_UPDATED_AT = "updated_at"
    NUM_COLS = 9
    TABLE_NAME = "posts"
    VAL_STATUS_PUBLIC = "1"
    VAL_TYPE_POST = "1"
    VAL_TYPE_COMMENT = "2"
    VAL_TYPE_REPOST = "3"
)

type Post struct {
    Id string
    ParentId string
    Status string
    Type string
    LangCode string
    CountryCode string
    Text string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type AddEditValues struct {
    Id *string
    ParentId *string
    Status *string
    Type *string
    LangCode *string
    CountryCode *string
    Text *string
    CreatedAt *time.Time
}
