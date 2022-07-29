//////////////////////////////////////////////////////////////////////
// util.go
//////////////////////////////////////////////////////////////////////
package query

import (
    "context"
    "database/sql"
    "math/rand"
    "strconv"
    "strings"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

type AssignmentList struct {
    Assignment []*Assignment
}

type Assignment struct {
    Column string
    Value interface{}
}

type InnerJoinTable struct {
    SourceTableName string
    SourceCoulmn string
    ToTableName string
    ToCoulmn string
}   

type OuterJoinTable struct {
    SourceTableName string
    SourceCoulmn string
    ToTableName string
    ToCoulmn string
    IsLeft bool
}

type WhereCondition struct {
    Where *WhereConditionParam
    And []*WhereConditionParam
    Or []*WhereConditionParam
    AndNestedOr []*WhereConditionParam
}

type JoinCondition struct {
    InnerJoinAnd []*WhereConditionParam
    InnerJoinAndNestedOr []*WhereConditionParam
    OuterJoinAnd []*WhereConditionParam
    OuterJoinAndNestedOr []*WhereConditionParam
}

type WhereConditionParam struct {
    BetweenEnd string
    BetweenStart string
    Column string
    Value interface{}
    LikeFlag bool
    InFlag bool
    LessThanFlag bool
    LessThanOrEqualToFlag bool
    GreaterThanFlag bool
    GreaterThanOrEqualToFlag bool
    EqualToFlag bool
    NotEqualToFlag bool
}

type Row struct {
    Columns []*Column
    SpecifiedColumns []string
}

type Column struct {
    Name string
    Value interface{}
}


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
        row, err := QueryRow(db, tx, ctx, query, nil)
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
// Generate query for INNER JOIN clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForInnerJoin(innerJoinTables []*InnerJoinTable, jc *JoinCondition) (string, []interface{}) {
    buf := make([]byte, 0)
    args := make([]interface{}, 0)

    if len(innerJoinTables) == 0 {
        return string(buf[:]), args
    }

    // INNER JOIN
    for _, o := range innerJoinTables {
        buf = append(buf, " INNER JOIN " + o.ToTableName + " ON " + o.SourceTableName + "." + o.SourceCoulmn + "=" + o.ToTableName + "." + o.ToCoulmn...)
    }

    // AND
    if len(jc.InnerJoinAnd) > 0 {
        for _, param := range jc.InnerJoinAnd {
            if param.LikeFlag {
                buf = append(buf, " AND " + param.Column + " LIKE ?"...)
            } else if param.InFlag {
                v, ok := param.Value.(string)
                if ok {
                    arr := strings.Split(v, ",")
                    buf = append(buf, " AND " + param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                    for _, vv := range arr {
                        args = append(args, vv)
                    }
                } else {
                    buf = append(buf, " AND " + param.Column + " = ?"...)
                    args = append(args, v)
                }
            } else if param.LessThanFlag {
                buf = append(buf, " AND " + param.Column + " < ?"...)
            } else if param.LessThanOrEqualToFlag {
                buf = append(buf, " AND " + param.Column + " <= ?"...)
            } else if param.GreaterThanFlag {
                buf = append(buf, " AND " + param.Column + " > ?"...)
            } else if param.GreaterThanOrEqualToFlag {
                buf = append(buf, " AND " + param.Column + " >= ?"...)
            } else if param.EqualToFlag {
                buf = append(buf, " AND " + param.Column + " = ?"...)
            } else if param.NotEqualToFlag {
                buf = append(buf, " AND " + param.Column + " != ?"...)
            } else {
                buf = append(buf, " AND " + param.Column + " = ?"...)
            }
            if !param.InFlag {
                args = append(args, param.Value)
            }
        }
    }

    // AND(XXX OR XXX)
    if len(jc.InnerJoinAndNestedOr) > 0 {
        buf = append(buf, " AND ("...)
        for i, param := range jc.InnerJoinAndNestedOr {
            if i > 0 {
                buf = append(buf, " OR "...)
            }
            if param.LikeFlag {
                buf = append(buf, param.Column + " LIKE ?"...)
            } else if param.InFlag {
                v, ok := param.Value.(string)
                if ok {
                    arr := strings.Split(v, ",")
                    buf = append(buf, param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                    for _, vv := range arr {
                        args = append(args, vv)
                    }
                } else {
                    buf = append(buf, param.Column + " = ?"...)
                    args = append(args, v)
                }
            } else if param.LessThanFlag {
                buf = append(buf, param.Column + " < ?"...)
            } else if param.LessThanOrEqualToFlag {
                buf = append(buf, param.Column + " <= ?"...)
            } else if param.GreaterThanFlag {
                buf = append(buf, param.Column + " > ?"...)
            } else if param.GreaterThanOrEqualToFlag {
                buf = append(buf, param.Column + " >= ?"...)
            } else if param.EqualToFlag {
                buf = append(buf, param.Column + " = ?"...)
            } else if param.NotEqualToFlag {
                buf = append(buf, param.Column + " != ?"...)
            } else {
                buf = append(buf, param.Column + " = ?"...)
            }
            if !param.InFlag {
                args = append(args, param.Value)
            }
        }
        buf = append(buf, ")"...)
    }

    return string(buf[:]), args
}


//////////////////////////////////////////////////////////////////////
// Generate query for LIMIT clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForLimit(limit, offset int) string {
    if limit > 0 && offset > 0 {
        return  " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
    } else if limit > 0 {
        return " LIMIT " + strconv.Itoa(limit)
    } else {
        return ""
    }
}


//////////////////////////////////////////////////////////////////////
// Generate query for ORDER BY clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForOrderBy(orderBy string, orderDesc bool) string {
    if orderBy == "" {
        return ""
    }
    if orderDesc {
        return " ORDER BY " + orderBy + " DESC"
    } else {
        return " ORDER BY " + orderBy + " ASC"
    }
}


//////////////////////////////////////////////////////////////////////
// Generate query for ORDER BY clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForOrderByRand() string {
    return " ORDER BY RAND()"
}


//////////////////////////////////////////////////////////////////////
// Generate query for OUTER JOIN clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForOuterJoin(outerJoinTables []*OuterJoinTable, jc *JoinCondition) (string, []interface{}) {
    buf := make([]byte, 0)
    args := make([]interface{}, 0)

    if len(outerJoinTables) == 0 {
        return string(buf[:]), args
    }

    // OUTER JOIN
    for _, o := range outerJoinTables {
        if o.IsLeft {
            buf = append(buf, " LEFT OUTER JOIN " + o.ToTableName + " ON " + o.SourceTableName + "." + o.SourceCoulmn + "=" + o.ToTableName + "." + o.ToCoulmn...)
        } else {
            buf = append(buf, " RIGHT OUTER JOIN " + o.ToTableName + " ON " + o.SourceTableName + "." + o.SourceCoulmn + "=" + o.ToTableName + "." + o.ToCoulmn...)
        }
    }

    // AND
    if len(jc.OuterJoinAnd) > 0 {
        for _, param := range jc.OuterJoinAnd {
            if param.LikeFlag {
                buf = append(buf, " AND " + param.Column + " LIKE ?"...)
            } else if param.InFlag {
                v, ok := param.Value.(string)
                if ok {
                    arr := strings.Split(v, ",")
                    buf = append(buf, " AND " + param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                    for _, vv := range arr {
                        args = append(args, vv)
                    }
                } else {
                    buf = append(buf, " AND " + param.Column + " = ?"...)
                    args = append(args, v)
                }
            } else if param.LessThanFlag {
                buf = append(buf, " AND " + param.Column + " < ?"...)
            } else if param.LessThanOrEqualToFlag {
                buf = append(buf, " AND " + param.Column + " <= ?"...)
            } else if param.GreaterThanFlag {
                buf = append(buf, " AND " + param.Column + " > ?"...)
            } else if param.GreaterThanOrEqualToFlag {
                buf = append(buf, " AND " + param.Column + " >= ?"...)
            } else if param.EqualToFlag {
                buf = append(buf, " AND " + param.Column + " = ?"...)
            } else if param.NotEqualToFlag {
                buf = append(buf, " AND " + param.Column + " != ?"...)
            } else {
                buf = append(buf, " AND " + param.Column + " = ?"...)
            }
            if !param.InFlag {
                args = append(args, param.Value)
            }
        }
    }

    // AND(XXX OR XXX)
    if len(jc.OuterJoinAndNestedOr) > 0 {
        buf = append(buf, " AND ("...)
        for i, param := range jc.OuterJoinAndNestedOr {
            if i > 0 {
                buf = append(buf, " OR "...)
            }
            if param.LikeFlag {
                buf = append(buf, param.Column + " LIKE ?"...)
            } else if param.InFlag {
                v, ok := param.Value.(string)
                if ok {
                    arr := strings.Split(v, ",")
                    buf = append(buf, param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                    for _, vv := range arr {
                        args = append(args, vv)
                    }
                } else {
                    buf = append(buf, param.Column + " = ?"...)
                    args = append(args, v)
                }
            } else if param.LessThanFlag {
                buf = append(buf, param.Column + " < ?"...)
            } else if param.LessThanOrEqualToFlag {
                buf = append(buf, param.Column + " <= ?"...)
            } else if param.GreaterThanFlag {
                buf = append(buf, param.Column + " > ?"...)
            } else if param.GreaterThanOrEqualToFlag {
                buf = append(buf, param.Column + " >= ?"...)
            } else if param.EqualToFlag {
                buf = append(buf, param.Column + " = ?"...)
            } else if param.NotEqualToFlag {
                buf = append(buf, param.Column + " != ?"...)
            } else {
                buf = append(buf, param.Column + " = ?"...)
            }
            if !param.InFlag {
                args = append(args, param.Value)
            }
        }
        buf = append(buf, ")"...)
    }

    return string(buf[:]), args
}


//////////////////////////////////////////////////////////////////////
// Generate query for SET clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForSet(a *AssignmentList) (string, []interface{}) {
    buf := make([]byte, 0)
    args := make([]interface{}, 0)

    if len(a.Assignment) > 0 {
        for i, v := range a.Assignment {
            if i > 0 {
                buf = append(buf, ","...)
            }
            buf = append(buf, " " + v.Column + " = ?"...)
            args = append(args, v.Value)
        }
    }
    return string(buf[:]), args
}


//////////////////////////////////////////////////////////////////////
// Generate query for WHERE clause.
//////////////////////////////////////////////////////////////////////
func GenerateQueryForWhere(wc *WhereCondition) (string, []interface{}) {
    buf := make([]byte, 0)
    args := make([]interface{}, 0)

    // WHERE
    if wc.Where != nil {
        if wc.Where.LikeFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " LIKE ?"...)
        } else if wc.Where.InFlag {
            v, ok := wc.Where.Value.(string)
            if ok {
                arr := strings.Split(v, ",")
                buf = append(buf, " WHERE " + wc.Where.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                for _, vv := range arr {
                    args = append(args, vv)
                }
            } else {
                buf = append(buf, " WHERE " + wc.Where.Column + " = ?"...)
                args = append(args, v)
            }
        } else if wc.Where.LessThanFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " < ?"...)
        } else if wc.Where.LessThanOrEqualToFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " <= ?"...)
        } else if wc.Where.GreaterThanFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " > ?"...)
        } else if wc.Where.GreaterThanOrEqualToFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " >= ?"...)
        } else if wc.Where.EqualToFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " = ?"...)
        } else if wc.Where.NotEqualToFlag {
            buf = append(buf, " WHERE " + wc.Where.Column + " != ?"...)
        } else {
            buf = append(buf, " WHERE " + wc.Where.Column + " = ?"...)
        }
        if !wc.Where.InFlag {
            args = append(args, wc.Where.Value)
        }

        // AND
        if len(wc.And) > 0 {
            for _, param := range wc.And {
                if param.LikeFlag {
                    buf = append(buf, " AND " + param.Column + " LIKE ?"...)
                } else if param.InFlag {
                    v, ok := param.Value.(string)
                    if ok {
                        arr := strings.Split(v, ",")
                        buf = append(buf, " AND " + param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                        for _, vv := range arr {
                            args = append(args, vv)
                        }
                    } else {
                        buf = append(buf, " AND " + param.Column + " = ?"...)
                        args = append(args, v)
                    }
                } else if param.LessThanFlag {
                    buf = append(buf, " AND " + param.Column + " < ?"...)
                } else if param.LessThanOrEqualToFlag {
                    buf = append(buf, " AND " + param.Column + " <= ?"...)
                } else if param.GreaterThanFlag {
                    buf = append(buf, " AND " + param.Column + " > ?"...)
                } else if param.GreaterThanOrEqualToFlag {
                    buf = append(buf, " AND " + param.Column + " >= ?"...)
                } else if param.EqualToFlag {
                    buf = append(buf, " AND " + param.Column + " = ?"...)
                } else if param.NotEqualToFlag {
                    buf = append(buf, " AND " + param.Column + " != ?"...)
                } else if param.BetweenStart != "" && param.BetweenEnd != "" {
                    buf = append(buf, " AND " + param.Column + " BETWEEN ? AND ?"...)
                } else {
                    buf = append(buf, " AND " + param.Column + " = ?"...)
                }
                if param.BetweenStart != "" && param.BetweenEnd != "" { 
                    args = append(args, param.BetweenStart)
                    args = append(args, param.BetweenEnd)
                } else if !param.InFlag {
                    args = append(args, param.Value)
                }
            }
        }

        // OR
        if len(wc.Or) > 0 {
            for _, param := range wc.Or {
                if param.LikeFlag {
                    buf = append(buf, " OR " + param.Column + " LIKE ?"...)
                } else if param.InFlag {
                    v, ok := param.Value.(string)
                    if ok {
                        arr := strings.Split(v, ",")
                        buf = append(buf, " OR " + param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                        for _, vv := range arr {
                            args = append(args, vv)
                        }
                    } else {
                        buf = append(buf, " OR " + param.Column + " = ?"...)
                        args = append(args, v)
                    }
                } else if param.LessThanFlag {
                    buf = append(buf, " OR " + param.Column + " < ?"...)
                } else if param.LessThanOrEqualToFlag {
                    buf = append(buf, " OR " + param.Column + " <= ?"...)
                } else if param.GreaterThanFlag {
                    buf = append(buf, " OR " + param.Column + " > ?"...)
                } else if param.GreaterThanOrEqualToFlag {
                    buf = append(buf, " OR " + param.Column + " >= ?"...)
                } else if param.EqualToFlag {
                    buf = append(buf, " OR " + param.Column + " = ?"...)
                } else if param.NotEqualToFlag {
                    buf = append(buf, " OR " + param.Column + " != ?"...)
                } else {
                    buf = append(buf, " OR " + param.Column + " = ?"...)
                }
                if !param.InFlag {
                    args = append(args, param.Value)
                }
            }
        }

        // AND(XXX OR XXX)
        if len(wc.AndNestedOr) > 0 {
            buf = append(buf, " AND ("...)
            for i, param := range wc.AndNestedOr {
                if i > 0 {
                    buf = append(buf, " OR "...)
                }   
                if param.LikeFlag {
                    buf = append(buf, param.Column + " LIKE ?"...)
                } else if param.InFlag {
                    v, ok := param.Value.(string)
                    if ok {
                        arr := strings.Split(v, ",")
                        buf = append(buf, param.Column + " IN (?" + strings.Repeat(",?", len(arr)-1) + ")"...)
                        for _, vv := range arr {
                            args = append(args, vv)
                        }   
                    } else {
                        buf = append(buf, param.Column + " = ?"...)
                        args = append(args, v)
                    }   
                } else if param.LessThanFlag {
                    buf = append(buf, param.Column + " < ?"...)
                } else if param.LessThanOrEqualToFlag {
                    buf = append(buf, param.Column + " <= ?"...)
                } else if param.GreaterThanFlag {
                    buf = append(buf, param.Column + " > ?"...)
                } else if param.GreaterThanOrEqualToFlag {
                    buf = append(buf, param.Column + " >= ?"...)
                } else if param.EqualToFlag {
                    buf = append(buf, param.Column + " = ?"...)
                } else if param.NotEqualToFlag {
                    buf = append(buf, param.Column + " != ?"...)
                } else {
                    buf = append(buf, param.Column + " = ?"...)
                }   
                if !param.InFlag {
                    args = append(args, param.Value)
                }   
            }   
            buf = append(buf, ")"...)
        }
    }

    return string(buf[:]), args
}


//////////////////////////////////////////////////////////////////////
// Exec.
//////////////////////////////////////////////////////////////////////
func Exec(db *sql.DB, tx *sql.Tx, ctx context.Context, rawQuery string, rawArgs []interface{}) (sql.Result, error) {
    var stmt *sql.Stmt
    var err error
    if tx != nil {
        if ctx != nil {
            stmt, err = tx.PrepareContext(ctx, rawQuery)
        } else {
            stmt, err = tx.Prepare(rawQuery)
        }
    } else {
        if ctx != nil {
            stmt, err = db.PrepareContext(ctx, rawQuery)
        } else {
            stmt, err = db.Prepare(rawQuery)
        }
    }
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    if ctx != nil {
        return stmt.ExecContext(ctx, rawArgs...)
    } else {
        return stmt.Exec(rawArgs...)
    }
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func Query(db *sql.DB, tx *sql.Tx, ctx context.Context, rawQuery string, rawArgs []interface{}) (*sql.Rows, error) {
    var stmt *sql.Stmt
    var err error
    if tx != nil {
        if ctx != nil {
            stmt, err = tx.PrepareContext(ctx, rawQuery)
        } else {
            stmt, err = tx.Prepare(rawQuery)
        }
    } else {
        if ctx != nil {
            stmt, err = db.PrepareContext(ctx, rawQuery)
        } else {
            stmt, err = db.Prepare(rawQuery)
        }
    }
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    if ctx != nil {
        return stmt.QueryContext(ctx, rawArgs...)
    } else {
        return stmt.Query(rawArgs...)
    }
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func QueryRow(db *sql.DB, tx *sql.Tx, ctx context.Context, rawQuery string, rawArgs []interface{}) (*sql.Row, error) {
    var stmt *sql.Stmt
    var err error
    if tx != nil {
        if ctx != nil {
            stmt, err = tx.PrepareContext(ctx, rawQuery)
        } else {
            stmt, err = tx.Prepare(rawQuery)
        }
    } else {
        if ctx != nil {
            stmt, err = db.PrepareContext(ctx, rawQuery)
        } else {
            stmt, err = db.Prepare(rawQuery)
        }
    }
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    var row *sql.Row
    if ctx != nil {
        row = stmt.QueryRowContext(ctx, rawArgs...)
    } else {
        row = stmt.QueryRow(rawArgs...)
    }
    return row, nil
}


//////////////////////////////////////////////////////////////////////
// AssignmentList: Append Assignment.
//////////////////////////////////////////////////////////////////////
func (a *AssignmentList) Append(column string, value interface{}) *AssignmentList {
    a.Assignment = append(a.Assignment, &Assignment{
        Column: column,
        Value: value,
    })
    return a
}


//////////////////////////////////////////////////////////////////////
// Append AND clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAnd(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND LIKE clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndLike(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, true, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND IN clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndIn(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, true, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndLessThan(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, true, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndLessThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, false, true, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndGreaterThan(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, false, false, true, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndGreaterThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, false, false, false, true, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, false, false, false, false, true, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNotEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAnd(column, value, false, false, false, false, false, false, false, true)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) appendAnd(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    wc.And = append(wc.And, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Append AND BETWEEN clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) appendAndBetween(column string, betweenStart, betweenEnd string, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    wc.And = append(wc.And, &WhereConditionParam{
        BetweenEnd: betweenEnd,
        BetweenStart: betweenStart,
        Column: column,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOr(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR LIKE clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrLike(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, true, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR IN clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrIn(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, true, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrLessThan(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, true, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrLessThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, false, true, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrGreaterThan(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, false, false, true, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrGreaterThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, false, false, false, true, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, false, false, false, false, true, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendAndNestedOrNotEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendAndNestedOr(column, value, false, false, false, false, false, false, false, true)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append AND nested OR clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) appendAndNestedOr(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    wc.AndNestedOr = append(wc.AndNestedOr, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Append OR clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOr(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR LIKE clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrLike(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, true, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR IN clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrIn(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, true, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrLessThan(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, true, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrLessThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, false, true, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrGreaterThan(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, false, false, true, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrGreaterThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, false, false, false, true, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, false, false, false, false, true, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) AppendOrNotEqualTo(column string, value interface{}) *WhereCondition {
    wc.appendOr(column, value, false, false, false, false, false, false, false, true)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Append OR clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) appendOr(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    wc.Or = append(wc.Or, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Set WHERE clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhere(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE LIKE clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereLike(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, true, false, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE IN clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereIn(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, true, false, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereLessThan(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, true, false, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereLessThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, false, true, false, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereGreaterThan(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, false, false, true, false, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereGreaterThanOrEqualTo(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, false, false, false, true, false, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereEqualTo(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, false, false, false, false, true, false)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) SetWhereNotEqualTo(column string, value interface{}) *WhereCondition {
    wc.setWhere(column, value, false, false, false, false, false, false, false, true)
    return wc
}


//////////////////////////////////////////////////////////////////////
// Set WHERE clause.
//////////////////////////////////////////////////////////////////////
func (wc *WhereCondition) setWhere(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    wc.Where = &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    }
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAnd(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND LIKE clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndLike(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, true, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND IN clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndIn(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, true, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndLessThan(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, true, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndLessThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, false, true, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndGreaterThan(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, false, false, true, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndGreaterThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, false, false, false, true, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, false, false, false, false, true, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNotEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAnd(column, value, false, false, false, false, false, false, false, true)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) appendInnerJoinAnd(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    jc.InnerJoinAnd = append(jc.InnerJoinAnd, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOr(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR LIKE clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrLike(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, true, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR IN clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrIn(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, true, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrLessThan(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, true, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrLessThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, false, true, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrGreaterThan(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, false, false, true, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrGreaterThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, false, false, false, true, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, false, false, false, false, true, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendInnerJoinAndNestedOrNotEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendInnerJoinAndNestedOr(column, value, false, false, false, false, false, false, false, true)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN AND nested OR clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) appendInnerJoinAndNestedOr(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    jc.InnerJoinAndNestedOr = append(jc.InnerJoinAndNestedOr, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAnd(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND LIKE clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndLike(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, true, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND IN clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndIn(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, true, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndLessThan(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, true, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndLessThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, false, true, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndGreaterThan(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, false, false, true, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndGreaterThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, false, false, false, true, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, false, false, false, false, true, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNotEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAnd(column, value, false, false, false, false, false, false, false, true)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) appendOuterJoinAnd(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    jc.OuterJoinAnd = append(jc.OuterJoinAnd, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOr(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR LIKE clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrLike(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, true, false, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR IN clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrIn(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, true, false, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR < (less than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrLessThan(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, true, false, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR <= (less than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrLessThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, false, true, false, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR > (greater than) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrGreaterThan(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, false, false, true, false, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR >= (greater than or equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrGreaterThanOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, false, false, false, true, false, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR = (equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, false, false, false, false, true, false)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR != (not equal to) clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) AppendOuterJoinAndNestedOrNotEqualTo(column string, value interface{}) *JoinCondition {
    jc.appendOuterJoinAndNestedOr(column, value, false, false, false, false, false, false, false, true)
    return jc
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN AND nested OR clause.
//////////////////////////////////////////////////////////////////////
func (jc *JoinCondition) appendOuterJoinAndNestedOr(column string, value interface{}, likeFlag, inFlag, lessThanFlag, lessThanOrEqualToFlag, greaterThanFlag, greaterThanOrEqualToFlag, equalToFlag, notEqualToFlag bool) {
    jc.OuterJoinAndNestedOr = append(jc.OuterJoinAndNestedOr, &WhereConditionParam{
        Column: column,
        Value: value,
        LikeFlag: likeFlag,
        InFlag: inFlag,
        LessThanFlag: lessThanFlag,
        LessThanOrEqualToFlag: lessThanOrEqualToFlag,
        GreaterThanFlag: greaterThanFlag,
        GreaterThanOrEqualToFlag: greaterThanOrEqualToFlag,
        EqualToFlag: equalToFlag,
        NotEqualToFlag: notEqualToFlag,
    })
}
