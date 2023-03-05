package connection

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func New(opts ...Option) (*amqp.Connection, error) {
	o := options{
		protocol:  "amqp",
		user:      "guest",
		password:  "guest",
		host:      "rabbitmq",
		port:      5673,
		heartbeat: 10 * time.Second,
		locale:    "en_US",
	}

	for _, opt := range opts {
		opt(&o)
	}

	// Define RabbitMQ server URL.
	// amqp://guest:guest@message-broker:5672/
	amqpServerURL := fmt.Sprintf("%s://%s:%s@%s:%d/", o.protocol, o.user, o.password, o.host, o.port)

	// Create a new RabbitMQ connection.
	return amqp.DialConfig(
		amqpServerURL,
		amqp.Config{
			Heartbeat: o.heartbeat,
			Locale:    o.locale,
		},
	)
}
