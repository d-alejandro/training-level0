package main

import (
	"encoding/json"
	"github.com/d-alejandro/training-level0/internal/memory_cache"
	"github.com/d-alejandro/training-level0/internal/migrations"
	"github.com/d-alejandro/training-level0/internal/models"
	"github.com/d-alejandro/training-level0/internal/nats_streaming"
	"github.com/d-alejandro/training-level0/internal/providers"
	"github.com/d-alejandro/training-level0/internal/repositories"
	"github.com/d-alejandro/training-level0/internal/use_cases"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	gormDB                 *gorm.DB
	orderModelStoreUseCase *use_cases.OrderModelStoreUseCase
	orderModelLoadUseCase  *use_cases.OrderModelLoadUseCase
	cache                  memory_cache.CacheInterface
)

func main() {
	bootProviders()
	runMigrations()
	initBindings()
	initCache()

	clientID := viper.GetString("CLIENT_ID_SUBSCRIBER")
	stanConnector := nats_streaming.NewStanConnector(clientID)

	stanConnection := stanConnector.CreateConnection()
	defer stanConnector.CloseConnection(stanConnection)

	stanSubscriber := nats_streaming.NewStanSubscriber(stanConnection)
	stanSubscriber.Subscribe(getSubscriberFunction)

	time.Sleep(60 * time.Second)
}

func bootProviders() {
	envReaderProvider := providers.NewEnvReaderProvider()
	envReaderProvider.InitViper()

	databaseProvider := providers.NewDatabaseProvider()
	gormDB = databaseProvider.InitGorm()
}

func runMigrations() {
	orderModelsTableMigration := migrations.NewOrderModelsTableMigration(gormDB)
	err := orderModelsTableMigration.Migrate()
	if err != nil {
		log.Fatal("Failed to complete migrations.")
	}
}

func initBindings() {
	orderModelStoreRepository := repositories.NewOrderModelStoreRepository(gormDB)
	orderModelStoreUseCase = use_cases.NewOrderModelStoreUseCase(orderModelStoreRepository)

	newOrderModelLoadRepository := repositories.NewOrderModelLoadRepository(gormDB)
	orderModelLoadUseCase = use_cases.NewOrderModelLoadUseCase(newOrderModelLoadRepository)
}

func initCache() {
	orderModels := orderModelLoadUseCase.Execute()
	cache = memory_cache.NewCache(orderModels)
}

func getSubscriberFunction(message *stan.Msg) {
	var model models.Model

	if err := json.Unmarshal(message.Data, &model); err != nil {
		log.Println("Incorrect json structure.")
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(model); err != nil {
		log.Println("Validation error.")
		return
	}

	if err := orderModelStoreUseCase.Execute(&model); err != nil {
		log.Println("Error saving to database.")
		return
	}

	if err := cache.SetModel(model.OrderUID, &model); err != nil {
		log.Println("Error saving to cache.")
		return
	}

	log.Println("The message is received.")
}
