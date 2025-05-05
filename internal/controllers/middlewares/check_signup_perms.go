package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckSignupPerms(c *gin.Context) {
	roleID := c.GetUint(UserRoleIDCtx)

	if roleID != 1 && roleID != 3 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Permission denied"})
		return
	}

	c.Next()
}
