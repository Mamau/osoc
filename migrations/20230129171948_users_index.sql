-- +goose Up
-- +goose StatementBegin
CREATE index firstName_lastName on users(first_name,last_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP index firstName_lastName;
-- +goose StatementEnd
