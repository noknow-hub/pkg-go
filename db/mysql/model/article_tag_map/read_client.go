//////////////////////////////////////////////////////////////////////
// read_client.go
//////////////////////////////////////////////////////////////////////
package article_tag_map

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelTag "github.com/noknow-hub/pkg-go/db/mysql/model/tag"
)

type ReadClient struct {
    BaseClient *myQuery.SelectClient
}
type ReadClientWithArticleAndTag struct {
    BaseClient *myQuery.SelectClient
    RefArticleTable string
    RefTagTable string
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
// New ReadClient with reference articles and tags table.
//////////////////////////////////////////////////////////////////////
func (c *ReadClient) NewReadClientWithArticleAndTag(refArticleTable, refTagTable string) *ReadClientWithArticleAndTag {
    return &ReadClientWithArticleAndTag{
        BaseClient: c.BaseClient,
        RefArticleTable: refArticleTable,
        RefTagTable: refTagTable,
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
func (c *ReadClient) Run() (*ArticleTagMap, *myQuery.SelectResult, error) {
    var articleTagMap *ArticleTagMap
    c.BaseClient.SetLimit(1)
    result, err := c.BaseClient.Run()
    if err != nil {
        return articleTagMap, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return articleTagMap, result, err
    }
    articleTagMap = &ArticleTagMap{}
    if err := scanArticleTagMap(result.Rows[0], articleTagMap); err != nil {
        return articleTagMap, result, err
    }
    return articleTagMap, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *ReadClientWithArticleAndTag) Run() (*ArticleTagMap, *myQuery.SelectResult, error) {
    var articleTagMap *ArticleTagMap
    c.BaseClient.SetLimit(1)
    c.BaseClient.
        AppendInnerJoinTables(c.BaseClient.TableName, COL_ARTICLE_ID, c.RefArticleTable, nkwMysqlModelArticle.COL_ID).
        AppendInnerJoinTables(c.BaseClient.TableName, COL_TAG_ID, c.RefTagTable, nkwMysqlModelTag.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return articleTagMap, result, err
    }
    if result != nil && len(result.Rows) != 1 {
        return articleTagMap, result, err
    }
    articleTagMap = &ArticleTagMap{
        Article: &nkwMysqlModelArticle.Article{},
        Tag: &nkwMysqlModelTag.Tag{},
    }
    if err := scanArticleTagMapWithArticleAndTag(result.Rows[0], c.BaseClient.TableName, c.RefArticleTable, c.RefTagTable, articleTagMap); err != nil {
        return articleTagMap, result, err
    }
    return articleTagMap, result, nil
}

