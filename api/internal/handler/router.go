package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter wires all routes and middleware into a Gin engine.
func NewRouter(postHandler *PostHandler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:14200"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	{
		api.GET("/posts/:id", postHandler.GetPost)
		api.GET("/posts/:id/comments", postHandler.ListComments)
		api.POST("/posts/:id/comments", postHandler.AddComment)
	}

	return router
}
