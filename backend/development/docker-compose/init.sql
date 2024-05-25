CREATE DATABASE IF NOT EXISTS testDB;
USE testDB;
CREATE TABLE IF NOT EXISTS phone_numbers (
    msisdn VARCHAR(100) PRIMARY KEY,
    grade VARCHAR(100),
    type VARCHAR(100),
    reserved_at DATETIME,
    expires_at DATETIME
);
INSERT INTO phone_numbers(msisdn, grade, type, reserved_at, expires_at) values ("+33652412359", "0", "CELL", null, null);
INSERT INTO phone_numbers(msisdn, grade, type, reserved_at, expires_at) values ("+33734254637", "3", "CELL", CURRENT_DATE, null);
INSERT INTO phone_numbers(msisdn, grade, type, reserved_at, expires_at) values ("0652412359", "1", "CELL", CURRENT_DATE, CURRENT_DATE);
INSERT INTO phone_numbers(msisdn, grade, type, reserved_at, expires_at) values ("0152412359", "2", "FIX", CURRENT_DATE, CURRENT_DATE);