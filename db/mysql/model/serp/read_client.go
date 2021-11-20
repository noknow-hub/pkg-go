//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    mySelectStatement "github.com/noknow-hub/pkg-go/db/mysql/query/select_statement"
)

type ReadClient struct {
    BaseClient *mySelectStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDb(tableName string, db *sql.DB) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTx(tableName string, tx *sql.Tx) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Query() (*mySelectStatement.ResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) QueryRow() (*mySelectStatement.ResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Run() ([]*Serp, *mySelectStatement.Result, error) {
    var serps []*Serp
    c.BaseClient.SetLimit(1)
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
