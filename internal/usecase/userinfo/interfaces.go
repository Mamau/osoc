package userinfo

import (
	"context"
	"osoc/internal/api/http/v1/request"

	"osoc/pkg/healthcheck"

	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/userinfo.go -package=mocks
type (
	// UserService -.
	UserService interface {
		GetUser(ctx context.Context, id int) (entity.User, error)
		SearchUser(ctx context.Context, query *request.UserSearch) ([]entity.User, error)
	}
	// UserRepo -.
	UserRepo interface {
		GetUser(ctx context.Context, id int) (entity.User, error)
		MultiCreateUser(ctx context.Context, users []entity.SecureUser) error
		CreateUser(ctx context.Context, user entity.SecureUser) error
		SearchUsers(ctx context.Context, query *request.UserSearch) ([]entity.User, error)
		UpdateUser(ctx context.Context, user entity.User) error
		DeleteUser(ctx context.Context, id int) error
	}
	// UserDaemon -.
	UserDaemon interface {
		Run()
		Terminate(ctx context.Context) error
		Healthcheck(ctx context.Context) healthcheck.ProbeStatus
	}
)
