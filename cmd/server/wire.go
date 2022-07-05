//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
)

var ApplicationSet = wire.NewSet(
	ProvideConfig,
	ProvidePostgreDB,
	ProvideStorage,
)

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	wire.Build(
		ApplicationSet,
		wire.Struct(new(ApplicationContext), "*"),
	)
	return nil, nil, nil
}
