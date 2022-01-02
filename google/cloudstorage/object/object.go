//////////////////////////////////////////////////////////////////////
// object.go
//////////////////////////////////////////////////////////////////////
package object

import (
    "context"
    "golang.org/x/oauth2"
    "google.golang.org/api/option"
    "cloud.google.com/go/storage"
)


type Client struct {
    BucketName string
    TimeoutSec int
}


//////////////////////////////////////////////////////////////////////
// New Client.
//////////////////////////////////////////////////////////////////////
func NewClient(bucketName string) *Client {
    return &Client{
        BucketName: bucketName,
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewDeleteClient(object string) *DeleteClient {
    return &DeleteClient{
        Client: c,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with credentials.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewDeleteClientWithCredentials(object, credentialsFile string) *DeleteClient {
    return &DeleteClient{
        Client: c,
        CredentialFile: credentialsFile,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New DeleteClient with token source.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewDeleteClientWithTokenSource(object string, tokenSource oauth2.TokenSource) *DeleteClient {
    return &DeleteClient{
        Client: c,
        TokenSource: tokenSource,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New DownloadClient.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewDownloadClient(object string) *DownloadClient {
    return &DownloadClient{
        Client: c,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New DownloadClient with credentials.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewDownloadClientWithCredentials(object, credentialsFile string) *DownloadClient {
    return &DownloadClient{
        Client: c,
        CredentialFile: credentialsFile,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New DownloadClient with token source.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewDownloadClientWithTokenSource(object string, tokenSource oauth2.TokenSource) *DownloadClient {
    return &DownloadClient{
        Client: c,
        TokenSource: tokenSource,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New ListingClient.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewListingClient() *ListingClient {
    return &ListingClient{
        Client: c,
    }
}


//////////////////////////////////////////////////////////////////////
// New ListingClient with credentials.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewListingClientWithCredentials(credentialsFile string) *ListingClient {
    return &ListingClient{
        Client: c,
        CredentialFile: credentialsFile,
    }
}


//////////////////////////////////////////////////////////////////////
// New ListingClient with token source.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewListingClientWithTokenSource(tokenSource oauth2.TokenSource) *DeleteClient {
    return &DeleteClient{
        Client: c,
        TokenSource: tokenSource,
    }
}


//////////////////////////////////////////////////////////////////////
// New UploadClient with credentials.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewUploadClientWithCredentials(object, credentialsFile string) *UploadClient {
    return &UploadClient{
        Client: c,
        CredentialFile: credentialsFile,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New UploadClient with token source.
//////////////////////////////////////////////////////////////////////
func (c *Client) NewUploadClientWithTokenSource(object string, tokenSource oauth2.TokenSource) *UploadClient {
    return &UploadClient{
        Client: c,
        TokenSource: tokenSource,
        Object: object,
    }
}


//////////////////////////////////////////////////////////////////////
// New *storage.Client.
//////////////////////////////////////////////////////////////////////
func newStorageClient(credentialFile string, tokenSource oauth2.TokenSource) (*storage.Client, error) {
    var client *storage.Client
    var err error
    ctx := context.Background()
    if credentialFile != "" {
        client, err = storage.NewClient(ctx, option.WithCredentialsFile(credentialFile))
    } else if tokenSource != nil {
        client, err = storage.NewClient(ctx, option.WithTokenSource(tokenSource))
    } else {
        client, err = storage.NewClient(ctx, option.WithoutAuthentication())
    }
    if err != nil {
        return nil, err
    }
    return client, nil
}
