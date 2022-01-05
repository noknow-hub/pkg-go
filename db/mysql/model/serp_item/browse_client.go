//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package serp_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
)

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}

type BrowseClientWithSerp struct {
    BaseClient *myQuery.SelectClient
    RefSerpTableName string
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDb(tableName string, db *sql.DB) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTx(tableName string, tx *sql.Tx) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with reference serp table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithSerp(refSerpTable string) *BrowseClientWithSerp {
    return &BrowseClientWithSerp{
        BaseClient: c.BaseClient,
        RefSerpTableName: refSerpTable,
    }
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Query() (*myQuery.SelectResultQuery, error) {
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Run() ([]*SerpItem, *myQuery.SelectResult, error) {
    var serpItems []*SerpItem
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpItems, result, err
    }

    for _, row := range result.Rows {
        serpItem := &SerpItem{}
        if err := scanSerpItem(row, serpItem); err != nil {
            return serpItems, result, err
        }
        serpItems = append(serpItems, serpItem)
    }

    return serpItems, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithSerp) RunInJoin() ([]*SerpItemWithSerp, *myQuery.SelectResult, error) {
    var serpItems []*SerpItemWithSerp
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_SERP_ID, c.RefSerpTableName, mySerp.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpItems, result, err
    }

    for _, row := range result.Rows {
        serpItem := &SerpItemWithSerp{
            Serp: &mySerp.Serp{},
        }
        if err := scanSerpItemWithSerp(row, c.BaseClient.TableName, c.RefSerpTableName, serpItem); err != nil {
            return serpItems, result, err
        }
        serpItems = append(serpItems, serpItem)
    }

    return serpItems, result, nil
}
