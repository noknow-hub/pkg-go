//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package article_tag_map

import (
    "errors"
    "strings"
    myModelUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelTag "github.com/noknow-hub/pkg-go/db/mysql/model/tag"
)


//////////////////////////////////////////////////////////////////////
// Scan ArticleTagMap object.
//////////////////////////////////////////////////////////////////////
func scanArticleTagMap(row *myQuery.Row, articleTagMap *ArticleTagMap) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ARTICLE_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.ArticleId = val
            }
        } else if col.Name == COL_TAG_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.TagId = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Scan ArticleTagMap with article and tag  object.
//////////////////////////////////////////////////////////////////////
func scanArticleTagMapWithArticleAndTag(row *myQuery.Row, mapTable, articletable, tagTable string, articleTagMap *ArticleTagMap) error {
    var specifiedTblName string
    for index, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        if len(s) > 1 {
            specifiedTblName = strings.Join(s[:len(s)-1], ".")
        }
        col.Name = s[len(s)-1]
    
        if col.Name == COL_ARTICLE_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.ArticleId = val
            }
        } else if col.Name == COL_TAG_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.TagId = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_ID {
            val, err := myModelUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == articletable {
                    articleTagMap.Article.Id = val
                } else if specifiedTblName == tagTable {
                    articleTagMap.Tag.Id = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index >= NUM_COLS && index < nkwMysqlModelArticle.NUM_COLS {
                    articleTagMap.Article.Id = val
                } else if index >= nkwMysqlModelArticle.NUM_COLS && index < nkwMysqlModelTag.NUM_COLS {
                    articleTagMap.Tag.Id = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == nkwMysqlModelArticle.COL_STATUS {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Status = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_TITLE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Title = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Url = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_TEXT {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Text = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_LANG_CODE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.LangCode = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_EXCERPT {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Excerpt = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_THUMBNAIL_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.ThumbnailUrl = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_PASSWORD {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Password = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_TYPE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.Type = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_CREATED_AT {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.CreatedAt = val
            }
        } else if col.Name == nkwMysqlModelArticle.COL_UPDATED_AT {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Article.UpdatedAt = val
            }
        } else if col.Name == nkwMysqlModelTag.COL_NAME {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Tag.Name = val
            }
        } else if col.Name == nkwMysqlModelTag.COL_LABEL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                articleTagMap.Tag.Label = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}