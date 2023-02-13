package post

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"osoc/internal/entity"
	"osoc/pkg/mysql"
	"time"
)

type Repository struct {
	db *mysql.DB
}

func New(db *mysql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) DeletePost(ctx context.Context, id int) error {
	sqlQuery, args, err := r.db.Builder.Delete("posts").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = r.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (r *Repository) GetPost(ctx context.Context, id int) (entity.Post, error) {
	sqlQuery, args, err := r.db.Builder.
		Select("id", "title", "text", "user_id", "created_at").
		From("posts").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return entity.Post{}, err
	}

	var post entity.Post

	err = r.db.GetContext(ctx, &post, sqlQuery, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Post{}, entity.ErrNotFound
	}

	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
func (r *Repository) UpdatePost(ctx context.Context, post entity.Post) error {
	sqlQuery, args, err := r.db.Builder.Update("posts").
		Set("text", post.Text).
		Where(squirrel.Eq{"id": post.ID}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = r.db.ExecContext(ctx, sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
func (r *Repository) AddPost(ctx context.Context, post entity.Post) (int, error) {
	sqlQuery, args, err := r.db.Builder.Insert("posts").
		Columns("title", "text", "user_id", "created_at").
		Values(post.Title, post.Text, post.UserID, time.Now()).
		ToSql()

	if err != nil {
		return 0, err
	}
	res, err := r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}
func (r *Repository) Feeds(ctx context.Context, userId int, limit int, offset int) ([]entity.Post, error) {
	sqlQuery, args, err := r.db.Builder.
		Select("id", "title", "text", "user_id", "created_at").
		From("posts").
		LeftJoin("friends USING (user_id)").
		Where(squirrel.Eq{"user_id": userId}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()

	if err != nil {
		return nil, err
	}

	data := make([]entity.Post, 0, limit)
	if err = r.db.SelectContext(ctx, &data, sqlQuery, args...); err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, entity.ErrNotFound
	}
	return data, nil
}
