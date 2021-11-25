//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package organic

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
)

type ReadClient struct {
    BaseClient *myQuery.SelectClient
}

type ReadClientWithSerp struct {
    BaseClient *myQuery.SelectClient
    RefTable string
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDb(tableName string, db *sql.DB) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTx(tableName string, tx *sql.Tx) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithTxContext(tableName, tx, ctx),
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
func (c *ReadClient) Query() (*myQuery.SelectResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Run() (*Organic, *myQuery.SelectResult, error) {
    var organic *Organic
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return organic, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return organic, result, err
    }
    organic = &Organic{}
    if err := scanOrganic(result.Rows[0], organic); err != nil {
        return organic, result, err
    }
    return organic, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithSerp) RunInJoin() (*SerpOrganic, *myQuery.SelectResult, error) {
    var serpOrganic *SerpOrganic
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_SERP_ID, c.RefTable, mySerp.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpOrganic, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return serpOrganic, result, err
    }
    serp := &mySerp.Serp{}
    organic := &Organic{}
    if err := scanSerpOrganic(result.Rows[0], c.BaseClient.TableName, c.RefTable, organic, serp); err != nil {
        return serpOrganic, result, err
    }
    serpOrganic = &SerpOrganic{
        Serp: serp,
        Organics: []*Organic{organic},
    }
    return serpOrganic, result, nil
}
