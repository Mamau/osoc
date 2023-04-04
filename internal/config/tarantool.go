package config

type Tarantool struct {
	Host     string `envconfig:"TA_HOST" default:"tarantool"`
	Port     int    `envconfig:"TA_PORT" default:"3301"`
	User     string `envconfig:"TA_USER" default:"admin"`
	Password string `envconfig:"TA_PASSWORD" default:"admin"`
}

func GetTarantoolConfig(config *Config) Tarantool {
	return config.Tarantool
}
