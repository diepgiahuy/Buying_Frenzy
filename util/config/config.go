package config

type Config struct {
	PostgresConfig
	Logrus
	ServerConfig
}

type PostgresConfig struct {
	Host         string `envconfig:"POSTGRES_HOST"`
	User         string `envconfig:"POSTGRES_USER"`
	Password     string `envconfig:"POSTGRES_PASSWORD"`
	Db           string `envconfig:"POSTGRES_DB"`
	Port         string `envconfig:"POSTGRES_PORT"`
	MigrationUrl string `envconfig:"MIGRATION_URL"`
}

type Logrus struct {
	LogLevel string `envconfig:"LOG_LEVEL"`
}

type ServerConfig struct {
	Port string `envconfig:"HTTP_PORT"`
	Env  int    `nvconfig:"Env"`
}
