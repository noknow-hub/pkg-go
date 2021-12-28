//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package book

import (
    "time"
)

const (
    COL_ACKNOWLEDGMENTS = "acknowledgements"
    COL_AUTHOR_NOTE = "author_note"
    COL_CREATED_AT = "created_at"
    COL_FOREWORD = "foreword"
    COL_ID = "id"
    COL_INTRODUCTION = "introduction"
    COL_LANG_CODE = "lang_code"
    COL_PASSWORD = "password"
    COL_PREFACE = "preface"
    COL_STATUS = "status"
    COL_THUMBNAIL_URL = "thumbnail_url"
    COL_TITLE = "title"
    COL_TYPE = "type"
    COL_UPDATED_AT = "updated_at"
    COL_URL = "url"
    NUM_COLS = 15
    TABLE_NAME = "books"
    VAL_TYPE_ADVERTISEMENT = "ads"
    VAL_TYPE_PAID = "paid"
    VAL_STATUS_PUBLIC = "public"
    VAL_STATUS_PRIVATE = "private"
    VAL_STATUS_DELETED = "deleted"
    VAL_STATUS_REVIEW = "review"
)

type Book struct {
    Id string
    Status string
    Title string
    Url string
    LangCode string
    Introduction string
    ThumbnailUrl string
    Preface string
    Foreword string
    Acknowledgements string
    AuthorNote string
    Password string
    Type string
    CreatedAt time.Time
    UpdatedAt time.Time
}
