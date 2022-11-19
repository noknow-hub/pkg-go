//////////////////////////////////////////////////////////////////////
// browse_client.go
//////////////////////////////////////////////////////////////////////
package post_image

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}

type BrowseClientWithPost struct {
    BaseClient *myQuery.SelectClient
    RefPostTableName string
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
// New BrowseClient with reference post table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithPost(refPostTable string) *BrowseClientWithPost {
    return &BrowseClientWithPost{
        BaseClient: c.BaseClient,
        RefPostTableName: refPostTable,
    }
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := c.BaseClient.Count()
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
func (o *BrowseClient) Run() ([]*PostImage, *myQuery.SelectResult, error) {
    var postImages []*PostImage
    result, err := o.BaseClient.Run()
    if err != nil {
        return postImages, result, err
    }

    for _, row := range result.Rows {
        postImage := &PostImage{}
        if err := scan(row, postImage); err != nil {
            return postImages, result, err
        }
        postImages = append(postImages, postImage)
    }

    return postImages, result, nil
}


//////////////////////////////////////////////////////////////////////
// Count.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithPost) Count() (int64, *myQuery.SelectResultCount, error) {
    resultCount, err := c.BaseClient.Count()
    return resultCount.Count, resultCount, err
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClientWithPost) Query() (*myQuery.SelectResultQuery, error) {
    return o.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (o *BrowseClientWithPost) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return o.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithPost) RunInJoin() ([]*PostImageWithPost, *myQuery.SelectResult, error) {
    var postImages []*PostImageWithPost
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_POST_ID, c.RefPostTableName, nkwMysqlModelPost.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return postImages, result, err
    }

    for _, row := range result.Rows {
        postImage := &PostImageWithPost{
            Post: &nkwMysqlModelPost.Post{},
        }
        if err := scanWithPost(row, c.BaseClient.TableName, c.RefPostTableName, postImage); err != nil {
            return postImages, result, err
        }
        postImages = append(postImages, postImage)
    }

    return postImages, result, nil
}
