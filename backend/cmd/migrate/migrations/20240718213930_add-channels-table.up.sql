CREATE TABLE IF NOT EXISTS channels (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `createdBy` INT UNSIGNED NOT NULL,
  `imgLink` VARCHAR(255),
  
  PRIMARY KEY (id),
  FOREIGN KEY (createdBy) REFERENCES users(id) ON DELETE CASCADE,
  
  INDEX idx_createdBy (createdBy)
);