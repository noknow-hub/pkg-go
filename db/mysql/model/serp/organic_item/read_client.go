//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package organic_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    mySelectStatement "github.com/noknow-hub/pkg-go/db/mysql/query/select_statement"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
)

type ReadClient struct {
    BaseClient *mySelectStatement.Client
}

type ReadClientWithSerp struct {
    BaseClient *mySelectStatement.Client
    RefTable string
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDb(tableName string, db *sql.DB) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTx(tableName string, tx *sql.Tx) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: mySelectStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with reference serp table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithSerp(refTable string) *ReadClientWithSerp {
    return &ReadClientWithSerp{
        BaseClient: c.BaseClient,
        RefTable: refTable,
    }
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Query() (*mySelectStatement.ResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) QueryRow() (*mySelectStatement.ResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Run() (*OrganicItem, *mySelectStatement.Result, error) {
    var organicItem *OrganicItem
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return organicItem, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return organicItem, result, err
    }
    if err := scanOrganicItem(result.Rows[0], organicItem); err != nil {
        return organicItem, result, err
    }
    return organicItem, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithSerp) RunInJoin() (*SerpOrganicItem, *mySelectStatement.Result, error) {
    var serpOrganicItem *SerpOrganicItem
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_SERP_ID, c.RefTable, mySerp.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpOrganicItem, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return serpOrganicItem, result, err
    }
    serp := &mySerp.Serp{}
    organicItem := &OrganicItem{}
    if err := scanSerpOrganicItem(result.Rows[0], c.BaseClient.TableName, c.RefTable, organicItem, serp); err != nil {
        return serpOrganicItem, result, err
    }
    return serpOrganicItem, result, nil
}
