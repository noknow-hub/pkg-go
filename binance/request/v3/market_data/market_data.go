//////////////////////////////////////////////////////////////////////
// market_data.go
//////////////////////////////////////////////////////////////////////
package market_data

import (
    "encoding/json"
    "time"
    myAuthentication "github.com/noknow-hub/pkg-go/binance/authentication"
    myConstant "github.com/noknow-hub/pkg-go/binance/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResponseMarketData "github.com/noknow-hub/pkg-go/binance/response/v3/market_data"
)

const (
    SYMBOL = "symbol"
)

type Client struct {
}


//////////////////////////////////////////////////////////////////////
// New Client
//////////////////////////////////////////////////////////////////////
func NewClient() *Client {
    return &Client{}
}


//////////////////////////////////////////////////////////////////////
// Symbol Price Ticker
//////////////////////////////////////////////////////////////////////
func (c *Client) v3TickerPrice(symbol string) (*myHttpClient.Response, *myResponseMarketData.TickerPrice, error) {
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.AddUrlQueryData(SYMBOL, symbol)
    resp, err := httpClient.Get()
    if err != nil {
        return resp, nil, err
    }
    var result *myResponseMarketData.TickerPrice
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return resp, nil, err
    }

    return resp, result, nil
}


//////////////////////////////////////////////////////////////////////
// Symbol Price Ticker All
//////////////////////////////////////////////////////////////////////
func (c *Client) v3TickerPriceAll() (int, []*myResponseMarketData.TickerPrice, error) {
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    resp, err := httpClient.Get()
    if err != nil {
        return resp, nil, err
    }
    var result []*myResponseMarketData.TickerPrice
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return resp, nil, err
    }

    return resp, result, nil
}

