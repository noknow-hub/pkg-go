//////////////////////////////////////////////////////////////////////
// create_table_client.go
//////////////////////////////////////////////////////////////////////
package query

import (
    "context"
    "database/sql"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    myUtil "github.com/noknow-hub/pkg-go/db/mysql/query/util"
)

const (
    CHARSET_UTF8 = "utf8"
    DEFAULT_CHARSET = "utf8mb4"
    DEFAULT_CURRENT_TIMESTAMP = "CURRENT_TIMESTAMP"
    DEFAULT_CURRENT_TIMESTAMP_ON_UPDATE = "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"
    DEFAULT_ENGINE_INNODB = "InnoDB"
)

type CreateTableClient struct {
    Charset string
    ColumnDefinitions []*ColumnDefinition
    Comment string
    Constraints []*Constraint
    Ctx context.Context
    Db *sql.DB
    Engine string
    IndexKeys []string
    PrimaryKeys []string
    TableName string
    Tx *sql.Tx
    UniqueKeys []string
}

type ColumnDefinition struct {
    ColName string
    DataType string
    NotNull bool
    AutoIncrement bool
    Default string
    PrimaryKey bool
    UniqueKey bool
    Comment string
}

type Constraint struct {
    Symbol string
    ForeignKey string
    RefTableName string
    RefTblColName string
    OnDelete bool
    OnUpdate bool
}

type CreateTableResult struct {
    RawQuery string
    SqlResult sql.Result
}


//////////////////////////////////////////////////////////////////////
// New CreateTableClient with db object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDb(tableName string, db *sql.DB) *CreateTableClient {
    return &CreateTableClient{
        Charset: DEFAULT_CHARSET,
        Db: db,
        Engine: DEFAULT_ENGINE_INNODB,
        TableName: tableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New CreateTableClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        Charset: DEFAULT_CHARSET,
        Ctx: ctx,
        Db: db,
        Engine: DEFAULT_ENGINE_INNODB,
        TableName: tableName,
    }
}


//////////////////////////////////////////////////////////////////////
// New CreateTableClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTx(tableName string, tx *sql.Tx) *CreateTableClient {
    return &CreateTableClient{
        Charset: DEFAULT_CHARSET,
        Engine: DEFAULT_ENGINE_INNODB,
        TableName: tableName,
        Tx: tx,
    }
}


//////////////////////////////////////////////////////////////////////
// New CreateTableClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewCreateTableClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *CreateTableClient {
    return &CreateTableClient{
        Charset: DEFAULT_CHARSET,
        Ctx: ctx,
        Engine: DEFAULT_ENGINE_INNODB,
        TableName: tableName,
        Tx: tx,
    }
}


//////////////////////////////////////////////////////////////////////
// New ColumnDefinition.
//////////////////////////////////////////////////////////////////////
func NewColumnDefinition(colName, dataType string) *ColumnDefinition {
    return &ColumnDefinition{
        ColName: colName,
        DataType: dataType,
    }
}


//////////////////////////////////////////////////////////////////////
// Append constraint.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) AppendConstraint(symbol, foreignKey, refTableName, refTableColName string, onDelete, onUpdate bool) *CreateTableClient {
    c.Constraints = append(c.Constraints, &Constraint{
        Symbol: symbol,
        ForeignKey: foreignKey,
        RefTableName: refTableName,
        RefTblColName: refTableColName,
        OnDelete: onDelete,
        OnUpdate: onUpdate,
    })
    return c
}


//////////////////////////////////////////////////////////////////////
// Append column definition.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) AppendColumnDefinition(cd *ColumnDefinition) *CreateTableClient {
    c.ColumnDefinitions = append(c.ColumnDefinitions, cd)
    return c
}


func (c *ColumnDefinition) SetNotNull() *ColumnDefinition {
    c.NotNull = true
    return c
}
func (c *ColumnDefinition) SetAutoIncrement() *ColumnDefinition {
    c.AutoIncrement = true
    return c
}
func (c *ColumnDefinition) SetDefault(val string) *ColumnDefinition {
    c.Default = val
    return c
}
func (c *ColumnDefinition) SetPrimaryKey(primaryKeys []string) *ColumnDefinition {
    c.PrimaryKey = true
    return c
}
func (c *ColumnDefinition) SetUniqueKey() *ColumnDefinition {
    c.UniqueKey = true
    return c
}
func (c *ColumnDefinition) SetComment(val string) *ColumnDefinition {
    c.Comment = val
    return c
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) Run() (*CreateTableResult, error) {
    result := &CreateTableResult{}
    result.RawQuery = c.generateQuery()
    var err error
    result.SqlResult, err = myUtil.Exec(c.Db, c.Tx, c.Ctx, result.RawQuery, nil)
    return result, err
}


//////////////////////////////////////////////////////////////////////
// Set charset.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) SetCharset(charset string) *CreateTableClient {
    c.Charset = charset
    return c
}


//////////////////////////////////////////////////////////////////////
// Set comment.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) SetComment(comment string) *CreateTableClient {
    c.Comment = comment
    return c
}


//////////////////////////////////////////////////////////////////////
// Set engine.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) SetEngine(engine string) *CreateTableClient {
    c.Engine = engine
    return c
}


//////////////////////////////////////////////////////////////////////
// Set index keys.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) SetIndexKeys(indexKeys []string) *CreateTableClient {
    c.IndexKeys = indexKeys
    return c
}


//////////////////////////////////////////////////////////////////////
// Set primary keys.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) SetPrimaryKeys(primaryKeys []string) *CreateTableClient {
    c.PrimaryKeys = primaryKeys
    return c
}


//////////////////////////////////////////////////////////////////////
// Set unique keys.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) SetUniqueKeys(uniqueKeys []string) *CreateTableClient {
    c.UniqueKeys = uniqueKeys
    return c
}


//////////////////////////////////////////////////////////////////////
// Generate query.
//////////////////////////////////////////////////////////////////////
func (c *CreateTableClient) generateQuery() string {
    buf := make([]byte, 0)

    buf = append(buf, "CREATE TABLE IF NOT EXISTS " + c.TableName + " ("...)
    for i, o := range c.ColumnDefinitions {
        if i > 0 {
            buf = append(buf, ","...)
        }
        buf = append(buf, o.ColName + " " + o.DataType...)
        if o.NotNull {
            buf = append(buf, " NOT NULL"...)
        }
        if o.Default != "" {
             buf = append(buf, " DEFAULT " + o.Default...)
        }
        if o.AutoIncrement {
            buf = append(buf, " AUTO_INCREMENT"...)
        }
        if o.PrimaryKey {
            buf = append(buf, " PRIMARY KEY"...)
        }
        if o.UniqueKey {
            buf = append(buf, " UNIQUE KEY"...)
        }
        if o.Comment != "" {
            buf = append(buf, " COMMENT '" + o.Comment + "'"...)
        }
    }
    if len(c.PrimaryKeys) > 0 {
        buf = append(buf, ", PRIMARY KEY (" + strings.Join(c.PrimaryKeys, ",") + ")"...)
    }
    if len(c.IndexKeys) > 0 {
        buf = append(buf, ", INDEX (" + strings.Join(c.IndexKeys, ",") + ")"...)
    }
    if len(c.UniqueKeys) > 0 {
        buf = append(buf, ", UNIQUE KEY (" + strings.Join(c.UniqueKeys, ",") + ")"...)
    }
    if len(c.Constraints) > 0 {
        for _, c := range c.Constraints {
            buf = append(buf, ","...)
            if c.Symbol != "" {
                buf = append(buf, " CONSTRAINT " + c.Symbol...)
            }
            buf = append(buf, " FOREIGN KEY (" + c.ForeignKey + ")"...)
            buf = append(buf, " REFERENCES " + c.RefTableName + " (" + c.RefTblColName + ")"...)
            if c.OnDelete {
                buf = append(buf, " ON DELETE CASCADE"...)
            }
            if c.OnUpdate {
                buf = append(buf, " ON UPDATE CASCADE"...)
            }
        }
    }
    buf = append(buf, ")"...)
    if c.Engine != "" {
        buf = append(buf, " ENGINE=" + c.Engine...)
    } else {
        buf = append(buf, " ENGINE=" + DEFAULT_ENGINE_INNODB...)
    }
    if c.Charset != "" {
        buf = append(buf, " DEFAULT CHARSET=" + c.Charset...)
    } else {
        buf = append(buf, " DEFAULT CHARSET=" + DEFAULT_CHARSET...)
    }
    if c.Comment != "" {
        buf = append(buf, " COMMENT='" + c.Comment + "'"...)
    }
    return string(buf[:])
}
