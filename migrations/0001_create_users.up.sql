CREATE TABLE IF NOT EXISTS users (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     email VARCHAR NOT NULL UNIQUE,
     login VARCHAR NOT NULL,
     password VARCHAR NOT NULL,
     role VARCHAR NOT NULL,
     telegram_id VARCHAR,
     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
     updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
