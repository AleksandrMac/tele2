\c tele2_dbase
SET ROLE tele2;

CREATE TABLE group (
    id VARCHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,  
    deleted_at DATETIME,  
    name VARCHAR(200) UNIQUE NOT NULL,
    description VARCHAR(500) NOT NULL
);