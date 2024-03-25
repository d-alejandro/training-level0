package nats_streaming

import (
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
)

type StanConnector struct {
	stanClusterID string
	clientID      string
	natsOption    stan.Option
}

func NewStanConnector(clientID string) *StanConnector {
	stanClusterID := viper.GetString("STAN_CLUSTER_ID")

	natsUrl := viper.GetString("NATS_URL")
	natsOption := stan.NatsURL(natsUrl)

	return &StanConnector{
		stanClusterID: stanClusterID,
		clientID:      clientID,
		natsOption:    natsOption,
	}
}

func (connector *StanConnector) CreateConnection() stan.Conn {
	stanConnection, err := stan.Connect(
		connector.stanClusterID,
		connector.clientID,
		connector.natsOption,
	)

	if err != nil {
		log.Fatal(err)
	}

	return stanConnection
}

func (connector *StanConnector) CloseConnection(stanConnection stan.Conn) {
	err := stanConnection.Close()

	if err != nil {
		log.Fatal(err)
	}
}
