//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package post

import (
    "context"
    "database/sql"
    "time"
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
        COL_PARENT_ID,
        COL_STATUS,
        COL_TYPE,
        COL_LANG_CODE,
        COL_COUNTRY_CODE,
        COL_TEXT,
        COL_CREATED_AT,
    })
    for _, o := range valuesList {
        var vals []interface{}
        if o.Id != nil {
            vals = append(vals, *o.Id)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.ParentId != nil {
            vals = append(vals, *o.ParentId)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Status != nil {
            vals = append(vals, *o.Status)
        } else {
            vals = append(vals, VAL_STATUS_PUBLIC)
        }
        if o.Type != nil {
            vals = append(vals, *o.Type)
        } else {
            vals = append(vals, VAL_TYPE_POST)
        }
        if o.LangCode != nil {
            vals = append(vals, *o.LangCode)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.CountryCode != nil {
            vals = append(vals, *o.CountryCode)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.Text != nil {
            vals = append(vals, *o.Text)
        } else {
            vals = append(vals, sql.NullString{})
        }
        if o.CreatedAt != nil {
            vals = append(vals, *o.CreatedAt)
        } else {
            vals = append(vals, time.Now().UTC().Format("2006-01-02 15:04:05"))
        }
        c.BaseClient.AppendValues(vals)
    }
    return c.BaseClient.Run()
}
