-- Create the prexelposts table
CREATE TABLE prexelposts (
    uuid      SERIAL PRIMARY KEY,
    username  VARCHAR(255) NOT NULL,
    contact   VARCHAR(255),
    code      TEXT NOT NULL,
    date      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes:
CREATE INDEX idx_prexelposts_uuid ON prexelposts(uuid);
CREATE INDEX idx_prexelposts_username ON prexelposts(username);
CREATE INDEX idx_prexelposts_date ON prexelposts(date);