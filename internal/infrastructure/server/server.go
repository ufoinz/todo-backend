package server

import (
	"fmt"
	"log"
	nethttp "net/http"
	"time"
)

type Config struct {
	Port   int
	Router nethttp.Handler
}

func Start(cfg Config) error {
	server := &nethttp.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      cfg.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on port %d", cfg.Port)
	return server.ListenAndServe()
}
