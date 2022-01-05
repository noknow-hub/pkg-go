//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package serp_item

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
    RefSerpTableName string
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
func (c *ReadClient) NewReadClientWithSerp(refSerpTable string) *ReadClientWithSerp {
    return &ReadClientWithSerp{
        BaseClient: c.BaseClient,
        RefSerpTableName: refSerpTable,
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
func (c *ReadClient) Run() (*SerpItem, *myQuery.SelectResult, error) {
    var serpItem *SerpItem
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpItem, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return serpItem, result, err
    }
    serpItem = &SerpItem{}
    if err := scanSerpItem(result.Rows[0], serpItem); err != nil {
        return serpItem, result, err
    }
    return serpItem, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithSerp) RunInJoin() (*SerpItemWithSerp, *myQuery.SelectResult, error) {
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_SERP_ID, c.RefSerpTableName, mySerp.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return nil, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return nil, result, err
    }
    serpItem := &SerpItemWithSerp{
        Serp: &mySerp.Serp{},
    }
    if err := scanSerpItemWithSerp(result.Rows[0], c.BaseClient.TableName, c.RefSerpTableName, serpItem); err != nil {
        return nil, result, err
    }
    return serpItem, result, nil
}
