//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package account

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    myInsertStatement "github.com/noknow-hub/pkg-go/db/mysql/query/insert_statement"
)

type AddClient struct {
    BaseClient *myInsertStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Generate an ID.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) GenerateId() string {
    return myUtil.GenerateId(o.BaseClient.TableName, COL_ID, o.BaseClient.Db, o.BaseClient.Tx, o.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) Run() (*myInsertStatement.Result, error) {
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) RunWithAll(id, email string, status, nickName, firstName, lastName, middleName, nationalityCode, phoneNumber, age, gender, birthday, biography, password, publishableToken, secretToken, addressCountryCode, addressPostalCode, addressCity, address, addressOption, lastLoggedIn *string) (*myInsertStatement.Result, error) {
    var cols []string
    var vals []interface{}

    // id
    cols = append(cols, COL_ID)
    vals = append(vals, id)

    // email
    cols = append(cols, COL_EMAIL)
    vals = append(vals, email)

    // status
    if status != nil {
        cols = append(cols, COL_STATUS)
        vals = append(vals, *status)
    }

    // nickname
    if nickName != nil {
        cols = append(cols, COL_NICKNAME)
        vals = append(vals, *nickName)
    }

    // first_name
    if firstName != nil {
        cols = append(cols, COL_FIRST_NAME)
        vals = append(vals, *firstName)
    }

    // last_name
    if lastName != nil {
        cols = append(cols, COL_LAST_NAME)
        vals = append(vals, *lastName)
    }

    // middle_name
    if middleName != nil {
        cols = append(cols, COL_MIDDLE_NAME)
        vals = append(vals, *middleName)
    }

    // nationality_code
    if nationalityCode != nil {
        cols = append(cols, COL_NATIONALITY_CODE)
        vals = append(vals, *nationalityCode)
    }

    // phone_number
    if phoneNumber != nil {
        cols = append(cols, COL_PHONE_NUMBER)
        vals = append(vals, *phoneNumber)
    }

    // age
    if age != nil {
        cols = append(cols, COL_AGE)
        vals = append(vals, *age)
    }

    // gender
    if gender != nil {
        cols = append(cols, COL_GENDER)
        vals = append(vals, *gender)
    }

    // birthday
    if birthday != nil {
        cols = append(cols, COL_BIRTHDAY)
        vals = append(vals, *birthday)
    }

    // biography
    if biography != nil {
        cols = append(cols, COL_BIOGRAPHY)
        vals = append(vals, *biography)
    }

    // password
    if password != nil {
        cols = append(cols, COL_PASSWORD)
        vals = append(vals, *password)
    }

    // publishable_token
    if publishableToken != nil {
        cols = append(cols, COL_PUBLISHABLE_TOKEN)
        vals = append(vals, *publishableToken)
    }

    // secret_token
    if secretToken != nil {
        cols = append(cols, COL_SECRET_TOKEN)
        vals = append(vals, *secretToken)
    }

    // address_country_code
    if addressCountryCode != nil {
        cols = append(cols, COL_ADDRESS_COUNTRY_CODE)
        vals = append(vals, *addressCountryCode)
    }

    // address_postal_code
    if addressPostalCode != nil {
        cols = append(cols, COL_ADDRESS_POSTAL_CODE)
        vals = append(vals, *addressPostalCode)
    }

    // address_city
    if addressCity != nil {
        cols = append(cols, COL_ADDRESS_CITY)
        vals = append(vals, *addressCity)
    }

    // address
    if address != nil {
        cols = append(cols, COL_ADDRESS)
        vals = append(vals, *address)
    }

    // address_option
    if addressOption != nil {
        cols = append(cols, COL_ADDRESS_OPTION)
        vals = append(vals, *addressOption)
    }

    // last_logged_in
    if lastLoggedIn != nil {
        cols = append(cols, COL_LAST_LOGGED_IN)
        vals = append(vals, *lastLoggedIn)
    }

    o.BaseClient.
        SetColNames(cols).
        AppendValues(vals)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with required.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) RunWithRequired(id, email string) (*myInsertStatement.Result, error) {
    o.BaseClient.
        SetColNames([]string{COL_ID, COL_EMAIL}).
        AppendValues([]interface{}{id, email})
    return o.BaseClient.Run()
}
