// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"osoc/internal/api/http/v1"
	"osoc/internal/config"
	"osoc/internal/repository/friend"
	"osoc/internal/repository/post"
	"osoc/internal/repository/user"
	"osoc/internal/repository/webdata"
	"osoc/internal/serviceprovider"
	"osoc/internal/usecase/friends"
	"osoc/internal/usecase/posts"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/application"
	"osoc/pkg/broker/rabbit/consumer"
	"osoc/pkg/broker/rabbit/producer"
)

// Injectors from wire.go:

func newApp() (*application.App, func(), error) {
	configConfig, err := config.GetConfig()
	if err != nil {
		return nil, nil, err
	}
	buildVersion, err := application.GetBuildVersion()
	if err != nil {
		return nil, nil, err
	}
	logger := serviceprovider.NewLogger(configConfig, buildVersion)
	app := config.GetAppConfig(configConfig)
	daemon := userinfo.NewDaemon(logger, app)
	engine := serviceprovider.NewBaseRouter(configConfig, logger, buildVersion, daemon)
	mysql := config.GetMysqlConfig(configConfig)
	db, cleanup, err := serviceprovider.NewMysql(mysql)
	if err != nil {
		return nil, nil, err
	}
	secureRepo := user.NewSecureRepo(db)
	auth := secure.NewAuth(logger, secureRepo, app)
	tarantool := config.GetTarantoolConfig(configConfig)
	connection, cleanup2, err := serviceprovider.NewTarantool(tarantool)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tarantoolRepository := user.NewTarantool(connection)
	service := userinfo.NewService(tarantoolRepository, logger)
	repository := friend.New(db)
	friendsService := friends.NewService(repository, tarantoolRepository, logger)
	postRepository := post.New(db)
	redis := config.GetRedisConfig(configConfig)
	client, cleanup3, err := serviceprovider.NewRedis(redis, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cache := post.NewCacheRepository(client, logger)
	rabbit := config.GetRabbitConfig(configConfig)
	rmqProducer := producer.New(rabbit, logger)
	postsService := posts.NewService(postRepository, tarantoolRepository, logger, cache, rabbit, rmqProducer)
	webData := webdata.NewWebData(tarantoolRepository, logger)
	handler := v1.NewRouter(engine, configConfig, auth, service, friendsService, postsService, logger, webData, cache)
	server := serviceprovider.NewHttp(handler, configConfig, logger)
	promConfig := config.GetPrometheusConfig(configConfig)
	promServer := serviceprovider.NewPrometheus(promConfig, logger)
	rmqConsumer := consumer.New(logger, rabbit)
	postsConsumer := posts.NewConsumer(logger, rmqConsumer, rabbit, tarantoolRepository)
	applicationApp := createApp(server, promServer, configConfig, logger, postsConsumer)
	return applicationApp, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
