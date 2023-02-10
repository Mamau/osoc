package friends

import (
	"context"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/friends.go -package=mocks
type (
	// FrienderService -.
	FrienderService interface {
		AddFriend(ctx context.Context, userId int, friendId int) error
		DeleteFriend(ctx context.Context, userId int, friendId int) error
	}

	// FriendRepo -.
	FriendRepo interface {
		AddFriend(ctx context.Context, user entity.User, friend entity.User) error
		DeleteFriend(ctx context.Context, user entity.User, friend entity.User) error
	}
)
