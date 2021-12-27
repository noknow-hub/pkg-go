//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package image

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
func (o *AddClient) RunWithAll(url string, name, size, path, alt, _type, mimeType, label, link *string) (*myQuery.InsertResult, error) {
    var cols []string
    var vals []interface{}

    cols = append(cols, COL_URL)
    vals = append(vals, url)

    if name != nil {
        cols = append(cols, COL_NAME)
        vals = append(vals, *name)
    }
    if size != nil {
        cols = append(cols, COL_SIZE)
        vals = append(vals, *size)
    }
    if path != nil {
        cols = append(cols, COL_PATH)
        vals = append(vals, *path)
    }
    if alt != nil {
        cols = append(cols, COL_ALT)
        vals = append(vals, *alt)
    }
    if _type != nil {
        cols = append(cols, COL_TYPE)
        vals = append(vals, *_type)
    }
    if mimeType != nil {
        cols = append(cols, COL_MIME_TYPE)
        vals = append(vals, *mimeType)
    }
    if label != nil {
        cols = append(cols, COL_LABEL)
        vals = append(vals, *label)
    }
    if link != nil {
        cols = append(cols, COL_LINK)
        vals = append(vals, *link)
    }

    o.BaseClient.
        SetColNames(cols).
        AppendValues(vals)
    return o.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with required.
//////////////////////////////////////////////////////////////////////
func (o *AddClient) RunWithRequired(url string) (*myQuery.InsertResult, error) {
    o.BaseClient.
        SetColNames([]string{COL_URL}).
        AppendValues([]interface{}{url})
    return o.BaseClient.Run()
}
