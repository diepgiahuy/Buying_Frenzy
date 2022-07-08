package cmd

import (
	"github.com/diepgiahuy/Buying_Frenzy/pkg/api"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/diepgiahuy/Buying_Frenzy/util/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
	"log"
)

func ProvideConfig() (*config.Config, error) {
	cfg := config.Config{}
	err := godotenv.Load("./util/config/dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	err = envconfig.Process("", &cfg)
	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}

func ProvidePostgreDB(cfg *config.Config) *gorm.DB {
	return storage.NewDB(cfg.Host, cfg.User, cfg.Password, cfg.Db, cfg.PostgresConfig.Port)
}

func ProvideStorage(db *gorm.DB) *storage.Repo {
	return storage.NewRepo(db)
}

func ProvideHandler(cfg *config.Config, repo *storage.Repo) *api.GinServer {
	return api.NewServer(&cfg.ServerConfig, repo)
}
