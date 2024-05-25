CREATE TABLE Rankings (
    user_id VARCHAR(255) PRIMARY KEY,
    like_count INT NOT NULL DEFAULT 0,
    dislike_count INT NOT NULL DEFAULT 0,
    attractiveness_score DECIMAL(5,2) NOT NULL DEFAULT 0.00,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);