DROP DATABASE go_pmb;
CREATE DATABASE go_pmb;
USE go_pmb;

CREATE TABLE IF NOT EXISTS `registrants` (
    `registrant_id` VARCHAR(255) PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `payment_url` VARCHAR(100) NOT NULL,
    `program` VARCHAR(10) NOT NULL,
    `bill` INT NOT NULL,
    `payment_status` VARCHAR(50) DEFAULT "pending",
    `created_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `last_updated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX `USERNAME_IDX` ON `registrants` (`username`);
CREATE INDEX `PASSWORD_IDX` ON `registrants` (`password`);
CREATE INDEX `PAYMENT_STATUS_IDX` ON `registrants` (`payment_status`);