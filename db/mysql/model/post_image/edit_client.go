//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package post_image

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
// Run with values.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithValues(currentId string, values *AddEditValues) (*myQuery.UpdateResult, error) {
    if values.Id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *values.Id)
    }
    if values.PostId != nil {
        c.BaseClient.AssignmentList.Append(COL_POST_ID, *values.PostId)
    }
    if values.Url != nil {
        c.BaseClient.AssignmentList.Append(COL_URL, *values.Url)
    }
    if values.Name != nil {
        c.BaseClient.AssignmentList.Append(COL_NAME, *values.Name)
    }
    if values.Size != nil {
        c.BaseClient.AssignmentList.Append(COL_SIZE, *values.Size)
    }
    if values.Path != nil {
        c.BaseClient.AssignmentList.Append(COL_PATH, *values.Path)
    }
    if values.Alt != nil {
        c.BaseClient.AssignmentList.Append(COL_ALT, *values.Alt)
    }
    if values.Type != nil {
        c.BaseClient.AssignmentList.Append(COL_TYPE, *values.Type)
    }
    if values.MimeType != nil {
        c.BaseClient.AssignmentList.Append(COL_MIME_TYPE, *values.MimeType)
    }
    if values.Label != nil {
        c.BaseClient.AssignmentList.Append(COL_LABEL, *values.Label)
    }
    if values.CreatedAt != nil {
        c.BaseClient.AssignmentList.Append(COL_CREATED_AT, *values.CreatedAt)
    }
    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
