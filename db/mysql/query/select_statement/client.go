//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package select_statement

import (
    "context"
    "database/sql"
    "strings"
    _ "github.com/go-sql-driver/mysql"
   myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

type Client struct {
    Columns []string
    Ctx context.Context
    Db *sql.DB
    Ignore bool
    InnerJoinTables []*myUtil.InnerJoinTable
    JoinCondition *myUtil.JoinCondition
    Limit int
    Offset int
    OrderBy string
    OrderDesc bool
    OrderRand bool
    OuterJoinTables []*myUtil.OuterJoinTable
    TableName string
    Tx *sql.Tx
    WhereCondition *myUtil.WhereCondition
}

type Result struct {
    RawArgs []interface{}
    RawQuery string
    Rows []*myUtil.Row
}

type ResultCount struct {
    Count int64
    RawArgs []interface{}
    RawQuery string
}

type ResultQuery struct {
    RawArgs []interface{}
    RawQuery string
    Rows *sql.Rows
}

type ResultQueryRow struct {
    RawArgs []interface{}
    RawQuery string
    Row *sql.Row
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewClientWithDb(tableName string, db *sql.DB) *Client {
    return &Client{
        Db: db,
        JoinCondition: &myUtil.JoinCondition{},
        TableName: tableName,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *Client {
    return &Client{
        Ctx: ctx,
        Db: db,
        JoinCondition: &myUtil.JoinCondition{},
        TableName: tableName,
        WhereCondition: &myUtil.WhereCondition{},
    }
}

//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewClientWithTx(tableName string, tx *sql.Tx) *Client {
    return &Client{
        JoinCondition: &myUtil.JoinCondition{},
        TableName: tableName,
        Tx: tx,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *Client {
    return &Client{
        Ctx: ctx,
        JoinCondition: &myUtil.JoinCondition{},
        TableName: tableName,
        Tx: tx,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// Append column.
//////////////////////////////////////////////////////////////////////
func (c *Client) AppendColumn(column string) *Client {
    c.Columns = append(c.Columns, column)
    return c
}


//////////////////////////////////////////////////////////////////////
// Append INNER JOIN clause.
//////////////////////////////////////////////////////////////////////
func (c *Client) AppendInnerJoinTables(sourceTableName, sourceCoulmn, toTableName, toCoulmn string) *Client {
    c.InnerJoinTables = append(c.InnerJoinTables, &myUtil.InnerJoinTable{
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
func (c *Client) AppendOuterJoinTables(sourceTableName, sourceCoulmn, toTableName, toCoulmn string, isLeft bool) *Client {
    c.OuterJoinTables = append(c.OuterJoinTables, &myUtil.OuterJoinTable{
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
func (c *Client) Count() (*ResultCount, error) {
    result := &ResultCount{}
    result.RawQuery, result.RawArgs = c.generateQueryForCount()
    row, err := myUtil.QueryRow(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
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
func (c *Client) Query() (*ResultQuery, error) {
    result := &ResultQuery{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.Rows, err = myUtil.Query(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *Client) QueryRow() (*ResultQueryRow, error) {
    result := &ResultQueryRow{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.Row, err = myUtil.QueryRow(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run() (*Result, error) {
    result := &Result{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    rows, err := myUtil.Query(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
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
        row := &myUtil.Row{}

        if err := rows.Scan(dests...); err != nil {
            return result, err
        }

        for i, col := range cols {
            row.Columns = append(row.Columns, &myUtil.Column{
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
func (c *Client) SetLimit(limit int) *Client {
    c.Limit = limit
    return c
}


//////////////////////////////////////////////////////////////////////
// Set OFFSET clause.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetOffset(offset int) *Client {
    c.Offset = offset
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER BY clause.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetOrderBy(orderBy string) *Client {
    c.OrderBy = orderBy
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER DESC.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetOrderDesc(isDesc bool) *Client {
    c.OrderDesc = isDesc
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER RAND.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetOrderRand() *Client {
    c.OrderRand = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *Client) generateQuery() (string, []interface{}) {
    args := make([]interface{}, 0)
    buf := make([]byte, 0)

    // SELECT
    if len(c.Columns) == 0 {
        buf = append(buf, "SELECT * FROM " + c.TableName...)
    } else {
        buf = append(buf, "SELECT " + strings.Join(c.Columns, ",") + " FROM " + c.TableName...)
    }

    // INNER JOIN
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForInnerJoin(c.InnerJoinTables, c.JoinCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // OUTER JOIN
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForOuterJoin(c.OuterJoinTables, c.JoinCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // WHERE
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForWhere(c.WhereCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // ORDER BY
    if c.OrderRand {
        buf = append(buf, myUtil.GenerateQueryForOrderByRand()...)
    } else {
        if tmpBuf := myUtil.GenerateQueryForOrderBy(c.OrderBy, c.OrderDesc); tmpBuf != "" {
            buf = append(buf, tmpBuf...)
        }
    }

    // LIMIT
    if tmpBuf := myUtil.GenerateQueryForLimit(c.Limit, c.Offset); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
    }

    return string(buf[:]), args
}


//////////////////////////////////////////////////////////////////////
// Generate query for count.
//////////////////////////////////////////////////////////////////////
func (c *Client) generateQueryForCount() (string, []interface{}) {
    args := make([]interface{}, 0)
    buf := make([]byte, 0)

    // SELECT
    buf = append(buf, "SELECT COUNT(*) FROM " + c.TableName...)

    // INNER JOIN
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForInnerJoin(c.InnerJoinTables, c.JoinCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // OUTER JOIN
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForOuterJoin(c.OuterJoinTables, c.JoinCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // WHERE
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForWhere(c.WhereCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    // LIMIT
    if tmpBuf := myUtil.GenerateQueryForLimit(c.Limit, c.Offset); tmpBuf != "" {
        buf = append(buf, tmpBuf...)
    }

    return string(buf[:]), args
}
