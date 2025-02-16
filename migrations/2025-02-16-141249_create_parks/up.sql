-- Your SQL goes here

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS parks (
    park_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    park_name VARCHAR(64) NOT NULL,
    address VARCHAR(166) NOT NULL,
    price_ph DOUBLE PRECISION NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('available', 'full', 'closed')),
    available_spots_count INT NOT NULL,
    total_spots_count INT NOT NULL,
    electro_charging_available BOOLEAN NOT NULL DEFAULT FALSE,
    rating DOUBLE PRECISION DEFAULT 0,
    park_balance DOUBLE PRECISION DEFAULT 0,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
);
