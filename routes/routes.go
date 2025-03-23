package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailkassar11/go-crud/controllers"
)

func SetupUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", controllers.CreateUser)
		users.GET("", controllers.FindAllUser)
		users.GET("/:id", controllers.FindOneUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
