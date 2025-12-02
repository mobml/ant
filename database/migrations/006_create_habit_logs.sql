CREATE TABLE IF NOT EXISTS habit_logs (
    id BLOB PRIMARY KEY NOT NULL,
    habit_id BLOB NOT NULL,
    log_date TIMESTAMP NOT NULL,
    value NUMERIC NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (habit_id) REFERENCES habits(id)
);

CREATE INDEX IF NOT EXISTS idx_habit_logs_habit_date
    ON habit_logs (habit_id, log_date);
