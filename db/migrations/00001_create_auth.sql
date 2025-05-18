-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE role_type AS ENUM ('user', 'admin');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role role_type NOT NULL,
    PRIMARY KEY (user_id, role)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS role_type;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
