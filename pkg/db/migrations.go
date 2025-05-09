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

	//if userDBConn == nil {
	//	logger.Error.Printf("[db.Migrate] Error because users database connection is nil")
	//
	//	return errors.New("users database connection is not initialized")
	//}
	//
	//err := userDBConn.AutoMigrate(
	//	&models2.User{},
	//	&models2.Admin{},
	//)
	//if err != nil {
	//	logger.Error.Printf("[db.Migrate] Error migrating users tables: %v", err)
	//
	//	return err
	//}

	err := dbConn.AutoMigrate(
		&models2.Role{},
		&models2.User{},
		&models2.Office{},
		&models2.OfficeUser{},
		&models2.CardSales{},
		&models2.MobileBankSales{},
		&models2.CardTurnovers{},
		&models2.OperatingActive{},
		&models2.ServiceQuality{},
		&models2.History{},

		&models2.KnowledgeBase{},
		&models2.Knowledge{},
		&models2.KnowledgeDocs{},
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
