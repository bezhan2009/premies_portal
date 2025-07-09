package automation

import (
	"context"
	"premiesPortal/internal/app/grpc/gen/tus"
	"premiesPortal/internal/app/service/validators"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/pkg/errs"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func UploadTusData(c *gin.Context) {
	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		controllers.HandleError(c, errs.ErrInvalidMonth)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		controllers.HandleError(c, errs.ErrInvalidYear)
		return
	}

	filePath := c.GetString(UploadedAutomationFilePath)

	if filePath == "" {
		controllers.HandleError(c, errs.ErrInvalidFilePath)
		return
	}

	client := grpc.GetClient()

	resp, err := client.UploadTusData(context.Background(), &tus.TusUploadRequest{
		Month:    int32(month),
		Year:     int32(year),
		FilePath: filePath,
	})
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
