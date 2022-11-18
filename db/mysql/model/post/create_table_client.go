//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package post

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type CreateTableClient struct {
    *myQuery.CreateTableClient
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{ myQuery.NewCreateTableClientWithDb(tableName, db) }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{ myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx) }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{ myQuery.NewCreateTableClientWithTx(tableName, tx) }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{ myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx) }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    c.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Post ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PARENT_ID, "BIGINT UNSIGNED").
                SetComment("Parent post ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_STATUS, "TINYINT(1) UNSIGNED").
                SetNotNull().
                SetDefault(VAL_STATUS_PUBLIC).
                SetComment("Status")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TYPE, "TINYINT(1) UNSIGNED").
                SetNotNull().
                SetDefault(VAL_TYPE_POST).
                SetComment("Post type")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_LANG_CODE, "VARCHAR(2)").
                SetComment("Language code with 2 digits")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_COUNTRY_CODE, "VARCHAR(2)").
                SetComment("Country code with 2 digits")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TEXT, "LONGTEXT").
                SetNotNull().
                SetComment("Text")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_CREATED_AT, "DATETIME").
                SetNotNull().
                SetDefault("CURRENT_TIMESTAMP").
                SetComment("Created at.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_UPDATED_AT, "DATETIME").
                SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP").
                SetComment("Updated at.")).
        AppendConstraint("", COL_PARENT_ID, c.TableName, COL_ID, true, true).
        SetPrimaryKeys([]string{COL_ID}).
        SetComment(c.TableName + " table")
    return c.CreateTableClient.Run()
}
