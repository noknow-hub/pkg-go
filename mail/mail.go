//////////////////////////////////////////////////////////////////////
// mail.go
//////////////////////////////////////////////////////////////////////
package mail

import (
    "html/template"
    "math/rand"
    "time"
)

const (
    CHARSET_ISO_2022_JP = "iso-2022-jp"
    CHARSET_US_ASCII = "us-ascii"
    CHARSET_UTF8 = "UTF-8"
    CONTENT_TYPE_TEXT_HTML = "text/html"
    CONTENT_TYPE_TEXT_PLAIN = "text/plain"
    CONTENT_TYPE_TEXT_RICHTEXT = "text/richtext"
    CONTENT_TYPE_TEXT_X_WHATEVER = "text/x-whatever"
    MIME_VERSION_1_0 = "1.0"
)


type Attachment struct {
    Base64EncodedData string
    ContentType string
    FileName string
}

type AuthConfig struct {
    Crammd5Auth *CRAMMD5Auth
    PlainAuth *PlainAuth
}

type CRAMMD5Auth struct {
    UserName string
    Secret string
}

type PlainAuth struct {
    UserName string
    Password string
    Host string
}

type Body struct {
    ContentType string
    Charset string
    Data string
}


//////////////////////////////////////////////////////////////////////
// Set funcMap for SafeHtml.
//////////////////////////////////////////////////////////////////////
func SetFuncMapSafeHtml(f template.FuncMap) {
    f["SafeHtml"] = func(s string) template.HTML {
        return template.HTML(s)
    }
}


//////////////////////////////////////////////////////////////////////
// Generate a radom value for boundary.
//////////////////////////////////////////////////////////////////////
func genBoundary() string {
    charset := "1234567890abcdefghijklmnopqrstuvwxyz"
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, 32)
    for i := range b {
        b[i] = charset[r.Intn(len(charset))]
    }
    return string(b)
}
