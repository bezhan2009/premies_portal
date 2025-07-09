package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserOperatorOrChairman(c *gin.Context) {
	roleID := c.GetUint(UserRoleIDCtx)

	if roleID != 3 && roleID != 9 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Permission denied",
		})
		return
	}

	c.Next()
}
