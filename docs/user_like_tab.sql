CREATE DATABASE `unsplash_db`;

CREATE TABLE `unsplash_db`.`user_like_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `image_id` bigint(20) unsigned NOT NULL,
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`),
    INDEX (`user_id`)
);