DROP
DATABASE IF EXISTS dentaloffice;
CREATE
DATABASE dentaloffice;
USE
dentaloffice;

CREATE TABLE dentists(
                         `id`        INT NOT NULL PRIMARY KEY auto_increment,
                         name        VARCHAR(50) NOT NULL,
                         surname     VARCHAR(50) NOT NULL,
                         enrollment  VARCHAR(50) NOT NULL
);


CREATE TABLE dentists
(
    `id`       INT         NOT NULL PRIMARY KEY auto_increment,
    name       VARCHAR(50) NOT NULL,
    surname    VARCHAR(50) NOT NULL,
    enrollment VARCHAR(50) NOT NULL
);

CREATE TABLE patients
(
    `id`          INT         NOT NULL PRIMARY KEY auto_increment,
    name          VARCHAR(50) NOT NULL,
    surname       VARCHAR(50) NOT NULL,
    address       VARCHAR(50) NOT NULL,
    dni           VARCHAR(50) NOT NULL,
    dischargeDate VARCHAR(50) NOT NULL
);

CREATE TABLE shifts
(
    `id` INT         NOT NULL PRIMARY KEY auto_increment,
    data VARCHAR(50) NOT NULL,
    time VARCHAR(50) NOT NULL
);

INSERT INTO dentists (name, surname, enrollment)
VALUES ('zoe', 'tira', '123');

INSERT INTO dentists (name, surname, enrollment)
VALUES ('grego', 'garcia', '456');