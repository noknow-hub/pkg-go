//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package account_post_map

import (
    nkwMysqlModelAccount "github.com/noknow-hub/pkg-go/db/mysql/model/account"
    nkwMysqlModelPost "github.com/noknow-hub/pkg-go/db/mysql/model/post"
)

const (
    COL_ACCOUNT_ID = "account_id"
    COL_POST_ID = "lang_code"
    NUM_COLS = 2
    TABLE_NAME = "account_post_map"
)

type AccountPostMap struct {
    AccountId string
    PostId string
}

type AccountPostMapWithAccount struct {
    AccountId string
    PostId string
    Account *nkwMysqlModelAccount.Account
}

type AccountPostMapWithPost struct {
    AccountId string
    PostId string
    Post *nkwMysqlModelPost.Post
}

type AddEditValues struct {
    AccountId *string
    PostId *string
}
