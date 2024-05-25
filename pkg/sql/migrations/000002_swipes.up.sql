CREATE TABLE Swipes (
    swipe_id INT AUTO_INCREMENT PRIMARY KEY,
    swiper_id INT,
    swipee_id INT,
    swipe_type ENUM('YES', 'NO'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (swiper_id) REFERENCES Users(user_id),
    FOREIGN KEY (swipee_id) REFERENCES Users(user_id)
);