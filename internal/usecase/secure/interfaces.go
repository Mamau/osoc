package secure

import (
	"context"

	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/usersecure.go -package=mocks
type (
	// UserSecureRepo -.
	UserSecureRepo interface {
		GetUserByName(ctx context.Context, firstName string) (entity.SecureUser, error)
		GetUserById(ctx context.Context, id int) (entity.SecureUser, error)
		CreateUser(ctx context.Context, user *entity.SecureUser) (int64, error)
	}
)
