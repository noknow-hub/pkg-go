//////////////////////////////////////////////////////////////////////
// browse_client.go
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

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}

type BrowseClientWithAccount struct {
    BaseClient *myQuery.SelectClient
    RefAccountTableName string
}

type BrowseClientWithPost struct {
    BaseClient *myQuery.SelectClient
    RefPostTableName string
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
// New BrowseClient with account table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithAccount(refAccountTable string) *BrowseClientWithAccount {
    return &BrowseClientWithAccount{
        BaseClient: c.BaseClient,
        RefAccountTableName: refAccountTable,
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with post table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithPost(refPostTable string) *BrowseClientWithPost {
    return &BrowseClientWithPost{
        BaseClient: c.BaseClient,
        RefPostTableName: refPostTable,
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
func (o *BrowseClient) Query() (*myQuery.SelectResultQuery, error) {
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) Run() ([]*AccountPostMap, *myQuery.SelectResult, error) {
    var accountPostMaps []*AccountPostMap
    result, err := o.BaseClient.Run()
    if err != nil {
        return accountPostMaps, result, err
    }

    for _, row := range result.Rows {
        accountPostMap := &AccountPostMap{}
        if err := scan(row, accountPostMap); err != nil {
            return accountPostMaps, result, err
        }
        accountPostMaps = append(accountPostMaps, accountPostMap)
    }

    return accountPostMaps, result, nil
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithAccount) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClientWithAccount) Query() (*myQuery.SelectResultQuery, error) {
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClientWithAccount) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithAccount) Run() ([]*AccountPostMapWithAccount, *myQuery.SelectResult, error) {
    var accountPostMaps []*AccountPostMapWithAccount
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_ACCOUNT_ID, c.RefAccountTableName, nkwMysqlModelAccount.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return accountPostMaps, result, err
    }

    for _, row := range result.Rows {
        accountPostMap := &AccountPostMapWithAccount{
            Account: &nkwMysqlModelAccount.Account{},
        }
        if err := scanWithAccount(row, c.BaseClient.TableName, c.RefAccountTableName, accountPostMap); err != nil {
            return accountPostMaps, result, err
        }
        accountPostMaps = append(accountPostMaps, accountPostMap)
    }

    return accountPostMaps, result, nil
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithPost) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClientWithPost) Query() (*myQuery.SelectResultQuery, error) {
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClientWithPost) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithPost) Run() ([]*AccountPostMapWithPost, *myQuery.SelectResult, error) {
    var accountPostMaps []*AccountPostMapWithPost
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_POST_ID, c.RefPostTableName, nkwMysqlModelPost.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return accountPostMaps, result, err
    }

    for _, row := range result.Rows {
        accountPostMap := &AccountPostMapWithPost{
            Post: &nkwMysqlModelPost.Post{},
        }
        if err := scanWithPost(row, c.BaseClient.TableName, c.RefPostTableName, accountPostMap); err != nil {
            return accountPostMaps, result, err
        }
        accountPostMaps = append(accountPostMaps, accountPostMap)
    }

    return accountPostMaps, result, nil
}
