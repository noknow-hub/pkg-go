//////////////////////////////////////////////////////////////////////
// insert_client.go
//////////////////////////////////////////////////////////////////////
package query

import (
    "context"
    "database/sql"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

type InsertClient struct {
    ColNames []string
    Ctx context.Context
    Db *sql.DB
    Tx *sql.Tx
    Ignore bool
    TableName string
    Values [][]interface{}
}

type InsertResult struct {
    RawQuery string
    RawArgs []interface{}
    SqlResult sql.Result
}


//////////////////////////////////////////////////////////////////////
// New InsertClient with db object.
//////////////////////////////////////////////////////////////////////
func NewInsertClientWithDb(tableName string, db *sql.DB) *InsertClient {
    return &InsertClient{
        Db: db,
        TableName: tableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New InsertClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewInsertClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *InsertClient {
    return &InsertClient{
        Ctx: ctx,
        Db: db,
        TableName: tableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New InsertClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewInsertClientWithTx(tableName string, tx *sql.Tx) *InsertClient {
    return &InsertClient{
        TableName: tableName,
        Tx: tx,
    }
}


//////////////////////////////////////////////////////////////////////
// New InsertClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewInsertClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *InsertClient {
    return &InsertClient{
        Ctx: ctx,
        TableName: tableName,
        Tx: tx,
    }
}


//////////////////////////////////////////////////////////////////////
// Append VALUES clause.
//////////////////////////////////////////////////////////////////////
func (c *InsertClient) AppendValues(valueList []interface{}) *InsertClient {
    c.Values = append(c.Values, valueList)
    return c
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *InsertClient) Run() (*InsertResult, error) {
    result := &InsertResult{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.SqlResult, err = myUtil.Exec(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Set column names.
//////////////////////////////////////////////////////////////////////
func (c *InsertClient) SetColNames(colNames []string) *InsertClient {
    c.ColNames = colNames
    return c
}


//////////////////////////////////////////////////////////////////////
// Set IGNORE clause.
//////////////////////////////////////////////////////////////////////
func (c *InsertClient) SetIgnore() *InsertClient {
    c.Ignore = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Set table name.
//////////////////////////////////////////////////////////////////////
func (c *InsertClient) SetTableName(tableName string) *InsertClient {
    c.TableName = tableName
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *InsertClient) generateQuery() (string, []interface{}) {
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
