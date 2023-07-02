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

CREATE TABLE patients
(
    id          INT         NOT NULL PRIMARY KEY auto_increment,
    name          VARCHAR(50) NOT NULL,
    surname       VARCHAR(50) NOT NULL,
    address       VARCHAR(50) NOT NULL,
    dni           VARCHAR(50) NOT NULL,
    dischargeDate VARCHAR(50) NOT NULL
);

CREATE TABLE shifts
(
    id INT         NOT NULL PRIMARY KEY auto_increment,
    data VARCHAR(50) NOT NULL,
    time VARCHAR(50) NOT NULL,
    dentist_id int not null,
    patient_id int not null

);

ALTER TABLE shifts ADD CONSTRAINT FK_SHIFT_DENTIST
    FOREIGN KEY (dentist_id) REFERENCES dentists(id);

ALTER TABLE shifts ADD CONSTRAINT FK_SHIFT_PATIENT
    FOREIGN KEY (patient_id) REFERENCES patients(id);


INSERT INTO dentists (name, surname, enrollment)
VALUES ('Zoe', 'Tira', '123');

INSERT INTO dentists (name, surname, enrollment)
VALUES ('Gregorio', 'Garcia', '456');


INSERT INTO patients (name, surname, address, dni, dischargeDate)
VALUES ('Zaira', 'Tira', 'Pedro Zenteno', '4041', '12-12');

INSERT INTO patients (name, surname, address, dni, dischargeDate)
VALUES ('Zora', 'Tira', 'Pedro Zenteno', '4642', '15-11');

INSERT INTO shifts (data, time, dentist_id, patient_id)
VALUES ('1-1', '12:22', 1, 1);

INSERT INTO shifts (data, time, dentist_id, patient_id)
VALUES ('3-5', '21:15', 2, 2);