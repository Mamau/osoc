-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts
    ADD CONSTRAINT fk_posts_users
        FOREIGN KEY (user_id) REFERENCES users (id)
            ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts
DROP FOREIGN KEY fk_posts_users;
-- +goose StatementEnd
