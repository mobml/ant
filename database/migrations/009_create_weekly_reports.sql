CREATE TABLE IF NOT EXISTS weekly_reports (
    id BLOB PRIMARY KEY NOT NULL,
    plan_id BLOB NOT NULL,
    week_start DATE NOT NULL,
    week_end DATE NOT NULL,
    report_md TEXT,
    generated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (plan_id) REFERENCES plans(id)
);
