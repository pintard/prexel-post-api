-- @block
-- Create Table
CREATE TABLE prexel_users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    service VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    contact VARCHAR(255),
    contact_url VARCHAR(255),
    date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create Indexes
CREATE INDEX idx_prexelusers_id ON prexelusers(id);
CREATE INDEX idx_prexelusers_date ON prexelusers(date);