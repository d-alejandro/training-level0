package main

import (
	"fmt"
	"github.com/d-alejandro/training-level0/internal/migrations"
	"github.com/d-alejandro/training-level0/internal/nats_streaming"
	"github.com/d-alejandro/training-level0/internal/providers"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	envReaderProvider := providers.NewEnvReaderProvider()
	envReaderProvider.InitViper()

	databaseProvider := providers.NewDatabaseProvider()
	gorm := databaseProvider.InitGorm()

	orderModelsTableMigration := migrations.NewOrderModelsTableMigration(gorm)
	err := orderModelsTableMigration.Migrate()
	if err != nil {
		log.Fatal("Failed to complete migrations.")
	}

	clientID := viper.GetString("CLIENT_ID_SUBSCRIBER")
	stanConnector := nats_streaming.NewStanConnector(clientID)

	stanConnection := stanConnector.CreateConnection()
	defer stanConnector.CloseConnection(stanConnection)

	stanSubscriber := nats_streaming.NewStanSubscriber(stanConnection)
	stanSubscriber.Subscribe(func(message *stan.Msg) {
		fmt.Println(string(message.Data))
	})

	time.Sleep(60 * time.Second)
}
