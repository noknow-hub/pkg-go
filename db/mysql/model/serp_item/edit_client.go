//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package serp_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type EditClient struct {
    BaseClient *myQuery.UpdateClient
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDb(tableName string, db *sql.DB) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTx(tableName string, tx *sql.Tx) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) Run() (*myQuery.UpdateResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by "id".
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunById(id string) (*myQuery.UpdateResult, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by COL_ID.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllById(currentId string, id, serpId, ranking, domain, url, title, groupRanking, _type, timestamp, snippet, relatedUrl, cacheUrl *string) (*myQuery.UpdateResult, error) {
    if id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if serpId != nil {
        c.BaseClient.AssignmentList.Append(COL_SERP_ID, *serpId)
    }
    if ranking != nil {
        c.BaseClient.AssignmentList.Append(COL_RANKING, *ranking)
    }
    if domain != nil {
        c.BaseClient.AssignmentList.Append(COL_DOMAIN, *domain)
    }
    if url != nil {
        c.BaseClient.AssignmentList.Append(COL_URL, *url)
    }
    if title != nil {
        c.BaseClient.AssignmentList.Append(COL_TITLE, *title)
    }
    if groupRanking != nil {
        c.BaseClient.AssignmentList.Append(COL_GROUP_RANKING, *groupRanking)
    }
    if _type != nil {
        c.BaseClient.AssignmentList.Append(COL_TYPE, *timestamp)
    }
    if timestamp != nil {
        c.BaseClient.AssignmentList.Append(COL_TIMESTAMP, *timestamp)
    }
    if snippet != nil {
        c.BaseClient.AssignmentList.Append(COL_SNIPPET, *snippet)
    }
    if relatedUrl != nil {
        c.BaseClient.AssignmentList.Append(COL_RELATED_URL, *relatedUrl)
    }
    if cacheUrl != nil {
        c.BaseClient.AssignmentList.Append(COL_CACHE_URL, *cacheUrl)
    }

    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
