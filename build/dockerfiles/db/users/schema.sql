CREATE DATABASE users;
use users;

CREATE TABLE users(id VARCHAR(36) NOT NULL, username VARCHAR(50) NOT NULL, name VARCHAR(100) NOT NULL, email VARCHAR(100), password VARCHAR(100), created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id), UNIQUE KEY unique_username(username), UNIQUE KEY unique_email(email));
