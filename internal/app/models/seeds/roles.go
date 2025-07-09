package seeds

import (
	"errors"
	"gorm.io/gorm"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/logger"
)

func SeedRoles(db *gorm.DB) error {
	// Определяем стандартные ралли
	roles := []models.Role{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "Worker"},
		{ID: 3, Name: "Operator"},
		{ID: 4, Name: "Head"},
		{ID: 5, Name: "Director"},
		{ID: 6, Name: "Card Seller"},
		{ID: 7, Name: "Universal"},
		{ID: 8, Name: "Credit seller"},
		{ID: 9, Name: "Chairman"},
	}

	for _, role := range roles {
		// Проверяем, существует ли роль
		var existingRole models.Role
		if err := db.First(&existingRole, "name = ?", role.Name).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Если роль не найдена, создаем её
				db.Create(&role)
			} else {
				// Обработка других ошибок
				logger.Error.Printf("[seeds.SeedRoles] Error seeding roles: %v", err)

				return err
			}
		}
	}

	return nil
}
