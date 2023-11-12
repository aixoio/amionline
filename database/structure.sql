-- MySQL

START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS `amionline`;

USE `amionline`;

CREATE TABLE IF NOT EXISTS `events` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `success` BOOLEAN NOT NULL,
    `time_ms` DOUBLE NOT NULL,
    `target_ip` VARCHAR(255) NOT NULL,
    `time_of_request` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

COMMIT;