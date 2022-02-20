//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package serp_item

import (
    "errors"
    "strconv"
    "strings"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/util"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
)


//////////////////////////////////////////////////////////////////////
// Scan SerpItem object.
//////////////////////////////////////////////////////////////////////
func scanSerpItem(row *myQuery.Row, serpItem *SerpItem) error {
    for _, col := range row.Columns {
        s := strings.Split(col.Name, ".")
        col.Name = s[len(s)-1]

        if col.Name == COL_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Id = val
            }
        } else if col.Name == COL_SERP_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.SerpId = val
            }
        } else if col.Name == COL_RANKING {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Ranking = val
            }
        } else if col.Name == COL_GROUP_RANKING {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.GroupRanking = val
            }
        } else if col.Name == COL_DOMAIN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Domain = val
            }
        } else if col.Name == COL_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Url = val
            }
        } else if col.Name == COL_TITLE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Title = val
            }
        } else if col.Name == COL_TYPE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Type = val
            }
        } else if col.Name == COL_TIMESTAMP {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                serpItem.Timestamp = val
            }
        } else if col.Name == COL_SNIPPET {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Snippet = val
            }
        } else if col.Name == COL_RELATED_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.RelatedUrl = val
            }
        } else if col.Name == COL_CACHE_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.CacheUrl = val
            }
        } else if col.Name == COL_CREATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                serpItem.CreatedAt = val
            }
        } else if col.Name == COL_UPDATED_AT {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                serpItem.UpdatedAt = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Scan SerpItem with Serp object.
//////////////////////////////////////////////////////////////////////
func scanSerpItemWithSerp(row *myQuery.Row, serpItemTable, serpTable string, serpItem *SerpItemWithSerp) error {
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
                if specifiedTblName == serpItemTable {
                    serpItem.Id = val
                } else if specifiedTblName == serpTable {
                    serpItem.Serp.Id = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    serpItem.Id = val
                } else if index >= NUM_COLS && index < NUM_COLS + mySerp.NUM_COLS {
                    serpItem.Serp.Id = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == COL_SERP_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.SerpId = val
            }
        } else if col.Name == COL_RANKING {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Ranking = val
            }
        } else if col.Name == COL_GROUP_RANKING {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.GroupRanking = val
            }
        } else if col.Name == COL_DOMAIN {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Domain = val
            }
        } else if col.Name == COL_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Url = val
            }
        } else if col.Name == COL_TITLE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Title = val
            }
        } else if col.Name == COL_TYPE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Type = val
            }
        } else if col.Name == COL_TIMESTAMP {
            if val, err := myUtil.ConvertInterfaceToTime(col.Value); err != nil {
                return err
            } else {
                serpItem.Timestamp = val
            }
        } else if col.Name == COL_SNIPPET {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Snippet = val
            }
        } else if col.Name == COL_RELATED_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.RelatedUrl = val
            }
        } else if col.Name == COL_CACHE_URL {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.CacheUrl = val
            }
        } else if col.Name == COL_CREATED_AT {
            val, err := myUtil.ConvertInterfaceToTime(col.Value)
            if err != nil {
                return err
            }
            if specifiedTblName != "" {
                if specifiedTblName == serpItemTable {
                    serpItem.CreatedAt = val
                } else if specifiedTblName == serpTable {
                    serpItem.Serp.CreatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS { 
                    serpItem.CreatedAt = val 
                } else if index >= NUM_COLS && index < NUM_COLS + mySerp.NUM_COLS {
                    serpItem.Serp.CreatedAt = val
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
                if specifiedTblName == serpItemTable {
                    serpItem.UpdatedAt = val
                } else if specifiedTblName == serpTable {
                    serpItem.Serp.UpdatedAt = val
                } else {
                    return errors.New("Unknown the table name. Table: " + specifiedTblName + ", column name: " + col.Name)
                }
            } else {
                if index < NUM_COLS {
                    serpItem.UpdatedAt = val
                } else if index >= NUM_COLS && index < NUM_COLS + mySerp.NUM_COLS {
                    serpItem.Serp.UpdatedAt = val
                } else {
                    return errors.New("Unknown the column index. Index: " + strconv.FormatInt(int64(index), 10) + ", column name: " + col.Name)
                }
            }
        } else if col.Name == mySerp.COL_KEYWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.Keyword = val
            }
        } else if col.Name == mySerp.COL_OBJECT_ID {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.ObjectId = val
            }
        } else if col.Name == mySerp.COL_COUNTRY_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.CountryCode = val
            }
        } else if col.Name == mySerp.COL_LANG_CODE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.LangCode = val
            }
        } else if col.Name == mySerp.COL_DEVICE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.Device = val
            }
        } else if col.Name == mySerp.COL_TOTAL_RESULTS {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.TotalResults = val
            }
        } else if col.Name == mySerp.COL_SEARCH_ENGINE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.SearchEngine = val
            }
        } else if col.Name == mySerp.COL_SEARCH_TYPE {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.SearchType = val
            }
        } else if col.Name == mySerp.COL_NUM_OF_SEARCHES_FOR_KEYWORD {
            if val, err := myUtil.ConvertInterfaceToString(col.Value); err != nil {
                return err
            } else {
                serpItem.Serp.NumOfSearchesForKeyword = val
            }
        } else {
            return errors.New("Unknown column. Name: " + col.Name)
        }
    }
    return nil
}
