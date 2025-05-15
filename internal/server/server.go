package server

import (
	"context"
	"net/http"
	"premiesPortal/internal/security"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	serverParams := security.AppSettings.ServerParams

	s.httpServer = &http.Server{
		Addr:           serverParams.Addr + port,
		Handler:        handler,
		MaxHeaderBytes: serverParams.MaxHeaderMBs * 1024 * 1024,
		ReadTimeout:    time.Duration(serverParams.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(serverParams.WriteTimeout) * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
