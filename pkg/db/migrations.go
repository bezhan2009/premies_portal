package db

import (
	"errors"
	models2 "premiesPortal/internal/app/models"
	"premiesPortal/internal/app/models/seeds"
)

func Migrate() error {
	if dbConn == nil {
		//logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models2.Role{},
		&models2.User{},
		&models2.Worker{},

		&models2.Office{},
		&models2.OfficeUser{},

		&models2.CardSales{},
		&models2.CardDetails{},
		&models2.MobileBankSales{},
		&models2.MobileBankDetails{},
		&models2.CardTurnovers{},
		&models2.ServiceQuality{},
		&models2.Overdraft{},

		&models2.KnowledgeBase{},
		&models2.Knowledge{},
		&models2.KnowledgeDocs{},

		&models2.Test{},
		&models2.Question{},
		&models2.Option{},
		&models2.Answer{},
	)
	if err != nil {
		//logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	err = seeds.SeedRoles(dbConn)
	if err != nil {
		return err
	}

	return nil
}
