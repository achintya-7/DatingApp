CREATE TABLE Rankings (
    user_id INT PRIMARY KEY,
    like_count INT,
    dislike_count INT,
    attractiveness_score DECIMAL(5,2),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);