package main

import (
	"os"
	"osoc/internal/config"
	"osoc/pkg/application"
	"osoc/pkg/log"
	"osoc/pkg/transport/http"
	"osoc/pkg/transport/prom"
)

var id, _ = os.Hostname()

func createApp(
	hs *http.Server,
	prom *prom.Server,
	c *config.Config,
	logger log.Logger,
) *application.App {
	return application.New(
		application.ID(id),
		application.Name(c.App.Name),
		application.Location(c.App.TZ),
		application.Logger(logger),
		application.Servers(
			hs,
			prom,
		),
	)
}

func main() {
	a, cleanup, err := newApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = a.Run(); err != nil {
		panic(err)
	}
}
