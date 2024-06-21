package validators

import (
	"github.com/JP-Cano/sports-management-app/src/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ValidateBindJSON(entity interface{}, c *gin.Context) {
	if err := c.ShouldBindJSON(&entity); err != nil {
		log.Printf("Error binding JSON: %v", err.Error())
		utils.ErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}
}
