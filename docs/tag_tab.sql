CREATE DATABASE `unsplash_db`;

CREATE TABLE `unsplash_db`.`tag_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(32) DEFAULT '',
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`)
);