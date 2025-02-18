CREATE DATABASE IF NOT EXISTS socialapp;

USE socialapp;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nickname varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created timestamp default current_timestamp()
) ENGINE=INNODB