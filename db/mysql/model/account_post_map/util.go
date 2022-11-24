//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package account_post_map

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)


//////////////////////////////////////////////////////////////////////
// scan
//////////////////////////////////////////////////////////////////////
func scan(row *myQuery.Row, o *AccountPostMap) (err error) {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ACCOUNT_ID {
            o.AccountId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_POST_ID {
            o.PostId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// scan with Account object.
//////////////////////////////////////////////////////////////////////
func scanWithAccount(row *myQuery.Row, accountPostMapTable, accountTable string, o *AccountPostMapWithAccount) (err error) {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ACCOUNT_ID {
            o.AccountId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_POST_ID {
            o.PostId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ID {
            o.Account.Id, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_EMAIL {
            o.Account.Email, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_STATUS {
            o.Account.Status, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_FIRST_NAME {
            o.Account.FirstName, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_LAST_NAME {
            o.Account.LastName, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_MIDDLE_NAME {
            o.Account.MiddleName, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_NICKNAME {
            o.Account.NickName, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_NATIONALITY_CODE {
            o.Account.NationalityCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_PHONE_NUMBER {
            o.Account.PhoneNumber, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_AGE {
            o.Account.Age, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_GENDER {
            o.Account.Gender, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_BIRTHDAY {
            o.Account.Birthday, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_PASSWORD {
            o.Account.Password, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_PUBLISHABLE_TOKEN {
            o.Account.PublishableToken, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_SECRET_TOKEN {
            o.Account.SecretToken, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_COUNTRY_CODE {
            o.Account.AddressCountryCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_POSTAL_CODE {
            o.Account.AddressPostalCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_CITY {
            o.Account.AddressCity, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS {
            o.Account.Address, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_ADDRESS_OPTION {
            o.Account.AddressOption, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_BIOGRAPHY {
            o.Account.Biography, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_LAST_LOGGED_IN {
            o.Account.LastLoggedIn, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_CREATED_AT {
            o.Account.CreatedAt, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelAccount.COL_UPDATED_AT {
            o.Account.UpdatedAt, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// scan with Post object.
//////////////////////////////////////////////////////////////////////
func scanWithPost(row *myQuery.Row, accountPostMapTable, postTable string, o *AccountPostMapWithPost) (err error) {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ACCOUNT_ID {
            o.AccountId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_POST_ID {
            o.PostId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_ID {
            o.Post.Id, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_PARENT_ID {
            o.Post.ParentId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_STATUS {
            o.Post.Status, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_TYPE {
            o.Post.Type, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_LANG_CODE {
            o.Post.LangCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_COUNTRY_CODE {
            o.Post.CountryCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_TEXT {
            o.Post.Text, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_CREATED_AT {
            o.Post.CreatedAt, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == nkwMysqlModelPost.COL_UPDATED_AT {
            o.Post.UpdatedAt, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
