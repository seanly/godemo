package rest

import (
	"github.com/gin-gonic/gin"
	"godemo/pkg/api"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"net/http"
)

type engine struct {
	*gin.Engine

	mode string
}

const (
	ApiRootPath = "/sapis"
)

var GroupVersion = schema.GroupVersion{Group: api.GroupName, Version: "v1alpha1"}

// NewEngine create engine object and add routes
func NewEngine(mode string) *engine {
	s := &engine{mode: mode}
	gin.SetMode(s.mode)

	g := gin.New()
	s.Engine = g

	// add middleware
	g.Use(gin.Recovery())
	// add routers
	s.injectRoutes()
	s.printRouters()

	return s
}

func (s *engine) injectRoutes()  {

	g := s.Engine

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Invalid path: %s", c.Request.URL.Path)
	})

	g.HandleMethodNotAllowed = true
	g.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "Method not allowed: %s %s", c.Request.Method, c.Request.URL.Path)
	})

	apiRouters := g.Group("api")
	s.injectRouterGroup(apiRouters)
}

func (s *engine) printRouters() {
	g := s.Engine

	// debug: list route
	for _, router := range g.Routes() {
		log.Printf("path: %s", router.Path)
	}
}