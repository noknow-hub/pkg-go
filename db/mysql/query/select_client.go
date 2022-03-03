//////////////////////////////////////////////////////////////////////
// select_client.go
//////////////////////////////////////////////////////////////////////
package query

import (
    "context"
    "database/sql"
    "strings"
    _ "github.com/go-sql-driver/mysql"
)

type SelectClient struct {
    Columns []string
    Ctx context.Context
    Db *sql.DB
    Ignore bool
    InnerJoinTables []*InnerJoinTable
    JoinCondition *JoinCondition
    Limit int
    Offset int
    OrderBy string
    OrderDesc bool
    OrderRand bool
    OuterJoinTables []*OuterJoinTable
    TableName string
    Tx *sql.Tx
    WhereCondition *WhereCondition
}

type SelectResult struct {
    RawArgs []interface{}
    RawQuery string
    Rows []*Row
}

type SelectResultCount struct {
    Count int64
    RawArgs []interface{}
    RawQuery string
}

type SelectResultQuery struct {
    RawArgs []interface{}
    RawQuery string
    Rows *sql.Rows
}

type SelectResultQueryRow struct {
    RawArgs []interface{}
    RawQuery string
    Row *sql.Row
}


//////////////////////////////////////////////////////////////////////
// New SelectClient with db object.
//////////////////////////////////////////////////////////////////////
func NewSelectClientWithDb(tableName string, db *sql.DB) *SelectClient {
    return &SelectClient{
        Db: db,
        JoinCondition: &JoinCondition{},
        TableName: tableName,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New SelectClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewSelectClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *SelectClient {
    return &SelectClient{
        Ctx: ctx,
        Db: db,
        JoinCondition: &JoinCondition{},
        TableName: tableName,
        WhereCondition: &WhereCondition{},
    }
}

//////////////////////////////////////////////////////////////////////
// New SelectClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewSelectClientWithTx(tableName string, tx *sql.Tx) *SelectClient {
    return &SelectClient{
        JoinCondition: &JoinCondition{},
        TableName: tableName,
        Tx: tx,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New SelectClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewSelectClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *SelectClient {
    return &SelectClient{
        Ctx: ctx,
        JoinCondition: &JoinCondition{},
        TableName: tableName,
        Tx: tx,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// Append column.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) AppendColumn(column string) *SelectClient {
    c.Columns = append(c.Columns, column)
    return c
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN clause.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) AppendInnerJoinTables(sourceTableName, sourceCoulmn, toTableName, toCoulmn string) *SelectClient {
    c.InnerJoinTables = append(c.InnerJoinTables, &InnerJoinTable{
        SourceTableName: sourceTableName,
        SourceCoulmn: sourceCoulmn,
        ToTableName: toTableName,
        ToCoulmn: toCoulmn,
    })
    return c
}


//////////////////////////////////////////////////////////////////////
// Append OUTER JOIN clause.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) AppendOuterJoinTables(sourceTableName, sourceCoulmn, toTableName, toCoulmn string, isLeft bool) *SelectClient {
    c.OuterJoinTables = append(c.OuterJoinTables, &OuterJoinTable{
        SourceTableName: sourceTableName,
        SourceCoulmn: sourceCoulmn,
        ToTableName: toTableName,
        ToCoulmn: toCoulmn,
        IsLeft: isLeft,
    })
    return c
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) Count() (*SelectResultCount, error) {
    result := &SelectResultCount{}
    result.RawQuery, result.RawArgs = c.generateQueryForCount()
    row, err := QueryRow(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    if err != nil {
        return result, err
    }

    if err := row.Scan(&result.Count); err != nil {
        return result, err
    }
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) Query() (*SelectResultQuery, error) {
    result := &SelectResultQuery{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.Rows, err = Query(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) QueryRow() (*SelectResultQueryRow, error) {
    result := &SelectResultQueryRow{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.Row, err = QueryRow(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) Run() (*SelectResult, error) {
    result := &SelectResult{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    rows, err := Query(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    defer rows.Close()
    if err != nil {
        return result, err
    }

    cols, err := rows.Columns()
    if err != nil {
        return result, err
    }
    numCols := len(cols)
    vals := make([]interface{}, numCols)
    dests := make([]interface{}, numCols)
    for i, _ := range dests {
        dests[i] = &vals[i]
    }
    for rows.Next() {
        row := &Row{
            SpecifiedColumns: c.Columns,
        }

        if err := rows.Scan(dests...); err != nil {
            return result, err
        }

        for i, col := range cols {
            row.Columns = append(row.Columns, &Column{
                Name: col,
                Value: vals[i],
            })
        }
        result.Rows = append(result.Rows, row)
    }

    return result, err
}


//////////////////////////////////////////////////////////////////////
// Set LIMIT clause.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) SetLimit(limit int) *SelectClient {
    c.Limit = limit
    return c
}


//////////////////////////////////////////////////////////////////////
// Set OFFSET clause.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) SetOffset(offset int) *SelectClient {
    c.Offset = offset
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER BY clause.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) SetOrderBy(orderBy string) *SelectClient {
    c.OrderBy = orderBy
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER DESC.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) SetOrderDesc(isDesc bool) *SelectClient {
    c.OrderDesc = isDesc
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER RAND.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) SetOrderRand() *SelectClient {
    c.OrderRand = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) generateQuery() (string, []interface{}) {
    args := make([]interface{}, 0)
    buf := make([]byte, 0)

    // SELECT
    if len(c.Columns) == 0 {
        buf = append(buf, "SELECT * FROM " + c.TableName...)
    } else {
        buf = append(buf, "SELECT " + strings.Join(c.Columns, ",") + " FROM " + c.TableName...)
    }

    // INNER JOIN
    if tmpBuf, tmpArgs := GenerateQueryForInnerJoin(c.InnerJoinTables, c.JoinCondition); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // OUTER JOIN
    if tmpBuf, tmpArgs := GenerateQueryForOuterJoin(c.OuterJoinTables, c.JoinCondition); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // WHERE
    if tmpBuf, tmpArgs := GenerateQueryForWhere(c.WhereCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // ORDER BY
    if c.OrderRand {
        buf = append(buf, GenerateQueryForOrderByRand()...)
    } else {
        if tmpBuf := GenerateQueryForOrderBy(c.OrderBy, c.OrderDesc); tmpBuf != "" {
            buf = append(buf, tmpBuf...)
        }
    }

    // LIMIT
    if tmpBuf := GenerateQueryForLimit(c.Limit, c.Offset); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
    }

    return string(buf[:]), args
}


//////////////////////////////////////////////////////////////////////
// Generate query for count.
//////////////////////////////////////////////////////////////////////
func (c *SelectClient) generateQueryForCount() (string, []interface{}) {
    args := make([]interface{}, 0)
    buf := make([]byte, 0)

    // SELECT
    buf = append(buf, "SELECT COUNT(*) FROM " + c.TableName...)

    // INNER JOIN
    if tmpBuf, tmpArgs := GenerateQueryForInnerJoin(c.InnerJoinTables, c.JoinCondition); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // OUTER JOIN
    if tmpBuf, tmpArgs := GenerateQueryForOuterJoin(c.OuterJoinTables, c.JoinCondition); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // WHERE
    if tmpBuf, tmpArgs := GenerateQueryForWhere(c.WhereCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // LIMIT
    if tmpBuf := GenerateQueryForLimit(c.Limit, c.Offset); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
    }

    return string(buf[:]), args
}
