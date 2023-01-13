package v1

import (
	"context"
	"github.com/twitchtv/twirp"
	v1 "osoc/api/v1"
	"osoc/internal/entity"
	"osoc/internal/usecase/secure"
	"osoc/pkg/log"
)

type AuthCtrl struct {
	logger      log.Logger
	authService *secure.Auth
}

func NewAuthCtrl(logger log.Logger, authService *secure.Auth) *AuthCtrl {
	return &AuthCtrl{
		logger:      logger,
		authService: authService,
	}
}
func (a *AuthCtrl) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshOkResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, twirp.InvalidArgument.Error(err.Error())
	}
	tokens, err := a.authService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}

	return &v1.RefreshOkResponse{
		Token:        tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
func (a *AuthCtrl) Authorization(ctx context.Context, req *v1.AuthRequest) (*v1.AuthOKResponse, error) {
	tokens, err := a.authService.LoginUser(ctx, req.FirstName, req.Password)
	if err != nil {
		return nil, twirp.NotFoundError(err.Error())
	}

	return &v1.AuthOKResponse{
		Token:        tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (a *AuthCtrl) Registration(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterOkResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, twirp.InvalidArgument.Error(err.Error())
	}

	user := &entity.SecureUser{
		User: entity.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Sex:       req.Sex,
			Age:       int(req.Age),
			Interests: req.Interests,
		},
		Password: req.Password,
	}

	tokens, err := a.authService.RegisterUser(ctx, user)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return &v1.RegisterOkResponse{
		Token:        tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
