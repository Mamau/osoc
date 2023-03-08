package friends

import (
	"context"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interface.go -destination=../../mocks/friends.go -package=mocks
type (
	// FriendRepo -.
	FriendRepo interface {
		AddFriend(ctx context.Context, user entity.User, friend entity.User) error
		DeleteFriend(ctx context.Context, user entity.User, friend entity.User) error
	}
)
