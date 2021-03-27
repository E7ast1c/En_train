package server

import (
	"context"
	"net/http"
	"time"

	"en_train/internal/config"
)

type HttpServer struct {
	httpServer *http.Server

}

func (s *HttpServer) Run(config config.ApiConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + config.Port,
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MB as default
	}

	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}