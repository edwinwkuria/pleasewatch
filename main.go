package main

import (
	"pleasewatch/routes"
)

func main() {
	r := routes.SetupRouter()

	// Run the server
	r.Run(":8080")
}
