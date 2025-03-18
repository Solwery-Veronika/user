CREATE TABLE IF NOT EXISTS participants (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    name TEXT,
    surname TEXT,
    birthday TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);