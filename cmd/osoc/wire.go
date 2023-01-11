//go:build wireinject

package main

import (
	"osoc/internal/serviceprovider"

	"github.com/google/wire"

	"osoc/pkg/application"
)

func newApp() (*application.App, func(), error) {
	panic(wire.Build(
		serviceprovider.ProviderSet,
		createApp,
	))
}
