package routes

import (
	webcontrollers "pleasewatch/controllers/web"
	movieservice "pleasewatch/services/model"

	"github.com/gin-gonic/gin"
)

func SetupWebRouter(r *gin.Engine) {
	r.LoadHTMLGlob("./templates/*")

	// Define routes
	ms := movieservice.NewMovieManager()
	webcontroller := webcontrollers.NewHomeController(ms)
	r.GET("/", webcontroller.GetMovies)
}
