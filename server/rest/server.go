package rest

import (
	"github.com/gin-gonic/gin"
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

	g := gin.New()
	s.Engine = g

	// add middleware
	g.Use(gin.Recovery())
	// add routers
	s.addRouters()

	return s
}