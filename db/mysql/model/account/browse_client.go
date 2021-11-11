//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package account

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
func (o *BrowseClient) Count() (int64, *mySelectStatement.ResultCount, error) {
    resultCount, err := o.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) Query() (*mySelectStatement.ResultQuery, error) {
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) QueryRow() (*mySelectStatement.ResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) Run() ([]*Account, *mySelectStatement.Result, error) {
    var accounts []*Account
    result, err := o.BaseClient.Run()
    if err != nil { 
        return accounts, result, err
    }

    for _, row := range result.Rows {
        account := &Account{}
        if err := scanAccount(row, account); err != nil {
            return accounts, result, err
        }
        accounts = append(accounts, account)
    }

    return accounts, result, nil
}
