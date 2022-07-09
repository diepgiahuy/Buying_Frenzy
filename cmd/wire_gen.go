// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	config, err := ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	postgresStore := ProvidePostgreDB(config)
	ginServer := ProvideHandler(config, postgresStore)
	applicationContext := &ApplicationContext{
		Ctx:         ctx,
		Db:          postgresStore,
		httpHandler: ginServer,
	}
	return applicationContext, func() {
	}, nil
}

// wire.go:

var ApplicationSet = wire.NewSet(
	ProvideConfig,
	ProvidePostgreDB,
	ProvideHandler,
)
