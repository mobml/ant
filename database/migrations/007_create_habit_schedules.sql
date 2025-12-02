CREATE TABLE IF NOT EXISTS habit_schedules (
    id BLOB PRIMARY KEY NOT NULL,
    habit_id BLOB NOT NULL,
    day_of_week INTEGER NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (habit_id) REFERENCES habits(id)
);
