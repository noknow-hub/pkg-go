//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package util

import (
    "database/sql"
    "reflect"
    "strconv"
    "time"
    _ "github.com/go-sql-driver/mysql"
)


//////////////////////////////////////////////////////////////////////
// Convert interface to bool.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToBool(value interface{}) (bool, error) {
    if value == nil {
        return false, nil
    }
    nb := &sql.NullBool{}
    if err := nb.Scan(value); err != nil {
        return false, err
    }
    return nb.Bool, nil
}


//////////////////////////////////////////////////////////////////////
// Convert interface to float64.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToFloat64(value interface{}) (float64, error) {
    if value == nil {
        return 0, nil
    }
    if reflect.ValueOf(value).CanFloat() {
        nf := &sql.NullFloat64{}
        if err := nf.Scan(value); err != nil {
            return 0, err
        }
        return nf.Float64, nil
    } else {
        ns := &sql.NullString{}
        if err := ns.Scan(value); err != nil {
            return 0, err
        }
        f, err := strconv.ParseFloat(ns.String, 64)
        if err != nil {
            return 0, err
        }
        return f, nil
    }
}


//////////////////////////////////////////////////////////////////////
// Convert interface to int64.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToInt64(value interface{}) (int64, error) {
    if value == nil || !reflect.ValueOf(value).CanInt() {
        return 0, nil
    }
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
    if value == nil {
        return "", nil
    }
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
    if value == nil {
        return time.Time{}, nil
    }
    nt := &sql.NullTime{}
    if err := nt.Scan(value); err != nil {
        return time.Time{}, err
    }
    return nt.Time, nil
}


//////////////////////////////////////////////////////////////////////
// Convert interface to uint64.
//////////////////////////////////////////////////////////////////////
func ConvertInterfaceToUint64(value interface{}) (uint64, error) {
    if value == nil {
        return 0, nil
    }
    ns := &sql.NullString{}
    if err := ns.Scan(value); err != nil {
        return 0, err
    }
    i, err := strconv.ParseUint(ns.String, 10, 64)
    if err != nil {
        return 0, err
    }
    return i, nil
}


