package main

import (
	api "dis-test/internal/api/ports/entity"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	runAPIServer()
}

func runAPIServer() {
	db, err := initDB()
	if err != nil {
		logrus.Fatalf("Error occurred while opening db: %s", err.Error())
	}

	err = db.AutoMigrate(&api.Discipline{}) //TODO move to migrations
	if err != nil {
		logrus.Fatalf("Error occurred while automigrating: %s", err.Error())
	}
	//repos := repositories.NewRepository(db)
	//services := service.NewService(repos)
	//handlers := handler.NewHandler(services)
	//
	//server := apiserver.NewServer("8080", handlers.InitRoutes())
	//if err := server.Run(); err != nil {
	//	logrus.Fatalf("Error occurred while running http server: %s", err.Error())
	//}
}

func initDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
