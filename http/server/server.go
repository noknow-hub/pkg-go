//////////////////////////////////////////////////////////////////////
// server.go
//////////////////////////////////////////////////////////////////////
package server

import (
    "net"
    "net/http"
    "net/http/fcgi"
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
    IsFcgi bool
    ServerMux *http.ServeMux
    SocketPath string
    UserId int
}

type TcpServer struct {
    Address string
    CertFile string
    IsFcgi bool
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
func (s *UnixServer) Run() error {
    listener, err := net.Listen("unix", s.SocketPath)
    if err != nil {
        return err
    }
    defer listener.Close()
    if s.UserId != 0 && s.GroupId != 0 {
        if err := os.Chown(s.SocketPath, s.UserId, s.GroupId); err != nil {
            return err
        }
        if err := os.Chmod(s.SocketPath, SOCKET_FILE_MODE); err != nil {
            return err
        }
    }
    go func() {
        shutdown(listener)
    }()
    if s.IsFcgi {
        if err := fcgi.Serve(listener, s.ServerMux); err != nil {
            return err 
        }   
    } else {
        if err := http.Serve(listener, s.ServerMux); err != nil {
            return err
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set FCGI.
//////////////////////////////////////////////////////////////////////
func (s *UnixServer) SetFcgi() *UnixServer {
    s.IsFcgi = true
    return s
}


//////////////////////////////////////////////////////////////////////
// Set owner.
//////////////////////////////////////////////////////////////////////
func (s *UnixServer) SetOwner(userId, groupId int) *UnixServer {
    s.UserId = userId
    s.GroupId = groupId
    return s
}


//////////////////////////////////////////////////////////////////////
// Run with TCP.
//////////////////////////////////////////////////////////////////////
func (s *TcpServer) Run() error {
    listener, err := net.Listen("tcp", s.Address + ":" + s.Port)
    if err != nil {
       return err
    }
    defer listener.Close()
    go func(){
        shutdown(listener)
    }()
    if s.IsFcgi {
        if err := fcgi.Serve(listener, s.ServerMux); err != nil {
            return err
        }
    } else {
        if s.CertFile != "" && s.KeyFile != "" {
            if err := http.ServeTLS(listener, s.ServerMux, s.CertFile, s.KeyFile); err != nil {
                return err
            }
        } else {
            if err := http.Serve(listener, s.ServerMux); err != nil {
                return err
            }
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set FCGI.
//////////////////////////////////////////////////////////////////////
func (s *TcpServer) SetFcgi() *TcpServer {
    s.IsFcgi = true
    return s
}


//////////////////////////////////////////////////////////////////////
// Set TLS configuration.
//////////////////////////////////////////////////////////////////////
func (s *TcpServer) SetTls(certFile, keyFile string) *TcpServer {
    s.CertFile = certFile
    s.KeyFile = keyFile
    return s
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
