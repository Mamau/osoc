package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/usecase/posts"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
)

type postRoutes struct {
	logger      log.Logger
	postService posts.PostService
}

func newPostRoutes(group *gin.RouterGroup, l log.Logger, ps posts.PostService) {
	p := &postRoutes{
		logger:      l,
		postService: ps,
	}

	group.GET("/get/:id", p.retrievePost)
	group.POST("/create", p.createPost)
	group.PUT("/update", p.updatePost)
	group.PUT("/delete/:id", p.deletePost)
}
func (p *postRoutes) retrievePost(c *gin.Context) {
	var req request.RetrievePost
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := p.postService.GetPost(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}
func (p *postRoutes) createPost(c *gin.Context) {
	var req request.Post
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt(jwt.XUserIDKey)
	if err := p.postService.CreatePost(c.Request.Context(), userID, req.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}
func (p *postRoutes) updatePost(c *gin.Context) {
	var req request.UpdatePost
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "update post"})
}
func (p *postRoutes) deletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "delete post"})
}
