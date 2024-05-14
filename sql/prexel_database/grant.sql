-- @block
-- Create User: psql -U postgres -d prexelpostdb
CREATE USER prexeluser WITH PASSWORD 'password';
-- @block
-- Grant Permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO prexeluser;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL ON TABLES TO prexeluser;