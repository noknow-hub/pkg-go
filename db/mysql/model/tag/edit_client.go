//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package article

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
// Run by "id".
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunById(id string) (*myQuery.UpdateResult, error) {
    o.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by "id".
//////////////////////////////////////////////////////////////////////
func (o *EditClient) RunWithAllById(currentId string, id, status, title, url, text, langCode, excerpt, thumbnailUrl, password, _type *string) (*myQuery.UpdateResult, error) {
    if id != nil {
        o.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if status != nil {
        o.BaseClient.AssignmentList.Append(COL_STATUS, *status)
    }
    if title != nil {
        o.BaseClient.AssignmentList.Append(COL_TITLE, *title)
    }
    if url != nil {
        o.BaseClient.AssignmentList.Append(COL_URL, *url)
    }
    if text != nil {
        o.BaseClient.AssignmentList.Append(COL_TEXT, *text)
    }
    if langCode != nil {
        o.BaseClient.AssignmentList.Append(COL_LANG_CODE, *langCode)
    }
    if excerpt != nil {
        o.BaseClient.AssignmentList.Append(COL_EXCERPT, *excerpt)
    }
    if thumbnailUrl != nil {
        o.BaseClient.AssignmentList.Append(COL_THUMBNAIL_URL, *thumbnailUrl)
    }
    if password != nil {
        o.BaseClient.AssignmentList.Append(COL_PASSWORD, *password)
    }
    if _type != nil {
        o.BaseClient.AssignmentList.Append(COL_TYPE, *_type)
    }
    o.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return o.BaseClient.Run()
}
