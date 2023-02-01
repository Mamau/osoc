package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"
	"time"

	"github.com/Masterminds/squirrel"
	"osoc/pkg/mysql"
)

type Repository struct {
	db      *mysql.DB
	slaveDB *mysql.SlaveMysql
}

func New(db *mysql.DB, slaveDB *mysql.SlaveMysql) *Repository {
	return &Repository{
		db:      db,
		slaveDB: slaveDB,
	}
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
func (u *Repository) SearchUsers(ctx context.Context, query *request.UserSearch) ([]entity.User, error) {
	buildQuery := u.slaveDB.Builder.
		Select("id", "first_name", "last_name", "age", "sex", "interests", "created_at").
		From("users")

	if query.FirstName != "" {
		buildQuery = buildQuery.Where(squirrel.Like{"first_name": fmt.Sprintf("%s%%", query.FirstName)})
	}
	if query.LastName != "" {
		buildQuery = buildQuery.Where(squirrel.Like{"last_name": fmt.Sprintf("%s%%", query.LastName)})
	}
	buildQuery.OrderBy("id")
	sqlQuery, args, err := buildQuery.ToSql()
	if err != nil {
		return nil, err
	}

	var users []entity.User

	if err = u.db.SelectContext(ctx, &users, sqlQuery, args...); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, entity.ErrNotFound
	}

	return users, nil
}
func (u *Repository) MultiCreateUser(ctx context.Context, users []entity.SecureUser) error {
	builder := u.db.Builder.Insert("users").
		Columns("first_name", "last_name", "age", "sex", "interests", "password", "created_at")

	for _, v := range users {
		user := v
		builder = builder.Values(user.FirstName, user.LastName, user.Age, user.Sex, user.Interests, user.Password, time.Now())
	}

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	if _, err = u.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (u *Repository) CreateUser(ctx context.Context, user entity.SecureUser) error {
	sqlQuery, args, err := u.db.Builder.Insert("users").
		Columns("first_name", "last_name", "age", "sex", "interests", "password", "created_at").
		Values(user.FirstName, user.LastName, user.Age, user.Sex, user.Interests, user.Password, time.Now()).
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
