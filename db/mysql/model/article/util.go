//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package article

import (
    "errors"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan article object.
//////////////////////////////////////////////////////////////////////
func scanArticle(row *myQuery.Row, article *Article) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Id = val
            }
        } else if col.Name == COL_STATUS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Status = val
            }
        } else if col.Name == COL_TITLE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Title = val
            }
        } else if col.Name == COL_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Url = val
            }
        } else if col.Name == COL_TEXT {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Text = val
            }
        } else if col.Name == COL_LANG_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.LangCode = val
            }
        } else if col.Name == COL_EXCERPT {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Excerpt = val
            }
        } else if col.Name == COL_THUMBNAIL_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.ThumbnailUrl = val
            }
        } else if col.Name == COL_PASSWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Password = val
            }
        } else if col.Name == COL_TYPE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                article.Type = val
            }
        } else if col.Name == COL_CREATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                article.CreatedAt = val
            }
        } else if col.Name == COL_UPDATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                article.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
