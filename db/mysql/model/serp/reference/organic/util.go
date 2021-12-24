//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package organic

import (
    "errors"
    "strconv"
    "strings"
    myModelUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
    myQueryUtil "github.com/noknow-hub/pkg-go/db/mysql/query"
)


//////////////////////////////////////////////////////////////////////
// Scan Organic object.
//////////////////////////////////////////////////////////////////////
func scanOrganic(row *myQueryUtil.Row, organic *Organic) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Id = val
            }
        } else if col.Name == COL_SERP_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.SerpId = val
            }
        } else if col.Name == COL_RANKING {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Ranking = val
            }
        } else if col.Name == COL_GROUP_RANKING {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.GroupRanking = val
            }
        } else if col.Name == COL_DOMAIN {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Domain = val
            }
        } else if col.Name == COL_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Url = val
            }
        } else if col.Name == COL_TITLE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Title = val
            }
        } else if col.Name == COL_TIMESTAMP {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                organic.Timestamp = val
            }
        } else if col.Name == COL_SNIPPET {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Snippet = val
            }
        } else if col.Name == COL_RELATED_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.RelatedUrl = val
            }
        } else if col.Name == COL_CACHE_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.CacheUrl = val
            }
        } else if col.Name == COL_CREATED_AT {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                organic.CreatedAt = val
            }
        } else if col.Name == COL_UPDATED_AT {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                organic.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Scan Serp and Organic object.
//////////////////////////////////////////////////////////////////////
func scanSerpOrganic(row *myQueryUtil.Row, organicTable, refTable string, organic *Organic, serp *mySerp.Serp) error {
    var specifiedTblName string
    for index, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        if len(s) > 1 {
            specifiedTblName = strings.Join(s[:len(s)-1], ".")
        }
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            val, err := myModelUtil.ConvertInterfaceToString(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == organicTable {
                    organic.Id = val
                } else if specifiedTblName == refTable {
                    serp.Id = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    organic.Id = val
                } else if index >= NUM_COLS && index < mySerp.NUM_COLS {
                    serp.Id = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == COL_SERP_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.SerpId = val
            }
        } else if col.Name == COL_RANKING {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Ranking = val
            }
        } else if col.Name == COL_GROUP_RANKING {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.GroupRanking = val
            }
        } else if col.Name == COL_DOMAIN {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Domain = val
            }
        } else if col.Name == COL_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Url = val
            }
        } else if col.Name == COL_TITLE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Title = val
            }
        } else if col.Name == COL_TIMESTAMP {
            if val, err := myModelUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                organic.Timestamp = val
            }
        } else if col.Name == COL_SNIPPET {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.Snippet = val
            }
        } else if col.Name == COL_RELATED_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.RelatedUrl = val
            }
        } else if col.Name == COL_CACHE_URL {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                organic.CacheUrl = val
            }
        } else if col.Name == COL_CREATED_AT {
            val, err := myModelUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == organicTable {
                    organic.CreatedAt = val
                } else if specifiedTblName == refTable {
                    serp.CreatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    organic.CreatedAt = val
                } else if index >= NUM_COLS && index < mySerp.NUM_COLS {
                    serp.CreatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == COL_UPDATED_AT {
            val, err := myModelUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == organicTable {
                    organic.UpdatedAt = val
                } else if specifiedTblName == refTable {
                    serp.UpdatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    organic.UpdatedAt = val
                } else if index >= NUM_COLS && index < mySerp.NUM_COLS {
                    serp.UpdatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == mySerp.COL_KEYWORD {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.Keyword = val
            }
        } else if col.Name == mySerp.COL_OBJECT_ID {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.ObjectId = val
            }
        } else if col.Name == mySerp.COL_COUNTRY_CODE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.CountryCode = val
            }
        } else if col.Name == mySerp.COL_LANG_CODE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.LangCode = val
            }
        } else if col.Name == mySerp.COL_DEVICE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.Device = val
            }
        } else if col.Name == mySerp.COL_TOTAL_RESULTS {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.TotalResults = val
            }
        } else if col.Name == mySerp.COL_SEARCH_ENGINE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.SearchEngine = val
            }
        } else if col.Name == mySerp.COL_SEARCH_TYPE {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.SearchType = val
            }
        } else if col.Name == mySerp.COL_NUM_OF_SEARCHES_FOR_KEYWORD {
            if val, err := myModelUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serp.NumOfSearchesForKeyword = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
