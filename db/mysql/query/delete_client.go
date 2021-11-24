//////////////////////////////////////////////////////////////////////
// delete_client.go
//////////////////////////////////////////////////////////////////////
package query

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DeleteClient struct {
    Ctx context.Context
    Db *sql.DB
    Tx *sql.Tx
    Ignore bool
    Limit int
    TableName string
    WhereCondition *WhereCondition
}

type DeleteResult struct {
    RawQuery string
    RawArgs []interface{}
    SqlResult sql.Result
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDb(tableName string, db *sql.DB) *DeleteClient {
    return &DeleteClient{
        Db: db,
        TableName: tableName,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        Ctx: ctx,
        Db: db,
        TableName: tableName,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTx(tableName string, tx *sql.Tx) *DeleteClient {
    return &DeleteClient{
        TableName: tableName,
        Tx: tx,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        Ctx: ctx,
        TableName: tableName,
        Tx: tx,
        WhereCondition: &WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) Run() (*DeleteResult, error) {
    result := &DeleteResult{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.SqlResult, err = Exec(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Set IGNORE clause.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) SetIgnore() *DeleteClient {
    c.Ignore = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) generateQuery() (string, []interface{}) {
    args := make([]interface{}, 0)
    buf := make([]byte, 0)

    // DELETE
    if c.Ignore {
        buf = append(buf, "DELETE IGNORE FROM " + c.TableName...)
    } else {
        buf = append(buf, "DELETE FROM " + c.TableName...)
    }

    if tmpBuf, tmpArgs := GenerateQueryForWhere(c.WhereCondition); tmpBuf != "" && len(tmpArgs) > 0 {
        buf = append(buf, tmpBuf...)
        args = append(args, tmpArgs...)
    }

    return string(buf[:]), args
}
