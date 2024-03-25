package nats_streaming

import (
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
)

type StanPublisher struct {
	stanConnection stan.Conn
}

func NewStanPublisher(stanConnection stan.Conn) *StanPublisher {
	return &StanPublisher{stanConnection}
}

func (stanPublisher *StanPublisher) Publish(data string) {
	subject := viper.GetString("STAN_CONNECTION_SUBJECT")

	stanErr := stanPublisher.stanConnection.Publish(
		subject,
		[]byte(data),
	)

	if stanErr != nil {
		log.Fatal(stanErr)
	}
}
