CREATE DATABASE `unsplash_db`;

CREATE TABLE `unsplash_db`.`user_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(32) DEFAULT '',
    `email_address` varchar(32) DEFAULT '',
    `hashed_password` blob,
    `salt` varchar(32) DEFAULT '',
    `last_name` varchar(32) DEFAULT '',
    `first_name` varchar(32) DEFAULT '',
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`,`username`, `email_address`),
    INDEX (username)
);