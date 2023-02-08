-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS friends(
    user_id INT NOT NULL,
    friend_id  INT NOT NULL,
    INDEX idx_friend_id (friend_id),
    INDEX idx_user_id (user_id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friends;
-- +goose StatementEnd
