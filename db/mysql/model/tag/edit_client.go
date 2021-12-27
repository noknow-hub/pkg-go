//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package tag

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
func (o *EditClient) Run() (*myQuery.UpdateResult, error) {
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_SLUG.
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunBySlug(slug string) (*myQuery.UpdateResult, error) {
    o.BaseClient.WhereCondition.SetWhere(COL_SLUG, slug)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by COL_SLUG.
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunWithAllBySlug(currentSlug string, slug, name, parentSlug *string) (*myQuery.UpdateResult, error) {
    if slug != nil {
        o.BaseClient.AssignmentList.Append(COL_SLUG, *slug)
    }
    if name != nil {
        o.BaseClient.AssignmentList.Append(COL_NAME, *name)
    }
    if parentSlug != nil {
        o.BaseClient.AssignmentList.Append(COL_PARENT_SLUG, *parentSlug)
    }
    o.BaseClient.WhereCondition.SetWhere(COL_SLUG, currentSlug)
    return o.BaseClient.Run()
}
