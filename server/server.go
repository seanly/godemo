package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type engine struct {
	*gin.Engine

	mode string
}

// NewEngine create engine object and add routes
func NewEngine(mode string) *engine {
	s := &engine{mode: mode}
	gin.SetMode(s.mode)

	g := s.Engine
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return s
}

// Serve http serve
func Serve(ctx context.Context) error {

	engine := NewEngine("test")
	server := &http.Server{
		Addr: ":3000", Handler: engine,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start http server, error: %s", err)
		return err
	}

	return nil
}
