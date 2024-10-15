-- +goose Up
-- +goose StatementBegin

CREATE table user (
    id serial primary key,
    name VARCHAR(20) not null,
    email VARCHAR(30) not null,
    role VARCHAR(20) not null,
    password VARCHAR(20) not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;

-- +goose StatementEnd
