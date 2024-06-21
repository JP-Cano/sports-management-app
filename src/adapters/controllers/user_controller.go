package controllers

import (
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/validators"
	"github.com/JP-Cano/sports-management-app/src/services"
	"github.com/JP-Cano/sports-management-app/src/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user entities.CreateUserDto
	validators.ValidateBindJSON(&user, c)

	createdUser, err := u.UserService.CreateUser(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, createdUser)
}

func (u *UserController) GetAllUser(c *gin.Context) {
	users, err := u.UserService.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err.Error())
		utils.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, users)
}
