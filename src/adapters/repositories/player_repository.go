package repositories

import (
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"gorm.io/gorm"
)

type Player struct {
	DB *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *Player {
	return &Player{DB: db}
}

func (p *Player) SavePlayer(player entities.Player) (entities.Player, error) {
	if err := p.DB.Create(&player).Error; err != nil {
		return entities.Player{}, err
	}
	return player, nil
}

func (p *Player) BeginTransaction() (*gorm.DB, error) {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}
