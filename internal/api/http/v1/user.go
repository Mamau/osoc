package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
)

type userRoutes struct {
	logger  log.Logger
	service userinfo.UserService
}

func newUserRoutes(group *gin.RouterGroup, logger log.Logger, service userinfo.UserService) {
	u := &userRoutes{
		logger:  logger,
		service: service,
	}
	group.GET("/user", u.GetUser)
}

func (u *userRoutes) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	claim, ok := jwt.FromContext(ctx)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	user, err := u.service.GetUser(ctx, claim.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
