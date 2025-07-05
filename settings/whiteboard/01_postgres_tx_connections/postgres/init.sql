-- Initialize database schema for transaction connection demo

-- Create a simple table to demonstrate transactions
CREATE TABLE IF NOT EXISTS demo_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    value INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some initial data
INSERT INTO demo_table (name, value) VALUES 
    ('initial_record_1', 100),
    ('initial_record_2', 200),
    ('initial_record_3', 300);

-- Create an index for better performance
CREATE INDEX IF NOT EXISTS idx_demo_table_name ON demo_table(name);

-- Enable pg_stat_statements extension if not already enabled
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;