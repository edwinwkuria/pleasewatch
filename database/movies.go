package database

import (
	"context"
	"fmt"
	"pleasewatch/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadMovies(c *gin.Context) []models.Movie {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	// Access the database and collection
	db := client.Database("todelete")
	collection := db.Collection("movies")

	//query db
	cursor, error := collection.Find(ctx, bson.D{})
	if error != nil {
		fmt.Println(error)
	}

	defer cursor.Close(ctx)

	// Loop through the documents and process them
	var results []models.Movie
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println(err)
	}
	// return
	return results
}
