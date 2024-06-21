package entities

import (
	"github.com/google/uuid"
	"time"
)

type Team struct {
	Id        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId    *uuid.UUID `gorm:"type:uuid" json:"userId" binding:"uuid"`
	User      User
	Name      string    `gorm:"size:100;not null" json:"name"`
	Sport     string    `gorm:"size:50;not null" json:"sport"`
	Category  string    `gorm:"size:50;not null" json:"category"`
	Sex       string    `gorm:"size:10;not null" json:"sex"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
