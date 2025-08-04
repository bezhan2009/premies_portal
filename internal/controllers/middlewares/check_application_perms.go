package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUsersApplicationPerms(c *gin.Context) {
	roleID := c.GetUint(UserRoleIDCtx)

	if roleID != 3 && roleID != 10 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Permission denied",
		})
		return
	}

	c.Next()
}
