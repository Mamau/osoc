package serviceprovider

import (
	nh "net/http"

	"osoc/internal/config"

	"osoc/pkg/log"
	"osoc/pkg/transport/http"
)

func NewHttp(handler nh.Handler, conf *config.Config, logger log.Logger) *http.Server {
	return http.New(
		http.Logger(logger),
		http.Handler(handler),
		http.Addr(conf.App.Port),
	)
}
