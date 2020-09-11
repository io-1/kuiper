CREATE DATABASE devices;
use devices;

/* FIXME: add id and change device_id to mac
   FIXME: add unique index for mac
   FIXME: set primary key to mac
   FIXME: update devices service to use mac instead of device_id */
CREATE TABLE bat_cave_device_settings(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, deep_sleep_delay INT NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

