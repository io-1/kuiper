CREATE DATABASE interactions;
use interactions;

CREATE TABLE bh1750_state(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL, 
    intensity INT NOT NULL, 
    PRIMARY KEY (id), 
    UNIQUE KEY `unique_bh1750_state_mac`(mac)
);

CREATE TABLE hdc1080_state(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,  
    temp FLOAT NOT NULL, 
    humidity FLOAT NOT NULL, 
    PRIMARY KEY (id), 
    UNIQUE KEY `unique_hdc1080_state_mac`(mac)
);

CREATE TABLE stats_state(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL, 
    voltage FLOAT NOT NULL, 
    connection_time INT NOT NULL, 
    rssi INT NOT NULL, 
    PRIMARY KEY(id), 
    UNIQUE KEY `unique_stats_state_mac`(mac)
);

CREATE TABLE interactions(
    id VARCHAR(36) NOT NULL, 
    name VARCHAR(100) NOT NULL, 
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE scheduled_conditions(
    id VARCHAR(36) NOT NULL, 
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE bh1750_conditions(
    id VARCHAR(36) NOT NULL, 
    interaction VARCHAR(36) NOT NULL,
    mac VARCHAR(12) NOT NULL,

    /* add measurement_operator table that has a list of all the measurement_operations by id (int) */ 
    measurement_operator VARCHAR(25) NOT NULL,
    intensity_value INT NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    INDEX `idx_bh1750_conditions_interaction`(interaction),
    CONSTRAINT `fk_bh1750_conditions_interaction` FOREIGN KEY(interaction) REFERENCES interactions(id),
    PRIMARY KEY(id)
);

CREATE TABLE measurement_conditions(
    id VARCHAR(36) NOT NULL, 
    interaction VARCHAR(36) NOT NULL,
    mac VARCHAR(12) NOT NULL,

    /* add sensor table that has a list of all the sensors by id (int) */
    sensor VARCHAR(25) NOT NULL,

    /* add measurement table that has a list of all the measurements by id (int) */
    measurement VARCHAR(25) NOT NULL,

    /* add measurement_operator table that has a list of all the measurement_operations by id (int) */ 
    measurement_operator VARCHAR(25) NOT NULL,
    measurement_value INT NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    INDEX `idx_measurement_conditions_interaction`(interaction),
    CONSTRAINT `fk_measurement_conditions_interaction` FOREIGN KEY(interaction) REFERENCES interactions(id),
    PRIMARY KEY(id)
);

CREATE TABLE actions(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,
    action VARCHAR(6) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);
