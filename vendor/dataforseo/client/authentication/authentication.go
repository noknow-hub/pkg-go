//////////////////////////////////////////////////////////////////////
// authentication.go
//////////////////////////////////////////////////////////////////////
package authentication

import (
    "encoding/base64"
)


type Authentication struct {
    Login string
    Password string
    Cred string
}


//////////////////////////////////////////////////////////////////////
// New Authentication.
//////////////////////////////////////////////////////////////////////
func NewAuthentication(login, password string) *Authentication {
    cred := base64.StdEncoding.EncodeToString([]byte(login + ":" + password))
    return &Authentication{
        Login: login,
        Password: password,
        Cred: cred,
    }
}
