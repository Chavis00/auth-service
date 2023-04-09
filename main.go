package main

import (
	database "authservice/database"
	"authservice/routes"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	app := SetupRouter()
	logrus.Fatal(app.Run(":" + os.Getenv("SERVER_PORT")))
}

func SetupRouter() *gin.Engine {
	// Obtener db
	db := database.SetupDatabase()
	// Iniciar Aplicacion
	app := gin.Default()
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.BestCompression))
	// INICIO RUTAS DE AUTH
	routes.InitAuthRoutes(db, app)

	return app
}
