//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package book

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
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
func (c *BrowseClient) Run() ([]*Book, *myQuery.SelectResult, error) {
    var books []*Book
    result, err := c.BaseClient.Run()
    if err != nil { 
        return books, result, err
    }

    for _, row := range result.Rows {
        book := &Book{}
        if err := scanBook(row, book); err != nil {
            return books, result, err
        }
        books = append(books, book)
    }

    return books, result, nil
}
