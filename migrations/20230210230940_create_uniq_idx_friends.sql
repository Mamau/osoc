-- +goose Up
-- +goose StatementBegin
ALTER TABLE friends
    ADD UNIQUE user_friend_uniq_idx(user_id, friend_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE friends
DROP INDEX user_friend_uniq_idx;
-- +goose StatementEnd
