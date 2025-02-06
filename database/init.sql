CREATE TABLE IF NOT EXISTS ping_status (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(15) NOT NULL,
    ping_time TIMESTAMP NOT NULL,
    success BOOLEAN NOT NULL,
    last_successful_ping TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_ip_ping
ON ping_status(ip_address);

CREATE INDEX IF NOT EXISTS idx_ip
ON ping_status(ip_address);