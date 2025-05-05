package middlewares

import "github.com/gin-gonic/gin"

func CheckUserKnowledgePerms(c *gin.Context) {
	roleID := c.GetUint(UserRoleIDCtx)

	if roleID != 3 && roleID != 1 {
		c.AbortWithStatusJSON(403, gin.H{"message": "Permission denied"})
		return
	}

	c.Next()
}
