// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"osoc/internal/api/http/v1"
	"osoc/internal/config"
	"osoc/internal/repository/dialog"
	"osoc/internal/repository/friend"
	"osoc/internal/repository/post"
	"osoc/internal/repository/user"
	"osoc/internal/repository/webdata"
	"osoc/internal/serviceprovider"
	dialog2 "osoc/internal/usecase/dialog"
	"osoc/internal/usecase/friends"
	"osoc/internal/usecase/posts"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/application"
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
	repository := user.New(db)
	service := userinfo.NewService(repository, logger)
	friendRepository := friend.New(db)
	friendsService := friends.NewService(friendRepository, repository, logger)
	postRepository := post.New(db)
	redis := config.GetRedisConfig(configConfig)
	client, cleanup2, err := serviceprovider.NewRedis(redis, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	cache := post.NewCacheRepository(client, logger)
	postsService := posts.NewService(postRepository, repository, logger, cache)
	webData := webdata.NewWebData(repository, logger)
	proxyMysql := config.GetProxyMysqlConfig(configConfig)
	mysqlProxyMysql, cleanup3, err := serviceprovider.NewProxyMysql(proxyMysql)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dialogRepository := dialog.New(db, mysqlProxyMysql)
	dialogService := dialog2.NewService(logger, dialogRepository)
	handler := v1.NewRouter(engine, configConfig, auth, service, friendsService, postsService, logger, webData, cache, dialogService)
	server := serviceprovider.NewHttp(handler, configConfig, logger)
	promConfig := config.GetPrometheusConfig(configConfig)
	promServer := serviceprovider.NewPrometheus(promConfig, logger)
	applicationApp := createApp(server, promServer, configConfig, logger)
	return applicationApp, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
