package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func KnowledgeDocsExists(c *gin.Context) {
	knowledgeDocStrId := c.Param("id")
	knowledgeDocId, err := strconv.Atoi(knowledgeDocStrId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errs.ErrInvalidID.Error(),
		})
		return
	}

	_, err = service.GetKnowledgeDocByID(uint(knowledgeDocId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": errs.ErrKnowledgeDocNotFound.Error(),
		})
		return
	}

	c.Next()
}
