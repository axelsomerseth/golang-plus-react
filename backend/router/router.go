package router

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/axelsomerseth/golang-plus-react/backend/controllers"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.ForwardedByClientIP = true
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.AllowAll())

	// TODO: check this in deployment. i.e. 192.168.1.2
	router.SetTrustedProxies([]string{""})

	// Serve React App
	// TODO: Issues to be resolved
	// router.Static("/web", "./web/build")
	// router.Static("/static", "./web/build/static")

	// Works
	router.StaticFS("/web", gin.Dir(filepath.Join("web", "build"), false))
	router.StaticFS("/static", gin.Dir(filepath.Join("web", "build", "static"), false))
	router.StaticFile("/manifest.json", filepath.Join("web", "build", "manifest.json"))
	router.StaticFile("/logo192.png", filepath.Join("web", "build", "logo192.png"))
	router.StaticFile("/logo512.png", filepath.Join("web", "build", "logo512.png"))
	router.StaticFile("/favicon.icon", filepath.Join("web", "build", "favicon.ico"))
	router.StaticFile("/robots.txt", filepath.Join("web", "build", "robots.txt"))

	// Health check
	router.GET("/health", controllers.Status)

	// TODO: define routes here
	// api := router.Group("api")

	return router
}
