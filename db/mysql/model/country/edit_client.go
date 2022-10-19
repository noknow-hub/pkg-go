//////////////////////////////////////////////////////////////////////
// edit_client.go
//////////////////////////////////////////////////////////////////////
package country

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type EditClient struct {
    BaseClient *myQuery.UpdateClient
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDb(tableName string, db *sql.DB) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTx(tableName string, tx *sql.Tx) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New EditClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewEditClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *EditClient {
    return &EditClient{
        BaseClient: myQuery.NewUpdateClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run
//////////////////////////////////////////////////////////////////////
func (c *EditClient) Run() (*myQuery.UpdateResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run by COL_COUNTRY_CODE
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunByCountryCode(countryCode string) (*myQuery.UpdateResult, error) {
    c.BaseClient.WhereCondition.SetWhere(COL_COUNTRY_CODE, countryCode)
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run with all by COL_COUNTRY_CODE
//////////////////////////////////////////////////////////////////////
func (c *EditClient) RunWithAllByCountryCode(currentCountryCode string, countryCode, ar, de, en, es, fr, ja, pt, ru, zhCn, zhTw, continent, status *string) (*myQuery.UpdateResult, error) {
    if countryCode != nil {
        c.BaseClient.AssignmentList.Append(COL_COUNTRY_CODE, *countryCode)
    }
    if ar != nil {
        c.BaseClient.AssignmentList.Append(COL_AR, *ar)
    }
    if de != nil {
        c.BaseClient.AssignmentList.Append(COL_DE, *de)
    }
    if en != nil {
        c.BaseClient.AssignmentList.Append(COL_EN, *en)
    }
    if es != nil {
        c.BaseClient.AssignmentList.Append(COL_ES, *es)
    }
    if fr != nil {
        c.BaseClient.AssignmentList.Append(COL_FR, *fr)
    }
    if ja != nil {
        c.BaseClient.AssignmentList.Append(COL_JA, *ja)
    }
    if pt != nil {
        c.BaseClient.AssignmentList.Append(COL_PT, *pt)
    }
    if ru != nil {
        c.BaseClient.AssignmentList.Append(COL_RU, *ru)
    }
    if zhCn != nil {
        c.BaseClient.AssignmentList.Append(COL_ZH_CN, *zhCn)
    }
    if zhTw != nil {
        c.BaseClient.AssignmentList.Append(COL_ZH_TW, *zhTw)
    }
    if continent != nil {
        c.BaseClient.AssignmentList.Append(COL_CONTINENT, *continent)
    }
    if status != nil {
        c.BaseClient.AssignmentList.Append(COL_STATUS, *status)
    }
    c.BaseClient.WhereCondition.SetWhere(COL_COUNTRY_CODE, currentCountryCode)
    return c.BaseClient.Run()
}
