package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/config"
	"osoc/internal/repository/webdata"
	"osoc/internal/usecase/secure"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
)

func NewRouter(
	engine *gin.Engine,
	conf *config.Config,
	authService *secure.Auth,
	s userinfo.UserService,
	logger log.Logger,
	wd *webdata.WebData,
) http.Handler {
	commonGroup := engine.Group("/api/v1")
	commonGroup.GET("/", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	{
		newAuthRoutes(commonGroup, logger, authService)
	}

	secureGroup := commonGroup.Group("/user")
	{
		newUserRoutes(secureGroup, logger, s, conf, wd)
	}

	return engine
}
