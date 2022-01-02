//////////////////////////////////////////////////////////////////////
// delete_client.go
//////////////////////////////////////////////////////////////////////
package object

import (
    "context"
    "time"
    "golang.org/x/oauth2"
)


type DeleteClient struct {
    *Client
    CredentialFile string
    Object string
    TokenSource oauth2.TokenSource
}


//////////////////////////////////////////////////////////////////////
// Set timeout second.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) SetTimeout(sec int) *DeleteClient {
    c.TimeoutSec = sec
    return c
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *DeleteClient) Run() error {
    client, err := newStorageClient(c.CredentialFile, c.TokenSource)
    if err != nil {
        return err
    }
    defer client.Close()
    ctx := context.Background()
    var cancel context.CancelFunc
    if c.TimeoutSec > 0 {
        ctx, cancel = context.WithTimeout(ctx, time.Second * time.Duration(c.TimeoutSec))
        defer cancel()
    }
    obj := client.Bucket(c.BucketName).Object(c.Object)
    return obj.Delete(ctx)
}
