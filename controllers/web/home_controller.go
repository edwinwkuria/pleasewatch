package webcontrollers

import (
	"net/http"
	moviemanager "pleasewatch/interfaces/model"
	"pleasewatch/models"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	moviemanager moviemanager.MovieManager
}

func NewHomeController(mm moviemanager.MovieManager) *HomeController {
	return &HomeController{moviemanager: mm}
}

// Add Movie
// Get Movie
// Get Movies
func (hc HomeController) GetMovies(c *gin.Context) {
	mov := models.Movie{
		ID:   1,
		Name: "Hello World",
	}
	hc.moviemanager.Add(mov)
	movies := hc.moviemanager.Get()
	c.HTML(http.StatusOK, "index.tmpl", movies)
}

//Edit Movie
//Delete Movie
