//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package article

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
                SetComment("Article ID.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_STATUS, "VARCHAR(10)").
                SetNotNull().
                SetDefault(VAL_STATUS_PUBLIC).
                SetComment("Status.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TITLE, "VARCHAR(255)").
                SetNotNull().
                SetComment("Title.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_URL, "VARCHAR(255)").
                SetNotNull().
                SetComment("URL.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TEXT, "LONGTEXT").
                SetNotNull().
                SetComment("Text.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_LANG_CODE, "VARCHAR(2)").
                SetComment("Language code with 2 digits.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_EXCERPT, "VARCHAR(255)").
                SetComment("Excerpt.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_THUMBNAIL_URL, "VARCHAR(255)").
                SetComment("Thumbnail image URL.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PASSWORD, "VARCHAR(255)").
                SetComment("Password to access it.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(Type, "VARCHAR(10)").
                SetDefault(VAL_TYPE_DEFAULT).
                SetComment("Type.")).
        AppendColumnDefinition(  
            myQuery.NewColumnDefinition(COL_CREATED_AT, "DATETIME").
                SetNotNull().
                SetDefault("CURRENT_TIMESTAMP").
                SetComment("Created at.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_UPDATED_AT, "DATETIME").
                SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP").
                SetComment("Updated at.")).
        SetPrimaryKeys([]string{COL_ID}).
        SetUniqueKeys([]string{COL_URL, COL_LANG_CODE}).
        SetIndexKeys([]string{COL_STATUS, COL_TITLE, COL_URL, COL_LANG_CODE}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
