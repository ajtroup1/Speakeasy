CREATE TABLE IF NOT EXISTS messages (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` TEXT NOT NULL,
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `createdBy` INT UNSIGNED NOT NULL,
  `createdIn` INT UNSIGNED NOT NULL,
  
  PRIMARY KEY (id),
  FOREIGN KEY (createdBy) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (createdIn) REFERENCES channels(id) ON DELETE CASCADE,
  
  INDEX idx_createdBy (createdBy),
  INDEX idx_createdIn (createdIn)
);