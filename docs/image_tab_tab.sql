CREATE DATABASE `unsplash_db`;

CREATE TABLE `unsplash_db`.`image_tag_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `image_id` bigint(20) unsigned NOT NULL,
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`),
    INDEX (`image_id`)
);