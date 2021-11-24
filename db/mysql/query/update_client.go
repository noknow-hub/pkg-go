//////////////////////////////////////////////////////////////////////
// update_client.go
//////////////////////////////////////////////////////////////////////
package query

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
   myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

type UpdateClient struct {
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

type UpdateResult struct {
    HasMore bool
    RawQuery string
    RawArgs []interface{}
    SqlResult sql.Result
}


//////////////////////////////////////////////////////////////////////
// New UpdateClient with db object.
//////////////////////////////////////////////////////////////////////
func NewUpdateClientWithDb(tableName string, db *sql.DB) *UpdateClient {
    return &UpdateClient{
        AssignmentList: &myUtil.AssignmentList{},
        Db: db,
        TableName: tableName,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New UpdateClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewUpdateClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *UpdateClient {
    return &UpdateClient{
        AssignmentList: &myUtil.AssignmentList{},
        Ctx: ctx,
        Db: db,
        TableName: tableName,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New UpdateClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewUpdateClientWithTx(tableName string, tx *sql.Tx) *UpdateClient {
    return &UpdateClient{
        AssignmentList: &myUtil.AssignmentList{},
        TableName: tableName,
        Tx: tx,
        WhereCondition: &myUtil.WhereCondition{},
    }
}


//////////////////////////////////////////////////////////////////////
// New UpdateClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewUpdateClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *UpdateClient {
    return &UpdateClient{
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
func (c *UpdateClient) Run() (*UpdateResult, error) {
    result := &UpdateResult{}
    result.RawQuery, result.RawArgs = c.generateQuery()
    var err error
    result.SqlResult, err = myUtil.Exec(c.Db, c.Tx, c.Ctx, result.RawQuery, result.RawArgs)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Set IGNORE clause.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) SetIgnore() *UpdateClient {
    c.Ignore = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Set LIMIT clause.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) SetLimit(limit int) *UpdateClient {
    c.Limit = limit
    return c
}


//////////////////////////////////////////////////////////////////////
// Set OFFSET clause.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) SetOffset(offset int) *UpdateClient {
    c.Offset = offset
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER BY clause.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) SetOrderBy(orderBy string) *UpdateClient {
    c.OrderBy = orderBy
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER DESC.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) SetOrderDesc(isDesc bool) *UpdateClient {
    c.OrderDesc = isDesc
    return c
}


//////////////////////////////////////////////////////////////////////
// Set ORDER RAND.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) SetOrderRand() *UpdateClient {
    c.OrderRand = true
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *UpdateClient) generateQuery() (string, []interface{}) {
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
