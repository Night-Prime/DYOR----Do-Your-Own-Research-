-- Migration to update all users missing fields based on the current model

-- Add missing columns if they don't exist
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'avatar') THEN
        ALTER TABLE users ADD COLUMN avatar VARCHAR(255);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'email') THEN
        ALTER TABLE users ADD COLUMN email VARCHAR(100) UNIQUE;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'role') THEN
        ALTER TABLE users ADD COLUMN role VARCHAR(50) DEFAULT 'user';
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'password') THEN
        ALTER TABLE users ADD COLUMN password VARCHAR(255) NOT NULL;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'created_at') THEN
        ALTER TABLE users ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'updated_at') THEN
        ALTER TABLE users ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'deleted_at') THEN
        ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP;
    END IF;
END $$;

-- Update default values for existing rows
UPDATE users
SET role = COALESCE(role, 'user'),
    created_at = COALESCE(created_at, CURRENT_TIMESTAMP),
    updated_at = COALESCE(updated_at, CURRENT_TIMESTAMP);