-- Create URLs table
CREATE TABLE IF NOT EXISTS urls (
                                    id BIGSERIAL PRIMARY KEY,
                                    short_code VARCHAR(10) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    user_id BIGINT,
    title VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE,
                             is_active BOOLEAN DEFAULT TRUE,
                             is_custom BOOLEAN DEFAULT FALSE,
                             password_hash VARCHAR(255),
    click_count BIGINT DEFAULT 0
    );

CREATE INDEX idx_urls_short_code ON urls(short_code);
CREATE INDEX idx_urls_user_id ON urls(user_id);
CREATE INDEX idx_urls_created_at ON urls(created_at);
CREATE INDEX idx_urls_expires_at ON urls(expires_at);

-- Create Users table
CREATE TABLE IF NOT EXISTS users (
                                     id BIGSERIAL PRIMARY KEY,
                                     email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    api_key VARCHAR(64) UNIQUE,
    is_premium BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_login_at TIMESTAMP WITH TIME ZONE,
                             daily_url_count INTEGER DEFAULT 0,
                             daily_click_count INTEGER DEFAULT 0,
                             last_reset_date DATE DEFAULT CURRENT_DATE
                             );

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_api_key ON users(api_key);

-- Add foreign key constraint
ALTER TABLE urls ADD CONSTRAINT fk_urls_user_id
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
