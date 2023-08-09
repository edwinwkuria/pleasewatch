package webcontrollers

import (
	"net/http"
	"pleasewatch/models"

	"github.com/gin-gonic/gin"
)

var movies []models.Movie

// Add Movie
// Get Movie
// Get Movies
func GetMovies(c *gin.Context) {
	mov := models.Movie{
		ID:   1,
		Name: "Hello World",
	}
	movies = append(movies, mov)
	c.HTML(http.StatusOK, "index.tmpl", movies)
}

//Edit Movie
//Delete Movie
