//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package account

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myUpdateStatement "github.com/noknow-hub/pkg-go/db/mysql/query/update_statement"
)

type EditClient struct {
    BaseClient *myUpdateStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDb(tableName string, db *sql.DB) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTx(tableName string, tx *sql.Tx) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *EditClient) Run() (*myUpdateStatement.Result, error) {
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by "id".
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunById(id string) (*myUpdateStatement.Result, error) {
    o.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by "id".
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunWithAllById(currentId string, id, email, status, nickName, firstName, lastName, middleName, nationalityCode, phoneNumber, age, gender, birthday, biography, password, publishableToken, secretToken, addressCountryCode, addressPostalCode, addressCity, address, addressOption, lastLoggedIn *string) (*myUpdateStatement.Result, error) {
    if id != nil {
        o.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if email != nil {
        o.BaseClient.AssignmentList.Append(COL_EMAIL, *email)
    }
    if status != nil {
        o.BaseClient.AssignmentList.Append(COL_STATUS, *status)
    }
    if nickName != nil {
        o.BaseClient.AssignmentList.Append(COL_NICKNAME, *nickName)
    }
    if firstName != nil {
        o.BaseClient.AssignmentList.Append(COL_FIRST_NAME, *firstName)
    }
    if lastName != nil {
        o.BaseClient.AssignmentList.Append(COL_LAST_NAME, *lastName)
    }
    if middleName != nil {
        o.BaseClient.AssignmentList.Append(COL_MIDDLE_NAME, *middleName)
    }
    if nationalityCode != nil {
        o.BaseClient.AssignmentList.Append(COL_NATIONALITY_CODE, *nationalityCode)
    }
    if phoneNumber != nil {
        o.BaseClient.AssignmentList.Append(COL_PHONE_NUMBER, *phoneNumber)
    }
    if age != nil {
        o.BaseClient.AssignmentList.Append(COL_AGE, *age)
    }
    if gender != nil {
        o.BaseClient.AssignmentList.Append(COL_GENDER, *gender)
    }
    if birthday != nil {
        o.BaseClient.AssignmentList.Append(COL_BIRTHDAY, *birthday)
    }
    if biography != nil {
        o.BaseClient.AssignmentList.Append(COL_BIOGRAPHY, *biography)
    }
    if password != nil {
        o.BaseClient.AssignmentList.Append(COL_PASSWORD, *password)
    }
    if publishableToken != nil {
        o.BaseClient.AssignmentList.Append(COL_PUBLISHABLE_TOKEN, *publishableToken)
    }
    if secretToken != nil {
        o.BaseClient.AssignmentList.Append(COL_SECRET_TOKEN, *secretToken)
    }
    if addressCountryCode != nil {
        o.BaseClient.AssignmentList.Append(COL_ADDRESS_COUNTRY_CODE, *addressCountryCode)
    }
    if addressPostalCode != nil {
        o.BaseClient.AssignmentList.Append(COL_ADDRESS_POSTAL_CODE, *addressPostalCode)
    }
    if addressCity != nil {
        o.BaseClient.AssignmentList.Append(COL_ADDRESS_CITY, *addressCity)
    }
    if address != nil {
        o.BaseClient.AssignmentList.Append(COL_ADDRESS, *address)
    }
    if addressOption != nil {
        o.BaseClient.AssignmentList.Append(COL_ADDRESS_OPTION, *addressOption)
    }
    if lastLoggedIn != nil {
        o.BaseClient.AssignmentList.Append(COL_LAST_LOGGED_IN, *lastLoggedIn)
    }
    o.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return o.BaseClient.Run()
}
