package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"strconv"
)

func GetKnowledgeDocsByKnowledgeID(c *gin.Context) {
	knowledgeStrId := c.Param("id")
	knowledgeId, err := strconv.Atoi(knowledgeStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	knowledge, err := service.GetKnowledgeDocsByKnowledgeID(uint(knowledgeId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, knowledge)
}

func GetKnowledgeDocsByID(c *gin.Context) {
	knowledgeDocStrId := c.Param("id")
	knowledgeDocId, err := strconv.Atoi(knowledgeDocStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	knowledgeDoc, err := service.GetKnowledgeDocByID(uint(knowledgeDocId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, knowledgeDoc)
}

func CreateKnowledgeDoc(c *gin.Context) {
	var knowledgeDoc models.KnowledgeDocs
	if err := c.ShouldBindJSON(&knowledgeDoc); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateKnowledgeDocs(knowledgeDoc)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge doc created successfully"})
}

func UpdateKnowledgeDoc(c *gin.Context) {
	knowledgeDocStrId := c.Param("id")
	knowledgeDocId, err := strconv.Atoi(knowledgeDocStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var knowledgeDoc models.KnowledgeDocs
	if err := c.ShouldBindJSON(&knowledgeDoc); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	knowledgeDoc.ID = uint(knowledgeDocId)

	err = service.UpdateKnowledgeDocs(knowledgeDoc)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Knowledge Doc Success"})
}

func DeleteKnowledgeDoc(c *gin.Context) {
	knowledgeDocStrId := c.Param("id")
	knowledgeDocId, err := strconv.Atoi(knowledgeDocStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteKnowledgeDocs(uint(knowledgeDocId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Knowledge Doc Successfully"})
}
