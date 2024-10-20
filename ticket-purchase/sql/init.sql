CREATE DATABASE IF NOT EXISTS eventim;

USE eventim;

CREATE TABLE IF NOT EXISTS tickets (
    id int NOT NULL AUTO_INCREMENT,
    available boolean NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO tickets (available) VALUES
    (TRUE), (TRUE), (TRUE), (TRUE), (TRUE),
    (TRUE), (TRUE), (TRUE), (TRUE), (TRUE),
    (TRUE), (TRUE), (TRUE), (TRUE), (TRUE),
    (TRUE), (TRUE), (TRUE), (TRUE), (TRUE);