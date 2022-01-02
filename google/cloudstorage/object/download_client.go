//////////////////////////////////////////////////////////////////////
// download_client.go
//////////////////////////////////////////////////////////////////////
package object

import (
    "context"
    "io"
    "time"
    "golang.org/x/oauth2"
)


type DownloadClient struct {
    *Client
    CredentialFile string
    Object string
    TokenSource oauth2.TokenSource
}


//////////////////////////////////////////////////////////////////////
// Set timeout second.
//////////////////////////////////////////////////////////////////////
func (c *DownloadClient) SetTimeout(sec int) *DownloadClient {
    c.TimeoutSec = sec
    return c
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *DownloadClient) Run() ([]byte, error) {
    client, err := newStorageClient(c.CredentialFile, c.TokenSource)
    if err != nil {
        return nil, err
    }
    defer client.Close()
    ctx := context.Background()
    var cancel context.CancelFunc
    if c.TimeoutSec > 0 {
        ctx, cancel = context.WithTimeout(ctx, time.Second * time.Duration(c.TimeoutSec))
        defer cancel()
    }
    r, err := client.Bucket(c.BucketName).Object(c.Object).NewReader(ctx)
    if err != nil {
        return nil, err
    }
    defer r.Close()
    return io.ReadAll(r)
}
