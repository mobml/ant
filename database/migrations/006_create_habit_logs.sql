CREATE TABLE IF NOT EXISTS habit_logs (
    id BLOB PRIMARY KEY NOT NULL,
    habit_id BLOB NOT NULL,
    log_date DATE NOT NULL,
    value NUMERIC NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
);
