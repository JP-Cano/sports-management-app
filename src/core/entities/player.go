package entities

import (
	"github.com/google/uuid"
	"time"
)

type Player struct {
	Id          uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	TeamId      *uuid.UUID `gorm:"type:uuid" json:"teamId" binding:"uuid"`
	Team        Team
	Name        string    `gorm:"size:100;not null" json:"name" binding:"required,max=100"`
	Age         int       `gorm:"not null" json:"age" binding:"required"`
	Address     string    `gorm:"size:100;not null" json:"address" binding:"required,max=100"`
	Email       string    `gorm:"size:100;not null" json:"email" binding:"required,max=100, email"`
	Phone       string    `gorm:"size:20;not null" json:"phone" binding:"required,max=15"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required,date"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type CreatePlayerDto struct {
	TeamId      *uuid.UUID `json:"teamId" binding:"uuid"`
	Name        string     `json:"name" binding:"required,max=100"`
	Age         int        `binding:"required"`
	Address     string     `json:"address" binding:"required,max=100"`
	Email       string     `json:"email" binding:"required,max=100, email"`
	Phone       string     `json:"phone" binding:"required,max=15"`
	DateOfBirth time.Time  `binding:"required,date"`
}

type UpdatePlayerDto struct {
	Name        string    `json:"name" binding:"required,max=100"`
	Age         int       `binding:"required"`
	Address     string    `json:"address" binding:"required,max=100"`
	Email       string    `json:"email" binding:"required,max=100, email"`
	Phone       string    `json:"phone" binding:"required,max=15"`
	DateOfBirth time.Time `binding:"required,date"`
}
