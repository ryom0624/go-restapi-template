
-- +migrate Up
CREATE TABLE `user` (
    `id` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `birthday` bigint NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `first_name` varchar(255) NOT NULL,
    `last_name_kana` varchar(255) NOT NULL,
    `first_name_kana` varchar(255) NOT NULL,
    `sex` int NOT NULL,
    `prefecture` varchar(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +migrate Down

DROP TABLE `user`;