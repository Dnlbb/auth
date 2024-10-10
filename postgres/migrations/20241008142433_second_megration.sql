-- +goose Up
-- +goose StatementBegin
CREATE table Users (
    id serial primary key,
    name VARCHAR(20) not null,
    email VARCHAR(30) not null,
    role VARCHAR(20) not null,
    password VARCHAR(20) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table Users;
-- +goose StatementEnd
