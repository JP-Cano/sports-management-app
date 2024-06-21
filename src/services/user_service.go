package services

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/repositories"
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"github.com/JP-Cano/sports-management-app/src/utils"
	"github.com/google/uuid"
	"log"
)

type UserService interface {
	CreateUser(user entities.CreateUserDto) (entities.UserDto, error)
	GetAllUsers() ([]entities.UserDto, error)
	SearchUsers(value string) ([]entities.UserDto, error)
	DeleteUsers(id uuid.UUID) error
	UpdateUser(id uuid.UUID, data entities.UpdateUserDto) error
	GetUserById(id uuid.UUID) (entities.UserDto, error)
}

type User struct {
	repository *repositories.User
}

func NewUserService(userRepository *repositories.User) *User {
	return &User{
		repository: userRepository,
	}
}

func (u *User) CreateUser(user entities.CreateUserDto) (entities.UserDto, error) {
	log.Printf("Creating a new User with email: %s", user.Email)
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err.Error())
		return entities.UserDto{}, err
	}
	newUser := entities.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Password: hashedPassword,
	}

	createdUser, err := u.repository.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err.Error())
		return entities.UserDto{}, err
	}

	return u.toUserDto(createdUser), nil
}

func (u *User) GetAllUsers() ([]entities.UserDto, error) {
	log.Println("Fetching all users")
	users, err := u.repository.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching all users: %v", err.Error())
		return nil, err
	}
	return u.toUsersDto(users), nil
}

func (u *User) SearchUsers(value string) ([]entities.UserDto, error) {
	if value == "" {
		return u.GetAllUsers()
	}
	users, err := u.repository.SearchUser(value)
	if err != nil {
		log.Printf("Error searching users: %v", err.Error())
		return nil, err
	}
	return u.toUsersDto(users), nil
}

func (u *User) DeleteUsers(id uuid.UUID) error {
	if err := u.repository.DeleteUser(id); err != nil {
		log.Printf("Error deleting user: %v", err.Error())
		return err
	}
	log.Printf("User with Id: %s deleted successfully", id.String())
	return nil
}

func (u *User) UpdateUser(id uuid.UUID, data entities.UpdateUserDto) error {
	if err := u.repository.UpdateUser(id, data); err != nil {
		log.Printf("Error updating user: %v", err.Error())
		return err
	}
	log.Printf("User with Id: %s updated successfully", id.String())
	return nil
}

func (u *User) GetUserById(id uuid.UUID) (entities.UserDto, error) {
	user, err := u.repository.GetUserById(id)
	if err != nil {
		log.Printf("Error getting user: %v", err.Error())
		return entities.UserDto{}, err
	}
	log.Printf("User with Id: %s retrieved successfully", id.String())
	return u.toUserDto(user), nil
}

func (u *User) toUsersDto(users []entities.User) []entities.UserDto {
	var usersDto []entities.UserDto
	for _, user := range users {
		usersDto = append(usersDto, entities.UserDto{
			Id:        user.Id,
			Name:      user.Name,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return usersDto
}

func (u *User) toUserDto(user entities.User) entities.UserDto {
	return entities.UserDto{
		Id:        user.Id,
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
