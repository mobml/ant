CREATE TABLE IF NOT EXISTS plans (
    id BLOB PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    start_date TIMESTAMP NOT NULL,
    duration INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
