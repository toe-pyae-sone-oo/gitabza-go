package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv  *http.Server
	addr string

	// used for graceful shutdown
	quit chan os.Signal // wait for system interrupt signals
	ttf  time.Duration  // time to finish currently handling requests
}

func New(addr string, ttf time.Duration) *Server {
	// TODO: add routes here
	r := gin.Default()
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	return &Server{
		srv:  srv,
		addr: addr,
		quit: make(chan os.Signal, 1),
		ttf:  ttf,
	}
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	signal.Notify(s.quit, syscall.SIGINT, syscall.SIGTERM)
	<-s.quit

	s.gracefulShutdown(ctx)
}

func (s *Server) gracefulShutdown(ctx context.Context) {
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, s.ttf)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v\n", err)
	}

	log.Println("server exiting...")
}
