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
	}
)
