package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"osoc/internal/entity"
	"osoc/pkg/mysql"
)

type SecureRepo struct {
	db *mysql.DB
}

func NewSecureRepo(db *mysql.DB) *SecureRepo {
	return &SecureRepo{db: db}
}

func (u *SecureRepo) GetUserByName(ctx context.Context, firstName string) (entity.SecureUser, error) {
	sqlQuery, args, err := u.db.Builder.
		Select("id", "first_name", "last_name", "age", "sex", "interests", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"first_name": firstName}).
		ToSql()

	if err != nil {
		return entity.SecureUser{}, err
	}

	var user entity.SecureUser

	err = u.db.GetContext(ctx, &user, sqlQuery, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.SecureUser{}, entity.ErrNotFound
	}

	if err != nil {
		return entity.SecureUser{}, err
	}

	return user, nil
}
