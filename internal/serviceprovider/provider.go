package serviceprovider

import (
	v1 "osoc/internal/api/http/v1"
	"osoc/internal/config"
	dr "osoc/internal/repository/dialog"
	"osoc/internal/repository/friend"
	"osoc/internal/repository/post"
	"osoc/internal/repository/user"
	"osoc/internal/repository/webdata"
	"osoc/internal/usecase/dialog"
	"osoc/internal/usecase/friends"
	"osoc/internal/usecase/posts"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/broker/rabbit/consumer"
	"osoc/pkg/broker/rabbit/producer"

	"github.com/google/wire"
	"osoc/pkg/application"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(userinfo.UserService), new(*userinfo.Service)), userinfo.NewService,
	wire.Bind(new(v1.FriendService), new(*friends.Service)), friends.NewService,
	wire.Bind(new(v1.PostService), new(*posts.Service)), posts.NewService,
	wire.Bind(new(v1.PostCache), new(*post.Cache)), post.NewCacheRepository,
	wire.Bind(new(v1.DialogProvider), new(*dialog.Service)), dialog.NewService,
	//wire.Bind(new(userinfo.UserRepo), new(*user.Repository)), user.New,
	wire.Bind(new(userinfo.UserRepo), new(*user.TarantoolRepository)), user.NewTarantool,
	wire.Bind(new(dialog.MessageStorage), new(*dr.Repository)), dr.New,
	wire.Bind(new(friends.FriendRepo), new(*friend.Repository)), friend.New,
	wire.Bind(new(posts.PostRepo), new(*post.Repository)), post.New,
	wire.Bind(new(secure.UserSecureRepo), new(*user.SecureRepo)), user.NewSecureRepo,
	wire.Bind(new(userinfo.UserDaemon), new(*userinfo.Daemon)), userinfo.NewDaemon,
	NewBaseRouter,
	application.GetBuildVersion,
	config.GetPrometheusConfig,
	config.GetConfig,
	config.GetAppConfig,
	config.GetRabbitConfig,
	config.GetMysqlConfig,
	config.GetProxyMysqlConfig,
	config.GetRedisConfig,
	config.GetTarantoolConfig,
	NewRabbitConnection,
	NewHttp,
	producer.New,
	consumer.New,
	posts.NewConsumer,
	NewTarantool,
	NewMysql,
	NewProxyMysql,
	NewRedis,
	NewPrometheus,
	webdata.NewWebData,
	NewLogger,
	v1.NewRouter,
	secure.NewAuth,
)
