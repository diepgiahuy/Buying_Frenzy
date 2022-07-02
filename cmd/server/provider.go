package main

import (
	"github.com/diepgiahuy/Buying_Frenzy/infra"
	"github.com/diepgiahuy/Buying_Frenzy/storage"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
	"log"
)

func ProvideConfig() (infra.Config, error) {
	cfg := infra.Config{}
	err := godotenv.Load("./infra/dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	err = envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ProvidePostgreDB(cfg infra.Config) *gorm.DB {
	return storage.NewDB(cfg.Host, cfg.User, cfg.Password, cfg.Db, cfg.Port)
}
