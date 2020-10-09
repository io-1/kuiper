CREATE DATABASE interactions;
use interactions;

CREATE TABLE bh1750_sensors(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, intensity INT NOT NULL, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

CREATE TABLE hdc1080_sensors(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, temp FLOAT NOT NULL, humidity FLOAT NOT NULL, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

CREATE TABLE stats_sensors(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, voltage FLOAT NOT NULL, connection_time INT NOT NULL, rssi INT NOT NULL, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

CREATE TABLE interactions(id VARCHAR(36) NOT NULL, name VARCHAR(100) NOT NULL, conditions VARCHAR(36) NOT NULL, actions VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));

CREATE TABLE scheduled_conditions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));

CREATE TABLE device_conditions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));

CREATE TABLE actions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));
