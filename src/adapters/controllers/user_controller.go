package controllers

import (
	"errors"
	"github.com/JP-Cano/sports-management-app/src/application/services"
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"github.com/JP-Cano/sports-management-app/src/core/exceptions"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/utils"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/validators"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if err := validators.BindJSON(&user, c); err != nil {
		e := err.Error()
		utils.ErrorResponse(c, http.StatusBadRequest, &e)
		return
	}
	createdUser, err := u.UserService.CreateUser(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, createdUser)
}

func (u *UserController) GetAllUser(c *gin.Context) {
	users, err := u.UserService.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err.Error())
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, users)
}

func (u *UserController) GetUserById(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	user, err := u.UserService.GetUserById(id)
	if err != nil {
		if errors.Is(err, exceptions.NotFound) {
			e := err.Error()
			utils.ErrorResponse(c, http.StatusNotFound, &e)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var user entities.UpdateUserDto
	id := uuid.MustParse(c.Param("id"))

	if err := validators.BindJSON(&user, c); err != nil {
		e := err.Error()
		utils.ErrorResponse(c, http.StatusBadRequest, &e)
		return
	}

	if err := u.UserService.UpdateUser(id, user); err != nil {
		if errors.Is(err, exceptions.NotFound) {
			e := err.Error()
			utils.ErrorResponse(c, http.StatusNotFound, &e)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := u.UserService.DeleteUsers(id); err != nil {
		if errors.Is(err, exceptions.NotFound) {
			e := err.Error()
			utils.ErrorResponse(c, http.StatusNotFound, &e)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (u *UserController) SearchUser(c *gin.Context) {
	value := c.Query("q")
	users, err := u.UserService.SearchUsers(value)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, users)
}
