CREATE TABLE channel_users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    userID INT UNSIGNED NOT NULL,
    channelID INT UNSIGNED NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    userRole INT DEFAULT 0, -- 0=Base tier, 1=VIP tier, 2=Channel Admin
    FOREIGN KEY (userID) REFERENCES users(id),
    FOREIGN KEY (channelID) REFERENCES channels(id),
    INDEX (userID),
    INDEX (channelID)
);