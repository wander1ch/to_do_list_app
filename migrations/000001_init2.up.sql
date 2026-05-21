CREATE SCHEMA todolist;

CREATE TABLE todolist.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) > 3),
    phone VARCHAR(15) CHECK (
        phone ~ '^\+?[1-9]\d{1,14}$'
        AND char_length(phone) BETWEEN 10 AND 15
    )
);

CREATE TABLE todolist.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(100) NOT NULL CHECK (char_length(title) > 0),
    description VARCHAR(1000) CHECK (char_length(description) BETWEEN 1 AND 1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    completed_at TIMESTAMPTZ,
    CHECK (
        (completed = FALSE AND completed_at IS NULL)
        OR (completed = TRUE AND completed_at IS NOT NULL AND completed_at >= created_at) 
    ),
    author_id INT NOT NULL REFERENCES todolist.users(id)
);