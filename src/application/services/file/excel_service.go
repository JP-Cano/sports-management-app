package file

import (
	"bytes"
	"github.com/JP-Cano/sports-management-app/src/application/services"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/worker"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"log"
)

type ExcelService interface {
	ProcessExcel(data []byte) error
}

type Excel struct {
	playerService services.PlayerService
}

func NewPlayerService(playerService services.PlayerService) *Excel {
	return &Excel{playerService: playerService}
}

func (e *Excel) ProcessExcel(data []byte) error {
	log.Println("Start to process excel file")
	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		log.Printf("Excel Reader Error: %v\n", err)
		return err
	}

	rows, err := f.GetRows("Players")
	if err != nil {
		log.Printf("Get rows Error: %v\n", err)
		return err
	}

	rowChan := make(chan []string)
	errorChan := make(chan error)
	workerPool := worker.New(10, e.playerService.CreatePlayerBatch, errorChan)
	go workerPool.Start(rowChan)

	tx, err := e.playerService.BeginTransaction()
	if err != nil {
		log.Printf("Error begin transaction: %v", err)
		return err
	}

	e.rollBackTrx(tx)

	for _, row := range rows {
		rowChan <- row
	}
	close(rowChan)
	workerPool.Wait()
	close(errorChan)

	err = e.validateErrorChannel(err, errorChan, tx)
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (e *Excel) validateErrorChannel(err error, errorChan chan error, tx *gorm.DB) error {
	for err = range errorChan {
		if err != nil {
			e.rollBackTrx(tx)
			log.Printf("Error processing rows: %v", err)
			return err
		}
	}
	return nil
}

func (e *Excel) rollBackTrx(tx *gorm.DB) {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
}
