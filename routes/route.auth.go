package routes

import (
	handlers "authservice/handlers/auth"
	repositories "authservice/repositories/auth"
	services "authservice/services/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/

	pingRepository := repositories.NewRepositoryPing(db)
	pingService := services.NewServicePing(pingRepository)
	pingHandler := handlers.NewHandlerPing(pingService)

	registerRepository := repositories.NewRepositoryRegister(db)
	registerService := services.NewServiceRegister(registerRepository)
	registerHandler := handlers.NewHandlerRegister(registerService)

	activationRepository := repositories.NewRepositoryActivation(db)
	activationService := services.NewServiceActivation(activationRepository)
	activationHandler := handlers.NewHandlerActivation(activationService)

	LoginRepository := repositories.NewRepositoryLogin(db)
	loginService := services.NewServiceLogin(LoginRepository)
	loginHandler := handlers.NewHandlerLogin(loginService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/users/ping", pingHandler.PingHandler)
	groupRoute.POST("/users/register", registerHandler.RegisterHandler)
	groupRoute.POST("/users/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/uesrs/login", loginHandler.LoginHandler)

}
