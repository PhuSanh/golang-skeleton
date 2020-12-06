-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint unsigned AUTO_INCREMENT NOT NULL,
    `username` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `role_id` tinyint,
    `is_active` bool NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_username` (`username`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +migrate Down
DROP TABLE IF EXISTS `users`;
