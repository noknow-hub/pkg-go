//////////////////////////////////////////////////////////////////////
// create_table_client.go
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

type CreateTableClient struct {
    *myQuery.CreateTableClient
    RefAccountTableName string
    RefPostTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refAccountTableName, refPostTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDb(tableName, db),
        RefAccountTableName: refAccountTableName,
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refAccountTableName, refPostTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx),
        RefAccountTableName: refAccountTableName,
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refAccountTableName, refPostTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTx(tableName, tx),
        RefAccountTableName: refAccountTableName,
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refAccountTableName, refPostTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx),
        RefAccountTableName: refAccountTableName,
        RefPostTableName: refPostTableName,
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
            myQuery.NewColumnDefinition(COL_POST_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Post ID")).
        AppendConstraint("", COL_ACCOUNT_ID, c.RefAccountTableName, nkwMysqlModelAccount.COL_ID, true, true).
        AppendConstraint("", COL_POST_ID, c.RefPostTableName, nkwMysqlModelPost.COL_ID, true, true).
        SetPrimaryKeys([]string{COL_ACCOUNT_ID, COL_POST_ID}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
