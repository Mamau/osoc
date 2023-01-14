package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/config"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
)

func NewRouter(
	engine *gin.Engine,
	conf *config.Config,
	authService *secure.Auth,
	s userinfo.UserService,
	logger log.Logger,
) http.Handler {
	commonGroup := engine.Group("/api/v1")

	commonGroup.GET("/", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	{
		newAuthRoutes(commonGroup, logger, authService)
	}

	// -------------------
	// SECURE ROUTES
	// -------------------
	secureGroup := commonGroup.Group("/secure")
	secureGroup.Use(jwt.New(
		jwt.HMACSecret([]byte(conf.App.AppJWTSecret)),
		jwt.Logger(logger),
	))
	{
		newUserRoutes(secureGroup, logger, s)
	}

	return engine
}
