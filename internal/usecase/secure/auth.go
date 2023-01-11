package secure

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"osoc/pkg/log"
)

type Auth struct {
	repo   UserSecureRepo
	logger log.Logger
}

func NewAuth(logger log.Logger, repo UserSecureRepo) *Auth {
	return &Auth{
		logger: logger,
		repo:   repo,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *Auth) CheckUser(ctx context.Context, firstName, password string) bool {
	user, err := a.repo.GetUserByName(ctx, firstName)
	if err != nil {
		a.logger.Err(err).Msgf("error while get user by name: %w", err)
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		a.logger.Err(err).Msgf("error while compare hash password: %w", err)
		return false
	}
	return true
}
