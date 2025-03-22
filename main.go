package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/suhailkassar11/go-crud/controllers"
	"github.com/suhailkassar11/go-crud/initializers"
)

func init() {
	// Load .env file
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("hello world")
	router := gin.Default()
	router.POST("/v1/user", controllers.CreateUser)
	router.GET("/v1/users", controllers.FindAllUser)
	router.GET("/v1/user/:id", controllers.FindOneUser)
	router.DELETE("/v1/user/:id", controllers.DeleteUser)
	router.Run() // listen and serve on 0.0.0.0:8080
}
