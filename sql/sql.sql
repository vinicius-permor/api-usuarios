CREATE DATABASE IF NOT EXISTS users;
USE users;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password varchar(100) not null
)ENGINE=INNODB;

