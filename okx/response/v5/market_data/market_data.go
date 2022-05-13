//////////////////////////////////////////////////////////////////////
// market_data.go
//////////////////////////////////////////////////////////////////////
package market_data

import (
    myResponse "github.com/noknow-hub/pkg-go/okx/response/v5"
)


import (
)


type IndexTickersResponse struct {
    *myResponse.General
    Data []*IndexTickersData  `json:"data"`
}


type IndexTickersData struct {
    InstId string  `json:"instId"`
    IdxPx string   `json:"idxPx"`
    High24h string `json:"high24h"`
    SodUtc0 string `json:"sodUtc0"`
    Open24h string `json:"open24h"`
    Low24h string  `json:"low24h"`
    SodUtc8 string `json:"sodUtc8"`
    Ts string      `json:"ts"`
}
