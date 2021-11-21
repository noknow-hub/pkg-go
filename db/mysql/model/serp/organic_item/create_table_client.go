//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package organic_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
    myCreateTableStatement "github.com/noknow-hub/pkg-go/db/mysql/query/create_table_statement"
)

type CreateTableClient struct {
    *myCreateTableStatement.Client
    RefTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        Client: myCreateTableStatement.NewClientWithDb(tableName, db),
        RefTableName: refTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        Client: myCreateTableStatement.NewClientWithDbContext(tableName, db, ctx),
        RefTableName: refTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        Client: myCreateTableStatement.NewClientWithTx(tableName, tx),
        RefTableName: refTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        Client: myCreateTableStatement.NewClientWithTxContext(tableName, tx, ctx),
        RefTableName: refTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myCreateTableStatement.Result, error) {
    c.
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_ID, "VARCHAR(50)").
                SetNotNull().
                SetComment("SERP organic item ID."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_SERP_ID, "VARCHAR(50)").
                SetNotNull().
                SetComment("SERP ID."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_RANK, "INT UNSIGNED").
                SetNotNull().
                SetComment("Rank."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_GROUP_RANK, "INT UNSIGNED").
                SetComment("Group rank."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_DOMAIN, "VARCHAR(150)").
                SetNotNull().
                SetComment("Domain."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_URL, "TEXT").
                SetNotNull().
                SetComment("URL."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_TITLE, "TEXT").
                SetNotNull().
                SetComment("Title."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_TIMESTAMP, "DATETIME").
                SetComment("Timestamp."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_SNIPPET, "TEXT").
                SetComment("Snippet."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_RELATED_URL, "TEXT").
                SetComment("Related search URL."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_CACHE_URL, "TEXT").
                SetComment("Cached URL."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_CREATED_AT, "DATETIME").
                SetNotNull().
                SetDefault("CURRENT_TIMESTAMP").
                SetComment("Created at."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_UPDATED_AT, "DATETIME").
                SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP").
                SetComment("Updated at."),
        ).
        AppendConstraint("", COL_SERP_ID, c.RefTableName, mySerp.COL_ID, true, true).
        SetPrimaryKeys([]string{COL_ID}).
        SetIndexKeys([]string{COL_RANK, COL_DOMAIN}).
        SetComment(c.TableName + " table.")
    return c.Client.Run()
}
