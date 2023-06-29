DROP
DATABASE IF EXISTS dentaloffice;
CREATE
DATABASE dentaloffice;
USE
dentaloffice;

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
    dischargeDate VARCHAR(50) NOT NULL,
);

CREATE TABLE shifts
(
    `id` INT         NOT NULL PRIMARY KEY auto_increment,
    data VARCHAR(50) NOT NULL,
    time VARCHAR(50) NOT NULL,
);

ALTER TABLE shifts ADD CONSTRAINT FK_SHIFT_DENTIST
FOREIGN KEY (dentist_id) REFERENCES dentists(id)

ALTER TABLE shifts ADD CONSTRAINT FK_SHIFT_PATIENT
FOREIGN KEY (patient_id) REFERENCES patients(id)

INSERT INTO dentists (name, surname, enrollment)
VALUES ('zoe', 'tira', '123');

INSERT INTO dentists (name, surname, enrollment)
VALUES ('grego', 'garcia', '456');