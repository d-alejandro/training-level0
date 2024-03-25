package main

import (
	"fmt"
	"github.com/d-alejandro/training-level0/internal/nats_streaming"
	"github.com/d-alejandro/training-level0/internal/providers"
	"github.com/d-alejandro/training-level0/internal/resources"
	"github.com/spf13/viper"
)

func main() {
	envReaderProvider := providers.NewEnvReaderProvider()
	envReaderProvider.InitViper()

	clientID := viper.GetString("CLIENT_ID_PUBLISHER")
	stanConnector := nats_streaming.NewStanConnector(clientID)

	stanConnection := stanConnector.CreateConnection()
	defer stanConnector.CloseConnection(stanConnection)

	stanPublisher := nats_streaming.NewStanPublisher(stanConnection)

	data := resources.GetModelJSON()
	stanPublisher.Publish(data)

	fmt.Println("Message published.")
}
