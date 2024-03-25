package nats_streaming

import (
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
)

type StanSubscriber struct {
	stanConnection stan.Conn
}

func NewStanSubscriber(stanConnection stan.Conn) *StanSubscriber {
	return &StanSubscriber{stanConnection}
}

func (stanSubscriber *StanSubscriber) Subscribe(function func(message *stan.Msg)) {
	subject := viper.GetString("STAN_CONNECTION_SUBJECT")
	durableName := viper.GetString("STAN_DURABLE_NAME")

	_, err := stanSubscriber.stanConnection.Subscribe(
		subject,
		function,
		stan.DurableName(durableName),
		stan.DeliverAllAvailable(),
	)

	if err != nil {
		log.Fatal(err)
	}
}
