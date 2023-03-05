package producer

import (
	"github.com/streadway/amqp"
	"osoc/internal/config"
	"osoc/pkg/broker/rabbit/connection"
	"osoc/pkg/log"
)

type RMQProducer struct {
	logger log.Logger
	conf   config.Rabbit
}

func New(conf config.Rabbit, l log.Logger) RMQProducer {
	return RMQProducer{
		conf:   conf,
		logger: l,
	}
}

func (x RMQProducer) PublishMessage(queue, contentType string, body []byte) {
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
	if contentType == "" {
		contentType = "text/plain"
	}

	ch, err := conn.Channel()
	if err != nil {
		x.logger.Err(err).Msg("error while get channel connection")
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
		x.logger.Err(err).Msg("error while queue declare")
		return
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		})

	if err != nil {
		x.logger.Err(err).Msg("error while publish")
		return
	}
}
