-- +goose Up
-- +goose StatementBegin
INSERT INTO users (first_name, last_name, age, sex, interests, password, created_at)
VALUES ('Willy', 'Barankin', 33, 'male', 'soccer', '$2a$14$odyxR448UtxohAJK6nHD5eu/Wy9OBy2AUdJEkIyyymjQiIfTX5sza', NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM users
WHERE first_name = 'Willy';
-- +goose StatementEnd
