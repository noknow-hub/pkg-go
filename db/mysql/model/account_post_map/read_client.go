//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package account_post_map

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)

type ReadClient struct {
    BaseClient *myQuery.SelectClient
}

type ReadClientWithAccount struct {
    BaseClient *myQuery.SelectClient
    RefAccountTableName string
}

type ReadClientWithPost struct {
    BaseClient *myQuery.SelectClient
    RefPostTableName string
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
// New ReadClient with account table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithAccount(refAccountTable string) *ReadClientWithAccount {
    return &ReadClientWithAccount{
        BaseClient: c.BaseClient,
        RefAccountTableName: refAccountTable,
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with post table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithPost(refPostTable string) *ReadClientWithPost {
    return &ReadClientWithPost{
        BaseClient: c.BaseClient,
        RefPostTableName: refPostTable,
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
func (c *ReadClient) Run() (*AccountPostMap, *myQuery.SelectResult, error) {
    var accountPostMap *AccountPostMap
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return accountPostMap, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return accountPostMap, result, err
    }
    accountPostMap = &AccountPostMap{}
    if err := scan(result.Rows[0], accountPostMap); err != nil {
        return accountPostMap, result, err
    }
    return accountPostMap, result, nil
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithAccount) Query() (*myQuery.SelectResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithAccount) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithAccount) Run() (*AccountPostMapWithAccount, *myQuery.SelectResult, error) {
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_ACCOUNT_ID, c.RefAccountTableName, nkwMysqlModelAccount.COL_ID)
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return nil, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return nil, result, err
    }
    accountPostMap := &AccountPostMapWithAccount{
        Account: &nkwMysqlModelAccount.Account{},
    }
    if err := scanWithAccount(result.Rows[0], c.BaseClient.TableName, c.RefAccountTableName, accountPostMap); err != nil {
        return accountPostMap, result, err
    }
    return accountPostMap, result, nil
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithPost) Query() (*myQuery.SelectResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithPost) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithPost) Run() (*AccountPostMapWithPost, *myQuery.SelectResult, error) {
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_POST_ID, c.RefPostTableName, nkwMysqlModelPost.COL_ID)
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return nil, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return nil, result, err
    }
    accountPostMap := &AccountPostMapWithPost{
        Post: &nkwMysqlModelPost.Post{},
    }
    if err := scanWithPost(result.Rows[0], c.BaseClient.TableName, c.RefPostTableName, accountPostMap); err != nil {
        return accountPostMap, result, err
    }
    return accountPostMap, result, nil
}
