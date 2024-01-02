
CREATE TYPE PROJECT_TYPE AS ENUM ('PRIVATE', 'PUBLIC');

CREATE TABLE  IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    cost INTEGER,
    duration INTEGER,
    project_type PROJECT_TYPE NOT NULL DEFAULT ('PRIVATE'),
    owner_id BIGINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(owner_id) REFERENCES users(id)
)