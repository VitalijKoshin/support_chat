CREATE TABLE IF NOT EXISTS `chat` (
    `chat_id` varchar(255) PRIMARY KEY,
    `chat_name` VARCHAR(255) NOT NULL,
    `user_id_support` varchar(255),
    `user_id_client` varchar(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS `message` (
    `message_id` varchar(255) PRIMARY KEY,
    `chat_id` varchar(255),
    `user_id` varchar(255),
    `message` TEXT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL
);

ALTER TABLE `chat`
    ADD CONSTRAINT `fk_chat_user_id_support` FOREIGN KEY (`user_id_support`) REFERENCES `user` (`user_id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_chat_user_id_client` FOREIGN KEY (`user_id_client`) REFERENCES `user` (`user_id`) ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE `message`
    ADD CONSTRAINT `fk_message_chat_id` FOREIGN KEY (`chat_id`) REFERENCES `chat` (`chat_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_message_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;

