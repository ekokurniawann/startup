CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    occupation VARCHAR(255),
    email VARCHAR(255),
    password_hash VARCHAR(255),
    avatar_file_name VARCHAR(255),
    role VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
