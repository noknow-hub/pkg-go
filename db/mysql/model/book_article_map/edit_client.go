//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package book_article_map

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
// Run with all by COL_BOOK_ID and COL_ARTICLE_ID.
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllByBookIdAndAccountId(currentBookId, currentAccountId string, bookId, articleId, part, chapter, section, subSection, introduction *string) (*myQuery.UpdateResult, error) {
    if bookId != nil {
        c.BaseClient.AssignmentList.Append(COL_BOOK_ID, *bookId)
    }
    if articleId != nil {
        c.BaseClient.AssignmentList.Append(COL_ARTICLE_ID, *articleId)
    }
    if part != nil {
        c.BaseClient.AssignmentList.Append(COL_PART, *part)
    }
    if chapter != nil {
        c.BaseClient.AssignmentList.Append(COL_CHAPTER, *chapter)
    }
    if section != nil {
        c.BaseClient.AssignmentList.Append(COL_SECTION, *section)
    }
    if subSection != nil {
        c.BaseClient.AssignmentList.Append(COL_SUB_SECTION, *subSection)
    }
    if introduction != nil {
        c.BaseClient.AssignmentList.Append(COL_INTRODUCTION, *introduction)
    }
    c.BaseClient.WhereCondition.
        SetWhere(COL_BOOK_ID, currentBookId).
        AppendAnd(COL_ARTICLE_ID, currentAccountId)
    return c.BaseClient.Run()
}
