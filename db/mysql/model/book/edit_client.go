//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package book

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
// Run by COL_ID.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunById(id string) (*myQuery.UpdateResult, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_ID, id)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by COL_ID.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllById(currentId string, id, status, title, url, langCode, introduction, thumbnailUrl, preface, foreword, acknowledgements, authorNote, password, _type *string) (*myQuery.UpdateResult, error) {
    if id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if status != nil {
        c.BaseClient.AssignmentList.Append(COL_STATUS, *status)
    }
    if title != nil {
        c.BaseClient.AssignmentList.Append(COL_TITLE, *title)
    }
    if url != nil {
        c.BaseClient.AssignmentList.Append(COL_URL, *url)
    }
    if langCode != nil {
        c.BaseClient.AssignmentList.Append(COL_LANG_CODE, *langCode)
    }
    if introduction != nil {
        c.BaseClient.AssignmentList.Append(COL_INTRODUCTION, *introduction)
    }
    if thumbnailUrl != nil {
        c.BaseClient.AssignmentList.Append(COL_THUMBNAIL_URL, *thumbnailUrl)
    }
    if preface != nil {
        c.BaseClient.AssignmentList.Append(COL_PREFACE, *preface)
    }
    if foreword != nil {
        c.BaseClient.AssignmentList.Append(COL_FOREWORD, *foreword)
    }
    if acknowledgements != nil {
        c.BaseClient.AssignmentList.Append(COL_ACKNOWLEDGMENTS, *acknowledgements)
    }
    if authorNote != nil {
        c.BaseClient.AssignmentList.Append(COL_AUTHOR_NOTE, *authorNote)
    }
    if password != nil {
        c.BaseClient.AssignmentList.Append(COL_PASSWORD, *password)
    }
    if _type != nil {
        c.BaseClient.AssignmentList.Append(COL_TYPE, *_type)
    }
    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
