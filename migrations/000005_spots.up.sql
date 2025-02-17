CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS spots (
    spot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    park_id UUID NOT NULL,
    spot_number INT NOT NULL,
    is_open BOOLEAN,
    is_electro_charging BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE
);
