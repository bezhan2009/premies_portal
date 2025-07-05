package automation

import (
	"context"
	"premiesPortal/internal/app/grpc/gen/tus"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/pkg/errs"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func UploadTusData(c *gin.Context) {
	filePath := c.GetString(UploadedAutomationFilePath)

	if filePath == "" {
		controllers.HandleError(c, errs.ErrInvalidFilePath)
		return
	}

	client := grpc.GetClient()

	resp, err := client.UploadTusData(context.Background(), &tus.TusUploadRequest{FilePath: filePath})
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
