//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package article

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
func (o *AddClient) GenerateId() string {
    return myQuery.GenerateId(o.BaseClient.TableName, COL_ID, o.BaseClient.Db, o.BaseClient.Tx, o.BaseClient.Ctx)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) Run() (*myQuery.InsertResult, error) {
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) RunWithAll(id, status, title, url, text string, langCode, excerpt, thumbnailUrl, password, _type *string) (*myQuery.InsertResult, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, []string{COL_ID, COL_STATUS, COL_TITLE, COL_URL, COL_TEXT}...)
    vals = append(vals, []interface{}{id, status, title, url, text}...)

    if langCode != nil {
        cols = append(cols, COL_LANG_CODE)
        vals = append(vals, *langCode)
    }
    if excerpt != nil {
        cols = append(cols, COL_EXCERPT)
        vals = append(vals, *excerpt)
    }
    if thumbnailUrl != nil {
        cols = append(cols, COL_THUMBNAIL_URL)
        vals = append(vals, *thumbnailUrl)
    }
    if password != nil {
        cols = append(cols, COL_PASSWORD)
        vals = append(vals, *password)
    }
    if _type != nil {
        cols = append(cols, COL_TYPE)
        vals = append(vals, *_type)
    }

    o.BaseClient.
        SetColNames(cols).
        AppendValues(vals)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with required.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) RunWithRequired(id, status, title, url, text string) (*myQuery.InsertResult, error) {
    o.BaseClient.
        SetColNames([]string{COL_ID, COL_STATUS, COL_TITLE, COL_URL, COL_TEXT}).
        AppendValues([]interface{}{id, status, title, url, text})
    return o.BaseClient.Run()
}
