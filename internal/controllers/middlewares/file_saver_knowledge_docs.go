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
	"premiesPortal/pkg/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveFileFromResponseKnowledgeDocs(c *gin.Context) {
	var filePath string
	var title string
	var knowledgeID int

	// Попытка взять файл из multipart формы
	formFile, err := c.FormFile("file")
	if err == nil && formFile != nil {
		// multipart form
		ext := "docs"

		// создаём директорию по расширению
		subfolder := strings.TrimPrefix(ext, ".")
		dir := filepath.Join("uploads", subfolder)
		if err := os.MkdirAll(dir, 0755); err != nil {
			logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] mkdir error: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
			return
		}

		base := strings.TrimSuffix(filepath.Base(formFile.Filename), ext)
		unixTime := time.Now().Unix()
		suffix := rand.Intn(10000)
		newName := fmt.Sprintf("%s_%d_%d%s", base, unixTime, suffix, ext)

		dstPath := filepath.Join(dir, newName)

		// сохраняем файл
		if err := c.SaveUploadedFile(formFile, dstPath); err != nil {
			logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] save uploaded file error: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
			return
		}

		filePath = utils.ReplaceBackslashWithSlash(dstPath)
		title = c.PostForm("title")
		knowledgeIDStr := c.PostForm("knowledge_id")
		if knowledgeIDStr != "" {
			knowledgeID, err = strconv.Atoi(knowledgeIDStr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidID.Error()})
				return
			}
		}
	} else {
		// иначе работаем по старой схеме - JSON
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

		subfolder := strings.TrimPrefix(ext, ".")
		dir := filepath.Join("uploads", subfolder)
		if err := os.MkdirAll(dir, 0755); err != nil {
			logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] mkdir error: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
			return
		}

		src, err := os.Open(srcPath)
		if err != nil {
			logger.Error.Printf("[middlewares.SaveFileFromResponseKnowledgeDocs] open error: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrFileNotFound.Error()})
			return
		}
		defer src.Close()

		base := strings.TrimSuffix(filepath.Base(srcPath), ext)
		unixTime := time.Now().Unix()
		suffix := rand.Intn(10000)
		newName := fmt.Sprintf("%s_%d_%d%s", base, unixTime, suffix, ext)

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

		filePath = utils.ReplaceBackslashWithSlash(dstPath)
		title = response.Title
		knowledgeID = int(response.KnowledgeID)
	}

	// Сохраняем данные в контекст
	c.Set(UploadedFilePath, filePath)
	c.Set(KnowledgeDocTitle, title)
	c.Set(KnowledgeDocKnowledgeID, knowledgeID)

	c.Next()
}
