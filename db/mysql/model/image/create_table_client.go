//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package image

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
            myQuery.NewColumnDefinition(COL_URL, "VARCHAR(255)").
                SetNotNull().
                SetComment("Image URL")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_NAME, "VARCHAR(255)").
                SetComment("Image name")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SIZE, "INT UNSIGNED").
                SetComment("Image size")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PATH, "VARCHAR(255)").
                SetComment("Image path")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ALT, "VARCHAR(255)").
                SetComment("Image alt")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TYPE, "VARCHAR(10)").
                SetComment("Image type")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_MIME_TYPE, "VARCHAR(200)").
                SetComment("Image mime type")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_LABEL, "VARCHAR(255)").
                SetComment("Image label")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_LINK, "VARCHAR(255)").
                SetComment("Link URL")).
        AppendColumnDefinition(  
            myQuery.NewColumnDefinition(COL_CREATED_AT, "DATETIME").
                SetNotNull().
                SetDefault("CURRENT_TIMESTAMP").
                SetComment("Created at.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_UPDATED_AT, "DATETIME").
                SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP").
                SetComment("Updated at.")).
        SetPrimaryKeys([]string{COL_URL}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
