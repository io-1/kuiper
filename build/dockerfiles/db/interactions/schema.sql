CREATE DATABASE interactions;
use interactions;

CREATE TABLE interactions(
    id VARCHAR(36) NOT NULL, 
    name VARCHAR(50) NOT NULL, 
    description VARCHAR(100) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE keypad_conditions(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,
    button_id INT NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE lamp_events(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    red int NOT NULL,
    green int NOT NULL,
    blue int NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE conditions_to_events(
    id VARCHAR(36) NOT NULL, 
    interaction_id VARCHAR(36) NOT NULL,
    condition_id VARCHAR(36) NOT NULL, 
    event_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    INDEX `idx_conditions_to_events_interaction`(interaction_id),
    INDEX `idx_conditions_to_events_condition`(condition_id),
    INDEX `idx_conditions_to_events_event`(event_id),
    CONSTRAINT `fk_conditions_to_events_interaction` FOREIGN KEY(interaction_id) REFERENCES interactions(id),
    CONSTRAINT `fk_conditions_to_events_condition` FOREIGN KEY(condition_id) REFERENCES keypad_conditions(id),
    CONSTRAINT `fk_conditions_to_events_event` FOREIGN KEY(event_id) REFERENCES lamp_events(id),
    PRIMARY KEY(id)
);

CREATE TABLE keypad_conditions_to_lamp_events (
    id VARCHAR(36) NOT NULL, 
    interaction_id VARCHAR(36) NOT NULL,
    condition_id VARCHAR(36) NOT NULL, 

    /* the event_type is the type of event to send up the right data to the front end */
    /* event_type VARCHAR(50) NOT NULL, */
    event_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    INDEX `idx_keypad_conditions_to_lamp_events_interaction`(interaction_id),
    INDEX `idx_keypad_conditions_to_lamp_events_condition`(condition_id),
    INDEX `idx_keypad_conditions_to_lamp_events_event`(event_id),
    CONSTRAINT `fk_keypad_conditions_to_lamp_events_interaction` FOREIGN KEY(interaction_id) REFERENCES interactions(id),
    CONSTRAINT `fk_keypad_conditions_to_lamp_events_condition` FOREIGN KEY(condition_id) REFERENCES keypad_conditions(id),
    CONSTRAINT `fk_keypad_conditions_to_lamp_events_event` FOREIGN KEY(event_id) REFERENCES lamp_events(id),
    PRIMARY KEY(id)
);

CREATE TABLE lamp_toggle_events(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE lamp_color_events(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,
    red int NOT NULL,
    green int NOT NULL,
    blue int NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

CREATE TABLE lamp_pulse_events(
    id VARCHAR(36) NOT NULL, 
    mac VARCHAR(12) NOT NULL,
    red int NOT NULL,
    green int NOT NULL,
    blue int NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, 
    PRIMARY KEY(id)
);

/* CREATE TABLE bh1750_state( */
/*     id VARCHAR(36) NOT NULL, */ 
/*     mac VARCHAR(12) NOT NULL, */ 
/*     intensity INT NOT NULL, */ 
/*     PRIMARY KEY (id), */ 
/*     UNIQUE KEY `unique_bh1750_state_mac`(mac) */
/* ); */

/* CREATE TABLE hdc1080_state( */
/*     id VARCHAR(36) NOT NULL, */ 
/*     mac VARCHAR(12) NOT NULL, */  
/*     temp FLOAT NOT NULL, */ 
/*     humidity FLOAT NOT NULL, */ 
/*     PRIMARY KEY (id), */ 
/*     UNIQUE KEY `unique_hdc1080_state_mac`(mac) */
/* ); */

/* CREATE TABLE stats_state( */
/*     id VARCHAR(36) NOT NULL, */ 
/*     mac VARCHAR(12) NOT NULL, */ 
/*     voltage FLOAT NOT NULL, */ 
/*     connection_time INT NOT NULL, */ 
/*     rssi INT NOT NULL, */ 
/*     PRIMARY KEY(id), */ 
/*     UNIQUE KEY `unique_stats_state_mac`(mac) */
/* ); */

/* CREATE TABLE scheduled_conditions( */
/*     id VARCHAR(36) NOT NULL, */ 
/*     created_at TIMESTAMP, */ 
/*     updated_at TIMESTAMP, */ 
/*     deleted_at TIMESTAMP, */ 
/*     PRIMARY KEY(id) */
/* ); */

/* CREATE TABLE bh1750_conditions( */
/*     id VARCHAR(36) NOT NULL, */ 
/*     interaction VARCHAR(36) NOT NULL, */
/*     mac VARCHAR(12) NOT NULL, */

/*     /1* add measurement_operator table that has a list of all the measurement_operations by id (int) *1/ */ 
/*     measurement_operator VARCHAR(25) NOT NULL, */
/*     intensity_value INT NOT NULL, */
/*     created_at TIMESTAMP, */ 
/*     updated_at TIMESTAMP, */ 
/*     deleted_at TIMESTAMP, */ 
/*     INDEX `idx_bh1750_conditions_interaction`(interaction), */
/*     CONSTRAINT `fk_bh1750_conditions_interaction` FOREIGN KEY(interaction) REFERENCES interactions(id), */
/*     PRIMARY KEY(id) */
/* ); */

/* CREATE TABLE measurement_conditions( */
/*     id VARCHAR(36) NOT NULL, */ 
/*     interaction VARCHAR(36) NOT NULL, */
/*     mac VARCHAR(12) NOT NULL, */

/*     /1* add sensor table that has a list of all the sensors by id (int) *1/ */
/*     sensor VARCHAR(25) NOT NULL, */

/*     /1* add measurement table that has a list of all the measurements by id (int) *1/ */
/*     measurement VARCHAR(25) NOT NULL, */

/*     /1* add measurement_operator table that has a list of all the measurement_operations by id (int) *1/ */ 
/*     measurement_operator VARCHAR(25) NOT NULL, */
/*     measurement_value INT NOT NULL, */
/*     created_at TIMESTAMP, */ 
/*     updated_at TIMESTAMP, */ 
/*     deleted_at TIMESTAMP, */ 
/*     INDEX `idx_measurement_conditions_interaction`(interaction), */
/*     CONSTRAINT `fk_measurement_conditions_interaction` FOREIGN KEY(interaction) REFERENCES interactions(id), */
/*     PRIMARY KEY(id) */
/* ); */

