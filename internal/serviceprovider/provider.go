package serviceprovider

import (
	v1 "osoc/internal/api/http/v1"
	"osoc/internal/config"
	"osoc/internal/repository/user"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"

	"github.com/google/wire"
	"osoc/pkg/application"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(userinfo.UserService), new(*userinfo.Service)), userinfo.NewService,
	wire.Bind(new(userinfo.UserRepo), new(*user.Repository)), user.New,
	wire.Bind(new(secure.UserSecureRepo), new(*user.SecureRepo)), user.NewSecureRepo,
	wire.Bind(new(userinfo.UserDaemon), new(*userinfo.Daemon)), userinfo.NewDaemon,
	NewBaseRouter,
	application.GetBuildVersion,
	config.GetPrometheusConfig,
	config.GetConfig,
	config.GetAppConfig,
	config.GetMysqlConfig,
	config.GetRedisConfig,
	NewHttp,
	NewMysql,
	NewRedis,
	NewPrometheus,
	NewLogger,
	v1.NewRouter,
	v1.NewUserCtrl,
	v1.NewAuthCtrl,
	secure.NewAuth,
)
