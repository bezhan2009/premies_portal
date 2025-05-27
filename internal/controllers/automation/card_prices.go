package automation

import (
	"context"
	"github.com/gin-gonic/gin"
	"premiesPortal/internal/app/grpc/gen/go/card_prices"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/pkg/errs"
)

func UploadCardPrices(c *gin.Context) {
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

	resp, err := client.UploadCardPricesData(context.Background(), &card_prices.CardPricesUploadRequest{FilePath: filePath.FilePath})
	if err != nil {
		controllers.HandleError(c, err)
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"message": resp.Resp,
	})
}
