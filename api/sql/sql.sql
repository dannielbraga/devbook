CREATE DATABASE  IF NOT EXISTS devbook; 

USE devbook;

DROP TABLE IF EXISTS usuarios; 

CREATE TABLE usuarios (
    id int auto_increment primary key, 
    nome varchar(200) not null, 
    nick varchar(50) not null unique, 
    email varchar(200) not null unique, 
    senha varchar(200) not null, 
    criadoem timestamp default current_timestamp()
) ENGINE=INNODB;