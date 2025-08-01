package routes

import (
	"os"
	_ "premiesPortal/docs"
	"premiesPortal/internal/controllers"
	"premiesPortal/internal/controllers/automation"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/internal/security"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.Static("/uploads", "./uploads")

	pingRoute := r.Group("/ping")
	{
		pingRoute.GET("/", controllers.Ping)
	}

	// user & users Маршруты для сущности пользователей
	r.GET("/user", middlewares.CheckUserAuthentication,
		controllers.GetMyDataUser)
	r.PATCH("/user", middlewares.CheckUserAuthentication,
		controllers.UpdateUsersPassword)

	usersRoute := r.Group("/users", middlewares.CheckUserAuthentication, middlewares.CheckUserNotWorker)
	{
		usersRoute.GET("", controllers.GetAllUsers)
		usersRoute.GET("/:id", controllers.GetUserByID)
	}

	// worker & workers маршруты для сущности
	r.GET("/worker", middlewares.CheckUserAuthentication,
		controllers.GetMyDataWorker)
	r.GET("/worker/card-details", middlewares.CheckUserAuthentication,
		controllers.GetMyCardDetailsWorker)
	r.GET("/worker/mb-details", middlewares.CheckUserAuthentication,
		controllers.GetAllWorkersMobileBankDetails)

	workerRoute := r.Group("/workers", middlewares.CheckUserAuthentication, middlewares.CheckUserNotWorker)
	{
		workerRoute.GET("", controllers.GetAllWorkers)
		workerRoute.GET("/:id", controllers.GetWorkerByID)
	}

	workersCardDetails := workerRoute.Group("card-details", middlewares.CheckUserAuthentication, middlewares.CheckUserNotWorker)
	{
		workersCardDetails.GET("", controllers.GetCardDetailsWorkers)
		workersCardDetails.GET("/stats", controllers.GetStatisticsCards)
		workersCardDetails.GET("/:id", controllers.GetCardDetailsWorkerByID)
	}

	// auth Маршруты для авторизаций
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", middlewares.CheckUserAuthentication, middlewares.CheckSignupPerms, controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
		auth.POST("/refresh", controllers.RefreshToken)

		auth.GET("/google", controllers.GoogleLogin)
		auth.GET("/google/callback", controllers.GoogleCallback)
	}

	// office Маршруты для офисов
	office := r.Group("/office", middlewares.CheckUserAuthentication, middlewares.CheckUserOperatorOrChairman)
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

	mobileBank := r.Group("mobile-bank", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)
	{
		mobileBank.POST("", controllers.AddMobileBankSale)
		mobileBank.PATCH("/:id", controllers.UpdateMobileBankSale)
		mobileBank.DELETE("/:id", controllers.DeleteMobileBankSale)
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

	r.GET("/worker/tests", middlewares.CheckUserAuthentication, controllers.GetTestsForWorker)

	tests := r.Group("/tests", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)
	{
		tests.GET("", controllers.GetAllTests)
		tests.GET("/:id", controllers.GetTestById)
		tests.POST("", controllers.CreateTest)
		tests.PATCH("/:id", controllers.UpdateTest)
		tests.DELETE("/:id", controllers.DeleteTest)
	}

	testQuestions := tests.Group("/questions")
	{
		testQuestions.POST("/:id", controllers.CreateTestQuestions)
		testQuestions.PATCH("/:id", controllers.UpdateTestQuestions)
		testQuestions.DELETE("/:id", controllers.DeleteTestQuestions)
	}

	testOptions := testQuestions.Group("/options")
	{
		testOptions.POST("/:id", controllers.CreateTestOptions)
		testOptions.PATCH("/:id", controllers.UpdateTestOptions)
		testOptions.DELETE("/:id", controllers.DeleteTestOptions)
	}

	testAnswers := r.Group("tests/answers", middlewares.CheckUserAuthentication)
	{
		testAnswers.GET("", middlewares.CheckUserOperator, controllers.GetTestAnswers)
		testAnswers.GET("/allow", controllers.AllowedAnswer)
		testAnswers.GET("/:id/single", controllers.GetTestAnswersByAnswerId)
		testAnswers.GET("/:id", controllers.GetTestAnswersByTestId)
		testAnswers.POST("", controllers.CreateTestAnswers)
	}

	automationRoutes := r.Group("automation", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)

	cardsAutomation := automationRoutes.Group("cards")
	{
		cardsAutomation.POST("", automation.UploadAutomationFile, automation.UploadCards)
		cardsAutomation.DELETE("", automation.CleanCards)
	}

	cardPricesAutomation := automationRoutes.Group("card-prices", automation.UploadAutomationFile)
	{
		cardPricesAutomation.POST("", automation.UploadCardPrices)
	}

	mobileBankAutomation := automationRoutes.Group("mobile-bank")
	{
		mobileBankAutomation.POST("", automation.UploadAutomationFile, automation.UploadMobileBankData)
		mobileBankAutomation.DELETE("", automation.CleanMobileBankTable)
	}

	tusAutomation := automationRoutes.Group("call-center")
	{
		tusAutomation.POST("", automation.UploadAutomationFile, automation.UploadTusData)
		tusAutomation.DELETE("", automation.CleanTusTable)
	}

	reportsAutomation := automationRoutes.Group("reports")
	{
		reportsAutomation.GET("", automation.CreateZIPReports)
		reportsAutomation.GET("/:id", automation.CreateExcelReport)
	}

	accountantAutomation := automationRoutes.Group("accountant")
	{
		accountantAutomation.GET("", automation.CreateXLSXAccountantReport)
	}

	return r
}
