//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package book_article_map

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type AddClient struct {
    BaseClient *myQuery.InsertClient
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) Run() (*myQuery.InsertResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithAll(bookId, articleId string, part, chapter, section, subSection, introduction *string) (*myQuery.InsertResult, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, []string{COL_BOOK_ID, COL_ARTICLE_ID}...)
    vals = append(vals, []interface{}{bookId, articleId}...)

    if part != nil {
        cols = append(cols, COL_PART)
        vals = append(vals, *part)
    }
    if chapter != nil {
        cols = append(cols, COL_CHAPTER)
        vals = append(vals, *chapter)
    }
    if section != nil {
        cols = append(cols, COL_SECTION)
        vals = append(vals, *section)
    }
    if subSection != nil {
        cols = append(cols, COL_SUB_SECTION)
        vals = append(vals, *subSection)
    }
    if introduction != nil {
        cols = append(cols, COL_INTRODUCTION)
        vals = append(vals, *introduction)
    }

    c.BaseClient.
        SetColNames(cols).
        AppendValues(vals)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with required.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithRequired(bookId, articleId string) (*myQuery.InsertResult, error) {
    c.BaseClient.
        SetColNames([]string{COL_BOOK_ID, COL_ARTICLE_ID}).
        AppendValues([]interface{}{bookId, articleId})
    return c.BaseClient.Run()
}
