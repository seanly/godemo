package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type engine struct {
	*gin.Engine

	mode string
}

// NewEngine create engine object and add routes
func NewEngine(mode string) *engine {
	s := &engine{mode: mode}
	gin.SetMode(s.mode)

	g := gin.New()
	defer func() {
		s.Engine = g
	}()

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return s
}

// Serve http serve
func Serve(ctx context.Context) error {

	address := ":8000"
	if port := os.Getenv("PORT"); port != "" {
		log.Printf("Environment variable PORT=%s", port)
		address = ":" + port
	}

	log.Printf("Serve[%s] started at %s", address, time.Now())

	engine := NewEngine("test")
	server := &http.Server{
		Addr: address, Handler: engine,
	}

	stopChan := make(chan struct{})

	go func() {
		defer close(stopChan)

		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to stop server, err: %s", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start http server, error: %s", err)
		return err
	}

	<-stopChan
	return nil
}
