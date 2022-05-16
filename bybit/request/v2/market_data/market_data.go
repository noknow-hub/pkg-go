//////////////////////////////////////////////////////////////////////
// market_data.go
//////////////////////////////////////////////////////////////////////
package market_data

import (
    "encoding/json"
    myConstant "github.com/noknow-hub/pkg-go/bybit/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResponseMarketData "github.com/noknow-hub/pkg-go/bybit/response/v2/market_data"
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
// Latest Information for all
//////////////////////////////////////////////////////////////////////
func (c *Client) LatestInformationForAll() (*myHttpClient.Response, *myResponseMarketData.LatestInformationResponse, error) {
    httpClient := myHttpClient.NewClient(myConstant.API_BASE_URL + myConstant.API_V2_PUBLIC_TICKER)
    resp, err := httpClient.Get()
    if err != nil {
        return resp, nil, err
    }
    var result *myResponseMarketData.LatestInformationResponse
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return resp, nil, err
    }

    return resp, result, nil
}


//////////////////////////////////////////////////////////////////////
// Latest Information for Symbol
//////////////////////////////////////////////////////////////////////
func (c *Client) LatestInformationForSymbol(symbol string) (*myHttpClient.Response, *myResponseMarketData.LatestInformationResponse, error) {
    httpClient := myHttpClient.NewClient(myConstant.API_BASE_URL + myConstant.API_V2_PUBLIC_TICKER)
    httpClient.Config.AddUrlQueryData(SYMBOL, symbol)
    resp, err := httpClient.Get()
    if err != nil {
        return resp, nil, err
    }
    var result *myResponseMarketData.LatestInformationResponse
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return resp, nil, err
    }

    return resp, result, nil
}


