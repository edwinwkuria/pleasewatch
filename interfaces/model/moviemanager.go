package moviemanager

import "pleasewatch/models"

type MovieManager interface {
	Add(movie models.Movie) bool
	Get() []models.Movie
}
