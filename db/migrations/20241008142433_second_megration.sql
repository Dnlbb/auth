-- +goose Up
-- +goose StatementBegin

CREATE TABLE users  (
    id BIGSERIAL primary key,
    name TEXT not null,
    email TEXT not null,
    role TEXT not null,
    password TEXT not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;

-- +goose StatementEnd
