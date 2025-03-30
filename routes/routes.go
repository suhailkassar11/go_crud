package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailkassar11/go-crud/controllers"
	"github.com/suhailkassar11/go-crud/middleware"
)

func SetupUserRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	protected := r.Group("/").Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", controllers.FindAllUser)
		protected.GET("/user/:id", controllers.FindOneUser)
		protected.PUT("/user/:id", controllers.UpdateUser)
		protected.DELETE("/user/:id", controllers.DeleteUser)
	}
}
