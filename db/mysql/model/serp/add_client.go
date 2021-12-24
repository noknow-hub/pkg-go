//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type AddClient struct {
    BaseClient *myQuery.InsertClient
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Generate an ID.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) GenerateId() string {
    return myQuery.GenerateId(c.BaseClient.TableName, COL_ID, c.BaseClient.Db, c.BaseClient.Tx, c.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) Run() (*myQuery.InsertResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithAll(id, keyword string, objectId, countryCode, langCode, device, totalResults, searchEngine, searchEngineType, numOfSearchesForKeyword *string) (*myQuery.InsertResult, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, COL_ID)
    vals = append(vals, id)
    cols = append(cols, COL_KEYWORD)
    vals = append(vals, keyword)
    if objectId != nil {
        cols = append(cols, COL_OBJECT_ID)
        vals = append(vals, *objectId)
    }
    if countryCode != nil {
        cols = append(cols, COL_COUNTRY_CODE)
        vals = append(vals, *countryCode)
    }
    if langCode != nil {
        cols = append(cols, COL_LANG_CODE)
        vals = append(vals, *langCode)
    }
    if device != nil {
        cols = append(cols, COL_DEVICE)
        vals = append(vals, *device)
    }
    if totalResults != nil {
        cols = append(cols, COL_TOTAL_RESULTS)
        vals = append(vals, *totalResults)
    }
    if searchEngine != nil {
        cols = append(cols, COL_SEARCH_ENGINE)
        vals = append(vals, *searchEngine)
    }
    if searchEngineType != nil {
        cols = append(cols, COL_SEARCH_TYPE)
        vals = append(vals, *searcType)
    }
    if numOfSearchesForKeyword != nil {
        cols = append(cols, COL_NUM_OF_SEARCHES_FOR_KEYWORD)
        vals = append(vals, *numOfSearchesForKeyword)
    }

    c.BaseClient.
        SetColNames(cols).
        AppendValues(vals)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with required.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithRequired(id, keyword string) (*myQuery.InsertResult, error) {
    c.BaseClient.
        SetColNames([]string{COL_ID, COL_KEYWORD}).
        AppendValues([]interface{}{id, keyword})
    return c.BaseClient.Run()
}
