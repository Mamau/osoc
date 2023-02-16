package dialog

import (
	"context"
	"osoc/internal/entity"
	"osoc/internal/errors"
	"osoc/pkg/log"
)

type Service struct {
	logger     log.Logger
	repository MessageProvider
}

func NewService(l log.Logger, r MessageProvider) *Service {
	return &Service{
		logger:     l,
		repository: r,
	}
}

func (s *Service) Messages(ctx context.Context, userID int) ([]entity.Message, error) {
	data, err := s.repository.GetList(ctx, userID)
	if err != nil {
		s.logger.Err(err).Msg("error while get messages")
		return nil, errors.SomethingWentWrong
	}
	return data, nil
}

func (s *Service) SaveMessage(ctx context.Context, userID int, authorID int, text string) error {
	message := entity.Message{
		UserID:   userID,
		AuthorID: authorID,
		Text:     text,
	}
	if err := s.repository.Save(ctx, message); err != nil {
		s.logger.Err(err).Msg("error while save dialog message")
		return errors.SomethingWentWrong
	}
	return nil
}
