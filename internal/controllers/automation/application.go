package automation

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/grpc/gen/application"
	upl "premiesPortal/internal/app/grpc/gen/upload_file"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/pkg/errs"
)

func CreateXLSXApplicationReport(c *gin.Context) {
	clientAutomation := grpc.GetClient()

	var appids models.ApplicationReportsRequest
	if err := c.ShouldBindJSON(&appids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrValidationFailed.Error()})
		return
	}

	// Шаг 1. Создаём XLSX-файл
	createResp, err := clientAutomation.CreateXLSXApplicationReport(context.Background(), &application.CreateXLSXApplicationRequest{
		ApplicationsIds: appids.ApplicationIDS,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errs.ErrSomethingWentWrong.Error(),
		})
		return
	}

	// Шаг 2. Извлекаем путь к файлу
	filePath, ok := createResp.Resp.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid type in CreateXLSXAccountantReport response",
		})
		return
	}

	// Шаг 3. Вызываем DownloadFile
	req := &upl.DownloadFileRequest{
		Path: filePath,
	}

	downloadResp, err := clientAutomation.DownloadFile(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrSomethingWentWrong):
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		case errors.Is(err, errs.ErrTempReportNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found or file is corrupted"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download file"})
		}

		return
	}

	// Шаг 4. Приводим Resp к []byte
	fileContent, ok := downloadResp.Resp.([]byte)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid type in DownloadFile response",
		})
		return
	}

	// Шаг 5. Выдаём файл пользователю
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=downloaded_file.zip")
	c.Header("Content-Type", "application/octet-stream")

	c.Data(http.StatusOK, "application/octet-stream", fileContent)
}
