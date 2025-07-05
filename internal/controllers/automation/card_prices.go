package automation

import (
	"context"
	"github.com/gin-gonic/gin"
	"premiesPortal/internal/app/grpc/gen/card_prices"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
	"premiesPortal/pkg/errs"
)

func UploadCardPrices(c *gin.Context) {
	filePath := c.GetString(UploadedAutomationFilePath)

	if filePath == "" {
		controllers.HandleError(c, errs.ErrInvalidFilePath)
		return
	}

	client := grpc.GetClient()

	resp, err := client.UploadCardPricesData(context.Background(), &card_prices.CardPricesUploadRequest{FilePath: filePath})
	if err != nil {
		controllers.HandleError(c, err)
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"message": resp.Resp,
	})
}
