//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package delete_statement

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
   myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

type Client struct {
    Ctx context.Context
    Db *sql.DB
    Tx *sql.Tx
    Ignore bool
    Limit int
    TableName string
    WhereCondition *myUtil.WhereCondition
}

type Result struct {
    RawQuery string
    RawArgs []interface{}
    SqlResult sql.Result
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewClientWithDb(tableName string, db *sql.DB) *Client {
    return &Client{
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
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *Client) generateQuery() (string, []interface{}) {
    args := make([]interface{}, 0)
    buf := make([]byte, 0)

    // DELETE
    if c.Ignore {
        buf = append(buf, "DELETE IGNORE FROM " + c.TableName...)
    } else {
        buf = append(buf, "DELETE FROM " + c.TableName...)
    }

    if tmpBuf, tmpArgs := myUtil.GenerateQueryForWhere(c.WhereCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    return string(buf[:]), args
}
