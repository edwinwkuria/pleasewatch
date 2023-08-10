package movieservice

import (
	"pleasewatch/models"
)

type MovieService struct {
	movies []models.Movie
}

func NewMovieManager() *MovieService {
	return &MovieService{}
}

func (ms *MovieService) Add(movie models.Movie) bool {
	ms.movies = append(ms.movies, movie)
	return true
}

func (ms *MovieService) Get() []models.Movie {
	return ms.movies
}
