package v1

import (
	"context"
	v1 "osoc/api/v1"
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

func (a *AuthCtrl) Authorization(ctx context.Context, req *v1.AuthRequest) (*v1.AuthOKResponse, error) {
	var token string
	if ok := a.authService.CheckUser(ctx, req.FirstName, req.Password); ok {
		token = "urtoken"
	}

	return &v1.AuthOKResponse{
		Token: token,
	}, nil
}
