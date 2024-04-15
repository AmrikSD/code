-- Migration number: 0001 	 2024-04-14T01:30:36.518Z

DROP TABLE IF EXISTS rsvp;

CREATE TABLE rsvp(
    rsvp_id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    plus_one INTEGER NOT NULL,
    plus_one_first_name TEXT,
    plus_one_last_name TEXT,
    coming_to_thailand INTEGER NOT NULL,
    coming_to_uk INTEGER NOT NULL,
    diet TEXT
);

