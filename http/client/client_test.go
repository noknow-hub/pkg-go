//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package client

import (
    "testing"
//    "bytes"
//    "io/ioutil"
//    "mime/multipart"
//    "net/http"
//    "strings"
//    "time"
)


func TestGenerateFormDataBody(t *testing.T) {
    url := "https://dummy.com"
    client := NewClient(url)
    client.Config.
        AddFormData("name1", "value1").
        AddFormData("name2", "value2")
    body, contentType, err := client.generateFormDataBody()

    t.Logf("[INFO] body: %s", body)
    t.Logf("[INFO] contentType: %s", contentType)
    if err != nil {
        t.Fatalf("[ERROR] %s", err)
    }
}
