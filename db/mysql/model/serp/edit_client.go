//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package serp

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
// Run with all by "id".
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllById(currentId string, id, keyword, objectId, countryCode, langCode, device, totalResults, searchEngine, searchType, numOfSearchesForKeyword *string) (*myQuery.UpdateResult, error) {
    if id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if keyword != nil {
        c.BaseClient.AssignmentList.Append(COL_KEYWORD, *keyword)
    }
    if objectId != nil {
        c.BaseClient.AssignmentList.Append(COL_OBJECT_ID, *objectId)
    }
    if countryCode != nil {
        c.BaseClient.AssignmentList.Append(COL_COUNTRY_CODE, *countryCode)
    }
    if langCode != nil {
        c.BaseClient.AssignmentList.Append(COL_LANG_CODE, *langCode)
    }
    if device != nil {
        c.BaseClient.AssignmentList.Append(COL_DEVICE, *device)
    }
    if totalResults != nil {
        c.BaseClient.AssignmentList.Append(COL_TOTAL_RESULTS, *totalResults)
    }
    if searchEngine != nil {
        c.BaseClient.AssignmentList.Append(COL_SEARCH_ENGINE, *searchEngine)
    }
    if searchType != nil {
        c.BaseClient.AssignmentList.Append(COL_SEARCH_TYPE, *searchType)
    }
    if numOfSearchesForKeyword != nil {
        c.BaseClient.AssignmentList.Append(COL_NUM_OF_SEARCHES_FOR_KEYWORD, *numOfSearchesForKeyword)
    }

    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
