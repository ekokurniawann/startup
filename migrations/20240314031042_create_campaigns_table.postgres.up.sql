CREATE TABLE campaigns (
    id SERIAL PRIMARY KEY,
    user_id INT,
    name VARCHAR(255),
    short_description VARCHAR(255),
    description TEXT,
    perks TEXT,
    backer_count INT,
    goal_amount INT,
    current_amount INT,
    slug VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
