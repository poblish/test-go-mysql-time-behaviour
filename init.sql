create database if not exists mydb;

USE mydb;

CREATE TABLE IF NOT EXISTS person (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(20) NOT NULL,
    `created_at` datetime(6) NOT NULL DEFAULT now(6),
    PRIMARY KEY (`id`)
);

insert into person values (null, 'Me', now(6));