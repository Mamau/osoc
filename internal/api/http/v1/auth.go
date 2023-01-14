package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"
	"osoc/internal/usecase/secure"
	"osoc/pkg/log"
)

type authRoutes struct {
	logger      log.Logger
	authService *secure.Auth
}

func newAuthRoutes(group *gin.RouterGroup, logger log.Logger, authService *secure.Auth) {
	a := &authRoutes{
		logger:      logger,
		authService: authService,
	}
	group.POST("/refresh", a.refresh)
	group.POST("/authorization", a.authorization)
	group.POST("/registration", a.registration)
}

func (a *authRoutes) refresh(c *gin.Context) {
	var r request.Refresh
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := a.authService.RefreshToken(c.Request.Context(), r.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tokens})
}

func (a *authRoutes) authorization(c *gin.Context) {
	var r request.Authorization
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := a.authService.LoginUser(c.Request.Context(), r.FirstName, r.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokens})
}

func (a *authRoutes) registration(c *gin.Context) {
	var req request.Registration
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entity.SecureUser{
		User: entity.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Sex:       req.Sex,
			Age:       req.Age,
			Interests: req.Interests,
		},
		Password: req.Password,
	}

	tokens, err := a.authService.RegisterUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokens})
}
