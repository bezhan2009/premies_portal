package automation

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/controllers"
)

func CreateZIPReports(c *gin.Context) {
	clientAutomation := grpc.GetClient()

	resp, err := clientAutomation.CreateZIPReports(context.Background(), &emptypb.Empty{})
	if err != nil {
		controllers.HandleError(c, err)
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"message": resp.Resp,
	})
}
