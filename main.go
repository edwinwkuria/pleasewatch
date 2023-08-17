package main

import (
	"pleasewatch/database"
	"pleasewatch/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Try connect to database
	database.TryConnect()
	//configure API routes
	routes.SetupAPIRouter(r)
	//configure API routes
	routes.SetupWebRouter(r)
	// Run the server
	r.Run(":8080")
}
