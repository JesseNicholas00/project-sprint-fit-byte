CREATE TABLE users
(
    user_id     UUID PRIMARY KEY,
    name        TEXT,
    email       TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL,
    preference  TEXT,
    weight_unit TEXT,
    height_unit TEXT,
    weight      INTEGER DEFAULT 0,
    height      INTEGER DEFAULT 0,
    image_uri   TEXT
);

CREATE INDEX IF NOT EXISTS users_email_hash_idx ON users USING HASH (email);