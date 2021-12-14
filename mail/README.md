# noknow-hub/pkg-go/mail

## Getting Started

### Import

```go
import (
    "github.com/noknow-hub/pkg-go/mail"
)
```

## Usage

### Basic usage

```go
host := "your.samtp.server"
port := 465
fromEmail := "from@example.com"
fromName := "Service"
toEmail := "to@example.com"
subject := "[Subject] Test email"
authName := "auth_user"
authPass := "auth_pass"

smtpClient := mail.NewSmtpClient(host, port, fromEmail, toEmail, subject).
    SetFromName(fromName).
    SetAuthPlain(authName, authPass, host).
    SetTlsConfig(host)

// Text body.


if err := smtpClient.Send(); err != nil {
    // Error handling.
}
```
