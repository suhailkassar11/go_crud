package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suhailkassar11/go-crud/initializers"
	"github.com/suhailkassar11/go-crud/routes"
)

func init() {
	// Load .env file
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	routes.SetupUserRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
