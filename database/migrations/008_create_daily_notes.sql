CREATE TABLE IF NOT EXISTS daily_notes (
    id BLOB PRIMARY KEY NOT NULL,
    note_date TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
