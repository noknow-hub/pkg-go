//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package util

import (
    "context"
    "database/sql"
    "math/rand"
    "strconv"
    "time"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)


//////////////////////////////////////////////////////////////////////
// Generate an id for the specific table.
//////////////////////////////////////////////////////////////////////
func GenerateId(tableName, idColName string, db *sql.DB, tx *sql.Tx, ctx context.Context) string {
    id := time.Now().UnixNano()
    r := rand.New(rand.NewSource(id))
    var cnt int64
    for {
        tmpId := strconv.FormatInt(id + r.Int63n(100), 10)
        query := "SELECT COUNT(*) FROM " + tableName + " WHERE " + idColName + "=" + tmpId + " LIMIT 1"
        row, err := myUtil.QueryRow(db, tx, ctx, query, nil)
        if err != nil {
            return tmpId
        }
        if err := row.Scan(&cnt); err != nil {
            return tmpId
        }
        if cnt == 0 {
            break
        } else {
            id = id + 1
        }
    }
    return strconv.FormatInt(id, 10)
}


//////////////////////////////////////////////////////////////////////
// Convert interface to float64.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToFloat64(value interface{}) (float64, error) {
    nf := &sql.NullFloat64{}
    if err := nf.Scan(value); err != nil {
        return 0, err
    }
    return nf.Float64, nil
}


//////////////////////////////////////////////////////////////////////
// Convert interface to int64.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToInt64(value interface{}) (int64, error) {
    ni := &sql.NullInt64{}
    if err := ni.Scan(value); err != nil {
        return 0, err
    }
    return ni.Int64, nil
}


//////////////////////////////////////////////////////////////////////
// Convert interface to string.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToString(value interface{}) (string, error) {
    ns := &sql.NullString{}
    if err := ns.Scan(value); err != nil {
        return "", err
    }
    return ns.String, nil
}


//////////////////////////////////////////////////////////////////////
// Convert interface to time.Time.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToTime(value interface{}) (time.Time, error) {
    nt := &sql.NullTime{}
    if err := nt.Scan(value); err != nil {
        return time.Time{}, err
    }
    return nt.Time, nil
}
