//////////////////////////////////////////////////////////////////////
// create_table_client.go
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

type CreateTableClient struct {
    *myQuery.CreateTableClient
    RefArticleTableName string
    RefTagTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refArticleTableName, refTagTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDb(tableName, db),
        RefArticleTableName: refArticleTableName,
        RefTagTableName: refTagTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refArticleTableName, refTagTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx),
        RefArticleTableName: refArticleTableName,
        RefTagTableName: refTagTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refArticleTableName, refTagTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTx(tableName, tx),
        RefArticleTableName: refArticleTableName,
        RefTagTableName: refTagTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refArticleTableName, refTagTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx),
        RefArticleTableName: refArticleTableName,
        RefTagTableName: refTagTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    c.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ARTICLE_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Article ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_TAG_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Tag ID")).
        SetPrimaryKeys([]string{COL_ARTICLE_ID, COL_TAG_ID}).
        AppendConstraint("", COL_ARTICLE_ID, c.RefArticleTableName, nkwMysqlModelArticle.COL_ID, true, true).
        AppendConstraint("", COL_TAG_ID, c.RefTagTableName, nkwMysqlModelTag.COL_ID, true, true).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
