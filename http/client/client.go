//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package client

import (
    "bytes"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "strings"
    "time"
)


type Client struct {
    Config *Config
}


//////////////////////////////////////////////////////////////////////
// New client.
//////////////////////////////////////////////////////////////////////
func NewClient(url string) *Client {
    return &Client{
        Config: &Config{
            Url: url,
        },
    }
}


//////////////////////////////////////////////////////////////////////
// HTTP DELETE request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Delete() (*Response, error) {
    return c.Do(http.MethodDelete)
}


//////////////////////////////////////////////////////////////////////
// HTTP request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Do(httpMethod string) (*Response, error) {
    var req *http.Request
    var err error

    // HTTP URL query
    if c.Config.UrlQueryData != nil {
        c.Config.Url = c.Config.Url + "?" + c.Config.UrlQueryData.Encode()
    }

    // HTTP message body
    if c.Config.XWwwFormUrlencodedData != nil {
        req, err = http.NewRequest(httpMethod, c.Config.Url, strings.NewReader(c.Config.XWwwFormUrlencodedData.Encode()))
        req.Header.Set("Content-Type", CONTENT_TYPE_X_WWW_FORM_URLENCODED)
    } else if len(c.Config.FormData) > 0 {
        body, contentType, err := c.generateFormDataBody()
        if err != nil {
            return nil, err
        }
        req, err = http.NewRequest(httpMethod, c.Config.Url, body)
        req.Header.Set("Content-Type", contentType)
    } else if c.Config.JsonData != nil {
        req, err = http.NewRequest(httpMethod, c.Config.Url, bytes.NewBuffer(c.Config.JsonData))
        req.Header.Set("Content-Type", CONTENT_TYPE_JSON)
    } else {
        req, err = http.NewRequest(httpMethod, c.Config.Url, nil)
    }
    if err != nil {
        return nil, err
    }

    // HTTP header
    if c.Config.Header != nil {
        for key, value := range c.Config.Header {
            req.Header.Set(key, value)
        }
    }

    client := &http.Client{}

    // Timeout
    if c.Config.TimeoutSec > 0 {
        client.Timeout = time.Second * time.Duration(c.Config.TimeoutSec)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    response := &Response{
        Body: respBody,
        ContentLength: resp.ContentLength,
        Header: resp.Header,
        Proto: resp.Proto,
        ProtoMajor: resp.ProtoMajor,
        ProtoMinor: resp.ProtoMinor,
        Status: resp.Status,
        StatusCode: resp.StatusCode,
        TransferEncoding: resp.TransferEncoding,
        Uncompressed: resp.Uncompressed,
        Trailer: resp.Trailer,
        Request: resp.Request,
        TLS: resp.TLS,
    }

    return response, nil
}


//////////////////////////////////////////////////////////////////////
// HTTP GET request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Get() (*Response, error) {
    return c.Do(http.MethodGet)
}


//////////////////////////////////////////////////////////////////////
// HTTP PATCH request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Patch() (*Response, error) {
    return c.Do(http.MethodPatch)
}


//////////////////////////////////////////////////////////////////////
// HTTP POST request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Post() (*Response, error) {
    return c.Do(http.MethodPost)
}


//////////////////////////////////////////////////////////////////////
// HTTP PUT request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Put() (*Response, error) {
    return c.Do(http.MethodPut)
}


//////////////////////////////////////////////////////////////////////
// Generate a body for "multipart/form-data".
//////////////////////////////////////////////////////////////////////
func (c *Client) generateFormDataBody() (*bytes.Buffer, string, error) {
    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    contentType := writer.FormDataContentType()
    for _, formData := range c.Config.FormData {
        if formData.FileName != "" && formData.FileData != nil {
            part, err := writer.CreateFormFile(formData.Name, formData.FileName)
            if err != nil {
                return nil, "", err
            }
            part.Write(formData.FileData)
        } else {
            writer.WriteField(formData.Name, formData.Value)
        }
    }
    if err := writer.Close(); err != nil {
        return nil, "", err
    }
    return body, contentType, nil
}
