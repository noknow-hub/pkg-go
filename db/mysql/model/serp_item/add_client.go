//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package serp_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type AddClient struct {
    BaseClient *myQuery.InsertClient
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Generate an ID.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) GenerateId() string {
    return myQuery.GenerateId(c.BaseClient.TableName, COL_ID, c.BaseClient.Db, c.BaseClient.Tx, c.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) Run() (*myQuery.InsertResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithAll(id, serpId, ranking, domain, url, title string, groupRanking, _type, timestamp, snippet, relatedUrl, cacheUrl *string) (*myQuery.InsertResult, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, COL_ID, COL_SERP_ID, COL_RANKING, COL_DOMAIN, COL_URL, COL_TITLE)
    vals = append(vals, id, serpId, ranking, domain, url, title)
    if groupRanking != nil {
        cols = append(cols, COL_GROUP_RANKING)
        vals = append(vals, *groupRanking)
    }
    if _type != nil {
        cols = append(cols, COL_TYPE)
        vals = append(vals, *_type)
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
func (c *AddClient) RunWithRequired(id, serpId, ranking, domain, url, title string) (*myQuery.InsertResult, error) {
    c.BaseClient.
        SetColNames([]string{COL_ID, COL_SERP_ID, COL_RANKING, COL_DOMAIN, COL_URL, COL_TITLE}).
        AppendValues([]interface{}{id, serpId, ranking, domain, url, title})
    return c.BaseClient.Run()
}
