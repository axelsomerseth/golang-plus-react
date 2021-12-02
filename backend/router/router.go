package router

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/axelsomerseth/varsity-dev-challenge/backend/controllers"
)

func Setup() *gin.Engine {
	router := gin.New()
	router.RedirectTrailingSlash = true
	router.ForwardedByClientIP = true
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.AllowAll())

	// TODO: check this in deployment.
	router.SetTrustedProxies([]string{"192.168.1.2"})

	router.GET("/health", controllers.Status)

	// TODO: define routes here
	// api := router.Group("api")

	return router
}
