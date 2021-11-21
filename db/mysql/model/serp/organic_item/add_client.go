//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package organic_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    myInsertStatement "github.com/noknow-hub/pkg-go/db/mysql/query/insert_statement"
)

type AddClient struct {
    BaseClient *myInsertStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Generate an ID.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) GenerateId() string {
    return myUtil.GenerateId(c.BaseClient.TableName, COL_ID, c.BaseClient.Db, c.BaseClient.Tx, c.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) Run() (*myInsertStatement.Result, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithAll(id, serpId, rank, domain, url, title string, groupRank, timestamp, snippet, relatedUrl, cacheUrl *string) (*myInsertStatement.Result, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, COL_ID, COL_SERP_ID, COL_RANK, COL_DOMAIN, COL_URL, COL_TITLE)
    vals = append(vals, id, serpId, rank, domain, url, title)
    if groupRank != nil {
        cols = append(cols, COL_GROUP_RANK)
        vals = append(vals, *groupRank)
    }
    if timestamp != nil {
        cols = append(cols, COL_TIMESTAMP)
        vals = append(vals, *timestamp)
    }
    if snippet != nil {
        cols = append(cols, COL_SNIPPET)
        vals = append(vals, *snippet)
    }
    if relatedUrl != nil {
        cols = append(cols, COL_RELATED_URL)
        vals = append(vals, *relatedUrl)
    }
    if cacheUrl != nil {
        cols = append(cols, COL_CACHE_URL)
        vals = append(vals, *cacheUrl)
    }

    c.BaseClient.
        SetColNames(cols).
        AppendValues(vals)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with required.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithRequired(id, serpId, rank, domain, url, title string) (*myInsertStatement.Result, error) {
    c.BaseClient.
        SetColNames([]string{COL_ID, COL_SERP_ID, COL_RANK, COL_DOMAIN, COL_URL, COL_TITLE}).
        AppendValues([]interface{}{id, serpId, rank, domain, url, title})
    return c.BaseClient.Run()
}
