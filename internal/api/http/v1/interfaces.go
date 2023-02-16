package v1

import (
	"context"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../../mocks/services.go -package=mocks
type (
	// FriendService -.
	FriendService interface {
		AddFriend(ctx context.Context, userId int, friendId int) error
		DeleteFriend(ctx context.Context, userId int, friendId int) error
	}
	// PostService -.
	PostService interface {
		PostList(ctx context.Context, userID int, feeds request.Feeds) ([]entity.Post, error)
		GetPost(ctx context.Context, id int) (entity.Post, error)
		CreatePost(ctx context.Context, userID int, text string) error
		UpdatePost(ctx context.Context, req request.UpdatePost) error
		DeletePost(ctx context.Context, id int) error
	}

	// PostCache -.
	PostCache interface {
		Save(ctx context.Context, userID int, post entity.Post) error
		GetFeeds(ctx context.Context, userID int) ([]entity.Post, error)
	}

	// DialogProvider -.
	DialogProvider interface {
		SaveMessage(ctx context.Context, userID int, authorID int, text string) error
		Messages(ctx context.Context, userID int) ([]entity.Message, error)
	}
)
