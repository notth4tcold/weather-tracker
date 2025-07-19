CREATE TABLE weather (
    id SERIAL PRIMARY KEY,
    city TEXT NOT NULL,
    temperature FLOAT,
    timestamp TIMESTAMPTZ DEFAULT NOW()
);