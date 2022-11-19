//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package post_image

import (
    "errors"
    "strconv"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)


//////////////////////////////////////////////////////////////////////
// scan
//////////////////////////////////////////////////////////////////////
func scan(row *myQuery.Row, o *PostImage) (err error) {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            o.Id, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_POST_ID {
            o.PostId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_URL {
            o.Url, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_NAME {
            o.Name, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_SIZE {
            o.Size, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_PATH {
            o.Path, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_ALT {
            o.Alt, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_TYPE {
            o.Type, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_MIME_TYPE {
            o.MimeType, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_LABEL {
            o.Label, err = myUtil.ConvertInterfaceToString(col.Value)
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


//////////////////////////////////////////////////////////////////////
// scan with Post object.
//////////////////////////////////////////////////////////////////////
func scanWithPost(row *myQuery.Row, postImageTable, postTable string, o *PostImageWithPost) (err error) {
    var specifiedTblName string
    for index, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        if len(s) > 1 {
            specifiedTblName = strings.Join(s[:len(s)-1], ".")
        }
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == postImageTable {
                    o.Id = val
                } else if specifiedTblName == postTable {
                    o.Post.Id = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    o.Id = val
                } else if index >= NUM_COLS && index < NUM_COLS + nkwMysqlModelPost.NUM_COLS {
                    o.Post.Id = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == COL_POST_ID {
            o.PostId, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_URL {
            o.Url, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_NAME {
            o.Name, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_SIZE {
            o.Size, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_PATH {
            o.Path, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_ALT {
            o.Alt, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_TYPE {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == postImageTable {
                    o.Type = val
                } else if specifiedTblName == postTable {
                    o.Post.Type = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    o.Type = val
                } else if index >= NUM_COLS && index < NUM_COLS + nkwMysqlModelPost.NUM_COLS {
                    o.Post.Type = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == COL_MIME_TYPE {
            o.MimeType, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_LABEL {
            o.Label, err = myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
        } else if col.Name == COL_CREATED_AT {
            val, err := myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == postImageTable {
                    o.CreatedAt = val
                } else if specifiedTblName == postTable {
                    o.Post.CreatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    o.CreatedAt = val
                } else if index >= NUM_COLS && index < NUM_COLS + nkwMysqlModelPost.NUM_COLS {
                    o.Post.CreatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == COL_UPDATED_AT {
            val, err := myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == postImageTable {
                    o.UpdatedAt = val
                } else if specifiedTblName == postTable {
                    o.Post.UpdatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                } 
            } else {
                if index < NUM_COLS {
                    o.UpdatedAt = val
                } else if index >= NUM_COLS && index < NUM_COLS + nkwMysqlModelPost.NUM_COLS {
                    o.Post.UpdatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
