//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package country

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan
//////////////////////////////////////////////////////////////////////
func scan(row *myQuery.Row, obj *Country) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_COUNTRY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.CountryCode = val
            }
        } else if col.Name == COL_AR {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Ar = val
            }
        } else if col.Name == COL_DE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.De = val
            }
        } else if col.Name == COL_EN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.En = val
            }
        } else if col.Name == COL_ES {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Es = val
            }
        } else if col.Name == COL_FR {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Fr = val
            }
        } else if col.Name == COL_JA {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Ja = val
            }
        } else if col.Name == COL_PT {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Pt = val
            }
        } else if col.Name == COL_RU {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Ru = val
            }
        } else if col.Name == COL_ZH_CN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.ZhCn = val
            }
        } else if col.Name == COL_ZH_TW {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.ZhTw = val
            }
        } else if col.Name == COL_CONTINENT {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Continent = val
            }
        } else if col.Name == COL_STATUS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                obj.Status = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
