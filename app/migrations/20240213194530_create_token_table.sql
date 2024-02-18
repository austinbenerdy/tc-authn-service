-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_tokens
(
    id varchar(255) UNIQUE NOT NULL,
    userId varchar(255) NOT NULL,
    token varchar(255) NOT NULL,
    expiration DATETIME NOT NULL,
    expired BOOL default false
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_tokens
-- +goose StatementEnd
