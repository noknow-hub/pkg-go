//////////////////////////////////////////////////////////////////////
// messages.go
//////////////////////////////////////////////////////////////////////
package messages

import (
    "encoding/json"
    "net/http"
    myConstant "github.com/noknow-hub/pkg-go/chatwork/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResponse "github.com/noknow-hub/pkg-go/chatwork/v2/response"
)

const (
    PARAM_BODY = "body"
    PARAM_SELF_UNREAD = "self_unread"
)

type Client struct {
    ApiToken string
    EndpointUrl string
}

type Message struct {          
    MessageId string    `json:"message_id,omitempty"`
    Account *Account    `json:"account,omitempty"`
    Body string         `json:"body,omitempty"`
    SendTime int64      `json:"send_time,omitempty"`
    UpdateTime int64    `json:"update_time,omitempty"`
}

type Account struct {
    AccountId int64          `json:"account_id,omitempty"`
    Name string              `json:"name,omitempty"`
    AvatarImageUrl string    `json:"avatar_image_url,omitempty"`
}


//////////////////////////////////////////////////////////////////////
// New Client.
//////////////////////////////////////////////////////////////////////
func NewClient(apiToken, roomId string) *Client {
    return &Client{
        ApiToken: apiToken,
        EndpointUrl: myConstant.ENDPOINT_URL_ROOMS + "/" + roomId + "/messages",
    }
}


//////////////////////////////////////////////////////////////////////
// POST
//////////////////////////////////////////////////////////////////////
func (c *Client) Post(body string, selfUnread bool) (int, *Message, *myResponse.Error) {
    errResp := &myResponse.Error{}
    tmpSelfUnread := "0"
    if selfUnread {
        tmpSelfUnread = "1"
    }
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.
        AddHeader(myConstant.HTTP_HEADER_TOKEN, c.ApiToken).
        AddFormData(PARAM_BODY, body).
        AddFormData(PARAM_SELF_UNREAD, tmpSelfUnread)
    resp, err := httpClient.Post()
    if err != nil {
        errResp.Errors = append(errResp.Errors, err.Error())
        return 0, nil, errResp
    }
    if resp == nil {
        errResp.Errors = append(errResp.Errors, "Empty response.")
        return 0, nil, errResp
    }
    if resp.StatusCode != http.StatusOK {
        if err := json.Unmarshal(resp.Body, &errResp); err != nil {
            errResp.Errors = append(errResp.Errors, err.Error())
            return resp.StatusCode, nil, errResp
        }
        return resp.StatusCode, nil, errResp
    }
    var message *Message
    if err := json.Unmarshal(resp.Body, &message); err != nil {
        errResp.Errors = append(errResp.Errors, err.Error())
        return resp.StatusCode, nil, errResp
    }
    return resp.StatusCode, message, nil
}

