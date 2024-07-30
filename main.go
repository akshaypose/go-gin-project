package main

import (
	"go-gin-project/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") // Start the Gin server
}
