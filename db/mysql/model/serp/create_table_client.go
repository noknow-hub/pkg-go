//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package serp

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
                SetComment("SERP ID."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_KEYWORD, "VARCHAR(200)").
                SetNotNull().
                SetComment("Keyword."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_OBJECT_ID, "VARCHAR(200)").
                SetComment("Object ID for referencing an external app."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_COUNTRY_CODE, "VARCHAR(2)").
                SetComment("Country code which is with 2 digits."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_LANG_CODE, "VARCHAR(2)").
                SetComment("Language code which is with 2 digits."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_DEVICE, "VARCHAR(20)").
                SetComment("Device."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TOTAL_RESULTS, "INT UNSIGNED").
                SetComment("Total number of search results."),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SEARCH_ENGINE, "VARCHAR(20)").
                SetComment("Search engine. (ex. google, bing. yahoo or baidu.)"),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SEARCH_TYPE, "VARCHAR(10)").
                SetComment("Search engine type. (ex. organic, images or map.)"),
        ).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_NUM_OF_SEARCHES_FOR_KEYWORD, "INT UNSIGNED").
                SetComment("Number of searches for this keyword."),
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
        SetPrimaryKeys([]string{COL_ID}).
        SetIndexKeys([]string{COL_KEYWORD, COL_OBJECT_ID, COL_COUNTRY_CODE, COL_DEVICE, COL_SEARCH_ENGINE, COL_SEARCH_TYPE}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
