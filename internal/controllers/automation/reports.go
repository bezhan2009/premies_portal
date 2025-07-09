package automation

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/grpc/gen/reports"
	upl "premiesPortal/internal/app/grpc/gen/upload_file"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
	"strconv"
)

func CreateZIPReports(c *gin.Context) {
	clientAutomation := grpc.GetClient()

	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		controllers.HandleError(c, errs.ErrInvalidYear)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		controllers.HandleError(c, errs.ErrInvalidYear)
		return
	}

	// Шаг 1. Создаём ZIP-файл
	createResp, err := clientAutomation.CreateZIPReports(context.Background(), &reports.CreateZIPReportsRequest{
		Month: int32(month),
		Year:  int32(year),
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
			"error": "Invalid type in CreateZIPReports response",
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
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
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

func CreateExcelReport(c *gin.Context) {
	workerUserIDStr := c.Param("id")
	workerUserID, err := strconv.Atoi(workerUserIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.ErrInvalidID.Error(),
		})
		return
	}

	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		controllers.HandleError(c, errs.ErrInvalidYear)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		controllers.HandleError(c, errs.ErrInvalidYear)
		return
	}

	_, err = repository.GetUserByID(uint(workerUserID))
	if err != nil {
		if !errors.Is(err, errs.ErrUserNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs.ErrRecordNotFound.Error(),
			})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": errs.ErrSomethingWentWrong.Error(),
		})
		return
	}

	clientAutomation := grpc.GetClient()

	createResp, err := clientAutomation.CreateExcelReport(context.Background(), &reports.CreateExcelReportRequest{
		OwnerId: int64(workerUserID),
		Month:   int32(month),
		Year:    int32(year),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errs.ErrSomethingWentWrong.Error(),
		})
		return
	}

	filePath, ok := createResp.Resp.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid type in CreateExcelReport response",
		})
		return
	}

	req := &upl.DownloadFileRequest{
		Path: filePath,
	}

	downloadResp, err := clientAutomation.DownloadFile(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrSomethingWentWrong):
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		case errors.Is(err, errs.ErrTempReportNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download file"})
		}
		return
	}

	fileContent, ok := downloadResp.Resp.([]byte)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid type in DownloadFile response",
		})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=downloaded_file.zip")
	c.Header("Content-Type", "application/octet-stream")

	c.Data(http.StatusOK, "application/octet-stream", fileContent)
}
