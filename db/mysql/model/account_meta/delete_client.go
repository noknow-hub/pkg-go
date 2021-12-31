//////////////////////////////////////////////////////////////////////
// delete_client.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type DeleteClient struct {
    BaseClient *myQuery.DeleteClient
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDb(tableName string, db *sql.DB) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTx(tableName string, tx *sql.Tx) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) Run() (*myQuery.DeleteResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_ACCOUNT_ID.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) RunByAccountId(accountId string) (*myQuery.DeleteResult, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_ACCOUNT_ID, accountId)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_ACCOUNT_ID and COL_META_KEY.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) RunByAccountIdAndMetaKey(accountId, metaKey string) (*myQuery.DeleteResult, error) {
    c.BaseClient.WhereCondition.
        SetWhere(COL_ACCOUNT_ID, accountId).
        AppendAnd(COL_META_KEY, metaKey)
    return c.BaseClient.Run()
}
