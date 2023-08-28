-- @block
-- Create Table
CREATE TABLE prexelposts (
    id SERIAL PRIMARY KEY,
    user_id INT8 NOT NULL REFERENCES prexelusers(id),
    code TEXT NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create Indexes
CREATE INDEX idx_prexelposts_id ON prexelposts(id);
CREATE INDEX idx_prexelposts_date ON prexelposts(date);