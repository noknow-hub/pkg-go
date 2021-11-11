//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package account

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
func (o *ReadClient) Query() (*mySelectStatement.ResultQuery, error) {
    o.BaseClient.SetLimit(1)
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *ReadClient) QueryRow() (*mySelectStatement.ResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *ReadClient) Run() ([]*Account, *mySelectStatement.Result, error) {
    var accounts []*Account
    o.BaseClient.SetLimit(1)
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
