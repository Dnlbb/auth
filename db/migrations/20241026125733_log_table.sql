-- +goose Up
-- +goose StatementBegin
CREATE TABLE log (
    id bigserial primary key,
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE log;
-- +goose StatementEnd
