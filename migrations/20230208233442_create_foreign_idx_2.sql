-- +goose Up
-- +goose StatementBegin
ALTER TABLE friends
    ADD CONSTRAINT fk_friends_users_2
        FOREIGN KEY (friend_id) REFERENCES users (id)
            ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE friends
DROP FOREIGN KEY fk_friends_users_2;
-- +goose StatementEnd
