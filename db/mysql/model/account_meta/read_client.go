//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
)

type ReadClient struct {
    BaseClient *myQuery.SelectClient
}
type ReadClientWithAccount struct {
    BaseClient *myQuery.SelectClient
    RefAccountTable string
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
// New ReadClient with reference account table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithAccount(refAccountTable string) *ReadClientWithAccount {
    return &ReadClientWithAccount{
        BaseClient: c.BaseClient,
        RefAccountTable: refAccountTable,
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
func (c *ReadClient) Run() (*AccountMeta, *myQuery.SelectResult, error) {
    var accountMeta *AccountMeta
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return accountMeta, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return accountMeta, result, err
    }
    accountMeta = &AccountMeta{}
    if err := scanAccountMeta(result.Rows[0], accountMeta); err != nil {
        return accountMeta, result, err
    }
    return accountMeta, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN account.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithAccount) Run() (*AccountMetaWithAccount, *myQuery.SelectResult, error) {
    var accountMeta *AccountMetaWithAccount
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_ACCOUNT_ID, c.RefAccountTable, nkwMysqlModelAccount.COL_ID)
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return accountMeta, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return accountMeta, result, err
    }
    accountMeta = &AccountMetaWithAccount{
        Account: &nkwMysqlModelAccount.Account{},
    }
    if err := scanAccountMetaWithAccount(result.Rows[0], c.BaseClient.TableName, c.RefAccountTable, accountMeta); err != nil {
        return accountMeta, result, err
    }
    return accountMeta, result, nil
}
