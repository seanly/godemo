package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

func Serve(ctx context.Context) error {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		return err
	}


	return nil
}
