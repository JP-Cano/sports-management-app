package utils

import (
	"github.com/JP-Cano/sports-management-app/src/core/exceptions"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, message *string) {
	if message == nil {
		internalServerError := exceptions.InternalServerError.Error()
		message = &internalServerError
	}
	c.JSON(code, gin.H{"status": "error", "message": message})
	return
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{"status": "success", "data": data})
	return
}
