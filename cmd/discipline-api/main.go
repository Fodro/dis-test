package main

import (
	entity "dis-test/internal/api/adapter/entity"
	appRepository "dis-test/internal/api/adapter/repository"
	appSerializer "dis-test/internal/api/app/serializer"
	appService "dis-test/internal/api/app/service"
	appHandler "dis-test/internal/api/port/handler"
	pkg "dis-test/pkg/responser"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
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

	err = db.AutoMigrate(&entity.Discipline{}) //TODO move to migrations
	if err != nil {
		logrus.Fatalf("Error occurred while automigrating: %s", err.Error())
	}
	r := mux.NewRouter()

	repo := appRepository.NewRepository(db)
	serializer := appSerializer.NewSerializer()
	service := appService.NewService(repo, serializer)
	responser := pkg.NewResponser()
	appHandler.NewHandler(r, service, responser)

	logrus.Fatal(http.ListenAndServe(":8080", r))
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
