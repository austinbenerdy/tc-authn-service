-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id       varchar(255) not null,
    email    varchar(255) not null,
    password longtext     not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
