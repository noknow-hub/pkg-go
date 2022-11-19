//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package post_image

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
// Run with values.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunWithValues(valuesList []*AddEditValues) (*myQuery.InsertResult, error) {
    c.BaseClient.SetColNames([]string{
        COL_ID,
        COL_POST_ID,
        COL_URL,
        COL_NAME,
        COL_SIZE,
        COL_PATH,
        COL_ALT,
        COL_TYPE,
        COL_MIME_TYPE,
        COL_LABEL,
        COL_CREATED_AT,
    })
    for _, o := range valuesList {
        var vals []interface{}
        if o.Id != nil {
            vals = append(vals, *o.Id)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.PostId != nil {
            vals = append(vals, *o.PostId)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Url != nil {
            vals = append(vals, *o.Url)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Name != nil {
            vals = append(vals, *o.Name)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Size != nil {
            vals = append(vals, *o.Size)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Path != nil {
            vals = append(vals, *o.Path)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Alt != nil {
            vals = append(vals, *o.Alt)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Type != nil {
            vals = append(vals, *o.Type)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.MimeType != nil {
            vals = append(vals, *o.MimeType)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Label != nil {
            vals = append(vals, *o.Label)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.CreatedAt != nil {
            vals = append(vals, *o.CreatedAt)
        } else {
            vals = append(vals, sql.NullTime{})
        }
        c.BaseClient.AppendValues(vals)
    }
    return c.BaseClient.Run()
}
