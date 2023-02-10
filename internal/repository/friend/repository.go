package friend

import (
	"context"
	"github.com/Masterminds/squirrel"
	"osoc/internal/entity"
	"osoc/pkg/mysql"
)

type Repository struct {
	db *mysql.DB
}

func New(db *mysql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (u *Repository) DeleteFriend(ctx context.Context, user entity.User, friend entity.User) error {
	sqlQuery, args, err := u.db.Builder.Delete("friends").
		Where(squirrel.Eq{"user_id": user.ID}).
		Where(squirrel.Eq{"friend_id": friend.ID}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (u *Repository) AddFriend(ctx context.Context, user entity.User, friend entity.User) error {
	sqlQuery, args, err := u.db.Builder.Insert("friends").
		Columns("user_id", "friend_id").
		Values(user.ID, friend.ID).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil

}
