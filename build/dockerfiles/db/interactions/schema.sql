CREATE DATABASE interactions;
use interactions;

CREATE TABLE bh1750_state(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, intensity INT NOT NULL, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

CREATE TABLE hdc1080_state(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, temp FLOAT NOT NULL, humidity FLOAT NOT NULL, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

CREATE TABLE stats_state(id VARCHAR(36) NOT NULL, mac VARCHAR(12) NOT NULL, voltage FLOAT NOT NULL, connection_time INT NOT NULL, rssi INT NOT NULL, PRIMARY KEY (id), UNIQUE KEY unique_mac(mac));

CREATE TABLE interactions(id VARCHAR(36) NOT NULL, name VARCHAR(100) NOT NULL, conditions VARCHAR(36) NOT NULL, actions VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));

CREATE TABLE scheduled_conditions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));

/* FIXME: each sensor or device needs to have a condition */
CREATE TABLE bh1750_conditions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));
CREATE TABLE hdc1080_conditions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));
CREATE TABLE stats_conditions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));

/* FIXME: each sensor or device needs to have a action */
CREATE TABLE actions(id VARCHAR(36) NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (id));
