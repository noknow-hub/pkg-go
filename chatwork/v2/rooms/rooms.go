//////////////////////////////////////////////////////////////////////
// rooms.go
//////////////////////////////////////////////////////////////////////
package rooms

import (
    "encoding/json"
    "net/http"
    myConstant "github.com/noknow-hub/pkg-go/chatwork/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myMessages "github.com/noknow-hub/pkg-go/chatwork/v2/rooms/messages"
    myResponse "github.com/noknow-hub/pkg-go/chatwork/v2/response"
)

type Client struct {
    ApiToken string
    EndpointUrl string
}

type Room struct {
    RoomId int64            `json:"room_id,omitempty"`
    Name string             `json:"name,omitempty"`
    Type string             `json:"type,omitempty"`
    Role string             `json:"role,omitempty"`
    Sticky bool             `json:"sticky,omitempty"`
    UnreadNum int64         `json:"unread_num,omitempty"`
    MentionNum int64        `json:"mention_num,omitempty"`
    MytaskNum int64         `json:"mytask_num,omitempty"`
    MessageNum int64        `json:"message_num,omitempty"`
    FileNum int64           `json:"file_num,omitempty"`
    TaskNum int64           `json:"task_num,omitempty"`
    IconPath string         `json:"icon_path,omitempty"`
    LastUpdateTime int64    `json:"last_update_time,omitempty"`
}


//////////////////////////////////////////////////////////////////////
// New Client.
//////////////////////////////////////////////////////////////////////
func NewClient(apiToken string) *Client {
    return &Client{
        ApiToken: apiToken,
        EndpointUrl: myConstant.ENDPOINT_URL_ROOMS,
    }
}


//////////////////////////////////////////////////////////////////////
// GET
//////////////////////////////////////////////////////////////////////
func (c *Client) Get() (int, []*Room, *myResponse.Error) {
    errResp := &myResponse.Error{}
    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.AddHeader(myConstant.HTTP_HEADER_TOKEN, c.ApiToken) 
    resp, err := httpClient.Get()
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
    var rooms []*Room
    if err := json.Unmarshal(resp.Body, &rooms); err != nil {
        errResp.Errors = append(errResp.Errors, err.Error())
        return resp.StatusCode, nil, errResp
    }
    return resp.StatusCode, rooms, nil
}


//////////////////////////////////////////////////////////////////////
// New Client for messages.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewMessagesClient(roomId string) *myMessages.Client {
    return myMessages.NewClient(c.ApiToken, roomId)
}

