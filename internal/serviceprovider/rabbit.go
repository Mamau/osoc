package serviceprovider

import (
	"github.com/streadway/amqp"
	"osoc/internal/config"
	"osoc/pkg/broker/rabbit/connection"
)

func NewRabbitConnection(conf config.Rabbit) (*amqp.Connection, func(), error) {
	con, err := connection.New(
		connection.Host(conf.Host),
		connection.User(conf.User),
		connection.Password(conf.Password),
		connection.Port(conf.Port),
		connection.Protocol(conf.Protocol),
	)
	if err != nil {
		return nil, nil, err
	}

	closeCon := func() {
		_ = con.Close()
	}

	return con, closeCon, nil
}
