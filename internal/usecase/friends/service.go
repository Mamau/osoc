package friends

import (
	"context"
	"osoc/internal/errors"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
)

type Service struct {
	friendRepo FriendRepo
	userRepo   userinfo.UserRepo
	logger     log.Logger
}

func NewService(repo FriendRepo, ur userinfo.UserRepo, logger log.Logger) *Service {
	return &Service{
		friendRepo: repo,
		userRepo:   ur,
		logger:     logger,
	}
}
func (s *Service) DeleteFriend(ctx context.Context, userId int, friendId int) error {
	user, err := s.userRepo.GetUser(ctx, userId)
	if err != nil {
		s.logger.Err(err).Msg("error while fetch user for delete")
		return errors.SomethingWentWrong
	}
	friend, err := s.userRepo.GetUser(ctx, friendId)
	if err != nil {
		s.logger.Err(err).Msg("error while fetch friend for delete")
		return errors.SomethingWentWrong
	}
	if err := s.friendRepo.DeleteFriend(ctx, user, friend); err != nil {
		s.logger.Err(err).Msg("error while delete friend")
		return errors.SomethingWentWrong
	}

	return nil
}
func (s *Service) AddFriend(ctx context.Context, userId int, friendId int) error {
	user, err := s.userRepo.GetUser(ctx, userId)
	if err != nil {
		s.logger.Err(err).Msg("error while fetch user")
		return errors.SomethingWentWrong
	}
	friend, err := s.userRepo.GetUser(ctx, friendId)
	if err != nil {
		s.logger.Err(err).Msg("error while fetch friend")
		return errors.SomethingWentWrong
	}
	if err := s.friendRepo.AddFriend(ctx, user, friend); err != nil {
		s.logger.Err(err).Msg("error while add friend")
		return errors.SomethingWentWrong
	}

	return nil
}
