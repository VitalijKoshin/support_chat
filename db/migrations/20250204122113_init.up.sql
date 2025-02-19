
CREATE TABLE IF NOT EXISTS `user` (
    `user_id` varchar(255) PRIMARY KEY,
    `nickname` VARCHAR(255) NOT NULL,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `is_public` BOOLEAN NOT NULL,
    `user_role` VARCHAR(255) NOT NULL,
    `created_by` varchar(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    `login_date` TIMESTAMP NULL
);