package config

type Mysql struct {
	Host      string `envconfig:"MY_HOST" default:"mysql"`
	Port      int    `envconfig:"MY_PORT" default:"3306"`
	User      string `envconfig:"MY_USER" default:"root"`
	Password  string `envconfig:"MY_PASSWORD" default:"root"`
	DbName    string `envconfig:"MY_DB_NAME" default:"osoc"`
	ParseTime bool   `envconfig:"MY_PARSE_TIME" default:"true"`
}

func GetMysqlConfig(config *Config) Mysql {
	return config.Mysql
}
