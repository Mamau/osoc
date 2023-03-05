package consumer

import (
	"github.com/streadway/amqp"
	"osoc/internal/config"
	"osoc/pkg/broker/rabbit/connection"
	"osoc/pkg/log"
)

type RMQConsumer struct {
	logger     log.Logger
	conf       config.Rabbit
	MsgHandler func(queue string, msg amqp.Delivery, err error)
}

func New(l log.Logger, conf config.Rabbit) RMQConsumer {
	return RMQConsumer{
		logger: l,
		conf:   conf,
	}
}

func (x RMQConsumer) Consume(queue string) {
	conn, err := connection.New(
		connection.Host(x.conf.Host),
		connection.User(x.conf.User),
		connection.Password(x.conf.Password),
		connection.Port(x.conf.Port),
		connection.Protocol(x.conf.Protocol),
	)
	if err != nil {
		x.logger.Err(err).Msg("error while connect to consumer")
		return
	}

	defer func() { _ = conn.Close() }()

	ch, err := conn.Channel()
	if err != nil {
		x.logger.Err(err).Msg("error while get channel connection for consumer")
		return
	}
	defer func() { _ = ch.Close() }()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		x.logger.Err(err).Msg("error while queue declare for consumer")
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		x.logger.Err(err).Msg("failed to register a consumer")
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			x.MsgHandler(queue, d, nil)
		}
	}()

	x.logger.Info().Msgf("started listening for messages on '%s' queue", queue)
	<-forever
}
