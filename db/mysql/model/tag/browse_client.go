//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package tag

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDb(tableName string, db *sql.DB) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTx(tableName string, tx *sql.Tx) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New BrowseClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewBrowseClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *BrowseClient {
    return &BrowseClient{
        BaseClient: myQuery.NewSelectClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := o.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) Query() (*myQuery.SelectResultQuery, error) {
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClient) Run() ([]*Tag, *myQuery.SelectResult, error) {
    var tags []*Tag
    result, err := o.BaseClient.Run()
    if err != nil { 
        return tags, result, err
    }

    for _, row := range result.Rows {
        tag := &Tag{}
        if err := scanTag(row, tag); err != nil {
            return tags, result, err
        }
        tags = append(tags, tag)
    }

    return tags, result, nil
}
