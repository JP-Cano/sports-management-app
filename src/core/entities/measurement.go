package entities

import (
	"github.com/google/uuid"
	"time"
)

type Measurement struct {
	Id              uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	PlayerId        uuid.UUID `gorm:"type:uuid;not null" json:"playerId"`
	Player          Player
	MeasurementDate time.Time `gorm:"autoCreateTime" json:"measurementDate"`
	Height          float64   `gorm:"type:decimal;not null" json:"height"`
	Weight          float64   `gorm:"type:decimal;not null" json:"weight"`
	Wingspan        float64   `gorm:"type:decimal;not null" json:"wingspan"`
	JumpHeight      float64   `gorm:"type:decimal;not null" json:"jumpHeight"`
	BMI             float64   `gorm:"type:decimal;not null" json:"bmi"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
