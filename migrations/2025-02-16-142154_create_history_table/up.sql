-- Your SQL goes here

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS history (
    history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    park_id UUID NOT NULL,
    spot_id UUID NOT NULL,
    entry_time TIMESTAMP NOT NULL DEFAULT NOW(),
    exit_time TIMESTAMP,
    total_cost DOUBLE PRECISION DEFAULT 0,
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'completed', 'canceled'))
);
