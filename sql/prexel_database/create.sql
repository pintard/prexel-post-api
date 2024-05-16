DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'prexel_user') THEN
        CREATE USER prexel_user WITH PASSWORD 'securepassword';
    END IF;
END $$;

DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'prexeldb') THEN
        CREATE DATABASE prexeldb;
    END IF;
END $$;