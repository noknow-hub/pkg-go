//////////////////////////////////////////////////////////////////////
// market_data.go
//////////////////////////////////////////////////////////////////////
package market_data

import (
    myResponseCommon "github.com/noknow-hub/pkg-go/bybit/response/v2/common"
)


type LatestInformationResponse struct {
    *myResponseCommon.Common
    Result []*LatestInformation  `json:"result"`
}


type LatestInformation struct {
    Symbol string                  `json:"symbol"`
    BidPrice string                `json:"bid_price"`
    AskPrice string                `json:"ask_price"`
    LastPrice string               `json:"last_price"`
    LastTickDirection string       `json:"last_tick_direction"`
    PrevPrice_24h string           `json:"prev_price_24h"`
    Price24hPcnt string            `json:"price_24h_pcnt"`
    HighPrice24h string            `json:"high_price_24h"`
    LowPrice24h string             `json:"low_price_24h"`
    PrevPrice1h string             `json:"prev_price_1h"`
    Price1hPcnt string             `json:"price_1h_pcnt"`
    MarkPrice string               `json:"mark_price"`
    IndexPrice string              `json:"index_price"`
    OpenInterest int               `json:"open_interest"`
    OpenValue string               `json:"open_value"`
    TotalTurnover string           `json:"total_turnover"`
    Turnover_24h string            `json:"turnover_24h"`
    TotalVolume string             `json:"total_volume"`
    Volume24h string               `json:"volume_24h"`
    FundingRate string             `json:"funding_rate"`
    PredictedFundingRate string    `json:"predicted_funding_rate"`
    NextFundingTime string         `json:"next_funding_time"`
    CountdownHour string           `json:"countdown_hour"`
    DeliveryFeeRate string         `json:"delivery_fee_rate"`
    PredictedDeliveryPrice string  `json:"predicted_delivery_price"`
    DeliveryTime string            `json:"delivery_time"`
}
