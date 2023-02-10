package posts

import (
	"context"
	"osoc/internal/entity"
	"osoc/internal/errors"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
)

type Service struct {
	postRepo PostRepo
	userRepo userinfo.UserRepo
	logger   log.Logger
}

func NewService(repo PostRepo, ur userinfo.UserRepo, logger log.Logger) *Service {
	return &Service{
		postRepo: repo,
		userRepo: ur,
		logger:   logger,
	}
}
func (s *Service) GetPost(ctx context.Context, id int) (entity.Post, error) {
	post, err := s.postRepo.GetPost(ctx, id)
	if err != nil {
		s.logger.Err(err).Msg("error while get post")
		return entity.Post{}, errors.SomethingWentWrong
	}
	return post, nil
}
func (s *Service) CreatePost(ctx context.Context, userID int, text string) error {
	post := entity.Post{
		Title:  "example",
		UserID: userID,
		Text:   text,
	}
	if err := s.postRepo.AddPost(ctx, post); err != nil {
		s.logger.Err(err).Msg("error while add post")
		return errors.SomethingWentWrong
	}
	return nil
}
