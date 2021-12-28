//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package book

import (
    "errors"
    "strings"
    myModelUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan book object.
//////////////////////////////////////////////////////////////////////
func scanBook(row *myQuery.Row, book *Book) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Id = val
            }
        } else if col.Name == COL_STATUS {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Status = val
            }
        } else if col.Name == COL_TITLE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Title = val
            }
        } else if col.Name == COL_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Url = val
            }
        } else if col.Name == COL_LANG_CODE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.LangCode = val
            }
        } else if col.Name == COL_INTRODUCTION {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Introduction = val
            }
        } else if col.Name == COL_THUMBNAIL_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.ThumbnailUrl = val
            }
        } else if col.Name == COL_PREFACE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Preface = val
            }
        } else if col.Name == COL_FOREWORD {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Foreword = val
            }
        } else if col.Name == COL_ACKNOWLEDGMENTS {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Acknowledgements = val
            }
        } else if col.Name == COL_AUTHOR_NOTE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.AuthorNote = val
            }
        } else if col.Name == COL_PASSWORD {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Password = val
            }
        } else if col.Name == COL_TYPE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                book.Type = val
            }
        } else if col.Name == COL_CREATED_AT {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                book.CreatedAt = val
            }
        } else if col.Name == COL_UPDATED_AT {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                book.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
