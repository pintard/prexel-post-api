-- @block
-- Create Database
CREATE DATABASE prexelpostdb;
-- @block
-- Create Table
CREATE TABLE prexelposts (
    uuid SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    contact VARCHAR(255),
    code TEXT NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create Indexes
CREATE INDEX idx_prexelposts_uuid ON prexelposts(uuid);
CREATE INDEX idx_prexelposts_username ON prexelposts(username);
CREATE INDEX idx_prexelposts_date ON prexelposts(date);
-- @block
-- Create User: psql - U postgres - d prexelpostdb
CREATE USER prexeluser WITH PASSWORD 'password';
-- @block
-- Grant Permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO prexeluser;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL ON TABLES TO prexeluser;