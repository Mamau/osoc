package config

type Redis struct {
	Host     string `envconfig:"RE_HOST" default:"redis"`
	Port     string `envconfig:"RE_PORT" default:"6379"`
	Channel  string `envconfig:"RE_CHANNEL" default:"some-channel"`
	Username string `envconfig:"RE_USERNAME" default:""`
	Password string `envconfig:"RE_PASSWORD" default:""`
	DB       int    `envconfig:"RE_DB" default:""`
}

func GetRedisConfig(config *Config) Redis {
	return config.Redis
}
