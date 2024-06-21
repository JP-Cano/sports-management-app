package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	LastName  string    `gorm:"size:100;not null" json:"lastName"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type UserDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserDto struct {
	Name     string `json:"name" binding:"required,max=100"`
	LastName string `json:"lastName" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,max=100,email"`
	Password string `json:"password" binding:"required,min=6,max=10"`
}

type UpdateUserDto struct {
	Name     string `json:"name" binding:"required,max=100"`
	LastName string `json:"lastName" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,max=100,email"`
}

// TODO: Add document number, date of birth, address, and other confidential fields
