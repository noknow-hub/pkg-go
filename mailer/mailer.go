// mailer.go
package mailer

import (
    "html/template"
    "math/rand"
    "time"
)

const (
    CharsetIso2022Jp         = "iso-2022-jp"
    CharsetUsAscii           = "us-ascii"
    CharsetUtf8              = "UTF-8"
    ContentTypeTextHtml      = "text/html"
    ContentTypeTextPlain     = "text/plain"
    ContentTypeTextRichText  = "text/richtext"
    ContentTypeTextXWhatever = "text/x-whatever"
    MimeVersion10            = "1.0"
)


type Attachment struct {
    Base64EncodedData string
    ContentType       string
    FileName          string
}

type AuthConfig struct {
    Crammd5Auth *CRAMMD5Auth
    PlainAuth   *PlainAuth
}

type CRAMMD5Auth struct {
    UserName string
    Secret   string
}

type PlainAuth struct {
    UserName string
    Password string
    Host     string
}

type Body struct {
    ContentType string
    Charset     string
    Data        string
}


// Set funcMap for SafeHtml
func SetFuncMapSafeHtml(f template.FuncMap) {
    f["SafeHtml"] = func(s string) template.HTML {
        return template.HTML(s)
    }
}


// Generate a radom value for boundary
func GenerateBoundary() string {
    charset := "1234567890abcdefghijklmnopqrstuvwxyz"
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, 32)
    for i := range b {
        b[i] = charset[r.Intn(len(charset))]
    }
    return string(b)
}
