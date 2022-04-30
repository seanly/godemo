package server

import (
	"context"
	"godemo/server/rest"
	"log"
	"net/http"
	"os"
	"time"
)

// Serve http serve
func Serve(ctx context.Context) error {

	address := ":8000"
	if port := os.Getenv("PORT"); port != "" {
		log.Printf("Environment variable PORT=%s", port)
		address = ":" + port
	}

	log.Printf("Serve[%s] started at %s", address, time.Now())

	engine := rest.NewEngine("test")
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
