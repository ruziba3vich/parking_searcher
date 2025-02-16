-- Your SQL goes here

CREATE TABLE IF NOT EXISTS spots (
    spot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    park_id UUID NOT NULL,
    spot_number INT NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('available', 'occupied', 'reserved')),
    is_electro_charging BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (park_id) REFERENCES parks(park_id) ON DELETE CASCADE
);
