CREATE TABLE IF NOT EXISTS `user_session` (
    `user_session_id` varchar(255) PRIMARY KEY,
    `user_id` varchar(255),
    `session_id` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL
);