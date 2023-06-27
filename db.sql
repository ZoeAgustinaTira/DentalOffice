DROP DATABASE IF EXISTS dentaloffice;
CREATE DATABASE dentaloffice;
USE dentaloffice;

CREATE TABLE dentists(
    `id`        INT NOT NULL PRIMARY KEY auto_increment;
    name        VARCHAR(50) NOT NULL;
    surname     VARCHAR(50) NOT NULL;
    enrollment  VARCHAR(50) NOT NULL;
)