//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    mySelectStatement "github.com/noknow-hub/pkg-go/db/mysql/query/select_statement"
)

type BrowseClient struct {
    BaseClient *mySelectStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDb(tableName string, db *sql.DB) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTx(tableName string, tx *sql.Tx) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Count() (int64, *mySelectStatement.ResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Query() (*mySelectStatement.ResultQuery, error) {
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) QueryRow() (*mySelectStatement.ResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Run() ([]*Serp, *mySelectStatement.Result, error) {
    var serps []*Serp
    result, err := c.BaseClient.Run()
    if err != nil {
        return serps, result, err
    }

    for _, row := range result.Rows {
        serp := &Serp{}
        if err := scanSerp(row, serp); err != nil {
            return serps, result, err
        }
        serps = append(serps, serp)
    }

    return serps, result, nil
}
