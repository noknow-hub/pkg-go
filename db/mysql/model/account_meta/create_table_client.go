//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
)

type CreateTableClient struct {
    *myQuery.CreateTableClient
    RefAccountTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refAccountTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDb(tableName, db),
        RefAccountTableName: refAccountTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refAccountTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx),
        RefAccountTableName: refAccountTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refAccountTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTx(tableName, tx),
        RefAccountTableName: refAccountTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refAccountTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx),
        RefAccountTableName: refAccountTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    c.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ACCOUNT_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Account ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_META_KEY, "VARCHAR(255)").
                SetNotNull().
                SetComment("Meta key")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_META_VALUE, "TEXT").
                SetNotNull().
                SetComment("Meta value")).
        SetPrimaryKeys([]string{COL_ACCOUNT_ID, COL_META_KEY}).
        AppendConstraint("", COL_ACCOUNT_ID, c.RefAccountTableName, nkwMysqlModelAccount.COL_ID, true, true).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
