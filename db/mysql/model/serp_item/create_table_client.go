//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package serp_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type CreateTableClient struct {
    *myQuery.CreateTableClient
    RefSerpTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refSerpTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDb(tableName, db),
        RefSerpTableName: refSerpTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refSerpTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx),
        RefSerpTableName: refSerpTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refSerpTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTx(tableName, tx),
        RefSerpTableName: refSerpTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refSerpTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx),
        RefSerpTableName: refSerpTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    c.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ID, "VARCHAR(50)").
                SetNotNull().
                SetComment("SERP organic item ID."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SERP_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("SERP ID."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_RANKING, "INT UNSIGNED").
                SetNotNull().
                SetComment("Ranking."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_GROUP_RANKING, "INT UNSIGNED").
                SetComment("Group ranking."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_DOMAIN, "VARCHAR(150)").
                SetNotNull().
                SetComment("Domain."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_URL, "TEXT").
                SetNotNull().
                SetComment("URL."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TITLE, "TEXT").
                SetNotNull().
                SetComment("Title."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TYPE, "VARCHAR(20)").
                SetComment("Type."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TIMESTAMP, "DATETIME").
                SetComment("Timestamp."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SNIPPET, "TEXT").
                SetComment("Snippet."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_RELATED_URL, "TEXT").
                SetComment("Related search URL."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_CACHE_URL, "TEXT").
                SetComment("Cached URL."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_CREATED_AT, "DATETIME").
                SetNotNull().
                SetDefault("CURRENT_TIMESTAMP").
                SetComment("Created at."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_UPDATED_AT, "DATETIME").
                SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP").
                SetComment("Updated at."),
        ).
        AppendConstraint("", COL_SERP_ID, c.RefSerpTableName, mySerp.COL_ID, true, true).
        SetPrimaryKeys([]string{COL_ID}).
        SetIndexKeys([]string{COL_RANKING, COL_DOMAIN, COL_TYPE}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
