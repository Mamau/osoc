package v1

import (
	"context"
	"github.com/twitchtv/twirp"
	v1 "osoc/api/v1"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
)

type UserCtrl struct {
	logger  log.Logger
	service userinfo.UserService
}

func NewUserCtrl(logger log.Logger, service userinfo.UserService) *UserCtrl {
	return &UserCtrl{
		logger:  logger,
		service: service,
	}
}

func (u *UserCtrl) GetUser(ctx context.Context, req *v1.UserRequest) (*v1.UserGetResponse, error) {
	if err := req.Validate(); err != nil {
		u.logger.Err(err).Msg("error while validate")
		return nil, twirp.InvalidArgument.Error(err.Error())
	}

	user, err := u.service.GetUser(ctx, int(req.GetId()))
	if err != nil {
		return nil, twirp.NotFoundError(err.Error())
	}

	usr := &v1.User{
		Id: uint64(user.ID),
	}

	return &v1.UserGetResponse{
		User: usr,
	}, nil
}

func (u *UserCtrl) DeleteUser(ctx context.Context, req *v1.UserRequest) (*v1.UserOkResponse, error) {
	return &v1.UserOkResponse{}, nil
}

func (u *UserCtrl) UpdateUser(ctx context.Context, req *v1.UserUpdateRequest) (*v1.UserOkResponse, error) {
	return &v1.UserOkResponse{}, nil

}

func (u *UserCtrl) CreateUser(ctx context.Context, req *v1.UserCreateRequest) (*v1.UserOkResponse, error) {
	return &v1.UserOkResponse{}, nil

}
