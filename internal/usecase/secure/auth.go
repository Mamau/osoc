package secure

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"osoc/internal/config"
	"osoc/internal/entity"
	"osoc/pkg/jwt"
	"osoc/pkg/log"
)

type Auth struct {
	repo   UserSecureRepo
	config config.App
	logger log.Logger
}

func NewAuth(logger log.Logger, repo UserSecureRepo, conf config.App) *Auth {
	return &Auth{
		logger: logger,
		repo:   repo,
		config: conf,
	}
}
func (a *Auth) RefreshToken(ctx context.Context, oldRefreshToken string) (entity.Tokens, error) {
	id, err := jwt.ExtractIDFromToken(oldRefreshToken, a.config.RefreshTokenSecret)
	if err != nil {
		return entity.Tokens{}, fmt.Errorf("error while extract id")
	}

	user, err := a.repo.GetUserById(ctx, id)
	if err == nil {
		return entity.Tokens{}, fmt.Errorf("user not found")
	}

	accessToken, err := jwt.CreateAccessToken(id, user.FirstName, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create access token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	refreshToken, err := jwt.CreateRefreshToken(id, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create refresh token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	return entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (a *Auth) RegisterUser(ctx context.Context, user *entity.SecureUser) (entity.Tokens, error) {
	if _, err := a.repo.GetUserByName(ctx, user.FirstName); err == nil {
		return entity.Tokens{}, fmt.Errorf("user with name %s already exists", user.FirstName)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		a.logger.Err(err).Msgf("error while generate hash password: %w", err)
		return entity.Tokens{}, fmt.Errorf("error while register")
	}
	user.Password = string(bytes)

	id, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		a.logger.Err(err).Msgf("error while create user: %w", err)
		return entity.Tokens{}, fmt.Errorf("error while create user")
	}
	accessToken, err := jwt.CreateAccessToken(int(id), user.FirstName, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create access token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	refreshToken, err := jwt.CreateRefreshToken(int(id), a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create refresh token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	return entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (a *Auth) LoginUserByID(ctx context.Context, id int) (entity.Tokens, error) {
	user, err := a.repo.GetUserById(ctx, id)
	if err != nil {
		a.logger.Err(err).Msgf("error while get user by id: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	accessToken, err := jwt.CreateAccessToken(user.ID, user.FirstName, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create access token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	refreshToken, err := jwt.CreateRefreshToken(user.ID, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create refresh token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	return entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (a *Auth) LoginUser(ctx context.Context, firstName, password string) (entity.Tokens, error) {
	user, err := a.repo.GetUserByName(ctx, firstName)
	if err != nil {
		a.logger.Err(err).Msgf("error while get user by name: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		a.logger.Err(err).Msgf("error while compare hash password: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	accessToken, err := jwt.CreateAccessToken(user.ID, user.FirstName, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create access token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	refreshToken, err := jwt.CreateRefreshToken(user.ID, a.config.AppJWTSecret)
	if err != nil {
		a.logger.Err(err).Msgf("error while create refresh token: %w", err)
		return entity.Tokens{}, fmt.Errorf("invalid credentials")
	}

	return entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
