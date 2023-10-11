package main

import (
	consumer "dis-test/internal/api/adapter/consumer"
	entity "dis-test/internal/api/adapter/entity"
	messageHandler "dis-test/internal/api/adapter/message_handler"
	appRepository "dis-test/internal/api/adapter/repository"
	_ "dis-test/internal/api/app/model"
	appSerializer "dis-test/internal/api/app/serializer"
	appService "dis-test/internal/api/app/service"
	appHandler "dis-test/internal/api/port/handler"
	pkg "dis-test/pkg/responser"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

//	@title			Discipline API
//	@version		1.0
//	@description	CRUD API for disciplines

// @host		localhost:8080
// @BasePath	/api
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db, err := initDB()
	if err != nil {
		logrus.Fatalf("Error occurred while opening db: %s", err.Error())
	}

	err = db.AutoMigrate(&entity.Discipline{})

	if err != nil {
		logrus.Fatalf("Error occurred while automigrating: %s", err.Error())
	}

	go runAPIServer(db)

	runMessageConsuming(db)
}

func runAPIServer(db *gorm.DB) {
	r := mux.NewRouter()

	repo := appRepository.NewRepository(db)
	serializer := appSerializer.NewSerializer()
	service := appService.NewService(repo, serializer)
	responser := pkg.NewResponser()
	appHandler.NewHandler(r, service, responser)

	r.PathPrefix("/api/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/redoc/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

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

func runMessageConsuming(db *gorm.DB) {
	repo := appRepository.NewRepository(db)
	serializer := appSerializer.NewSerializer()
	service := appService.NewService(repo, serializer)

	disciplineMessageHandler := messageHandler.NewDisciplineMessageHandler(service)
	disciplineConsumer := consumer.NewConsumer("disciplines", disciplineMessageHandler)

	config := getKafkaConfig()
	disciplineConsumer.Init(&config)
	disciplineConsumer.ReadMessages()
}

func getKafkaConfig() kafka.ConfigMap {
	m := make(map[string]kafka.ConfigValue)

	m["bootstrap.servers"] = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	m["group.id"] = os.Getenv("KAFKA_GROUP_ID")
	m["auto.offset.reset"] = "earliest"

	return m
}
