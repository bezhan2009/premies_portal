package automation

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	upl "premiesPortal/internal/app/grpc/gen/upload_file"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

const (
	UploadedAutomationFilePath = "uploaded_automation_file_path"
)

func UploadAutomationFile(c *gin.Context) {
	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		controllers.HandleError(c, fmt.Errorf("failed to get file: %v", err))
		return
	}

	// Открываем файл
	openedFile, err := file.Open()
	if err != nil {
		logger.Error.Printf("failed to open file: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, errs.ErrSomethingWentWrong)
		return
	}
	defer openedFile.Close()

	// Читаем содержимое файла
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		logger.Error.Printf("failed to read file: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, errs.ErrSomethingWentWrong)
		return
	}

	// Получаем gRPC клиент
	client := grpc.GetClient()

	// Создаем запрос с содержимым файла
	resp, err := client.UploadFile(context.Background(), &upl.UploadFileRequest{
		Filename:    file.Filename,
		FileContent: fileBytes,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errs.ErrSomethingWentWrong)
		return
	}

	c.Set(UploadedAutomationFilePath, resp.Resp)

	c.Next()
}
