//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package image

import (
    "time"
)

const (
    COL_ALT = "alt"
    COL_CREATED_AT = "created_at"
    COL_LABEL = "label"
    COL_LINK = "link"
    COL_MIME_TYPE = "mime_type"
    COL_NAME = "name"
    COL_PATH = "path"
    COL_SIZE = "size"
    COL_TYPE = "type"
    COL_UPDATED_AT = "updated_at"
    COL_URL = "url"
    NUM_COLS = 11
    TABLE_NAME = "images"
    VAL_TYPE_DEFAULT = "default"
    VAL_TYPE_ADVERTISEMENT = "ads"
)

type Image struct {
    Url string
    Name string
    Size string
    Path string
    Alt string
    Type string
    MimeType string
    Label string
    Link string
    CreatedAt time.Time
    UpdatedAt time.Time
}
