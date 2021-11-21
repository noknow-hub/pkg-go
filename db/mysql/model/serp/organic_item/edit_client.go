//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package organic_item

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myUpdateStatement "github.com/noknow-hub/pkg-go/db/mysql/query/update_statement"
)

type EditClient struct {
    BaseClient *myUpdateStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDb(tableName string, db *sql.DB) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTx(tableName string, tx *sql.Tx) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myUpdateStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) Run() (*myUpdateStatement.Result, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by "id".
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunById(id string) (*myUpdateStatement.Result, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by "id".
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllById(currentId string, id, serpId, rank, domain, url, title, groupRank, timestamp, snippet, relatedUrl, cacheUrl *string) (*myUpdateStatement.Result, error) {
    if id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if serpId != nil {
        c.BaseClient.AssignmentList.Append(COL_SERP_ID, *serpId)
    }
    if rank != nil {
        c.BaseClient.AssignmentList.Append(COL_RANK, *rank)
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
    if groupRank != nil {
        c.BaseClient.AssignmentList.Append(COL_GROUP_RANK, *groupRank)
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
