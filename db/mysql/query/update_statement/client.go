//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package update_statement

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
   myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

type Client struct {
    AssignmentList *myUtil.AssignmentList
    Ctx context.Context
    Db *sql.DB
    Ignore bool
    Limit int
    Offset int
    OrderBy string
    OrderDesc bool
    OrderRand bool
    TableName string
    Tx *sql.Tx
    WhereCondition *myUtil.WhereCondition
}

type Result struct {
    HasMore bool
    RawQuery string
    RawArgs []interface{}
    SqlResult sql.Result
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewClientWithDb(tableName string, db *sql.DB) *Client {
    return &Client{
        AssignmentList: &myUtil.AssignmentList{},
        Db: db,
        TableName: tableName,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *Client {
    return &Client{
        AssignmentList: &myUtil.AssignmentList{},
        Ctx: ctx,
        Db: db,
        TableName: tableName,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewClientWithTx(tableName string, tx *sql.Tx) *Client {
    return &Client{
        AssignmentList: &myUtil.AssignmentList{},
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
        AssignmentList: &myUtil.AssignmentList{},
        Ctx: ctx,
        TableName: tableName,
        Tx: tx,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run() (*Result, error) {
    result := &Result{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.SqlResult, err = myUtil.Exec(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Set IGNORE clause.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetIgnore() *Client {
    c.Ignore = true
    return c
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

    // UPDATE
    if c.Ignore {
        buf = append(buf, "UPDATE IGNORE " + c.TableName + " SET"...)
    } else {
        buf = append(buf, "UPDATE " + c.TableName + " SET"...)
    }

    // Assignment list.
    if tmpBuf, tmpArgs := myUtil.GenerateQueryForSet(c.AssignmentList); tmpBuf != "" && len(tmpArgs) > 0 {
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
