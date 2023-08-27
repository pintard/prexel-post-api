-- @block
-- Create Table
CREATE TABLE prexelposts (
    id SERIAL PRIMARY KEY,
    -- username VARCHAR(255) NOT NULL,
    -- contact VARCHAR(255),
    -- contact_url VARCHAR(1000),
    code TEXT NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create Indexes
CREATE INDEX idx_prexelposts_id ON prexelposts(id);
-- CREATE INDEX idx_prexelposts_username ON prexelposts(username);
CREATE INDEX idx_prexelposts_date ON prexelposts(date);