//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package image

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
// Run by COL_URL.
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunByUrl(url string) (*myQuery.UpdateResult, error) {
    o.BaseClient.WhereCondition.SetWhere(COL_URL, url)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by COL_URL.
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunWithAllById(currentUrl string, url, name, size, path, alt, _type, mimeType, label, link *string) (*myQuery.UpdateResult, error) {
    if url != nil {
        o.BaseClient.AssignmentList.Append(COL_URL, *url)
    }
    if name != nil {
        o.BaseClient.AssignmentList.Append(COL_NAME, *name)
    }
    if size != nil {
        o.BaseClient.AssignmentList.Append(COL_SIZE, *size)
    }
    if path != nil {
        o.BaseClient.AssignmentList.Append(COL_PATH, *path)
    }
    if alt != nil {
        o.BaseClient.AssignmentList.Append(COL_ALT, *alt)
    }
    if _type != nil {
        o.BaseClient.AssignmentList.Append(COL_TYPE, *_type)
    }
    if mimeType != nil {
        o.BaseClient.AssignmentList.Append(COL_MIME_TYPE, *mimeType)
    }
    if label != nil {
        o.BaseClient.AssignmentList.Append(COL_LABEL, *label)
    }
    if link != nil {
        o.BaseClient.AssignmentList.Append(COL_LINK, *link)
    }
    o.BaseClient.WhereCondition.SetWhere(COL_URL, currentUrl)
    return o.BaseClient.Run()
}
