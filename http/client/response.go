//////////////////////////////////////////////////////////////////////
// response.go
//////////////////////////////////////////////////////////////////////
package client

import (
    "crypto/tls"
    "net/http"
)


type Response struct {
    Body []byte
    ContentLength int64
    Header http.Header
    Proto string
    ProtoMajor int
    ProtoMinor int
    Status string
    StatusCode int
    TransferEncoding []string
    Uncompressed bool
    Trailer http.Header
    Request *http.Request
    TLS *tls.ConnectionState
}
