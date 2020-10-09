CREATE DATABASE devices;
use devices;

CREATE TABLE bat_cave_device_settings(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, deep_sleep_delay INT NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

