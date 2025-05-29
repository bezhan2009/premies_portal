package server

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"premiesPortal/internal/clients/automation_premies/grpc"
	"premiesPortal/internal/routes"
	"premiesPortal/internal/security"
	"premiesPortal/pkg/db"
	"premiesPortal/pkg/logger"
	"premiesPortal/pkg/utils"
)

var mainServer *Server

func ServiceStart() (err error) {
	gin.SetMode(security.AppSettings.AppParams.GinMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     security.AppSettings.Cors.AllowOrigins,
		AllowMethods:     security.AppSettings.Cors.AllowMethods,
		AllowHeaders:     security.AppSettings.Cors.AllowHeaders,
		ExposeHeaders:    security.AppSettings.Cors.ExposeHeaders,
		AllowCredentials: security.AppSettings.Cors.AllowCredentials,
	}))

	err = grpc.New(
		utils.Context,
		security.AppSettings.Clients.Premies.ClientAddress,
		security.AppSettings.Clients.Premies.Timeout,
		security.AppSettings.Clients.Premies.RetriesCount,
	)
	if err != nil {
		return err
	}

	mainServer = new(Server)
	go func() {
		if err := mainServer.Run(security.AppSettings.AppParams.PortRun, routes.InitRoutes(router)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting HTTP Service: %s", err)
		}
	}()

	return nil
}

func ServiceShutdown() {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("\n%s\n", yellow("Start of service termination"))

	// Закрытие соединения с БД
	err := db.CloseDBConn()
	if err != nil {
		strErr := fmt.Sprintf("Error closing database connection: %s", err.Error())
		fmt.Println(red(strErr))
		logger.Error.Println(strErr)
	}

	// Закрытие соединения с GRPC сервисом python для автоматизации
	err = grpc.GrpcConnClose()
	if err != nil {
		strErr := fmt.Sprintf("Error closing grpc service python automation connection: %s", err.Error())
		fmt.Println(red(strErr))
		logger.Error.Println(strErr)
	}

	// Корректное завершение HTTP-сервера
	if err = mainServer.Shutdown(context.Background()); err != nil {
		strErr := fmt.Sprintf("Error shutting down server: %s", err.Error())
		fmt.Println(red(strErr))
		logger.Error.Println(strErr)
	} else {
		strSuccess := "HTTP-service termination successfully"
		fmt.Println(green(strSuccess))
		logger.Info.Println(strSuccess)
	}
}
