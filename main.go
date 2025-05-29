package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"premiesPortal/configs"
	"premiesPortal/internal/security"
	"premiesPortal/internal/server"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
	"syscall"
)

// @title Active Bonk portal API
// @version 1.0.0

// @description API Server for portal Application
// @host localhost:7575
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	green := color.New(color.FgGreen).SprintFunc()

	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("example.env")
		if err != nil {
			panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
		}
	}

	security.AppSettings, err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}
	security.SetConnDB(security.AppSettings)

	err = logger.Init()
	if err != nil {
		panic(err)
	}

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	err = server.ServiceStart()
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	server.ServiceShutdown()

	logger.Info.Println("End of program completion")
	fmt.Println(green("End of program completion"))
}
