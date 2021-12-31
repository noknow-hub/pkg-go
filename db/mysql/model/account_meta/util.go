//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
)


//////////////////////////////////////////////////////////////////////
// Scan AccountMeta object.
//////////////////////////////////////////////////////////////////////
func scanAccountMeta(row *myQuery.Row, accountMeta *AccountMeta) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ACCOUNT_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.AccountId = val
            }
        } else if col.Name == COL_META_KEY {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.MetaKey = val
            }
        } else if col.Name == COL_META_VALUE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.MetaValue = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Scan AccountMeta with account object..
//////////////////////////////////////////////////////////////////////
func scanAccountMetaWithAccount(row *myQuery.Row, metaTable, accountTable string, accountMeta *AccountMetaWithAccount) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ACCOUNT_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.AccountId = val
            }
        } else if col.Name == COL_META_KEY {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.MetaKey = val
            }
        } else if col.Name == COL_META_VALUE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.MetaValue = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Id = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_EMAIL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Email = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_STATUS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Status = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_FIRST_NAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.FirstName = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_LAST_NAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.LastName = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_MIDDLE_NAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.MiddleName = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_NICKNAME {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.NickName = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_NATIONALITY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.NationalityCode = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_PHONE_NUMBER {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.PhoneNumber = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_AGE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Age = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_GENDER {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Gender = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_BIRTHDAY {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Birthday = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_PASSWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Password = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_PUBLISHABLE_TOKEN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.PublishableToken = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_SECRET_TOKEN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.SecretToken = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_COUNTRY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.AddressCountryCode = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_POSTAL_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.AddressPostalCode = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_CITY {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.AddressCity = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Address = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_OPTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.AddressOption = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_BIOGRAPHY {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.Biography = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_LAST_LOGGED_IN {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.LastLoggedIn = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_CREATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.CreatedAt = val
            }
        } else if col.Name == nkwMysqlModelAccount.COL_UPDATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                accountMeta.Account.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
