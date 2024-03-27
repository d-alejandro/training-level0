package main

import (
	"fmt"
	"github.com/d-alejandro/training-level0/internal/nats_streaming"
	"github.com/d-alejandro/training-level0/internal/providers"
	"github.com/d-alejandro/training-level0/internal/resources"
	"github.com/spf13/viper"
	"strconv"
)

var stanPublisher *nats_streaming.StanPublisher

func main() {
	envReaderProvider := providers.NewEnvReaderProvider()
	envReaderProvider.InitViper()

	clientID := viper.GetString("CLIENT_ID_PUBLISHER")
	stanConnector := nats_streaming.NewStanConnector(clientID)

	stanConnection := stanConnector.CreateConnection()
	defer stanConnector.CloseConnection(stanConnection)

	stanPublisher = nats_streaming.NewStanPublisher(stanConnection)

	publish("")

	for x := 1; x < 7; x++ {
		publish(strconv.Itoa(x))
	}

	fmt.Println("Messages published.")
}

func publish(value string) {
	data := resources.GetModelJSON(value)
	stanPublisher.Publish(data)
}
