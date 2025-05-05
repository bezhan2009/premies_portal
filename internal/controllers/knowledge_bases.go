package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetAllKnowledgeBases(c *gin.Context) {
	knowledgeBases, err := service.GetAllKnowledgeBases()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, knowledgeBases)
}

func CreateKnowledgeBase(c *gin.Context) {
	var kb models.KnowledgeBase
	if err := c.ShouldBindJSON(&kb); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateKnowledgeBase(kb)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge Base created successfully"})
}

func UpdateKnowledgeBase(c *gin.Context) {
	knowledgeBaseIDStr := c.Param("id")
	knowledgeBaseID, err := strconv.Atoi(knowledgeBaseIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var kb models.KnowledgeBase
	if err := c.ShouldBindJSON(&kb); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	kb.ID = uint(knowledgeBaseID)

	err = service.UpdateKnowledgeBase(kb)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge Base updated successfully"})
}
func DeleteKnowledgeBase(c *gin.Context) {
	knowledgeBaseIDStr := c.Param("id")
	knowledgeBaseID, err := strconv.Atoi(knowledgeBaseIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteKnowledgeBase(uint(knowledgeBaseID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge Base deleted successfully"})
}
