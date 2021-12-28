//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package book_article_map

import (
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelBook "github.com/noknow-hub/pkg-go/db/mysql/model/book"
)

const (
    COL_ARTICLE_ID = "article_id"
    COL_BOOK_ID = "book_id"
    COL_CHAPTER = "chapter"
    COL_INTRODUCTION = "introduction"
    COL_PART = "part"
    COL_SECTION = "section"
    COL_SUB_SECTION = "sub_section"
    NUM_COLS = 7
    TABLE_NAME = "book_article_map"
)

type BookArticleMap struct {
    BookId string
    ArticleId string
    Part string
    Chapter string
    Section string
    SubSection string
    Introduction string
    Book *nkwMysqlModelBook.Book
    Article *nkwMysqlModelArticle.Article
}
