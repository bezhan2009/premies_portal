package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	_ "premiesPortal/docs"
	"premiesPortal/internal/controllers"
	"premiesPortal/internal/controllers/automation"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/internal/security"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	googleProvider := security.AppSettings.ProvidersParams.GoogleProvider

	// Настройка провайдера
	if googleProvider.ClientSecret == "" || googleProvider.ClientID == "" {
		goth.UseProviders(
			google.New(
				os.Getenv("GOOGLE_CLIENT_ID"),
				os.Getenv("GOOGLE_CLIENT_SECRET"),
				os.Getenv("GOOGLE_REDIRECT"),
				"email", "profile",
			),
		)
	} else {
		goth.UseProviders(
			google.New(
				googleProvider.ClientID,
				googleProvider.ClientSecret,
				googleProvider.Redirect,
				"email", "profile",
			),
		)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pingRoute := r.Group("/ping")
	{
		pingRoute.GET("/", controllers.Ping)
	}

	// user & users Маршруты для сущности пользователей
	r.GET("/user", middlewares.CheckUserAuthentication,
		controllers.GetMyDataUser)

	usersRoute := r.Group("/users", middlewares.CheckUserAuthentication, middlewares.CheckUserNotWorker)
	{
		usersRoute.GET("", controllers.GetAllUsers)
		usersRoute.GET("/:id", controllers.GetUserByID)
	}

	// worker & workers маршруты для сущности
	r.GET("/worker", middlewares.CheckUserAuthentication,
		controllers.GetMyDataWorker)

	workerRoute := r.Group("/workers", middlewares.CheckUserAuthentication, middlewares.CheckUserNotWorker)
	{
		workerRoute.GET("", controllers.GetAllWorkers)
		workerRoute.GET("/:id", controllers.GetWorkerByID)
	}

	// auth Маршруты для авторизаций
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", middlewares.CheckUserAuthentication, middlewares.CheckSignupPerms, controllers.SignUp)
		auth.POST("/sign-up/temp", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
		auth.POST("/refresh", controllers.RefreshToken)

		auth.GET("/google", controllers.GoogleLogin)
		auth.GET("/google/callback", controllers.GoogleCallback)
	}

	// office Маршруты для офисов
	office := r.Group("/office", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)
	{
		office.GET("", controllers.GetAllOffices)
		office.GET("/:id", controllers.GetOfficeByID)

		office.POST("", controllers.CreateOffice)
		office.PATCH("/:id", controllers.UpdateOffice)
		office.DELETE("/:id", controllers.DeleteOffice)
	}

	// officeWorkers Маршруты для рабочих офисов
	officeWorkers := office.Group("/workers")
	{
		officeWorkers.GET("/:id", controllers.GetAllOfficeWorkers)
		officeWorkers.GET("single/:id", controllers.GetOfficeWorkerByID)

		officeWorkers.POST("", controllers.AddWorkerToOffice)
		officeWorkers.DELETE("/:id", controllers.DeleteUserFromOffice)
	}

	// knowledge маршруты для базы знаний
	knowledge := r.Group("/knowledge", middlewares.CheckUserAuthentication)
	{
		knowledge.POST("", middlewares.CheckUserKnowledgePerms, controllers.CreateKnowledge)

		knowledge.GET("/:id", controllers.GetKnowledgeByBaseID)
		knowledge.GET("single/:id", controllers.GetKnowledgeByID)
		knowledge.PATCH("/:id", middlewares.CheckUserKnowledgePerms, controllers.UpdateKnowledge)
		knowledge.DELETE("/:id", middlewares.CheckUserKnowledgePerms, controllers.DeleteKnowledge)
	}

	knowledgeDocs := knowledge.Group("/docs")
	{
		knowledgeDocs.POST("", middlewares.CheckUserKnowledgePerms, middlewares.SaveFileFromResponseKnowledgeDocs, controllers.CreateKnowledgeDoc)

		knowledgeDocs.GET("/:id", controllers.GetKnowledgeDocsByKnowledgeID)

		knowledgeDocs.PATCH("/:id",
			middlewares.CheckUserKnowledgePerms,
			middlewares.KnowledgeDocsExists,
			middlewares.SaveFileFromResponseKnowledgeDocs,
			controllers.UpdateKnowledgeDoc)

		knowledgeDocs.DELETE("/:id",
			middlewares.CheckUserKnowledgePerms,
			middlewares.KnowledgeDocsExists,
			middlewares.DeleteFileKnowledgeDocs,
			controllers.DeleteKnowledgeDoc)
	}

	knowledgeBases := knowledge.Group("/bases")
	{
		knowledgeBases.GET("", controllers.GetAllKnowledgeBases)
		knowledgeBases.GET("/:id", controllers.GetKnowledgeBaseByID)
		knowledgeBases.POST("", middlewares.CheckUserKnowledgePerms, controllers.CreateKnowledgeBase)
		knowledgeBases.PATCH("/:id", middlewares.CheckUserKnowledgePerms, controllers.UpdateKnowledgeBase)
		knowledgeBases.DELETE("/:id", middlewares.CheckUserKnowledgePerms, controllers.DeleteKnowledgeBase)
	}

	// cards Маршруты для карт
	cards := r.Group("/cards", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)

	cardSales := cards.Group("/sales")
	{
		cardSales.POST("", controllers.AddCardSales)
		cardSales.PATCH("/:id", controllers.UpdateCardSales)
		cardSales.DELETE("/:id", controllers.DeleteCardSales)
	}

	cardTurnovers := cards.Group("turnovers")
	{
		cardTurnovers.POST("", controllers.AddCardTurnovers)
		cardTurnovers.PATCH("/:id", controllers.UpdateCardTurnovers)
		cardTurnovers.DELETE("/:id", controllers.DeleteCardTurnovers)
	}

	// serviceQuality маршруты для качества обсулживания
	serviceQuality := r.Group("service-quality", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)
	{
		serviceQuality.POST("", controllers.AddServiceQuality)
		serviceQuality.PATCH("/:id", controllers.UpdateServiceQuality)
		serviceQuality.DELETE("/:id", controllers.DeleteServiceQuality)
	}

	automationRoutes := r.Group("automation", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)

	cardsAutomation := automationRoutes.Group("cards")
	{
		cardsAutomation.POST("", automation.UploadCards)
		cardsAutomation.DELETE("", automation.CleanCards)
	}

	cardPricesAutomation := automationRoutes.Group("card-prices")
	{
		cardPricesAutomation.POST("", automation.UploadCardPrices)
	}

	mobileBankAutomation := automationRoutes.Group("mobile-bank")
	{
		mobileBankAutomation.POST("", automation.UploadMobileBankData)
		mobileBankAutomation.DELETE("", automation.CleanMobileBankTable)
	}

	tusAutomation := automationRoutes.Group("call-center")
	{
		tusAutomation.POST("", automation.UploadTusData)
		tusAutomation.DELETE("", automation.CleanTusTable)
	}

	reportsAutomation := automationRoutes.Group("reports")
	{
		reportsAutomation.POST("", automation.CreateZIPReports)
	}

	return r
}
