package config

type Rabbit struct {
	Host        string `envconfig:"RA_HOST" default:"rabbitmq"`
	Port        int    `envconfig:"RA_PORT" default:"5673"`
	User        string `envconfig:"RA_USER" default:"guest"`
	Password    string `envconfig:"RA_PASSWORD" default:"guest"`
	Protocol    string `envconfig:"RA_PROTOCOL" default:"amqp"`
	PostChannel string `envconfig:"RA_POST_CHANNEL" default:"posts"`
}

func GetRabbitConfig(config *Config) Rabbit {
	return config.Rabbit
}
