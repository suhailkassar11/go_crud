package main

import (
	"github.com/suhailkassar11/go-crud/initializers"
	"github.com/suhailkassar11/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
