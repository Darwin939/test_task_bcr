-- Filename: 001_create_tables.sql

CREATE TABLE recipes (
    id          UUID PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    ingredients TEXT[] NOT NULL,
    steps       TEXT[] NOT NULL
);
