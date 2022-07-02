package main

import (
	"context"
	"github.com/google/wire"
)

var ApplicationSet = wire.NewSet(
	ProvideConfig,
	ProvidePostgreDB,
)

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	wire.Build(
		ApplicationSet,
		wire.Struct(new(ApplicationContext), "*"),
	)
	return nil, nil, nil
}
