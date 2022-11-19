//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package post_image

import (
    "time"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)


const (
    COL_ALT = "alt"
    COL_CREATED_AT = "created_at"
    COL_ID = "id"
    COL_LABEL = "label"
    COL_MIME_TYPE = "mime_type"
    COL_NAME = "name"
    COL_PATH = "path"
    COL_POST_ID = "post_id"
    COL_SIZE = "size"
    COL_TYPE = "type"
    COL_UPDATED_AT = "updated_at"
    COL_URL = "url"
    NUM_COLS = 12
    TABLE_NAME = "post_images"
    VAL_TYPE_NORMAL = "1"
)

type PostImage struct {
    Id string
    PostId string
    Url string
    Name string
    Size string
    Path string
    Alt string
    Type string
    MimeType string
    Label string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type PostImageWithPost struct {
    Id string
    PostId string
    Url string
    Name string
    Size string
    Path string
    Alt string
    Type string
    MimeType string
    Label string
    CreatedAt time.Time
    UpdatedAt time.Time
    Post *nkwMysqlModelPost.Post
}

type AddEditValues struct {
    Id *string
    PostId *string
    Url *string
    Name *string
    Size *string
    Path *string
    Alt *string
    Type *string
    MimeType *string
    Label *string
    CreatedAt *time.Time
}
