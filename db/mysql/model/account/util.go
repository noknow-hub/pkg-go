//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package account

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan account object.
//////////////////////////////////////////////////////////////////////
func scanAccount(row *myQuery.Row, account *Account) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Id = val
            }
        } else if col.Name == COL_EMAIL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Email = val
            }
        } else if col.Name == COL_STATUS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Status = val
            }
        } else if col.Name == COL_FIRST_NAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.FirstName = val
            }
        } else if col.Name == COL_LAST_NAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.LastName = val
            }
        } else if col.Name == COL_MIDDLE_NAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.MiddleName = val
            }
        } else if col.Name == COL_NICKNAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.NickName = val
            }
        } else if col.Name == COL_NATIONALITY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.NationalityCode = val
            }
        } else if col.Name == COL_PHONE_NUMBER {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.PhoneNumber = val
            }
        } else if col.Name == COL_AGE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Age = val
            }
        } else if col.Name == COL_GENDER {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Gender = val
            }
        } else if col.Name == COL_BIRTHDAY {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                account.Birthday = val
            }
        } else if col.Name == COL_PASSWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Password = val
            }
        } else if col.Name == COL_PUBLISHABLE_TOKEN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.PublishableToken = val
            }
        } else if col.Name == COL_SECRET_TOKEN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.SecretToken = val
            }
        } else if col.Name == COL_ADDRESS_COUNTRY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.AddressCountryCode = val
            }
        } else if col.Name == COL_ADDRESS_POSTAL_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.AddressPostalCode = val
            }
        } else if col.Name == COL_ADDRESS_CITY {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.AddressCity = val
            }
        } else if col.Name == COL_ADDRESS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Address = val
            }
        } else if col.Name == COL_ADDRESS_OPTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.AddressOption = val
            }
        } else if col.Name == COL_BIOGRAPHY {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                account.Biography = val
            }
        } else if col.Name == COL_LAST_LOGGED_IN {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                account.LastLoggedIn = val
            }
        } else if col.Name == COL_CREATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                account.CreatedAt = val
            }
        } else if col.Name == COL_UPDATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                account.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
