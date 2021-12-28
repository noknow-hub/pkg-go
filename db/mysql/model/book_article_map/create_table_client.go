//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package book_article_map

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
    nkwMysqlModelArticle "github.com/noknow-hub/pkg-go/db/mysql/model/article"
    nkwMysqlModelBook "github.com/noknow-hub/pkg-go/db/mysql/model/book"
)

type CreateTableClient struct {
    *myQuery.CreateTableClient
    RefBookTableName string
    RefArticleTableName string
}


//////////////////////////////////////////////////////////////////////
// New Client with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName, refBookTableName, refArticleTableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDb(tableName, db),
        RefBookTableName: refBookTableName,
        RefArticleTableName: refArticleTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName, refBookTableName, refArticleTableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithDbContext(tableName, db, ctx),
        RefBookTableName: refBookTableName,
        RefArticleTableName: refArticleTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName, refBookTableName, refArticleTableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTx(tableName, tx),
        RefBookTableName: refBookTableName,
        RefArticleTableName: refArticleTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Client with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName, refBookTableName, refArticleTableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        CreateTableClient: myQuery.NewCreateTableClientWithTxContext(tableName, tx, ctx),
        RefBookTableName: refBookTableName,
        RefArticleTableName: refArticleTableName,
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*myQuery.CreateTableResult, error) {
    c.
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_BOOK_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Book ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_ARTICLE_ID, "BIGINT UNSIGNED").
                SetNotNull().
                SetComment("Article ID")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_PART, "TINYINT(1) UNSIGNED").
                SetComment("Part")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_CHAPTER, "TINYINT(1) UNSIGNED").
                SetComment("Chapter")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SECTION, "TINYINT(1) UNSIGNED").
                SetComment("Section")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_SUB_SECTION, "TINYINT(1) UNSIGNED").
                SetComment("Sub section")).
        AppendColumnDefinition(
            myQuery.NewColumnDefinition(COL_INTRODUCTION, "TEXT").
                SetComment("Introduction")).
        SetPrimaryKeys([]string{COL_BOOK_ID, COL_ARTICLE_ID}).
        AppendConstraint("", COL_BOOK_ID, c.RefBookTableName, nkwMysqlModelBook.COL_ID, true, true).
        AppendConstraint("", COL_ARTICLE_ID, c.RefArticleTableName, nkwMysqlModelArticle.COL_ID, true, true).
        SetComment(c.TableName + " table.")
    return c.CreateTableClient.Run()
}
