-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    text       TEXT         NOT NULL,
    user_id    INT          NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    INDEX idx_post_user_id (user_id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
