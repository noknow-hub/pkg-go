//////////////////////////////////////////////////////////////////////
// market_data.go
//////////////////////////////////////////////////////////////////////
package market_data

import (
    "encoding/json"
    myConstant "github.com/noknow-hub/pkg-go/okx/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResponseMarketData "github.com/noknow-hub/pkg-go/okx/response/v5/market_data"
)

const (
    PARAM_INST_ID = "instId"
    PARAM_QUOTE_CCY = "quoteCcy"
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
// Get index tickers by "instId"
//////////////////////////////////////////////////////////////////////
func (c *Client) GetIndexTickersByInstId(instId string) (*myHttpClient.Response, *myResponseMarketData.IndexTickersResponse, error) {
    httpClient := myHttpClient.NewClient(myConstant.API_BASE_URL + myConstant.API_V5_MARKET_INDEX_TICKERS)
    httpClient.Config.AddUrlQueryData(PARAM_INST_ID, instId)
    resp, err := httpClient.Get()
    if err != nil {
        return resp, nil, err
    }
    var result *myResponseMarketData.IndexTickersResponse
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return resp, nil, err
    }

    return resp, result, nil
}


//////////////////////////////////////////////////////////////////////
// Get index tickers by "quoteCcy"
//////////////////////////////////////////////////////////////////////
func (c *Client) GetIndexTickersByQuoteCcy(quoteCcy string) (*myHttpClient.Response, *myResponseMarketData.IndexTickersResponse, error) {
    httpClient := myHttpClient.NewClient(myConstant.API_BASE_URL + myConstant.API_V5_MARKET_INDEX_TICKERS)
    httpClient.Config.AddUrlQueryData(PARAM_QUOTE_CCY, quoteCcy)
    resp, err := httpClient.Get()
    if err != nil {
        return resp, nil, err
    }
    var result *myResponseMarketData.IndexTickersResponse
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return resp, nil, err
    }

    return resp, result, nil
}

