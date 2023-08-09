package routes

import (
	webcontrollers "pleasewatch/controllers/web"

	"github.com/gin-gonic/gin"
)

func SetupWebRouter(r *gin.Engine) {
	r.LoadHTMLGlob("./templates/*")

	// Define routes
	r.GET("/", webcontrollers.GetMovies)
}
