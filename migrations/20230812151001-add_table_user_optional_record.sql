
-- +migrate Up

CREATE TABLE `optional_record` (
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `daily_record_id` BIGINT NOT NULL,
    `user_defined_id` BIGINT NOT NULL,
    `value` DOUBLE NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

ALTER TABLE `optional_record` ADD INDEX `daily_record_id` (`daily_record_id`);
ALTER TABLE `optional_record` ADD INDEX `user_defined_id` (`user_defined_id`);

-- +migrate Down
DROP TABLE `optional_record`;

