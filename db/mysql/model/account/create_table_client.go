//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package account

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
func (o *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    maxVarChar := "191"
    if o.Charset == myQuery.CHARSET_UTF8 {
        maxVarChar = "255"
    }

    o.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Account ID.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_EMAIL, "VARCHAR(" + maxVarChar + ")").
                SetNotNull().
                SetComment("Email address.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_STATUS, "TINYINT(1) UNSIGNED").
                SetNotNull().
                SetDefault(VAL_STATUS_DEFAULT).
                SetComment("Status. (e.g. 0: default, 10: initial stage for personal, 11: active for personal, 19: deactive for personal, 20: initial stage for business, 21: active for business, 29: deactive for business, 30: initial stage for manager, 31: active for manager, 39: deactive for manager)")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_NICKNAME, "VARCHAR(" + maxVarChar + ")").
                SetComment("Nickname.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_FIRST_NAME, "VARCHAR(50)").
                SetComment("First name.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_LAST_NAME, "VARCHAR(50)").
                SetComment("Last name.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_MIDDLE_NAME, "VARCHAR(50)").
                SetComment("Middle name.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_NATIONALITY_CODE, "VARCHAR(2)").
                SetComment("Nationality which is country code with 2 digits.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PHONE_NUMBER, "VARCHAR(30)").
                SetComment("Phone number.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_AGE, "TINYINT(1) UNSIGNED").
                SetComment("Age.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_GENDER, "TINYINT(1) UNSIGNED").
                SetComment("Gender. (e.g. 1: male, 2: female.)")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_BIRTHDAY, "DATETIME").
                SetComment("Birthday.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_BIOGRAPHY, "VARCHAR(255)").
                SetComment("Biography.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PASSWORD, "VARCHAR(255)").
                SetComment("Password.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PUBLISHABLE_TOKEN, "VARCHAR(100)").
                SetComment("Publishable token.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SECRET_TOKEN, "VARCHAR(100)").
                SetComment("Secret token.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ADDRESS_COUNTRY_CODE, "VARCHAR(2)").
                SetComment("Address country code with 2 digits.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ADDRESS_POSTAL_CODE, "VARCHAR(20)").
                SetComment("Address postal code.")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ADDRESS_CITY, "VARCHAR(50)").
                SetComment("Address city.")).
        AppendColumnDefinition(     
            myQuery.NewColumnDefinition(COL_ADDRESS, "VARCHAR(255)").
                SetComment("Address.")).
        AppendColumnDefinition( 
            myQuery.NewColumnDefinition(COL_ADDRESS_OPTION, "VARCHAR(255)").
                SetComment("Address option.")).
        AppendColumnDefinition(  
            myQuery.NewColumnDefinition(COL_LAST_LOGGED_IN, "DATETIME").
                SetComment("Last logged in at.")).
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
        SetUniqueKeys([]string{COL_EMAIL}).
        SetIndexKeys([]string{COL_NICKNAME}).
        SetComment(o.TableName + " table.")
    return o.CreateTableClient.Run()
}
