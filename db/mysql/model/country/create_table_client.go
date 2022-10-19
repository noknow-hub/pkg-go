//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package country

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
            myQuery.NewColumnDefinition(COL_COUNTRY_CODE, "VARCHAR(2)").
                SetNotNull().
                SetComment("Country code of 2 digits")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_AR, "VARCHAR(191)").
                SetNotNull().
                SetComment("Arabic")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_DE, "VARCHAR(191)").
                SetNotNull().
                SetComment("German")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_EN, "VARCHAR(191)").
                SetNotNull().
                SetComment("English")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ES, "VARCHAR(191)").
                SetNotNull().
                SetComment("Spanish")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_FR, "VARCHAR(191)").
                SetNotNull().
                SetComment("French")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_JA, "VARCHAR(191)").
                SetNotNull().
                SetComment("Japanese")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PT, "VARCHAR(191)").
                SetNotNull().
                SetComment("Portuguese")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_RU, "VARCHAR(191)").
                SetNotNull().
                SetComment("Russian")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ZH_CN, "VARCHAR(191)").
                SetNotNull().
                SetComment("Chinese (Simplified Chinese)")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ZH_TW, "VARCHAR(191)").
                SetNotNull().
                SetComment("Chinese (Traditional Chinese)")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_CONTINENT, "TINYINT(1) UNSIGNED").
                SetNotNull().
                SetComment("Continent, 1: Africa, 2: Asia, 3: Europe, 4: North America, 5: South America, 6: Australia / Oceania, 7: Antarctica")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_STATUS, "TINYINT(1) UNSIGNED").
                SetNotNull().
                SetComment("Status")).
        SetPrimaryKeys([]string{COL_COUNTRY_CODE}).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
