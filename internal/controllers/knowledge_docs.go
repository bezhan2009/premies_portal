package controllers

import (
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
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
	knowledgeDoc.Title = c.GetString(middlewares.KnowledgeDocTitle)
	knowledgeDoc.KnowledgeID = uint(c.GetInt(middlewares.KnowledgeDocKnowledgeID))
	knowledgeDoc.FilePath = c.GetString(middlewares.UploadedFilePath)

	err := service.CreateKnowledgeDocs(knowledgeDoc)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Knowledge doc created Successfully"})
}

func UpdateKnowledgeDoc(c *gin.Context) {
	knowledgeDocStrId := c.Param("id")
	knowledgeDocId, err := strconv.Atoi(knowledgeDocStrId)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var knowledgeDoc models.KnowledgeDocs

	knowledgeDoc.Title = c.GetString(middlewares.KnowledgeDocTitle)
	knowledgeDoc.KnowledgeID = c.GetUint(middlewares.KnowledgeDocKnowledgeID)
	knowledgeDoc.ID = uint(knowledgeDocId)

	newFilePath := c.GetString(middlewares.UploadedFilePath)
	if newFilePath != "" {
		knowledgeDoc.FilePath = newFilePath
	}

	err = service.UpdateKnowledgeDocs(knowledgeDoc)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Knowledge Doc Successfully"})
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
