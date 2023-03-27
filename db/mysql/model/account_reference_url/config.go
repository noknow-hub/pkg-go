//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package account_reference_url

import (
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)


const (
    COL_ACCOUNT_ID = "account_id"
    COL_TYPE = "type"
    COL_URL = "url"
    NUM_COLS = 3
    TABLE_NAME = "account_reference_url"
)


type AccountReferenceUrl struct {
    AccountId string
    PostId string
}


type AccountPostMapWithAccount struct {
    AccountId string
    Type string
    Url string
    Account *nkwMysqlModelAccount.Account
}


type AddEditValues struct {
    AccountId *string
    Type *string
    Url *string
}
