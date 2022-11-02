//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package account

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type AddClient struct {
    BaseClient *myQuery.InsertClient
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Generate an ID.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) GenerateId() string {
    return myQuery.GenerateId(o.BaseClient.TableName, COL_ID, o.BaseClient.Db, o.BaseClient.Tx, o.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) Run() (*myQuery.InsertResult, error) {
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) RunWithAll(id, email string, status, nickName, firstName, lastName, middleName, nationalityCode, phoneNumber, age, gender, birthday, biography, password, publishableToken, secretToken, addressCountryCode, addressPostalCode, addressCity, address, addressOption, lastLoggedIn *string) (*myQuery.InsertResult, error) {
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
func (c *AddClient) RunWithRequired(id, email string) (*myQuery.InsertResult, error) {
    c.BaseClient.
        SetColNames([]string{COL_ID, COL_EMAIL}).
        AppendValues([]interface{}{id, email})
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with values.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithValues(values []*AddEditValues) (*myQuery.InsertResult, error) {
    c.BaseClient.SetColNames([]string{COL_ID,COL_EMAIL,COL_STATUS,COL_NICKNAME,COL_FIRST_NAME,COL_LAST_NAME,COL_MIDDLE_NAME,COL_NATIONALITY_CODE,COL_PHONE_NUMBER,COL_AGE,COL_GENDER,COL_BIRTHDAY,COL_BIOGRAPHY,COL_PASSWORD,COL_PUBLISHABLE_TOKEN,COL_SECRET_TOKEN,COL_ADDRESS_COUNTRY_CODE,COL_ADDRESS_POSTAL_CODE,COL_ADDRESS_CITY,COL_ADDRESS,COL_ADDRESS_OPTION,COL_LAST_LOGGED_IN})

    for _, o := range values {
        var vals []interface{}
        if o.Id != nil {
            vals = append(vals, *o.Id)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Email != nil {
            vals = append(vals, *o.Email)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Status != nil {
            vals = append(vals, *o.Status)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.NickName != nil {
            vals = append(vals, *o.NickName)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.FirstName != nil {
            vals = append(vals, *o.FirstName)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.LastName != nil {
            vals = append(vals, *o.LastName)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.MiddleName != nil {
            vals = append(vals, *o.MiddleName)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.NationalityCode != nil {
            vals = append(vals, *o.NationalityCode)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.PhoneNumber != nil {
            vals = append(vals, *o.PhoneNumber)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Age != nil {
            vals = append(vals, *o.Age)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Gender != nil {
            vals = append(vals, *o.Gender)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Birthday != nil {
            vals = append(vals, *o.Birthday)
        } else {
            vals = append(vals, sql.NullTime{})
        }
        if o.Biography != nil {
            vals = append(vals, *o.Biography)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Password != nil {
            vals = append(vals, *o.Password)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.PublishableToken != nil {
            vals = append(vals, *o.PublishableToken)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.SecretToken != nil {
            vals = append(vals, *o.SecretToken)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.AddressCountryCode != nil {
            vals = append(vals, *o.AddressCountryCode)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.AddressPostalCode != nil {
            vals = append(vals, *o.AddressPostalCode)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.AddressCity != nil {
            vals = append(vals, *o.AddressCity)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Address != nil {
            vals = append(vals, *o.Address)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.AddressOption != nil {
            vals = append(vals, *o.AddressOption)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.LastLoggedIn != nil {
            vals = append(vals, *o.LastLoggedIn)
        } else {
            vals = append(vals, sql.NullTime{})
        }
        c.BaseClient.AppendValues(vals)
    }

    return c.BaseClient.Run()
}
