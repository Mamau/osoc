package serviceprovider

import (
	v1 "osoc/internal/api/http/v1"
	"osoc/internal/config"
	"osoc/internal/repository/friend"
	"osoc/internal/repository/post"
	"osoc/internal/repository/user"
	"osoc/internal/repository/webdata"
	"osoc/internal/usecase/friends"
	"osoc/internal/usecase/posts"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"

	"github.com/google/wire"
	"osoc/pkg/application"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(userinfo.UserService), new(*userinfo.Service)), userinfo.NewService,
	wire.Bind(new(friends.FrienderService), new(*friends.Service)), friends.NewService,
	wire.Bind(new(posts.PostService), new(*posts.Service)), posts.NewService,
	wire.Bind(new(userinfo.UserRepo), new(*user.Repository)), user.New,
	wire.Bind(new(friends.FriendRepo), new(*friend.Repository)), friend.New,
	wire.Bind(new(posts.PostRepo), new(*post.Repository)), post.New,
	wire.Bind(new(secure.UserSecureRepo), new(*user.SecureRepo)), user.NewSecureRepo,
	wire.Bind(new(userinfo.UserDaemon), new(*userinfo.Daemon)), userinfo.NewDaemon,
	NewBaseRouter,
	application.GetBuildVersion,
	config.GetPrometheusConfig,
	config.GetConfig,
	config.GetAppConfig,
	config.GetMysqlConfig,
	config.GetRedisConfig,
	post.NewCacheRepository,
	NewHttp,
	NewMysql,
	NewRedis,
	NewPrometheus,
	webdata.NewWebData,
	NewLogger,
	v1.NewRouter,
	secure.NewAuth,
)
