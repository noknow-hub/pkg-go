//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package post_image

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)

type ReadClient struct {
    BaseClient *myQuery.SelectClient
}

type ReadClientWithPost struct {
    BaseClient *myQuery.SelectClient
    RefPostTableName string
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDb(tableName string, db *sql.DB) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTx(tableName string, tx *sql.Tx) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewReadClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *ReadClient {
    return &ReadClient{
        BaseClient: myQuery.NewSelectClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New ReadClient with reference serp table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithPost(refPostTable string) *ReadClientWithPost {
    return &ReadClientWithPost{
        BaseClient: c.BaseClient,
        RefPostTableName: refPostTable,
    }
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Query() (*myQuery.SelectResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) Run() (*PostImage, *myQuery.SelectResult, error) {
    var postImage *PostImage
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return postImage, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return postImage, result, err
    }
    postImage = &PostImage{}
    if err := scan(result.Rows[0], postImage); err != nil {
        return postImage, result, err
    }
    return postImage, result, nil
}


//////////////////////////////////////////////////////////////////////
// Query.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithPost) Query() (*myQuery.SelectResultQuery, error) {
    c.BaseClient.SetLimit(1)
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithPost) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithPost) Run() (*PostImageWithPost, *myQuery.SelectResult, error) {
    c.BaseClient.AppendInnerJoinTables(c.BaseClient.TableName, COL_POST_ID, c.RefPostTableName, nkwMysqlModelPost.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return nil, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return nil, result, err
    }
    postImage := &PostImageWithPost{
        Post: &nkwMysqlModelPost.Post{},
    }
    if err := scanWithPost(result.Rows[0], c.BaseClient.TableName, c.RefPostTableName, postImage); err != nil {
        return nil, result, err
    }
    return postImage, result, nil
}
