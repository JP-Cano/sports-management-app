package routes

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpFile(router *gin.Engine, controller *controllers.FileController) {
	fileRoutes := router.Group("/api/v1/files")
	{
		fileRoutes.POST("/upload", controller.UploadExcel)
	}
}
