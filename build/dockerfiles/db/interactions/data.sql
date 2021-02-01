use interactions;

INSERT INTO interactions (id, name, description) VALUES ('e83fcdcf-6c94-4b9a-9001-6ba0380a814b', 'keypad test', 'keypad test');

INSERT INTO keypad_conditions (id, mac, button_id) VALUES ('4d55d25c-15cf-4737-9df4-21aa775127ad', '112233aabbcc', 1);

INSERT INTO lamp_pulse_events (id, mac, red, green, blue) VALUES ('3e5e1850-0e37-45c1-9785-94361e241750', 'aabbcc112233', 0, 0, 255);

INSERT INTO keypad_conditions_to_lamp_events (id, interaction_id, condition_id, event_id) VALUES ('', 'e83fcdcf-6c94-4b9a-9001-6ba0380a814b', '4d55d25c-15cf-4737-9df4-21aa775127ad', '3e5e1850-0e37-45c1-9785-94361e241750');
