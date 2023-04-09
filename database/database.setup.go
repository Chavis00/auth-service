package database

import (
	"authservice/models"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDsn() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser,
		dbPassword,
		dbName,
		dbHost,
		dbPort,
	)
}

func SetupDatabase() *gorm.DB {
	dsn := SetUpDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connect into Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connect into Database Successfully")
	}

	err = db.AutoMigrate(
		&models.ModelAuth{},
	//	&models.ModelStudent{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
