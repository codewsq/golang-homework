create database blog_db;



CREATE TABLE `users`
(
    `id`         bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NULL,
    `updated_at` datetime(3) NULL,
    `deleted_at` datetime(3) NULL,
    `username`   longtext    NOT NULL,
    `email`      longtext    NOT NULL,
    `password`   longtext    NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_users_deleted_at` (`deleted_at`),
    UNIQUE INDEX `idx_users_username` (`username`),
    UNIQUE INDEX `idx_users_email` (`email`)
)