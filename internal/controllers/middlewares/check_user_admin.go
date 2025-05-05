package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/pkg/errs"
)

func CheckUserNotWorker(c *gin.Context) {
	userRole := c.GetUint(UserRoleIDCtx)
	if userRole == 2 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": errs.ErrPermissionDenied})
		return
	}

	c.Next()
}
