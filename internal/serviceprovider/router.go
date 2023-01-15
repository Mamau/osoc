package serviceprovider

import (
	"osoc/internal/config"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/router/middleware/recoverer"

	"github.com/gin-gonic/gin"
	app "osoc/pkg/application"
	"osoc/pkg/log"
	"osoc/pkg/router"
	"osoc/pkg/router/middleware/logging"
	"osoc/pkg/router/middleware/servertiming"
	"osoc/pkg/router/middleware/timeout"
)

func NewBaseRouter(conf *config.Config, logger log.Logger, version app.BuildVersion, dm userinfo.UserDaemon) *gin.Engine {
	return router.New(
		router.Logger(logger),
		router.DocPath(conf.App.SwaggerFolder),
		router.BuildCommit(version.Commit),
		router.BuildTime(version.Time),
		router.ReadinessProbes(dm.Healthcheck),
		router.Middlewares(
			recoverer.New(
				recoverer.Logger(logger),
			),
			servertiming.New(),
			timeout.New(),
			logging.New(
				logging.Level(conf.App.LogLevel),
				logging.Env(conf.App.Environment),
				logging.Fallback(logger),
				logging.Prettify(conf.App.PrettyLogs),
			),
		))
}
