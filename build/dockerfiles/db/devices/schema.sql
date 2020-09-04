CREATE DATABASE devices;
use devices;

CREATE TABLE bat_cave_device_settings(device_id VARCHAR(12) NOT NULL, deep_sleep_delay INT NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (device_id));

/* CREATE DATABASE users; */
/* use users; */

/* CREATE TABLE users(username VARCHAR(50) NOT NULL, name VARCHAR(100) NOT NULL, email VARCHAR(100), password VARCHAR(100), created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (device_id)); */
