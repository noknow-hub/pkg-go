//////////////////////////////////////////////////////////////////////
// server.go
//////////////////////////////////////////////////////////////////////
package server

import (
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

const (
    ADDRESS_LOCAL_HOST = ""
    SOCKET_FILE_MODE = 0777
)

type UnixServer struct {
    GroupId int
    ServerMux *http.ServeMux
    SocketPath string
    UserId int
}

type TcpServer struct {
    Address string
    CertFile string
    KeyFile string
    Port string
    ServerMux *http.ServeMux
}


//////////////////////////////////////////////////////////////////////
// New UnixServer.
//////////////////////////////////////////////////////////////////////
func NewUnixServer(socketPath string, serverMux *http.ServeMux) *UnixServer {
    return &UnixServer{
        ServerMux: serverMux,
        SocketPath: socketPath,
    }
}


//////////////////////////////////////////////////////////////////////
// New TcpServer.
//////////////////////////////////////////////////////////////////////
func NewTcpServer(address, port string, serverMux *http.ServeMux) *TcpServer {
    return &TcpServer{
        Address: address,
        Port: port,
        ServerMux: serverMux,
    }
}


//////////////////////////////////////////////////////////////////////
// Run with UNIX socket.
//////////////////////////////////////////////////////////////////////
func (o *UnixServer) Run() error {
    listener, err := net.Listen("unix", o.SocketPath)
    if err != nil {
        return err
    }
    defer listener.Close()
    if o.UserId != 0 && o.GroupId != 0 {
        if err := os.Chown(o.SocketPath, o.UserId, o.GroupId); err != nil {
            return err
        }
        if err := os.Chmod(o.SocketPath, SOCKET_FILE_MODE); err != nil {
            return err
        }
    }
    go func() {
        shutdown(listener)
    }()
    if err := http.Serve(listener, o.ServerMux); err != nil {
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set owner.
//////////////////////////////////////////////////////////////////////
func (o *UnixServer) SetOwner(userId, groupId int) *UnixServer {
    o.UserId = userId
    o.GroupId = groupId
    return o
}


//////////////////////////////////////////////////////////////////////
// Run with TCP.
//////////////////////////////////////////////////////////////////////
func (o *TcpServer) Run() error {
    listener, err := net.Listen("tcp", o.Address + ":" + o.Port)
    if err != nil {
       return err
    }
    defer listener.Close()
    go func(){
        shutdown(listener)
    }()
    if o.CertFile != "" && o.KeyFile != "" {
        if err := http.ServeTLS(listener, o.ServerMux, o.CertFile, o.KeyFile); err != nil {
            return err
        }
    } else {
        if err := http.Serve(listener, o.ServerMux); err != nil {
            return err
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set TLS configuration.
//////////////////////////////////////////////////////////////////////
func (o *TcpServer) SetTls(certFile, keyFile string) *TcpServer {
    o.CertFile = certFile
    o.KeyFile = keyFile
    return o
}


//////////////////////////////////////////////////////////////////////
// Shutdown HTTP server.
//////////////////////////////////////////////////////////////////////
func shutdown(listener net.Listener) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    <-c
    listener.Close()
}
