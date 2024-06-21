package validators

import (
	"github.com/JP-Cano/sports-management-app/src/core/exceptions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"sync"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func registerValidators() {
	once.Do(func() {
		validate = validator.New()
		registerCustomValidators()
	})
}

func registerCustomValidators() {
	customValidators := map[string]validator.Func{
		"uuid": UUID,
	}
	for key, f := range customValidators {
		err := validate.RegisterValidation(key, f)
		if err != nil {
			log.Printf("Error registering custom validator %s: %v", key, err)
			return
		}
	}
}

func BindJSON(entity interface{}, c *gin.Context) error {
	registerValidators()
	if err := c.ShouldBindJSON(&entity); err != nil {
		log.Printf("Error binding JSON: %v", err.Error())
		return exceptions.Throw(exceptions.BadRequest)
	}
	return nil
}

func UUID(fl validator.FieldLevel) bool {
	id := fl.Field().String()
	_, err := uuid.Parse(id)
	return err == nil
}
