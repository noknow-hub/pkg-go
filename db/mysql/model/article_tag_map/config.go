//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package article_tag_map

import (
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelTag "github.com/noknow-hub/pkg-go/db/mysql/model/tag"
)

const (
    COL_ARTICLE_ID = "article_id"
    COL_TAG_ID = "tag_id"
    NUM_COLS = 2
    TABLE_NAME = "article_tag_map"
)

type ArticleTagMap struct {
    ArticleId string
    TagId string
    Article *nkwMysqlModelArticle.Article
    Tag *nkwMysqlModelTag.Tag
}
