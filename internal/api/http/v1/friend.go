package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/usecase/friends"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
)

type friendRoutes struct {
	logger  log.Logger
	service friends.FrienderService
}

func newFriendRoutes(group *gin.RouterGroup, l log.Logger, fs friends.FrienderService) {
	f := &friendRoutes{
		logger:  l,
		service: fs,
	}
	group.PUT("/add/:user_id", f.addFriend)
	group.PUT("/delete/:user_id", f.deleteFriend)
}
func (f *friendRoutes) addFriend(c *gin.Context) {
	var req request.AddFriendRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt(jwt.XUserIDKey)

	if err := f.service.AddFriend(c.Request.Context(), userID, req.FriendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}
func (f *friendRoutes) deleteFriend(c *gin.Context) {
	var req request.DeleteFriendRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt(jwt.XUserIDKey)
	if err := f.service.DeleteFriend(c.Request.Context(), userID, req.FriendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
