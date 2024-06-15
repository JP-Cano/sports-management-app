package routes

import (
	"github.com/JP-Cano/sports-management-app/src/health"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpHealthCheck(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/v1/health", health.New(db).Check())
}
