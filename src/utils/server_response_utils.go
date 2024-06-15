package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"status": "error", "message": message})
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	if code == 0 {
		code = http.StatusOK
	}
	c.JSON(code, gin.H{"status": "success", "data": data})
}
