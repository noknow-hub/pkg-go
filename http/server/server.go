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
type MultiHostsTcpServer struct {
    Address string
    CertFile string
    IsFcgi bool
    KeyFile string
    Port string
    *MultiHostServerMuxes
}
type MultiHostsUnixServer struct {
    GroupId int
    IsFcgi bool
    *MultiHostServerMuxes
    SocketPath string
    UserId int
}
type MultiHostServerMuxes struct {
    HostHandlers map[string]http.Handler
    RedirectUrlIfNotFound string
}


//////////////////////////////////////////////////////////////////////
// Wrap ServeHTTP
//////////////////////////////////////////////////////////////////////
func (o MultiHostServerMuxes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    handler, ok := o.HostHandlers[r.Host]
    if !ok {
        http.Redirect(w, r, o.RedirectUrlIfNotFound, http.StatusMovedPermanently)
        return
    }
    handler.ServeHTTP(w, r)
}


//////////////////////////////////////////////////////////////////////
// New MultiHostsTcpServer.
//////////////////////////////////////////////////////////////////////
func NewMultiHostsTcpServer(address, port string, serverMuxes map[string]*http.ServeMux, redirectUrlIfNotFound string) *MultiHostsTcpServer {
    multiHostServerMuxes := &MultiHostServerMuxes{
        HostHandlers: make(map[string]http.Handler),
        RedirectUrlIfNotFound: redirectUrlIfNotFound,
    }
    for host, mux := range serverMuxes {
        multiHostServerMuxes.HostHandlers[host] = mux
    }
    return &MultiHostsTcpServer{
        Address: address,
        Port: port,
        MultiHostServerMuxes: multiHostServerMuxes,
    }
}


//////////////////////////////////////////////////////////////////////
// New MultiHostsUnixServer.
//////////////////////////////////////////////////////////////////////
func NewMultiHostsUnixServer(socketPath string, serverMuxes map[string]*http.ServeMux, redirectUrlIfNotFound string) *MultiHostsUnixServer {
    multiHostServerMuxes := &MultiHostServerMuxes{
        HostHandlers: make(map[string]http.Handler),
        RedirectUrlIfNotFound: redirectUrlIfNotFound,
    }
    for host, mux := range serverMuxes {
        multiHostServerMuxes.HostHandlers[host] = mux
    }
    return &MultiHostsUnixServer{
        MultiHostServerMuxes: multiHostServerMuxes,
        SocketPath: socketPath,
    }
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
// Run with MultiHostsTcpServer
//////////////////////////////////////////////////////////////////////
func (s *MultiHostsTcpServer) Run() error {
    listener, err := net.Listen("tcp", s.Address + ":" + s.Port)
    if err != nil {
       return err
    }
    defer listener.Close()
    go func(){
        shutdown(listener)
    }()
    if s.IsFcgi {
        if err := fcgi.Serve(listener, s.MultiHostServerMuxes); err != nil {
            return err
        }
    } else {
        if s.CertFile != "" && s.KeyFile != "" {
            if err := http.ServeTLS(listener, s.MultiHostServerMuxes, s.CertFile, s.KeyFile); err != nil {
                return err
            }
        } else {
            if err := http.Serve(listener, s.MultiHostServerMuxes); err != nil {
                return err
            }
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set FCGI.
//////////////////////////////////////////////////////////////////////
func (s *MultiHostsTcpServer) SetFcgi() *MultiHostsTcpServer {
    s.IsFcgi = true
    return s
}


//////////////////////////////////////////////////////////////////////
// Set TLS configuration.
//////////////////////////////////////////////////////////////////////
func (s *MultiHostsTcpServer) SetTls(certFile, keyFile string) *MultiHostsTcpServer {
    s.CertFile = certFile
    s.KeyFile = keyFile
    return s
}


//////////////////////////////////////////////////////////////////////
// Run with MultiHostsUnixServer
//////////////////////////////////////////////////////////////////////
func (s *MultiHostsUnixServer) Run() error {
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
        if err := fcgi.Serve(listener, s.MultiHostServerMuxes); err != nil {
            return err 
        }   
    } else {
        if err := http.Serve(listener, s.MultiHostServerMuxes); err != nil {
            return err
        }
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Set FCGI.
//////////////////////////////////////////////////////////////////////
func (s *MultiHostsUnixServer) SetFcgi() *MultiHostsUnixServer {
    s.IsFcgi = true
    return s
}


//////////////////////////////////////////////////////////////////////
// Set owner.
//////////////////////////////////////////////////////////////////////
func (s *MultiHostsUnixServer) SetOwner(userId, groupId int) *MultiHostsUnixServer {
    s.UserId = userId
    s.GroupId = groupId
    return s
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
