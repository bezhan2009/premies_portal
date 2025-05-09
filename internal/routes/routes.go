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

	//depositRoute := r.Group("/deposit")
	//{
	//	depositRoute.GET("/portal", webSockets.WebSocketHandler)
	//	depositRoute.POST("", controllers.DesignDeposit)
	//}

	r.GET("/user", middlewares.CheckUserAuthentication,
		controllers.GetMyDataUser)

	usersRoute := r.Group("/users", middlewares.CheckUserAuthentication, middlewares.CheckUserNotWorker)
	{
		usersRoute.GET("", controllers.GetAllUsers)
		usersRoute.GET("/:id", controllers.GetUserByID)
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

	// knowledge маршруты для базы знаний
	knowledge := r.Group("/knowledge", middlewares.CheckUserAuthentication, middlewares.CheckUserKnowledgePerms)
	{
		knowledge.POST("", controllers.CreateKnowledge)

		knowledge.GET("/:id", controllers.GetKnowledgeByBaseID)
		knowledge.PATCH("/:id", controllers.UpdateKnowledge)
		knowledge.DELETE("/:id", controllers.DeleteKnowledge)
	}

	knowledgeDocs := knowledge.Group("/docs")
	{
		knowledgeDocs.POST("", middlewares.SaveFileFromResponse, controllers.CreateKnowledgeDoc)

		knowledgeDocs.GET("/:id", controllers.GetKnowledgeDocsByKnowledgeID)
		knowledgeDocs.PATCH("/:id", middlewares.SaveFileFromResponse, controllers.UpdateKnowledgeDoc)
		knowledgeDocs.DELETE("/:id", controllers.DeleteKnowledgeDoc)
	}

	knowledgeBases := knowledge.Group("/bases")
	{
		knowledgeBases.GET("", controllers.GetAllKnowledgeBases)
		knowledgeBases.POST("", controllers.CreateKnowledgeBase)
		knowledgeBases.PUT("/:id", controllers.UpdateKnowledgeBase)
		knowledgeBases.DELETE("/:id", controllers.DeleteKnowledgeBase)
	}

	// cards Маршруты для карт
	cards := r.Group("cards", middlewares.CheckUserAuthentication, middlewares.CheckUserOperator)

	cardSales := cards.Group("sales")
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

	// operatingActive маршруты активности

	operatingActive := r.Group("operating-active")
	{
		operatingActive.POST("", controllers.AddOperatingActive)
		operatingActive.PATCH("/:id", controllers.UpdateOperatingActive)
		operatingActive.DELETE("/:id", controllers.DeleteOperatingActive)
	}

	return r
}
