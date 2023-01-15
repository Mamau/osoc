package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/config"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
)

type userRoutes struct {
	logger  log.Logger
	service userinfo.UserService
}

// prefix routes - /api/v1/user
func newUserRoutes(group *gin.RouterGroup, logger log.Logger, service userinfo.UserService, conf *config.Config) {
	u := &userRoutes{
		logger:  logger,
		service: service,
	}
	group.GET(":id", u.GetUserById)
	group.GET("/search", u.SearchUser)

	// -------------------
	// SECURE ROUTES (work by jwt token)
	// -------------------
	group.Use(jwt.New(
		jwt.HMACSecret([]byte(conf.App.AppJWTSecret)),
		jwt.Logger(logger),
	))
	group.GET("", u.GetUser)
}
func (u *userRoutes) SearchUser(c *gin.Context) {
	var req request.UserSearch
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.FirstName == "" && req.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("need filter pls")})
		return
	}

	users, err := u.service.SearchUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
func (u *userRoutes) GetUserById(c *gin.Context) {
	var req request.UserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.GetUser(c.Request.Context(), req.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
func (u *userRoutes) GetUser(c *gin.Context) {
	userID := c.GetInt(jwt.XUserIDKey)

	user, err := u.service.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
