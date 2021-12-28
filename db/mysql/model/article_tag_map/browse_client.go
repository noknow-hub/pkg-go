//////////////////////////////////////////////////////////////////////
// browse_client.go
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

type BrowseClient struct {
    BaseClient *myQuery.SelectClient
}
type BrowseClientWithArticleAndTag struct {
    BaseClient *myQuery.SelectClient
    RefArticleTable string
    RefTagTable string
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
// New BrowseClient with reference article and tag table.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) NewBrowseClientWithArticleAndTag(refArticleTable, refTagTable string) *BrowseClientWithArticleAndTag {
    return &BrowseClientWithArticleAndTag{
        BaseClient: c.BaseClient,
        RefArticleTable: refArticleTable,
        RefTagTable: refTagTable,
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
func (c *BrowseClient) Query() (*myQuery.SelectResultQuery, error) {
    return c.BaseClient.Query()
}


//////////////////////////////////////////////////////////////////////
// QueryRow.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) QueryRow() (*myQuery.SelectResultQueryRow, error) {
    return c.BaseClient.QueryRow()
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClient) Run() ([]*ArticleTagMap, *myQuery.SelectResult, error) {
    var articleTagMaps []*ArticleTagMap
    result, err := c.BaseClient.Run()
    if err != nil { 
        return articleTagMaps, result, err
    }

    for _, row := range result.Rows {
        articleTagMap := &ArticleTagMap{}
        if err := scanArticleTagMap(row, articleTagMap); err != nil {
            return articleTagMaps, result, err
        }
        articleTagMaps = append(articleTagMaps, articleTagMap)
    }

    return articleTagMaps, result, nil
}


//////////////////////////////////////////////////////////////////////
// Run with INNER JOIN.
//////////////////////////////////////////////////////////////////////
func (c *BrowseClientWithArticleAndTag) Run() ([]*ArticleTagMap, *myQuery.SelectResult, error) {
    var articleTagMaps []*ArticleTagMap
    c.BaseClient.
        AppendInnerJoinTables(c.BaseClient.TableName, COL_ARTICLE_ID, c.RefArticleTable, nkwMysqlModelArticle.COL_ID).
        AppendInnerJoinTables(c.BaseClient.TableName, COL_TAG_ID, c.RefTagTable, nkwMysqlModelTag.COL_ID)
    result, err := c.BaseClient.Run()
    if err != nil {
        return articleTagMaps, result, err
    }

    for _, row := range result.Rows {
        articleTagMap := &ArticleTagMap{
            Article: &nkwMysqlModelArticle.Article{},
            Tag: &nkwMysqlModelTag.Tag{},
        }
        if err := scanArticleTagMapWithArticleAndTag(row, c.BaseClient.TableName, c.RefArticleTable, c.RefTagTable, articleTagMap); err != nil {
            return articleTagMaps, result, err
        }
        articleTagMaps = append(articleTagMaps, articleTagMap)
    }

    return articleTagMaps, result, nil
}
