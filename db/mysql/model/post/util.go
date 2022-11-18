//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package post

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// scan
//////////////////////////////////////////////////////////////////////
func scan(row *myQuery.Row, o *Post) (err error) {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            o.Id, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_PARENT_ID {
            o.ParentId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_STATUS {
            o.Status, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_TYPE {
            o.Type, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_LANG_CODE {
            o.LangCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_COUNTRY_CODE {
            o.CountryCode, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_TEXT {
            o.Text, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_CREATED_AT {
            o.CreatedAt, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_UPDATED_AT {
            o.UpdatedAt, err = myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
