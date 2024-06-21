package services

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/repositories"
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/utils"
	"gorm.io/gorm"
)

type PlayerService interface {
	CreatePlayerBatch(row []string) error
	BeginTransaction() (*gorm.DB, error)
}

type Player struct {
	repository *repositories.Player
}

func NewPlayerService(repository *repositories.Player) *Player {
	return &Player{repository}
}

func (p *Player) CreatePlayerBatch(row []string) error {
	newPlayer := entities.Player{
		Name:        row[0],
		LastName:    row[1],
		Age:         utils.ParseToInt(row[2]),
		Address:     row[3],
		Email:       row[4],
		Phone:       row[5],
		DateOfBirth: utils.ParseDate(row[6]),
		TeamId:      utils.ParseUUID(row[7]),
	}

	_, err := p.repository.SavePlayer(newPlayer)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) BeginTransaction() (*gorm.DB, error) {
	return p.repository.BeginTransaction()
}
