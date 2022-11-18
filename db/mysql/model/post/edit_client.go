//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package post

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
    if values.ParentId != nil {
        c.BaseClient.AssignmentList.Append(COL_PARENT_ID, *values.ParentId)
    }
    if values.Status != nil {
        c.BaseClient.AssignmentList.Append(COL_STATUS, *values.Status)
    }
    if values.Type != nil {
        c.BaseClient.AssignmentList.Append(COL_TYPE, *values.Type)
    }
    if values.LangCode != nil {
        c.BaseClient.AssignmentList.Append(COL_LANG_CODE, *values.LangCode)
    }
    if values.CountryCode != nil {
        c.BaseClient.AssignmentList.Append(COL_COUNTRY_CODE, *values.CountryCode)
    }
    if values.Text != nil {
        c.BaseClient.AssignmentList.Append(COL_TEXT, *values.Text)
    }
    if values.CreatedAt != nil {
        c.BaseClient.AssignmentList.Append(COL_CREATED_AT, *values.CreatedAt)
    }
    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
