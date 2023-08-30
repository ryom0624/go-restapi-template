
-- +migrate Up
CREATE TABLE `user_defined_record` (
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` varchar(255) NOT NULL,
    `item_name` varchar(255) NOT NULL,
    `unit_type` enum('1','2','3','4','5','6','7','8','9','10','99') NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +migrate Down
DROP TABLE `user_defined_record`;
