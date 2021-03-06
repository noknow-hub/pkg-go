//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan serp object.
//////////////////////////////////////////////////////////////////////
func scanSerp(row *myQuery.Row, serp *Serp) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.Id = val
            }
        } else if col.Name == COL_KEYWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.Keyword = val
            }
        } else if col.Name == COL_OBJECT_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.ObjectId = val
            }
        } else if col.Name == COL_COUNTRY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.CountryCode = val
            }
        } else if col.Name == COL_LANG_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.LangCode = val
            }
        } else if col.Name == COL_DEVICE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.Device = val
            }
        } else if col.Name == COL_TOTAL_RESULTS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.TotalResults = val
            }
        } else if col.Name == COL_SEARCH_ENGINE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.SearchEngine = val
            }
        } else if col.Name == COL_SEARCH_TYPE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.SearchType = val
            }
        } else if col.Name == COL_NUM_OF_SEARCHES_FOR_KEYWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.NumOfSearchesForKeyword = val
            }
        } else if col.Name == COL_CREATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                serp.CreatedAt = val
            }
        } else if col.Name == COL_UPDATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                serp.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
