package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"premiesPortal/internal/repository"
	"premiesPortal/pkg/errs"
	"strconv"
)

func DeleteFileKnowledgeDocs(c *gin.Context) {
	knowledgeDocStrId := c.Param("id")
	knowledgeDocId, err := strconv.Atoi(knowledgeDocStrId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errs.ErrInvalidID.Error(),
		})
		return
	}

	knowledgeDoc, err := repository.GetKnowledgeDocByID(uint(knowledgeDocId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errs.ErrKnowledgeDocNotFound.Error(),
		})
		return
	}

	err = os.Remove(knowledgeDoc.FilePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errs.ErrFileNotFound.Error(),
		})
		return
	}

	c.Next()
}
