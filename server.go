package server

import (
	"crypto/tls"
	"net"
	"net/http"
)

func (server *Server) Stop() {
	server.listener.Close()
	<-server.stopped
}

func (server *Server) Start() error {

	listener, err := net.Listen("tcp", server.Addr)

	if err != nil {
		return err
	}

	server.Addr = listener.Addr().String()

	if server.TLSConfig != nil {
		listener = tls.NewListener(listener, server.TLSConfig)
	}

	server.listener = listener

	go func() {
		server.Serve(server.listener)
		close(server.stopped)
	}()

	return nil
}

func Init(addressName string) *Server {
	newServer := &Server{
		stopped: make(chan bool),
	}
	newServer.Addr = addressName
	return newServer
}

type Server struct {
	http.Server
	listener net.Listener
	stopped  chan bool
}
