CREATE TABLE IF NOT EXISTS campaigns (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some sample data
INSERT INTO campaigns (name, description) VALUES
    ('Social Media', 'Let us unite and expand on social media'),
    ('Software architecture', 'Advanced software architecture with residuality'),
    ('AE turns 25', 'Something to celebrate: AE turns 25 today');