package posts

import (
	"context"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/posts.go -package=mocks
type (
	// PostService -.
	PostService interface {
		GetPost(ctx context.Context, id int) (entity.Post, error)
		CreatePost(ctx context.Context, userID int, text string) error
	}

	// PostRepo -.
	PostRepo interface {
		DeletePost(ctx context.Context, id int) error
		GetPost(ctx context.Context, id int) (entity.Post, error)
		UpdatePost(ctx context.Context, post entity.Post) error
		AddPost(ctx context.Context, post entity.Post) error
	}
)
