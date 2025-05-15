package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetKnowledgeByBaseID(c *gin.Context) {
	baseStrId := c.Param("id")
	baseId, err := strconv.Atoi(baseStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	knowledge, err := service.GetKnowledgeByBaseID(uint(baseId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, knowledge)
}

func GetKnowledgeByID(c *gin.Context) {
	knowledgeStrId := c.Param("id")
	knowledgeId, err := strconv.Atoi(knowledgeStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	knowledge, err := service.GetKnowledgeByID(uint(knowledgeId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, knowledge)
}

func CreateKnowledge(c *gin.Context) {
	var knowledge models.Knowledge
	if err := c.ShouldBindJSON(&knowledge); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	knowledge.ID = 0

	err := service.CreateKnowledgeTable(knowledge)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge created successfully"})
}

func UpdateKnowledge(c *gin.Context) {
	knowledgeStrId := c.Param("id")
	knowledgeId, err := strconv.Atoi(knowledgeStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var knowledge models.Knowledge
	if err := c.ShouldBindJSON(&knowledge); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	knowledge.ID = uint(knowledgeId)

	err = service.UpdateKnowledge(knowledge)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge updated successfully"})
}

func DeleteKnowledge(c *gin.Context) {
	knowledgeStrId := c.Param("id")
	knowledgeId, err := strconv.Atoi(knowledgeStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteKnowledge(uint(knowledgeId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge deleted successfully"})
}
