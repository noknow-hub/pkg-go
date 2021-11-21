//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package organic_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    mySelectStatement "github.com/noknow-hub/pkg-go/db/mysql/query/select_statement"
    mySerp "github.com/noknow-hub/pkg-go/db/mysql/model/serp"
)

type BrowseClient struct {
    BaseClient *mySelectStatement.Client
}

type BrowseClientWithSerp struct {
    BaseClient *mySelectStatement.Client
    RefTable string
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDb(tableName string, db *sql.DB) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTx(tableName string, tx *sql.Tx) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: mySelectStatement.NewClientWithTxContext(tableName, tx, ctx),
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
func (c *BrowseClient) Count() (int64, *mySelectStatement.ResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Query() (*mySelectStatement.ResultQuery, error) {
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) QueryRow() (*mySelectStatement.ResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Run() ([]*OrganicItem, *mySelectStatement.Result, error) {
    var organicItems []*OrganicItem
    result, err := c.BaseClient.Run()
    if err != nil {
        return organicItems, result, err
    }

    for _, row := range result.Rows {
        organicItem := &OrganicItem{}
        if err := scanOrganicItem(row, organicItem); err != nil {
            return organicItems, result, err
        }
        organicItems = append(organicItems, organicItem)
    }

    return organicItems, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithSerp) RunInJoin() ([]*SerpOrganicItem, *mySelectStatement.Result, error) {
    var serpOrganicItems []*SerpOrganicItem
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_SERP_ID, c.RefTable, mySerp.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return serpOrganicItems, result, err
    }

    for _, row := range result.Rows {
        serp := &mySerp.Serp{}
        organicItem := &OrganicItem{}
        if err := scanSerpOrganicItem(row, c.BaseClient.TableName, c.RefTable, organicItem, serp); err != nil {
            return serpOrganicItems, result, err
        }
        isSerp := false
        for i, o := range serpOrganicItems {
            if o.Serp.Id == serp.Id {
                isSerp = true
                serpOrganicItems[i].OrganicItems = append(serpOrganicItems[i].OrganicItems, organicItem)
                break
            }
        }
        if !isSerp {
            serpOrganicItem := &SerpOrganicItem{
                Serp: serp,
                OrganicItems: []*OrganicItem{organicItem},
            }
            serpOrganicItems = append(serpOrganicItems, serpOrganicItem)
        }
    }

    return serpOrganicItems, result, nil
}
