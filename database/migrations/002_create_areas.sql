CREATE TABLE IF NOT EXISTS areas (
    id BLOB PRIMARY KEY NOT NULL,
    plan_id BLOB NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (plan_id) REFERENCES plans(id)
);
