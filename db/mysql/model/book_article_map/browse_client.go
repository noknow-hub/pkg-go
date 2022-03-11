//////////////////////////////////////////////////////////////////////
// browse_client.go
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

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}
type BrowseClientWithBookAndArticle struct {
    BaseClient *myQuery.SelectClient
    RefArticleTable string
    RefBookTable string
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDb(tableName string, db *sql.DB) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTx(tableName string, tx *sql.Tx) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with reference book and article table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithBookAndArticle(refBookTable, refArticleTable string) *BrowseClientWithBookAndArticle {
    return &BrowseClientWithBookAndArticle{
        BaseClient: c.BaseClient,
        RefBookTable: refBookTable,
        RefArticleTable: refArticleTable,
    }
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Query() (*myQuery.SelectResultQuery, error) {
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Run() ([]*BookArticleMap, *myQuery.SelectResult, error) {
    var bookArticleMaps []*BookArticleMap
    result, err := c.BaseClient.Run()
    if err != nil { 
        return bookArticleMaps, result, err
    }

    for _, row := range result.Rows {
        bookArticleMap := &BookArticleMap{}
        if err := scanBookArticleMap(row, bookArticleMap); err != nil {
            return bookArticleMaps, result, err
        }
        bookArticleMaps = append(bookArticleMaps, bookArticleMap)
    }

    return bookArticleMaps, result, nil
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithBookAndArticle) Count() (int64, *myQuery.SelectResultCount, error) {
    c.BaseClient.
        AppendInnerJoinTables(c.BaseClient.TableName, COL_BOOK_ID, c.RefBookTable, nkwMysqlModelBook.COL_ID).
        AppendInnerJoinTables(c.BaseClient.TableName, COL_ARTICLE_ID, c.RefArticleTable, nkwMysqlModelArticle.COL_ID)
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithBookAndArticle) Run() ([]*BookArticleMap, *myQuery.SelectResult, error) {
    var bookArticleMaps []*BookArticleMap
    c.BaseClient.
        AppendInnerJoinTables(c.BaseClient.TableName, COL_BOOK_ID, c.RefBookTable, nkwMysqlModelBook.COL_ID).
        AppendInnerJoinTables(c.BaseClient.TableName, COL_ARTICLE_ID, c.RefArticleTable, nkwMysqlModelArticle.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return bookArticleMaps, result, err
    }

    for _, row := range result.Rows {
        bookArticleMap := &BookArticleMap{
            Book: &nkwMysqlModelBook.Book{},
            Article: &nkwMysqlModelArticle.Article{},
        }
        if err := scanBookArticleMapWithBookAndArticle(row, c.BaseClient.TableName, c.RefBookTable, c.RefArticleTable, bookArticleMap); err != nil {
            return bookArticleMaps, result, err
        }
        bookArticleMaps = append(bookArticleMaps, bookArticleMap)
    }

    return bookArticleMaps, result, nil
}
