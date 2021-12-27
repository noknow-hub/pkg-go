//////////////////////////////////////////////////////////////////////
// delete_client.go
//////////////////////////////////////////////////////////////////////
package article_tag_map

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type DeleteClient struct {
    BaseClient *myQuery.DeleteClient
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDb(tableName string, db *sql.DB) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTx(tableName string, tx *sql.Tx) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewDeleteClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *DeleteClient {
    return &DeleteClient{
        BaseClient: myQuery.NewDeleteClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) Run() (*myQuery.DeleteResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_ARTICLE_ID.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) RunByAccountId(accountId string) (*myQuery.DeleteResult, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_ARTICLE_ID, accountId)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_TAG_SLUG.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) RunByTagSlug(tagSlug string) (*myQuery.DeleteResult, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_TAG_SLUG, tagSlug)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_ARTICLE_ID and COL_TAG_SLUG.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) RunByAccountIdAndTagSlug(accountId, tagSlug string) (*myQuery.DeleteResult, error) {
    c.BaseClient.WhereCondition.
        SetWhere(COL_ARTICLE_ID, accountId).
        AppendAnd(COL_TAG_SLUG, tagSlug)
    return c.BaseClient.Run()
}
