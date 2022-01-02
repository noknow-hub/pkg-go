//////////////////////////////////////////////////////////////////////
// listing_client.go
//////////////////////////////////////////////////////////////////////
package object

import (
    "context"
    "time"
    "google.golang.org/api/iterator"
    "golang.org/x/oauth2"
    "cloud.google.com/go/storage"
)


type ListingClient struct {
    *Client
    CredentialFile string
    Delimiter string
    Prefix string
    TokenSource oauth2.TokenSource
}


//////////////////////////////////////////////////////////////////////
// Set prefix.
//////////////////////////////////////////////////////////////////////
func (c *ListingClient) SetPrefix(prefix string) *ListingClient {
    c.Prefix = prefix
    return c
}


//////////////////////////////////////////////////////////////////////
// Set delimiter.
//////////////////////////////////////////////////////////////////////
func (c *ListingClient) SetDelimiter(delimiter string) *ListingClient {
    c.Delimiter = delimiter
    return c
}


//////////////////////////////////////////////////////////////////////
// Set timeout second.
//////////////////////////////////////////////////////////////////////
func (c *ListingClient) SetTimeout(sec int) *ListingClient {
    c.TimeoutSec = sec
    return c
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *ListingClient) Run() ([]*storage.ObjectAttrs, error) {
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
    var q *storage.Query
    if c.Prefix != "" && c.Delimiter != "" {
        q = &storage.Query{
            Prefix: c.Prefix,
            Delimiter: c.Delimiter,
        }
    } else if c.Prefix != "" {
        q = &storage.Query{
            Prefix: c.Prefix,
        }
    } else if c.Delimiter != "" {
        q = &storage.Query{
            Delimiter: c.Delimiter,
        }
    }
    if q != nil {
        if err := q.SetAttrSelection([]string{"Name", "ContentType", "ContentLanguage", "ACL", "Size", "ContentEncoding", "MediaLink", "Metadata", "Generation", "Created", "Updated"}); err != nil {
            return nil, err
        }
    }
    it := client.Bucket(c.BucketName).Objects(ctx, q)
    var result []*storage.ObjectAttrs
    for {
        attrs, err := it.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, err
        }
        result = append(result, attrs)
    }
    return result, nil
}
