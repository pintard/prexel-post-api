-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE prexeldb TO prexel_user;

-- Connect to the prexeldb database
\c prexeldb;

-- Grant privileges on all tables in the public schema
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO prexel_user;

-- Grant privileges on all sequences in the public schema
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO prexel_user;

-- Grant privileges on all functions in the public schema
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO prexel_user;