//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package serp

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
func (c *EditClient) RunWithAllById(currentId string, id, keyword, countryCode, langCode, device, totalResults, searchEngine, searchEngineType, numOfSearchesForKeyword *string) (*myUpdateStatement.Result, error) {
    if id != nil {
        c.BaseClient.AssignmentList.Append(COL_ID, *id)
    }
    if keyword != nil {
        c.BaseClient.AssignmentList.Append(COL_KEYWORD, *keyword)
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
    if searchEngineType != nil {
        c.BaseClient.AssignmentList.Append(COL_SEARCH_ENGINE_TYPE, *searchEngineType)
    }
    if numOfSearchesForKeyword != nil {
        c.BaseClient.AssignmentList.Append(COL_NUM_OF_SEARCHES_FOR_KEYWORD, *numOfSearchesForKeyword)
    }

    c.BaseClient.WhereCondition.SetWhere(COL_ID, currentId)
    return c.BaseClient.Run()
}
