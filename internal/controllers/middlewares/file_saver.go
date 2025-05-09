package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const UploadedFilePath = "uploaded_file_path"

func SaveFileFromResponse(c *gin.Context) {
	// перехватываем ответ
	writer := &responseCatcher{ResponseWriter: c.Writer, body: &strings.Builder{}}
	c.Writer = writer

	c.Next()

	// парсим JSON-ответ
	var response struct {
		FilePath string `json:"file_path"`
	}
	if err := json.Unmarshal([]byte(writer.body.String()), &response); err != nil || response.FilePath == "" {
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
		fmt.Println("mkdir error:", err)
		return
	}

	// открываем исходный файл
	src, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("open error:", err)
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
		fmt.Println("create error:", err)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("copy error:", err)
	}

	c.Set(UploadedFilePath, dstPath)
}

// responseCatcher перехватывает тело ответа
type responseCatcher struct {
	gin.ResponseWriter
	body *strings.Builder
}

func (r *responseCatcher) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
