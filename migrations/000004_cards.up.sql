CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS cards (
    card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    card_number VARCHAR(16) UNIQUE NOT NULL,
    card_holder_name VARCHAR(128) NOT NULL,
    expiration_date DATE NOT NULL,
    balance DOUBLE PRECISION NOT NULL DEFAULT 0
);
