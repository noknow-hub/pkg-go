//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type EditClient struct {
    BaseClient *myQuery.UpdateClient
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDb(tableName string, db *sql.DB) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTx(tableName string, tx *sql.Tx) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) Run() (*myQuery.UpdateResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_ACCOUNT_ID and COL_META_KEY.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunByAccountIdAndMetaKey(accountId, metaKey string) (*myQuery.UpdateResult, error) {
    c.BaseClient.WhereCondition.
        SetWhere(COL_ACCOUNT_ID, accountId).
        AppendAnd(COL_META_KEY, metaKey)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by COL_ACCOUNT_ID and COL_META_KEY.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllByAccountIdAndMetaKey(currentAccountId, currentMetaKey string, accountId, metaKey, metaValue *string) (*myQuery.UpdateResult, error) {
    if accountId != nil {
        c.BaseClient.AssignmentList.Append(COL_ACCOUNT_ID, *accountId)
    }
    if metaKey != nil {
        c.BaseClient.AssignmentList.Append(COL_META_KEY, *metaKey)
    }
    if metaValue != nil {
        c.BaseClient.AssignmentList.Append(COL_META_VALUE, *metaValue)
    }
    c.BaseClient.WhereCondition.
        SetWhere(COL_ACCOUNT_ID, currentAccountId).
        AppendAnd(COL_META_KEY, currentMetaKey)
    return c.BaseClient.Run()
}
