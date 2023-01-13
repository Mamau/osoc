package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/twitchtv/twirp"
	"net/http"
	api "osoc/api/v1"
	"osoc/internal/config"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
	"osoc/pkg/router/middleware/recoverer"
)

func NewRouter(
	engine *gin.Engine,
	userCtrl *UserCtrl,
	authCtrl *AuthCtrl,
	conf *config.Config,
	logger log.Logger,
) http.Handler {

	engine.Use(recoverer.New())

	defaultPrefix := twirp.WithServerPathPrefix("")
	innerPrefix := twirp.WithServerPathPrefix("/inner")
	// -------------------
	// CREATE TWIRP ROUTES
	// -------------------
	userHandler := api.NewUserServiceServer(userCtrl, innerPrefix)
	authHandler := api.NewAuthServiceServer(authCtrl, defaultPrefix)

	// -------------------
	// SECURE ROUTES
	// -------------------
	inner := engine.Group("")
	inner.Use(jwt.New(
		jwt.HMACSecret([]byte(conf.App.AppJWTSecret)),
		jwt.Logger(logger),
	))
	inner.POST(userHandler.PathPrefix()+"*action", func(context *gin.Context) {
		userHandler.ServeHTTP(context.Writer, context.Request)
	})

	// -------------------
	// EXTERNAL ROUTES
	// -------------------
	engine.POST(authHandler.PathPrefix()+"*action", func(context *gin.Context) {
		authHandler.ServeHTTP(context.Writer, context.Request)
	})

	return engine
}
