package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/pkg/log"
)

type friendRoutes struct {
	logger log.Logger
}

func newFriendRoutes(group *gin.RouterGroup, l log.Logger) {
	f := &friendRoutes{
		logger: l,
	}
	group.PUT("/add/:user_id", f.addFriend)
	group.PUT("/delete/:user_id", f.deleteFriend)
}
func (f *friendRoutes) addFriend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "friend added"})
}
func (f *friendRoutes) deleteFriend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "friend deleted"})
}
