package serviceprovider

import (
	"osoc/internal/config"

	"osoc/pkg/log"
	"osoc/pkg/transport/prom"
)

func NewPrometheus(config config.PromConfig, logger log.Logger) *prom.Server {
	server := prom.New(
		prom.Logger(logger),
		prom.GuiPort(config.GuiPort),
		prom.Port(config.Port),
		prom.Handle(config.Handle),
	)

	return server
}
