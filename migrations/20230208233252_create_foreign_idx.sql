-- +goose Up
-- +goose StatementBegin
ALTER TABLE friends
    ADD CONSTRAINT fk_friends_users
        FOREIGN KEY (user_id) REFERENCES users (id)
            ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE friends
DROP FOREIGN KEY fk_friends_users;
-- +goose StatementEnd
