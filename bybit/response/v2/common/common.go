//////////////////////////////////////////////////////////////////////
// general.go
//////////////////////////////////////////////////////////////////////
package common

import (

)

type Common struct {
    RetCode int     `json:"ret_code"`
    RetMsg string   `json:"ret_msg"`
    ExtCode string  `json:"ext_code"`
    ExtInfo string  `json:"ext_info"`
}
