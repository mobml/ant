CREATE TABLE IF NOT EXISTS goals (
    id BLOB PRIMARY KEY NOT NULL,
    area_id BLOB NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (area_id) REFERENCES areas(id)
);
