package config

type Config struct {
	PostgresConfig
	Logrus
}

type PostgresConfig struct {
	Host     string `envconfig:"POSTGRES_HOST"`
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Db       string `envconfig:"POSTGRES_DB"`
	Port     string `envconfig:"POSTGRES_PORT"`
}

type Logrus struct {
	LogLevel string `envconfig:"LOG_LEVEL"`
}
