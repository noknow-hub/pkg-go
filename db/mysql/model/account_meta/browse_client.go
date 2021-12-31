//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
)

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}
type BrowseClientWithAccount struct {
    BaseClient *myQuery.SelectClient
    RefAccountTable string
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
// New BrowseClient with reference account table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithAccount(refAccountTable string) *BrowseClientWithAccount {
    return &BrowseClientWithAccount{
        BaseClient: c.BaseClient,
        RefAccountTable: refAccountTable,
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
func (c *BrowseClient) Run() ([]*AccountMeta, *myQuery.SelectResult, error) {
    var accountMetas []*AccountMeta
    result, err := c.BaseClient.Run()
    if err != nil { 
        return accountMetas, result, err
    }

    for _, row := range result.Rows {
        accountMeta := &AccountMeta{}
        if err := scanAccountMeta(row, accountMeta); err != nil {
            return accountMetas, result, err
        }
        accountMetas = append(accountMetas, accountMeta)
    }

    return accountMetas, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN account.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithAccount) Run() ([]*AccountMetaWithAccount, *myQuery.SelectResult, error) {
    var accountMetas []*AccountMetaWithAccount
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_ACCOUNT_ID, c.RefAccountTable, nkwMysqlModelAccount.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return accountMetas, result, err
    }

    for _, row := range result.Rows {
        accountMeta := &AccountMetaWithAccount{
            Account: &nkwMysqlModelAccount.Account{},
        }
        if err := scanAccountMetaWithAccount(row, c.BaseClient.TableName, c.RefAccountTable, accountMeta); err != nil {
            return accountMetas, result, err
        }
        accountMetas = append(accountMetas, accountMeta)
    }

    return accountMetas, result, nil
}
