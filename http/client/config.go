//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package client

import (
    "net/url"
)

type Config struct {
    Cred string
    Body []byte
    FormData []*MuitipartFormData
    Header map[string]string
    JsonData []byte
    TimeoutSec int
    Url string
    UrlQueryData url.Values
    XWwwFormUrlencodedData url.Values
}

type MuitipartFormData struct {
    FileContentType string
    FileData []byte
    FileName string
    Name string
    Value string
}


//////////////////////////////////////////////////////////////////////
// Add a form data.
//////////////////////////////////////////////////////////////////////
func (c *Config) AddFormData(name, value string) *Config {
    c.FormData = append(c.FormData, &MuitipartFormData{
        Name: name,
        Value: value,
    })
    return c
}


//////////////////////////////////////////////////////////////////////
// Add a form data with file.
//////////////////////////////////////////////////////////////////////
func (c *Config) AddFormDataWithFile(name, value, fileName, fileContentType string, fileData []byte) *Config {
    c.FormData = append(c.FormData, &MuitipartFormData{
        FileContentType: fileContentType,
        FileData: fileData,
        FileName: fileName,
        Name: name,
        Value: value,
    })
    return c
}


//////////////////////////////////////////////////////////////////////
// Add a header key-value pair.
//////////////////////////////////////////////////////////////////////
func (c *Config) AddHeader(key, value string) *Config {
    if c.Header == nil {
        c.Header = map[string]string{}
    }
    c.Header[key] = value
    return c
}


//////////////////////////////////////////////////////////////////////
// Add a header key-value pair for "Authorization".
//////////////////////////////////////////////////////////////////////
func (c *Config) AddHeaderAuthorization(value string) *Config {
    return c.AddHeader(HTTP_HEADER_AUTHORIZATION, value)
}


//////////////////////////////////////////////////////////////////////
// Add a header key-value pair for "Content-Type".
//////////////////////////////////////////////////////////////////////
func (c *Config) AddHeaderContentType(value string) *Config {
    return c.AddHeader(HTTP_HEADER_CONTENT_TYPE, value)
}


//////////////////////////////////////////////////////////////////////
// Add a URL query data.
//////////////////////////////////////////////////////////////////////
func (c *Config) AddUrlQueryData(key, value string) *Config {
    if c.UrlQueryData == nil {
        c.UrlQueryData = url.Values{}
    }
    c.UrlQueryData.Add(key, value)
    return c
}


//////////////////////////////////////////////////////////////////////
// Add a x-www-form-urlencoded data.
//////////////////////////////////////////////////////////////////////
func (c *Config) AddXWwwFormUrlencodedData(key, value string) *Config {
    if c.XWwwFormUrlencodedData == nil {
        c.XWwwFormUrlencodedData = url.Values{}
    }
    c.XWwwFormUrlencodedData.Add(key, value)
    return c
}


//////////////////////////////////////////////////////////////////////
// Set JsonData
//////////////////////////////////////////////////////////////////////
func (c *Config) SetJsonData(data []byte) *Config {
    c.JsonData = data
    return c
}
