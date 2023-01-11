package serviceprovider

import (
	"osoc/internal/config"

	"osoc/pkg/application"
	"osoc/pkg/log"
)

func NewLogger(conf *config.Config, buildVersion application.BuildVersion) log.Logger {
	return log.NewLogger(
		log.Env(conf.App.Environment),
		log.Level(conf.App.LogLevel),
		log.BuildCommit(buildVersion.Commit),
		log.BuildTime(buildVersion.Time),
		log.Prettify(conf.App.PrettyLogs),
	)
}
