//////////////////////////////////////////////////////////////////////
// delete_client.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myDeleteStatement "github.com/noknow-hub/pkg-go/db/mysql/query/delete_statement"
)

type DeleteClient struct {
    BaseClient *myDeleteStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDb(tableName string, db *sql.DB) *DeleteClient {
    return &DeleteClient{
        BaseClient: myDeleteStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        BaseClient: myDeleteStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTx(tableName string, tx *sql.Tx) *DeleteClient {
    return &DeleteClient{
        BaseClient: myDeleteStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        BaseClient: myDeleteStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) Run() (*myDeleteStatement.Result, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by "id".
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) RunById(id string) (*myDeleteStatement.Result, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return c.BaseClient.Run()
}
