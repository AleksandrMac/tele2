CREATE USER tele2
WITH PASSWORD 'tele2';

CREATE DATABASE tele2_dbase
    WITH OWNER tele2
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';