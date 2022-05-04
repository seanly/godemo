package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"

	_ "godemo/docs"
)

// @BasePath /api
// @title CruiseShip Service REST APIs
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @description The API doc is targeting for CruiseShip developers.
func (s *engine) injectRouterGroup(router *gin.RouterGroup) {
	router.GET("/ping", Pong)

	router.GET("/apidocs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}

// Pong
// @Router /api/ping [GET]
// @Summary check ping
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "map[string]string - {message: 'pong'}"
// @Description response ping check.
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}