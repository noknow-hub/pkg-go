//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package post_image

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)

type CreateTableClient struct {
    *myQuery.CreateTableClient
    RefPostTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refPostTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDb(tableName, db),
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refPostTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx),
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refPostTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTx(tableName, tx),
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refPostTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx),
        RefPostTableName: refPostTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    c.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Image ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_POST_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Post ID")).
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
            myQuery.NewColumnDefinition(COL_CREATED_AT, "DATETIME").
                SetNotNull().
                SetDefault("CURRENT_TIMESTAMP").
                SetComment("Created at.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_UPDATED_AT, "DATETIME").
                SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP").
                SetComment("Updated at.")).
        AppendConstraint("", COL_POST_ID, c.RefPostTableName, nkwMysqlModelPost.COL_ID, true, true).
        SetPrimaryKeys([]string{COL_ID}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
