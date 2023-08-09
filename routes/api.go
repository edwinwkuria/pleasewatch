package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupAPIRouter(r *gin.Engine) {

	// Define routes
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, API from different file!"})
	})
}
