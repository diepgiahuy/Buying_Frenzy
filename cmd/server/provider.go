package main

import (
	"github.com/diepgiahuy/Buying_Frenzy/infra/config"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
	"log"
)

func ProvideConfig() (config.Config, error) {
	cfg := config.Config{}
	err := godotenv.Load("./infra/config/dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	err = envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ProvidePostgreDB(cfg config.Config) *gorm.DB {
	return storage.NewDB(cfg.Host, cfg.User, cfg.Password, cfg.Db, cfg.Port)
}

func ProvideStorage(db *gorm.DB) *storage.Repo {
	return storage.NewRepo(db)
}
