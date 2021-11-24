//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package organic

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
    RefTable string
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
func (c *BrowseClient) NewBrowseClientWithSerp(refTable string) *BrowseClientWithSerp {
    return &BrowseClientWithSerp{
        BaseClient: c.BaseClient,
        RefTable: refTable,
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
func (c *BrowseClient) Run() ([]*Organic, *myQuery.SelectResult, error) {
    var organics []*Organic
    result, err := c.BaseClient.Run()
    if err != nil {
        return organics, result, err
    }

    for _, row := range result.Rows {
        organic := &Organic{}
        if err := scanOrganic(row, organic); err != nil {
            return organics, result, err
        }
        organics = append(organics, organic)
    }

    return organics, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithSerp) RunInJoin() ([]*SerpOrganic, *myQuery.SelectResult, error) {
    var serpOrganics []*SerpOrganic
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_SERP_ID, c.RefTable, mySerp.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpOrganics, result, err
    }

    for _, row := range result.Rows {
        serp := &mySerp.Serp{}
        organic := &Organic{}
        if err := scanSerpOrganic(row, c.BaseClient.TableName, c.RefTable, organicItem, serp); err != nil {
            return serpOrganics, result, err
        }
        isSerp := false
        for i, o := range serpOrganics {
            if o.Serp.Id == serp.Id {
                isSerp = true
                serpOrganics[i].Organics = append(serpOrganics[i].Organics, organic)
                break
            }
        }
        if !isSerp {
            serpOrganic := &SerpOrganic{
                Serp: serp,
                Organics: []*OrganicItem{organic},
            }
            serpOrganics = append(serpOrganics, serpOrganic)
        }
    }

    return serpOrganics, result, nil
}
