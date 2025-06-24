package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"premiesPortal/internal/app/models"
	"sort"
	"strings"
)

func parsePreloadQueryParams(c *gin.Context) models.WorkerPreloadOptions {
	return models.WorkerPreloadOptions{
		LoadCardTurnovers:  parseBoolQuery(c, "loadCardTurnovers"),
		LoadCardSales:      parseBoolQuery(c, "loadCardSales"),
		LoadServiceQuality: parseBoolQuery(c, "loadServiceQuality"),
		LoadMobileBank:     parseBoolQuery(c, "loadMobileBank"),
		LoadCardDetails:    parseBoolQuery(c, "loadCardDetails"),
		LoadUser:           parseBoolQuery(c, "loadUser"),
	}
}

func parseBoolQuery(c *gin.Context, key string) bool {
	val := strings.ToLower(c.Query(key))
	return val == "true" || val == "1" || val == "yes"
}

func GenerateRedisKeyFromQuery(c *gin.Context, prefix string) string {
	var parts []string

	// Добавим ID работника
	workerID := c.Param("id")
	if workerID != "" {
		parts = append(parts, fmt.Sprintf("worker:%s", workerID))
	}

	// Получаем все query-параметры и сортируем по ключу (для стабильности)
	queryMap := c.Request.URL.Query()
	keys := make([]string, 0, len(queryMap))
	for k := range queryMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Добавляем каждый query как key=value
	for _, key := range keys {
		vals := queryMap[key]
		if len(vals) > 0 {
			// Если параметр передан несколько раз, объединим через запятую
			joined := strings.Join(vals, ",")
			parts = append(parts, fmt.Sprintf("%s=%s", key, joined))
		}
	}

	// Собираем финальный ключ
	key := prefix + ":" + strings.Join(parts, ":")
	return key
}

func GenerateAllWorkersRedisKey(c *gin.Context, prefix string) string {
	var parts []string

	// Сначала — базовые параметры
	afterID := c.DefaultQuery("after", "0")
	parts = append(parts, fmt.Sprintf("after=%s", afterID))

	// Затем валидные query параметры
	queryMap := c.Request.URL.Query()
	keys := make([]string, 0, len(queryMap))
	for k := range queryMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Добавляем все query параметры (в отсортированном порядке)
	for _, key := range keys {
		// Пропускаем "after" — он уже добавлен выше
		if key == "after" {
			continue
		}

		vals := queryMap[key]
		if len(vals) > 0 {
			joined := strings.Join(vals, ",")
			parts = append(parts, fmt.Sprintf("%s=%s", key, joined))
		}
	}

	// Формируем ключ
	key := prefix + ":" + strings.Join(parts, ":")
	return key
}
