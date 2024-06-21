package routes

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpUser(router *gin.Engine, controller *controllers.UserController) {
	userRoutes := router.Group("api/v1/users")
	{
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.GET("/", controller.GetAllUser)
		userRoutes.GET("/:id", controller.GetUserById)
		userRoutes.PATCH("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
		userRoutes.GET("/search", controller.SearchUser)
	}
}
