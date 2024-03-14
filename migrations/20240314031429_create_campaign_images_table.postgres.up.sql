CREATE TABLE campaign_images (
    id SERIAL PRIMARY KEY,
    campaign_id INT,
    file_name VARCHAR(255),
    is_primary SMALLINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
);
