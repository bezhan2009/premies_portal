package middlewares

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	UploadedFilePath        = "uploaded_file_path"
	KnowledgeDocTitle       = "knowledge_doc_title"
	KnowledgeDocKnowledgeID = "knowledge_doc_knowledge_id"
)

func SaveFileFromResponseKnowledgeDocs(c *gin.Context) {
	// парсим JSON-ответ
	var response models.KnowledgeDocs
	if err := c.ShouldBindJSON(&response); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrFilePathIsRequired.Error()})
		return
	}

	srcPath := response.FilePath
	ext := strings.ToLower(filepath.Ext(srcPath))
	if ext == "" {
		ext = ".bin"
	}

	// создаём директорию по расширению
	subfolder := strings.TrimPrefix(ext, ".")
	dir := filepath.Join("uploads", subfolder)
	if err := os.MkdirAll(dir, 0755); err != nil {
		logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] mkdir error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}

	// открываем исходный файл
	src, err := os.Open(srcPath)
	if err != nil {
		logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] open error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrFileNotFound.Error()})
		return
	}
	defer src.Close()

	// создаём уникальное имя
	base := strings.TrimSuffix(filepath.Base(srcPath), ext)
	timestamp := time.Now().Format("20060102_150405")
	suffix := rand.Intn(10000)
	newName := fmt.Sprintf("%s_%s_%d%s", base, timestamp, suffix, ext)

	dstPath := filepath.Join(dir, newName)
	dst, err := os.Create(dstPath)
	if err != nil {
		logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] create error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] copy error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
		return
	}

	c.Set(UploadedFilePath, dstPath)
	c.Set(KnowledgeDocTitle, response.Title)
	c.Set(KnowledgeDocKnowledgeID, response.KnowledgeID)

	c.Next()
}
