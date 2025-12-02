CREATE TABLE IF NOT EXISTS habits (
    id BLOB PRIMARY KEY NOT NULL,
    goal_id BLOB NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    measure_type measure_type NOT NULL,
    measure_unit TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (goal_id) REFERENCES goals(id)
);
