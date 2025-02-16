-- Your SQL goes here

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS spots (
    spot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    park_id UUID NOT NULL,
    spot_number INT NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('available', 'occupied', 'reserved')),
    is_electro_charging BOOLEAN NOT NULL DEFAULT FALSE
);
