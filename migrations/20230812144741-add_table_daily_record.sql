
-- +migrate Up
CREATE TABLE `daily_record` (
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` varchar(255) NOT NULL,
    `weather` VARCHAR(255) NOT NULL,
    `memo` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +migrate Down

DROP TABLE `daily_record`;
