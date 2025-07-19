-- Create Analytics table
CREATE TABLE IF NOT EXISTS analytics (
                                         id BIGSERIAL PRIMARY KEY,
                                         url_id BIGINT REFERENCES urls(id) ON DELETE CASCADE,
    clicked_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ip_address INET,
    user_agent TEXT,
    referer TEXT,
    country VARCHAR(2),
    city VARCHAR(100),
    browser VARCHAR(50),
    os VARCHAR(50),
    device_type VARCHAR(20),
    utm_source VARCHAR(100),
    utm_medium VARCHAR(100),
    utm_campaign VARCHAR(100)
    );

CREATE INDEX idx_analytics_url_id ON analytics(url_id);
CREATE INDEX idx_analytics_clicked_at ON analytics(clicked_at);
CREATE INDEX idx_analytics_country ON analytics(country);
CREATE INDEX idx_analytics_device_type ON analytics(device_type);

-- Create URL Tags table
CREATE TABLE IF NOT EXISTS url_tags (
                                        id BIGSERIAL PRIMARY KEY,
                                        url_id BIGINT REFERENCES urls(id) ON DELETE CASCADE,
    tag VARCHAR(50) NOT NULL
    );

CREATE INDEX idx_url_tags_url_id ON url_tags(url_id);
CREATE INDEX idx_url_tags_tag ON url_tags(tag);
CREATE UNIQUE INDEX idx_url_tags_unique ON url_tags(url_id, tag);

-- Create Rate Limits table
CREATE TABLE IF NOT EXISTS rate_limits (
                                           id BIGSERIAL PRIMARY KEY,
                                           identifier VARCHAR(255) NOT NULL,
    endpoint VARCHAR(100) NOT NULL,
    count INTEGER DEFAULT 1,
    window_start TIMESTAMP WITH TIME ZONE DEFAULT NOW()
    );

CREATE INDEX idx_rate_limits_identifier_endpoint ON rate_limits(identifier, endpoint);
CREATE INDEX idx_rate_limits_window_start ON rate_limits(window_start);
