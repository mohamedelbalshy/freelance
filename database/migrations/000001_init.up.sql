CREATE TABLE  IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT uc_username UNIQUE (username)
)

-- create migration example =====>  migrate create -ext sql -dir database/migrations/ -seq init
-- apply migration example =====>  migrate -path database/migrations -database postgres://root:1234@127.0.0.1:5432/go_freelance?sslmode=disable up

