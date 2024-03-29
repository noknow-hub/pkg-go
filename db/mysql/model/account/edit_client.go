//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package account

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type EditClient struct {
    BaseClient *myQuery.UpdateClient
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDb(tableName string, db *sql.DB) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTx(tableName string, tx *sql.Tx) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *EditClient) Run() (*myQuery.UpdateResult, error) {
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by "id".
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunById(id string) (*myQuery.UpdateResult, error) {
    o.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by "id".
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunWithAllById(currentId string, id, email, status, nickName, firstName, lastName, middleName, nationalityCode, phoneNumber, age, gender, birthday, biography, password, publishableToken, secretToken, addressCountryCode, addressPostalCode, addressCity, address, addressOption, lastLoggedIn *string) (*myQuery.UpdateResult, error) {
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


//////////////////////////////////////////////////////////////////////
// Run with values.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithValues(currentId string, values *AddEditValues) (*myQuery.UpdateResult, error) {
    if values.Id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *values.Id)
    }
    if values.Email != nil {
        c.BaseClient.AssignmentList.Append(COL_EMAIL, *values.Email)
    }
    if values.Status != nil {
        c.BaseClient.AssignmentList.Append(COL_STATUS, *values.Status)
    }
    if values.NickName != nil {
        c.BaseClient.AssignmentList.Append(COL_NICKNAME, *values.NickName)
    }
    if values.FirstName != nil {
        c.BaseClient.AssignmentList.Append(COL_FIRST_NAME, *values.FirstName)
    }
    if values.LastName != nil {
        c.BaseClient.AssignmentList.Append(COL_LAST_NAME, *values.LastName)
    }
    if values.MiddleName != nil {
        c.BaseClient.AssignmentList.Append(COL_MIDDLE_NAME, *values.MiddleName)
    }
    if values.NationalityCode != nil {
        c.BaseClient.AssignmentList.Append(COL_NATIONALITY_CODE, *values.NationalityCode)
    }
    if values.PhoneNumber != nil {
        c.BaseClient.AssignmentList.Append(COL_PHONE_NUMBER, *values.PhoneNumber)
    }
    if values.Age != nil {
        c.BaseClient.AssignmentList.Append(COL_AGE, *values.Age)
    }
    if values.Gender != nil {
        c.BaseClient.AssignmentList.Append(COL_GENDER, *values.Gender)
    }
    if values.Birthday != nil {
        c.BaseClient.AssignmentList.Append(COL_BIRTHDAY, *values.Birthday)
    }
    if values.Biography != nil {
        c.BaseClient.AssignmentList.Append(COL_BIOGRAPHY, *values.Biography)
    }
    if values.Password != nil {
        c.BaseClient.AssignmentList.Append(COL_PASSWORD, *values.Password)
    }
    if values.PublishableToken != nil {
        c.BaseClient.AssignmentList.Append(COL_PUBLISHABLE_TOKEN, *values.PublishableToken)
    }
    if values.SecretToken != nil {
        c.BaseClient.AssignmentList.Append(COL_SECRET_TOKEN, *values.SecretToken)
    }
    if values.AddressCountryCode != nil {
        c.BaseClient.AssignmentList.Append(COL_ADDRESS_COUNTRY_CODE, *values.AddressCountryCode)
    }
    if values.AddressPostalCode != nil {
        c.BaseClient.AssignmentList.Append(COL_ADDRESS_POSTAL_CODE, *values.AddressPostalCode)
    }
    if values.AddressCity != nil {
        c.BaseClient.AssignmentList.Append(COL_ADDRESS_CITY, *values.AddressCity)
    }
    if values.Address != nil {
        c.BaseClient.AssignmentList.Append(COL_ADDRESS, *values.Address)
    }
    if values.AddressOption != nil {
        c.BaseClient.AssignmentList.Append(COL_ADDRESS_OPTION, *values.AddressOption)
    }
    if values.LastLoggedIn != nil {
        c.BaseClient.AssignmentList.Append(COL_LAST_LOGGED_IN, *values.LastLoggedIn)
    }
    
    
    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
