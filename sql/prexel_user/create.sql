-- @block
-- Create Table
CREATE TABLE prexelusers (
    id SERIAL PRIMARY KEY,
    -- username VARCHAR(255) NOT NULL,
    -- contact VARCHAR(255),
    -- contact_url VARCHAR(1000),
    date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create Indexes
CREATE INDEX idx_prexelusers_id ON prexelusers(id);
CREATE INDEX idx_prexelusers_date ON prexelusers(date);