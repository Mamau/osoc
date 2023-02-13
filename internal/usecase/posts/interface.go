package posts

import (
	"context"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/posts.go -package=mocks
type (
	// PostService -.
	PostService interface {
		PostList(ctx context.Context, userID int, feeds request.Feeds) ([]entity.Post, error)
		GetPost(ctx context.Context, id int) (entity.Post, error)
		CreatePost(ctx context.Context, userID int, text string) error
		UpdatePost(ctx context.Context, req request.UpdatePost) error
		DeletePost(ctx context.Context, id int) error
	}

	// PostRepo -.
	PostRepo interface {
		DeletePost(ctx context.Context, id int) error
		GetPost(ctx context.Context, id int) (entity.Post, error)
		UpdatePost(ctx context.Context, post entity.Post) error
		AddPost(ctx context.Context, post entity.Post) (int, error)
		Feeds(ctx context.Context, userId int, limit int, offset int) ([]entity.Post, error)
	}
)
