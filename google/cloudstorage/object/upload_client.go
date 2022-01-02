//////////////////////////////////////////////////////////////////////
// upload_client.go
//////////////////////////////////////////////////////////////////////
package object

import (
    "context"
    "encoding/base64"
    "errors"
    "io"
    "mime"
    "net/http"
    "os"
    "path"
    "strings"
    "time"
    "golang.org/x/oauth2"
    "cloud.google.com/go/storage"
)


type UploadClient struct {
    *Client
    CredentialFile string
    Object string
    TokenSource oauth2.TokenSource
}


//////////////////////////////////////////////////////////////////////
// Set timeout second.
//////////////////////////////////////////////////////////////////////
func (c *UploadClient) SetTimeout(sec int) *UploadClient {
    c.TimeoutSec = sec
    return c
}


//////////////////////////////////////////////////////////////////////
// Run with file path.
//////////////////////////////////////////////////////////////////////
func (c *UploadClient) RunWithFilePath(filePath string) (*storage.ObjectAttrs, error) {
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
    w := client.Bucket(c.BucketName).Object(c.Object).NewWriter(ctx)
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    fileType := mime.TypeByExtension(path.Ext(filePath))
    if fileType == "" {
        b := make([]byte, 512)
        if _, err := f.Read(b); err != nil {
            return nil, err
        }
        w.ContentType = http.DetectContentType(b)
    } else {
        w.ContentType = fileType
    }
    if _, err = io.Copy(w, f); err != nil {
        return nil, err
    }
    if err := w.Close(); err != nil {
        return nil, err
    }
    return w.Attrs(), nil
}


//////////////////////////////////////////////////////////////////////
// Run with data URI.
//////////////////////////////////////////////////////////////////////
func (c *UploadClient) RunWithDataUri(dataUri string) (*storage.ObjectAttrs, error) {
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
    w := client.Bucket(c.BucketName).Object(c.Object).NewWriter(ctx)
    dataUrlSplit := strings.Split(dataUri, ";base64,")
    if len(dataUrlSplit) != 2 {
        return nil, errors.New("UploadClient.DataUrl's format is inccorect. The syntax of DataURL is data:MEDIA_TYPE;base64,ENCODED_DATA.")
    }
    r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(dataUrlSplit[1]))
    w.ContentType = dataUrlSplit[0][5:]
    if _, err = io.Copy(w, r); err != nil {
        return nil, err
    }
    if err := w.Close(); err != nil {
        return nil, err
    }
    return w.Attrs(), nil
}
