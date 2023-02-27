package posts

import (
	"context"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"
	"osoc/internal/errors"
	"osoc/internal/repository/post"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/log"
	"time"
)

type Service struct {
	postRepo PostRepo
	userRepo userinfo.UserRepo
	logger   log.Logger
	cache    *post.Cache
}

func NewService(repo PostRepo, ur userinfo.UserRepo, logger log.Logger, cache *post.Cache) *Service {
	return &Service{
		postRepo: repo,
		userRepo: ur,
		logger:   logger,
		cache:    cache,
	}
}
func (s *Service) PostList(ctx context.Context, userID int, feeds request.Feeds) ([]entity.Post, error) {
	cachedData, err := s.cache.GetFeeds(ctx, userID)
	if err != nil {
		s.logger.Err(err).Msg("error while get cached posts list")
	}
	if len(cachedData) > 0 {
		return cachedData, nil
	}

	data, err := s.postRepo.Feeds(ctx, userID, feeds.Limit, feeds.Offset)
	if err != nil {
		s.logger.Err(err).Msg("error while get post list")
		return nil, errors.SomethingWentWrong
	}
	return data, nil
}
func (s *Service) DeletePost(ctx context.Context, id int) error {
	if err := s.postRepo.DeletePost(ctx, id); err != nil {
		s.logger.Err(err).Msg("error while delete post")
		return errors.SomethingWentWrong
	}
	return nil
}
func (s *Service) UpdatePost(ctx context.Context, req request.UpdatePost) error {
	data := entity.Post{
		ID:   req.ID,
		Text: req.Text,
	}

	if err := s.postRepo.UpdatePost(ctx, data); err != nil {
		s.logger.Err(err).Msg("error while update post")
		return errors.SomethingWentWrong
	}

	return nil
}
func (s *Service) GetPost(ctx context.Context, id int) (entity.Post, error) {
	p, err := s.postRepo.GetPost(ctx, id)
	if err != nil {
		s.logger.Err(err).Msg("error while get post")
		return entity.Post{}, errors.SomethingWentWrong
	}
	return p, nil
}
func (s *Service) CreatePost(ctx context.Context, userID int, text string) error {
	p := entity.Post{
		Title:     "example",
		UserID:    userID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	id, err := s.postRepo.AddPost(ctx, p)
	if err != nil {
		s.logger.Err(err).Msg("error while add post")
		return errors.SomethingWentWrong
	}
	p.ID = id

	if err := s.cache.Save(ctx, userID, p); err != nil {
		s.logger.Err(err).Msg("error save post to cache")
	}

	return nil
}
