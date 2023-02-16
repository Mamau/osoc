package posts

import (
	"context"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/posts.go -package=mocks
type (
	// PostRepo -.
	PostRepo interface {
		DeletePost(ctx context.Context, id int) error
		GetPost(ctx context.Context, id int) (entity.Post, error)
		UpdatePost(ctx context.Context, post entity.Post) error
		AddPost(ctx context.Context, post entity.Post) (int, error)
		Feeds(ctx context.Context, userId int, limit int, offset int) ([]entity.Post, error)
	}
)
