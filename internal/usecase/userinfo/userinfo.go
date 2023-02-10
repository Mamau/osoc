package userinfo

import (
	"context"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"

	"osoc/pkg/log"
)

type Service struct {
	repo   UserRepo
	logger log.Logger
}

func NewService(repo UserRepo, logger log.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) SearchUser(ctx context.Context, query *request.UserSearch) ([]entity.User, error) {
	users, err := s.repo.SearchUsers(ctx, query)
	if err != nil {
		log.AddContext(ctx, s.logger.Err(err)).
			Msgf("userinfo: could not get users by filter #%v", query)
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUser(ctx context.Context, id int) (entity.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		log.AddContext(ctx, s.logger.Err(err)).
			Msgf("userinfo: could not get user #%v", id)
		return entity.User{}, err
	}

	return user, nil
}
