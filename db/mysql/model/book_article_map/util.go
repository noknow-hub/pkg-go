//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package book_article_map

import (
    "errors"
    "strconv"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelBook "github.com/noknow-hub/pkg-go/db/mysql/model/book"
)


//////////////////////////////////////////////////////////////////////
// Scan BookArticleMap object.
//////////////////////////////////////////////////////////////////////
func scanBookArticleMap(row *myQuery.Row, bookArticleMap *BookArticleMap) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_BOOK_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.BookId = val
            }
        } else if col.Name == COL_ARTICLE_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.ArticleId = val
            }
        } else if col.Name == COL_PART {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Part = val
            }
        } else if col.Name == COL_CHAPTER {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Chapter = val
            }
        } else if col.Name == COL_SECTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Section = val
            }
        } else if col.Name == COL_SUB_SECTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.SubSection = val
            }
        } else if col.Name == COL_INTRODUCTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Introduction = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Scan BookArticleMap with book and article object.
//////////////////////////////////////////////////////////////////////
func scanBookArticleMapWithBookAndArticle(row *myQuery.Row, mapTable, bookTable, articletable string, bookArticleMap *BookArticleMap) error {
    var specifiedTblName string
    for index, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        if len(s) > 1 {
            specifiedTblName = strings.Join(s[:len(s)-1], ".")
        }
        col.Name = s[len(s)-1]

        if col.Name == COL_BOOK_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.BookId = val
            }
        } else if col.Name == COL_ARTICLE_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.ArticleId = val
            }
        } else if col.Name == COL_PART {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Part = val
            }
        } else if col.Name == COL_CHAPTER {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Chapter = val
            }
        } else if col.Name == COL_SECTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Section = val
            }
        } else if col.Name == COL_SUB_SECTION {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.SubSection = val
            }
        } else if col.Name == COL_INTRODUCTION {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == mapTable {
                    bookArticleMap.Introduction = val
                } else if specifiedTblName == bookTable {
                    bookArticleMap.Book.Introduction = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    bookArticleMap.Introduction = val
                } else if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Introduction = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_ID {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.Id = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.Id = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Id = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.Id = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_STATUS {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.Status = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.Status = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Status = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.Status = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_TITLE {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.Title = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.Title = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Title = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.Title = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_URL {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.Url = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.Url = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Url = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.Url = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_LANG_CODE {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.LangCode = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.LangCode = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.LangCode = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.LangCode = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_THUMBNAIL_URL {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.ThumbnailUrl = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.ThumbnailUrl = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.ThumbnailUrl = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.ThumbnailUrl = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_PREFACE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Book.Preface = val
            }
        } else if col.Name == nkwMysqlModelBook.COL_FOREWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Book.Foreword = val
            }
        } else if col.Name == nkwMysqlModelBook.COL_ACKNOWLEDGMENTS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Book.Acknowledgements = val
            }
        } else if col.Name == nkwMysqlModelBook.COL_AUTHOR_NOTE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Book.AuthorNote = val
            }
        } else if col.Name == nkwMysqlModelBook.COL_PASSWORD {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.Password = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.Password = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Password = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.Password = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_TYPE {
            val, err := myUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.Type = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.Type = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.Type = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.Type = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_CREATED_AT {
            val, err := myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.CreatedAt = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.CreatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.CreatedAt = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.CreatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelBook.COL_UPDATED_AT {
            val, err := myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == bookTable {
                    bookArticleMap.Book.UpdatedAt = val
                } else if specifiedTblName == articletable {
                    bookArticleMap.Article.UpdatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelBook.NUM_COLS {
                    bookArticleMap.Book.UpdatedAt = val
                } else if index >= nkwMysqlModelBook.NUM_COLS && index < nkwMysqlModelBook.NUM_COLS + nkwMysqlModelArticle.NUM_COLS {
                    bookArticleMap.Article.UpdatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelArticle.COL_TEXT {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Article.Text = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_EXCERPT {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                bookArticleMap.Article.Excerpt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
