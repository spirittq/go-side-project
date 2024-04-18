package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	notify     chan error
	timeout    time.Duration
}

func New(handler http.Handler, address string, timeout int) *Server {
	httpServer := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	server := &Server{
		httpServer: httpServer,
		notify:     make(chan error, 1),
		timeout:    time.Duration(timeout),
	}

	server.start()

	return server
}

func (server *Server) start() {
	go func() {
		server.notify <- server.httpServer.ListenAndServe()
		close(server.notify)
	}()
}

func (server *Server) Notify() <-chan error {
	return server.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
