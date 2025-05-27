package automation

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"premiesPortal/internal/app/grpc/gen/go/tus"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/pkg/errs"
)

func UploadTusData(c *gin.Context) {
	var filePath models.FilePath
	if err := c.ShouldBindJSON(&filePath); err != nil {
		controllers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	if filePath.FilePath == "" {
		controllers.HandleError(c, errs.ErrInvalidFilePath)
		return
	}

	client := grpc.GetClient()

	resp, err := client.UploadTusData(context.Background(), &tus.TusUploadRequest{FilePath: filePath.FilePath})
	if err != nil {
		controllers.HandleError(c, err)
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"message": resp.Resp,
	})
}

func CleanTusTable(c *gin.Context) {
	client := grpc.GetClient()

	resp, err := client.CleanTusTable(context.Background(), &emptypb.Empty{})
	if err != nil {
		controllers.HandleError(c, err)
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"message": resp.Resp,
	})
}
