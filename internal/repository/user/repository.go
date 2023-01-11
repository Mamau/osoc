package user

import (
	"context"
	"database/sql"
	"errors"
	"osoc/internal/entity"
	"time"

	"github.com/Masterminds/squirrel"
	"osoc/pkg/mysql"
)

type Repository struct {
	db *mysql.DB
}

func New(db *mysql.DB) *Repository {
	return &Repository{db: db}
}

func (u *Repository) DeleteUser(ctx context.Context, id int) error {
	sqlQuery, args, err := u.db.Builder.Delete("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (u *Repository) UpdateUser(ctx context.Context, user entity.User) error {
	sqlQuery, args, err := u.db.Builder.Update("users").
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("age", user.Age).
		Set("sex", user.Sex).
		Set("interests", user.Interests).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": user.ID}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (u *Repository) CreateUser(ctx context.Context, user entity.User) error {
	sqlQuery, args, err := u.db.Builder.Insert("users").
		Columns("first_name", "last_name", "age", "sex", "interests", "created_at").
		Values(user.FirstName, user.LastName, user.Age, user.Sex, user.Interests, time.Now()).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (u *Repository) GetUser(ctx context.Context, id int) (entity.User, error) {
	sqlQuery, args, err := u.db.Builder.
		Select("id", "first_name", "last_name", "age", "sex", "interests", "created_at").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return entity.User{}, err
	}

	var user entity.User

	err = u.db.GetContext(ctx, &user, sqlQuery, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.User{}, entity.ErrNotFound
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
