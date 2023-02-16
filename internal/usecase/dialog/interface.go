package dialog

import (
	"context"
	"osoc/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=../../mocks/message.go -package=mocks
type (
	// MessageProvider -.
	MessageProvider interface {
		Save(ctx context.Context, message entity.Message) error
		GetList(ctx context.Context, userID int) ([]entity.Message, error)
	}
)
