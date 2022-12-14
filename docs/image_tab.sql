CREATE DATABASE `unsplash_db`;

CREATE TABLE `unsplash_db`.`image_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `url` LONGTEXT,
    `description` LONGTEXT,
    `likes` bigint(10) DEFAULT 0,
    `downloads` bigint(10) DEFAULT 0,
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    INDEX (`created_at`),
    INDEX (`user_id`, `created_at`),
    INDEX (`id`)
);