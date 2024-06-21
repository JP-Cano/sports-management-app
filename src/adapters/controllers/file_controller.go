package controllers

import (
	"bytes"
	"github.com/JP-Cano/sports-management-app/src/application/services/file"
	"github.com/JP-Cano/sports-management-app/src/core/exceptions"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FileController struct {
	excelService file.ExcelService
}

func NewFileController(excelService file.ExcelService) *FileController {
	return &FileController{excelService: excelService}
}

func (f *FileController) UploadExcel(c *gin.Context) {
	var message string
	formFile, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Request FormFile error: %v", err.Error())
		message = exceptions.InternalServerError.Error()
		utils.ErrorResponse(c, http.StatusBadRequest, &message)
		return
	}

	buffer := bytes.NewBuffer(nil)
	if _, err = buffer.ReadFrom(formFile); err != nil {
		log.Printf("Error reading form file: %v", err.Error())
		message = exceptions.InternalServerError.Error()
		utils.ErrorResponse(c, http.StatusBadRequest, &message)
		return
	}

	err = f.excelService.ProcessExcel(buffer.Bytes())
	if err != nil {
		log.Printf("Error processing excel file: %v", err.Error())
		message = exceptions.InternalServerError.Error()
		utils.ErrorResponse(c, http.StatusInternalServerError, &message)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, map[string]string{
		"message": "File processed successfully",
	})
}
