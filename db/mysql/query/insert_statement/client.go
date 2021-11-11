//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package insert_statement

import (
    "context"
    "database/sql"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

type Client struct {
    ColNames []string
    Ctx context.Context
    Db *sql.DB
    Tx *sql.Tx
    Ignore bool
    TableName string
    Values [][]interface{}
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
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewClientWithTx(tableName string, tx *sql.Tx) *Client {
    return &Client{
        TableName: tableName,
        Tx: tx,
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
    }
}


//////////////////////////////////////////////////////////////////////
// Append VALUES clause.
//////////////////////////////////////////////////////////////////////
func (c *Client) AppendValues(valueList []interface{}) *Client {
    c.Values = append(c.Values, valueList)
    return c
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
// Set column names.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetColNames(colNames []string) *Client {
    c.ColNames = colNames
    return c
}


//////////////////////////////////////////////////////////////////////
// Set IGNORE clause.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetIgnore() *Client {
    c.Ignore = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Set table name.
//////////////////////////////////////////////////////////////////////
func (c *Client) SetTableName(tableName string) *Client {
    c.TableName = tableName
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *Client) generateQuery() (string, []interface{}) {
    buf := make([]byte, 0)
    args := make([]interface{}, 0)

    // INSERT
    if c.Ignore {
        buf = append(buf, "INSERT IGNORE INTO " + c.TableName...)
    } else {
        buf = append(buf, "INSERT INTO " + c.TableName...)
    }

    // COL NAMES
    if len(c.ColNames) > 0 {
        buf = append(buf, " (" + strings.Join(c.ColNames, ",") + ")"...)
    }

    // VALUES
    if len(c.Values) > 0 {
        buf = append(buf, " VALUES "...)
        for i, valueList := range c.Values {
            if i > 0 {
                buf = append(buf, ","...)
            }
            buf = append(buf, "("...)
            for ii, _ := range valueList {
                if ii > 0 {
                    buf = append(buf, ","...)
                }
                buf = append(buf, "?"...)
            }
            buf = append(buf, ")"...)
            args = append(args, valueList...)
        }
    }

    return string(buf[:]), args
}
