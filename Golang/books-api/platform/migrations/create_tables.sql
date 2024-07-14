-- Purpose: This file is used to create the database schema and seed the database with initial data.

-- create extension for uuid generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- set timezone to UTC
SET TIMEZONE = 'UTC';

-- Create users table
CREATE TABLE users
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name  VARCHAR(255)        NOT NULL,
    username   VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE authors
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    title      VARCHAR(255),
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);


-- Create book_attrs table
CREATE TABLE book_attrs
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    publisher   VARCHAR(255) NOT NULL,
    pages       INTEGER      NOT NULL,
    description TEXT,
    status      VARCHAR(50),
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

-- Create books table
CREATE TABLE books
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title         VARCHAR(255) NOT NULL,
    user_id       UUID REFERENCES users (id),
    author_id     UUID REFERENCES authors (id),
    book_attrs_id UUID REFERENCES book_attrs (id),
    created_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);


-- Example data insertion for users
INSERT INTO users (username, password, full_name)
VALUES ('mohrifkan', '$2a$12$Yzl3Oaj4T9//KPBYK/5mUelIDz5EFPKFtIqCjntLzcQT6ch7lCF/S', 'Moh Rifkan');

SELECT *
FROM users;

-- Example data insertion for authors
INSERT INTO authors (name, title)
VALUES ('Tan Malaka', 'Pahlawan Indonesia'),
       ('Oh Su Hyang', 'Dosen & Pakar Komunikasi Korea Selatan');

SELECT *
FROM authors;