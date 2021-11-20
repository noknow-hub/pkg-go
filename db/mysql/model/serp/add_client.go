//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/model/util"
    myInsertStatement "github.com/noknow-hub/pkg-go/db/mysql/query/insert_statement"
)

type AddClient struct {
    BaseClient *myInsertStatement.Client
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myInsertStatement.NewClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Generate an ID.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) GenerateId() string {
    return myUtil.GenerateId(c.BaseClient.TableName, COL_ID, c.BaseClient.Db, c.BaseClient.Tx, c.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) Run() (*myInsertStatement.Result, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithAll(id, keyword string, countryCode, langCode, device, totalResults, searchEngine, searchEngineType, numOfSearchesForKeyword *string) (*myInsertStatement.Result, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, COL_ID)
    vals = append(vals, id)
    cols = append(cols, COL_KEYWORD)
    vals = append(vals, keyword)
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
        cols = append(cols, COL_SEARCH_ENGINE_TYPE)
        vals = append(vals, *searchEngineType)
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
func (c *AddClient) RunWithRequired(id, keyword string) (*myInsertStatement.Result, error) {
    c.BaseClient.
        SetColNames([]string{COL_ID, COL_KEYWORD}).
        AppendValues([]interface{}{id, keyword})
    return c.BaseClient.Run()
}
