-- histories Table
CREATE TABLE IF NOT EXISTS histories
(
    id         SERIAL PRIMARY KEY,
    user_id    TEXT NOT NULL,
    disease_id TEXT NOT NULL,
    note       TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
