package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/repository/post"
	"osoc/internal/usecase/posts"
	"osoc/pkg/log"
	"osoc/pkg/router/middleware/auth/jwt"
)

type postRoutes struct {
	logger      log.Logger
	postService posts.PostService
	cache       *post.Cache
}

func newPostRoutes(group *gin.RouterGroup, l log.Logger, ps posts.PostService, cache *post.Cache) {
	p := &postRoutes{
		logger:      l,
		postService: ps,
		cache:       cache,
	}

	group.GET("/feed", p.feeds)
	group.GET("/get/:id", p.retrievePost)
	group.POST("/create", p.createPost)
	group.PUT("/update", p.updatePost)
	group.PUT("/delete/:id", p.deletePost)
}
func (p *postRoutes) feeds(c *gin.Context) {
	var req request.Feeds
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetInt(jwt.XUserIDKey)

	data, err := p.postService.PostList(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
func (p *postRoutes) retrievePost(c *gin.Context) {
	var req request.RetrievePost
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := p.postService.GetPost(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
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
	if err := p.postService.UpdatePost(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "post updated"})
}
func (p *postRoutes) deletePost(c *gin.Context) {
	var req request.DeletePost
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.postService.DeletePost(c.Request.Context(), req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "post deleted"})
}
