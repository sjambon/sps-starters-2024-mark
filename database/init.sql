CREATE TABLE IF NOT EXISTS campaigns (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some sample data
INSERT INTO campaigns (name, description) VALUES
    ('Summer Sale', 'Promotional campaign for summer products'),
    ('New Product Launch', 'Campaign to introduce our latest product'),
    ('Holiday Special', 'Special discounts for the holiday season');