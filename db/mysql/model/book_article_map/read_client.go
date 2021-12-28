//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package book_article_map

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelBook "github.com/noknow-hub/pkg-go/db/mysql/model/book"
)

type ReadClient struct {
    BaseClient *myQuery.SelectClient
}
type ReadClientWithBookAndArticle struct {
    BaseClient *myQuery.SelectClient
    RefBookTable string
    RefArticleTable string
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDb(tableName string, db *sql.DB) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTx(tableName string, tx *sql.Tx) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with reference articles and tags table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithBookAndArticle(refBookTable, refArticleTable string) *ReadClientWithBookAndArticle {
    return &ReadClientWithBookAndArticle{
        BaseClient: c.BaseClient,
        RefBookTable: refBookTable,
        RefArticleTable: refArticleTable,
    }
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Query() (*myQuery.SelectResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Run() (*BookArticleMap, *myQuery.SelectResult, error) {
    var bookArticleMap *BookArticleMap
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return bookArticleMap, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return bookArticleMap, result, err
    }
    bookArticleMap = &BookArticleMap{}
    if err := scanBookArticleMap(result.Rows[0], bookArticleMap); err != nil {
        return bookArticleMap, result, err
    }
    return bookArticleMap, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithBookAndArticle) Run() (*BookArticleMap, *myQuery.SelectResult, error) {
    var bookArticleMap *BookArticleMap
    c.BaseClient.SetLimit(1)
    c.BaseClient.
        AppendInnerJoinTables(c.BaseClient.TableName, COL_BOOK_ID, c.RefBookTable, nkwMysqlModelBook.COL_ID).
        AppendInnerJoinTables(c.BaseClient.TableName, COL_ARTICLE_ID, c.RefArticleTable, nkwMysqlModelArticle.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return bookArticleMap, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return bookArticleMap, result, err
    }
    bookArticleMap = &BookArticleMap{
        Book: &nkwMysqlModelBook.Book{},
        Article: &nkwMysqlModelArticle.Article{},
    }
    if err := scanBookArticleMapWithBookAndArticle(result.Rows[0], c.BaseClient.TableName, c.RefBookTable, c.RefArticleTable, bookArticleMap); err != nil {
        return bookArticleMap, result, err
    }
    return bookArticleMap, result, nil
}

