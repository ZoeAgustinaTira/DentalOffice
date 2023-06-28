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

INSERT INTO dentists (name, surname, enrollment)
VALUES ('zoe', 'tira', '123');

INSERT INTO dentists (name, surname, enrollment)
VALUES ('grego', 'garcia', '456');