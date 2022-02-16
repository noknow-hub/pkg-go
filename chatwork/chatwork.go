//////////////////////////////////////////////////////////////////////
// chatwork.go
//////////////////////////////////////////////////////////////////////
package chatwork

import (
    myV2Rooms "github.com/noknow-hub/pkg-go/chatwork/v2/rooms"
)


//////////////////////////////////////////////////////////////////////
// New v2 rooms client
//////////////////////////////////////////////////////////////////////
func NewV2RoomsClient(apiToken string) *myV2Rooms.Client {
    return myV2Rooms.NewClient(apiToken)
}
