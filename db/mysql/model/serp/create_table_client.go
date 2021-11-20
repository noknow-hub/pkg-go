//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myCreateTableStatement "github.com/noknow-hub/pkg-go/db/mysql/query/create_table_statement"
)

type CreateTableClient struct {
    *myCreateTableStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{ myCreateTableStatement.NewClientWithDb(tableName, db) }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{ myCreateTableStatement.NewClientWithDbContext(tableName, db, ctx) }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{ myCreateTableStatement.NewClientWithTx(tableName, tx) }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{ myCreateTableStatement.NewClientWithTxContext(tableName, tx, ctx) }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myCreateTableStatement.Result, error) {
    c.
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_ID, "VARCHAR(50)").
                SetNotNull().
                SetComment("SERP ID."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_KEYWORD, "VARCHAR(200)").
                SetNotNull().
                SetComment("Keyword."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_COUNTRY_CODE, "VARCHAR(2)").
                SetComment("Country code which is with 2 digits."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_LANG_CODE, "VARCHAR(2)").
                SetComment("Language code which is with 2 digits."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_DEVICE, "VARCHAR(20)").
                SetComment("Device."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_TOTAL_RESULTS, "INT UNSIGNED").
                SetComment("Total number of search results."),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_SEARCH_ENGINE, "VARCHAR(20)").
                SetComment("Search engine. (ex. google, bing. yahoo or baidu.)"),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_SEARCH_ENGINE_TYPE, "VARCHAR(10)").
                SetComment("Search engine type. (ex. organic, images or map.)"),
        ).
        AppendColumnDefinition(
            myCreateTableStatement.NewColumnDefinition(COL_NUM_OF_SEARCHES_FOR_KEYWORD, "INT UNSIGNED").
                SetComment("Number of searches for this keyword."),
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
        SetPrimaryKeys([]string{COL_ID}).
        SetIndexKeys([]string{COL_KEYWORD, COL_COUNTRY_CODE, COL_DEVICE, COL_SEARCH_ENGINE, COL_SEARCH_ENGINE_TYPE}).
        SetComment(c.TableName + " table.")
    return c.Client.Run()
}
