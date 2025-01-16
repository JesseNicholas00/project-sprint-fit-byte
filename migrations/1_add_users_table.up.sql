CREATE TABLE users
(
    user_id     UUID PRIMARY KEY,
    user_name   TEXT,
    email       TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL,
    preference  TEXT,
    weight_unit TEXT,
    height_unit TEXT,
    weight      INTEGER,
    height      INTEGER,
    image_uri   TEXT
);