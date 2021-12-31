//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package account_meta

import (
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
)

const (
    COL_ACCOUNT_ID = "account_id"
    COL_META_KEY = "meta_key"
    COL_META_VALUE = "meta_value"
    NUM_COLS = 3
    TABLE_NAME = "account_meta"
)

type AccountMeta struct {
    AccountId string
    MetaKey string
    MetaValue string
}

type AccountMetaWithAccount struct {
    AccountId string
    MetaKey string
    MetaValue string
    Account *nkwMysqlModelAccount.Account
}
