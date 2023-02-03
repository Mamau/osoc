package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/pkg/log"
)

type postRoutes struct {
	logger log.Logger
}

func newPostRoutes(group *gin.RouterGroup, l log.Logger) {
	p := &postRoutes{
		logger: l,
	}

	group.GET("/get/:id", p.retrievePost)
	group.POST("/create", p.createPost)
	group.PUT("/update", p.updatePost)
	group.PUT("/delete/:id", p.deletePost)
}
func (p *postRoutes) retrievePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "get post"})
}
func (p *postRoutes) createPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "create post"})
}
func (p *postRoutes) updatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "update post"})
}
func (p *postRoutes) deletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "delete post"})
}
