//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package tag

import (
    "errors"
    "strings"
    myModelUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan tag object.
//////////////////////////////////////////////////////////////////////
func scanTag(row *myQuery.Row, tag *Tag) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_SLUG {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                tag.Slug = val
            }
        } else if col.Name == COL_NAME {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                tag.Name = val
            }
        } else if col.Name == COL_PARENT_SLUG {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                tag.ParentSlug = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
