package health

import (
	"github.com/JP-Cano/sports-management-app/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Check struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Check {
	return &Check{DB: db}
}

func (hc *Check) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dbStatus string

		if err := hc.DB.Raw("SELECT 1").Scan(&dbStatus).Error; err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
			log.Printf("Error connecting to database: %v", err.Error())
			return
		}

		response := map[string]interface{}{"message": map[string]string{
			"Server":   "Up and running",
			"Database": "Up and running",
		}}
		utils.SuccessResponse(c, 0, response)
	}
}
