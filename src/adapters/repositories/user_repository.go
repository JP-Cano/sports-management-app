package repositories

import (
	"errors"
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"github.com/JP-Cano/sports-management-app/src/core/exceptions"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var userPublicProperties = []string{"id", "name", "last_name", "email", "created_at", "updated_at"}

type User struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(user entities.User) (entities.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u *User) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	if err := u.DB.Select(userPublicProperties).Order("created_at desc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) SearchUser(value string) ([]entities.User, error) {
	var users []entities.User
	if err := u.DB.Select(userPublicProperties).Where("name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", "%"+value+"%", "%"+value+"%", "%"+value+"%").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetUserById(id uuid.UUID) (entities.User, error) {
	var user entities.User
	if err := u.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, exceptions.NotFound
		}
		return entities.User{}, err
	}
	return user, nil
}

func (u *User) DeleteUser(id uuid.UUID) error {
	if err := u.DB.Delete(&entities.User{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.NotFound
		}
		return err
	}
	return nil
}

func (u *User) UpdateUser(id uuid.UUID, data entities.UpdateUserDto) error {
	userToUpdate, err := u.GetUserById(id)
	if err != nil {
		return err
	}

	if data.Name != "" {
		userToUpdate.Name = data.Name
	}

	if data.LastName != "" {
		userToUpdate.LastName = data.LastName
	}

	if data.Email != "" {
		userToUpdate.Email = data.Email
	}

	if err = u.DB.Save(&userToUpdate).Error; err != nil {
		return err
	}
	return nil
}
