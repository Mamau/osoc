package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"osoc/internal/entity"
	"osoc/pkg/mysql"
	"time"
)

type SecureRepo struct {
	db *mysql.DB
}

func NewSecureRepo(db *mysql.DB) *SecureRepo {
	return &SecureRepo{db: db}
}
func (u *SecureRepo) CreateUser(ctx context.Context, user *entity.SecureUser) (int64, error) {
	sqlQuery, args, err := u.db.Builder.Insert("users").
		Columns("first_name", "last_name", "age", "sex", "interests", "password", "created_at").
		Values(user.FirstName, user.LastName, user.Age, user.Sex, user.Interests, user.Password, time.Now()).
		ToSql()
	if err != nil {
		return 0, err
	}
	res, err := u.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (u *SecureRepo) GetUserById(ctx context.Context, id int) (entity.SecureUser, error) {
	var user entity.SecureUser
	sqlQuery, args, err := u.db.Builder.
		Select("id", "first_name", "last_name", "age", "sex", "interests", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return user, err
	}

	err = u.db.GetContext(ctx, &user, sqlQuery, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return user, entity.ErrNotFound
	}

	if err != nil {
		return user, err
	}

	return user, nil
}
func (u *SecureRepo) GetUserByName(ctx context.Context, firstName string) (entity.SecureUser, error) {
	var user entity.SecureUser
	sqlQuery, args, err := u.db.Builder.
		Select("id", "first_name", "last_name", "age", "sex", "interests", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"first_name": firstName}).
		ToSql()

	if err != nil {
		return user, err
	}

	err = u.db.GetContext(ctx, &user, sqlQuery, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return user, entity.ErrNotFound
	}

	if err != nil {
		return user, err
	}

	return user, nil
}
