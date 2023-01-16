//////////////////////////////////////////////////////////////////////
// smtp_client.go
//////////////////////////////////////////////////////////////////////
package mail

import (
    "bytes"
    "crypto/tls"
    "errors"
    "html/template"
    "net/smtp"
    "path"
    "strconv"
    "strings"
)

type SmtpClient struct {
    Attachments []*Attachment
    AuthConfig *AuthConfig
    Bcc []string
    BodyHtml *Body
    BodyText *Body
    Cc []string
    FromEmail string
    FromName string
    Host string
    MimeVersion string
    Port int
    Rcpts []string
    Subject string
    TlsConfig *tls.Config
    To []string
}


//////////////////////////////////////////////////////////////////////
// New SmtpClient.
//////////////////////////////////////////////////////////////////////
func NewSmtpClient(host string, port int, fromEmail, toEmail, subject string) *SmtpClient {
    return &SmtpClient{
        FromEmail: fromEmail,
        Host: host,
        MimeVersion: MIME_VERSION_1_0,
        Port: port,
        To: []string{ toEmail },
        Subject: subject,
    }
}


//////////////////////////////////////////////////////////////////////
// Set SMTP body text from files.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetBodyTextFromFiles(charset string, fileNames []string, params map[string]interface{}, funcMap template.FuncMap) error {
    var t *template.Template
    var err error

    if funcMap != nil {
        t, err = template.New(path.Base(fileNames[0])).Funcs(funcMap).ParseFiles(fileNames...)
    } else {
        t, err = template.New(path.Base(fileNames[0])).ParseFiles(fileNames...)
    }
    if err != nil {
        return err
    }
    buffer := new(bytes.Buffer)
    if err := t.Execute(buffer, params); err != nil {
        return err
    }
    c.BodyText = &Body{
        ContentType: CONTENT_TYPE_TEXT_PLAIN,
        Charset: charset,
        Data: buffer.String(),
    }

    return nil
}


//////////////////////////////////////////////////////////////////////
// Set SMTP body HTML from files.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetBodyHtmlFromFiles(charset string, fileNames []string, params map[string]interface{}, funcMap template.FuncMap) error {
    var t *template.Template
    var err error

    if funcMap != nil {
        t, err = template.New(path.Base(fileNames[0])).Funcs(funcMap).ParseFiles(fileNames...)
    } else {
        t, err = template.New(path.Base(fileNames[0])).ParseFiles(fileNames...)
    }
    if err != nil {
        return err
    }
    buffer := new(bytes.Buffer)
    if err := t.Execute(buffer, params); err != nil {
        return err
    }
    c.BodyHtml = &Body{
        ContentType: CONTENT_TYPE_TEXT_HTML,
        Charset: charset,
        Data: buffer.String(),
    }

    return nil
}


//////////////////////////////////////////////////////////////////////
// Set SMTP body text from string.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetBodyTextFromStr(charset string, text string, params map[string]interface{}, funcMap template.FuncMap) error {
    var t *template.Template
    var err error

    if funcMap != nil {
        t, err = template.New("t").Funcs(funcMap).Parse(text)
    } else {
        t, err = template.New("t").Parse(text)
    }
    if err != nil {
        return err
    }
    buffer := new(bytes.Buffer)
    if err := t.Execute(buffer, params); err != nil {
        return err
    }
    c.BodyText = &Body{
        ContentType: CONTENT_TYPE_TEXT_PLAIN,
        Charset: charset,
        Data: buffer.String(),
    }

    return nil
}


//////////////////////////////////////////////////////////////////////
// Set SMTP body HTML from string.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetBodyHtmlFromStr(charset string, text string, params map[string]interface{}, funcMap template.FuncMap) error {
    var t *template.Template
    var err error

    if funcMap != nil {
        t, err = template.New("t").Funcs(funcMap).Parse(text)
    } else {
        t, err = template.New("t").Parse(text)
    }
    if err != nil {
        return err
    }
    buffer := new(bytes.Buffer)
    if err := t.Execute(buffer, params); err != nil {
        return err
    }
    c.BodyHtml = &Body{
        ContentType: CONTENT_TYPE_TEXT_HTML,
        Charset: charset,
        Data: buffer.String(),
    }

    return nil
}


//////////////////////////////////////////////////////////////////////
// Set SMTP from name.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetFromName(name string) *SmtpClient {
    c.FromName = name
    return c
}


//////////////////////////////////////////////////////////////////////
// Append an attachment.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) AppendAttachment(base64EncodedData, contentType, fileName string) *SmtpClient {
    c.Attachments = append(c.Attachments, &Attachment{
        Base64EncodedData: base64EncodedData,
        ContentType: contentType,
        FileName: fileName,
    })
    return c
}


//////////////////////////////////////////////////////////////////////
// Append an email into to.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) AppendTo(to string) *SmtpClient {
    c.To = append(c.To, to)
    return c
}


//////////////////////////////////////////////////////////////////////
// Append an email into cc.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) AppendCc(cc string) *SmtpClient {
    c.Cc = append(c.Cc, cc)
    return c
}


//////////////////////////////////////////////////////////////////////
// Append an email into bcc.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) AppendBcc(bcc string) *SmtpClient {
    c.Bcc = append(c.Bcc, bcc)
    return c
}


//////////////////////////////////////////////////////////////////////
// Set SMTP auth for PLAIN.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetAuthPlain(userName, password, host string) *SmtpClient {
    c.AuthConfig = &AuthConfig{
        PlainAuth: &PlainAuth{
            UserName: userName,
            Password: password,
            Host: host,
        },
    }
    return c
}


//////////////////////////////////////////////////////////////////////
// Set SMTP auth for CRAM-MD5.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetAuthCramMd5(userName, secret string) *SmtpClient {
    c.AuthConfig = &AuthConfig{
        Crammd5Auth: &CRAMMD5Auth{
            UserName: userName,
            Secret: secret,
        },
    }
    return c
}


//////////////////////////////////////////////////////////////////////
// Set SMTP TLS.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetTlsConfig(serverName string) *SmtpClient {
    c.TlsConfig = &tls.Config{
        ServerName: serverName,
    }
    return c
}


//////////////////////////////////////////////////////////////////////
// Set certificate files into the TLS config.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetCertFiles(certFile string, keyFile string) error {
    if c.TlsConfig == nil {
        return errors.New("SmtpClient.TlsConfig is nil.")
    }
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return err
    }
    c.TlsConfig.Certificates = []tls.Certificate{cert}
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set certificate bytes into the TLS config.
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) SetCertBytes(certPem []byte, keyPem []byte) error {
    if c.TlsConfig == nil {
        return errors.New("SmtpClient.TlsConfig is nil.")
    }
    cert, err := tls.X509KeyPair(certPem, keyPem)
    if err != nil {
        return err
    }
    c.TlsConfig.Certificates = []tls.Certificate{cert}
    return nil
}


//////////////////////////////////////////////////////////////////////
// Send Email
//////////////////////////////////////////////////////////////////////
func (c *SmtpClient) Send() error {
    var sc *smtp.Client
    var err error

    // Connect.
    if c.TlsConfig != nil {
        conn, err := tls.Dial("tcp", c.Host + ":" + strconv.Itoa(c.Port), c.TlsConfig)
        if err != nil {
            return err
        }
        sc, err = smtp.NewClient(conn, c.Host)
        if err != nil {
            return err
        }
    } else {
        sc, err = smtp.Dial(c.Host + ":" + strconv.Itoa(c.Port))
        if err != nil {
            return err
        }
    }
    defer sc.Close()

    // Authentication.
    if c.AuthConfig != nil {
        if c.AuthConfig.Crammd5Auth != nil {
            a := smtp.CRAMMD5Auth(c.AuthConfig.Crammd5Auth.UserName, c.AuthConfig.Crammd5Auth.Secret)
            if err = sc.Auth(a); err != nil {
                return err
            }
        }
        if c.AuthConfig.PlainAuth != nil {
            a := smtp.PlainAuth("", c.AuthConfig.PlainAuth.UserName, c.AuthConfig.PlainAuth.Password, c.AuthConfig.PlainAuth.Host)
            if err = sc.Auth(a); err != nil {
                return err
            }
        }
    }

    // From.
    if err = sc.Mail(c.FromEmail); err != nil {
        return err
    }

    // Rcpt.
    for _, o := range c.To {
        if err = sc.Rcpt(o); err != nil {
            return err
        }
    }
    for _, o := range c.Cc {
        if err = sc.Rcpt(o); err != nil {
            return err
        }
    }
    for _, o := range c.Bcc {
        if err = sc.Rcpt(o); err != nil {
            return err
        }
    }

    // Data.
    if c.BodyHtml == nil && c.BodyText == nil {
        return errors.New("SmtpClient.BodyHtml or SmtpClient.BodyText is mandatory.")
    }

    wc, err := sc.Data()
    if err != nil {
        return err
    }
    body := make([]byte, 0)

    // Header.
    var from string
    if c.FromName == "" {
        from = c.FromEmail
    } else {
        from = c.FromName + " <" + c.FromEmail + ">"
    }
    body = append(body, "From: " + from + "\r\n"...)
    body = append(body, "To: " + strings.Join(c.To, ",") + "\r\n"...)
    if len(c.Cc) > 0 {
        body = append(body, "Cc: " + strings.Join(c.Cc, ",") + "\r\n"...)
    }
    body = append(body, "Subject: " + c.Subject + "\r\n"...)
    body = append(body, "MIME-version: " + c.MimeVersion + "\r\n"...)
    var mixedBoundary string
    if len(c.Attachments) > 0 {
        mixedBoundary = genBoundary()
        body = append(body, "Content-Type: multipart/mixed; boundary=\"" + mixedBoundary + "\"\r\n"...)
        body = append(body, "--" + mixedBoundary + "\r\n"...)
    }

    // Body.
    var alternativeBoundary string
    if c.BodyHtml != nil && c.BodyText != nil {
        alternativeBoundary = genBoundary()
        body = append(body, "Content-Type: multipart/alternative; boundary=\"" + alternativeBoundary + "\"\r\n"...)
    }
    if c.BodyText != nil {
        if c.BodyHtml != nil && c.BodyText != nil {
            body = append(body, "--" + alternativeBoundary + "\r\nContent-Type: " + c.BodyText.ContentType + "; charset=\"" + c.BodyText.Charset + "\"\r\n\r\n" + c.BodyText.Data + "\r\n"...)
        } else {
            body = append(body, "Content-Type: " + c.BodyText.ContentType + "; charset=\"" + c.BodyText.Charset + "\"\r\n\r\n" + c.BodyText.Data + "\r\n"...)
        }
    }
    if c.BodyHtml != nil {
        if c.BodyHtml != nil && c.BodyText != nil {
            body = append(body, "--" + alternativeBoundary + "\r\nContent-Type: " + c.BodyHtml.ContentType + "; charset=\"" + c.BodyHtml.Charset + "\"\r\n\r\n" + c.BodyHtml.Data + "\r\n"...)
        } else {
            body = append(body, "Content-Type: " + c.BodyHtml.ContentType + "; charset=\"" + c.BodyHtml.Charset + "\"\r\n\r\n" + c.BodyHtml.Data + "\r\n"...)
        }
    }
    if c.BodyHtml != nil && c.BodyText != nil {
        body = append(body, "--" + alternativeBoundary + "--\r\n"...)
    }
    if len(c.Attachments) > 0 {
        for _, attachment := range c.Attachments {
            body = append(body, "--" + mixedBoundary + "\r\nContent-Type: " + attachment.ContentType + "; name=\"" + attachment.FileName + "\"\r\n"...)
            body = append(body, "Content-Disposition: attachment; filename=\"" + attachment.FileName + "\"\r\n"...)
            body = append(body, "Content-Transfer-Encoding: base64\r\n\r\n"...)
            body = append(body, attachment.Base64EncodedData + "\r\n"...)
        }
        body = append(body, "--" + mixedBoundary + "--\r\n"...)
    }
    if _, err = wc.Write(body); err != nil {
        return err
    }
    if err = wc.Close(); err != nil {
        return err
    }

    // Quit.
    if err = sc.Quit(); err != nil {
        return err
    }

    return nil
}
