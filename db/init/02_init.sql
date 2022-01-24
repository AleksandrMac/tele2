\c tele2_dbase
SET ROLE tele2;

DROP TABLE IF EXISTS user CASCADE;
CREATE TABLE user (
    id VARCHAR(36) PRIMARY KEY,
    created_at DATE NOT NULL,
    updated_at DATE,    
    name VARCHAR(200) UNIQUE NOT NULL,
    email VARCHAR(200) NOT NULL
);

DROP TABLE IF EXISTS group CASCADE;
CREATE TABLE group (
    id VARCHAR(36) PRIMARY KEY,
    created_at DATE NOT NULL,
    updated_at DATE,    
    name VARCHAR(200) UNIQUE NOT NULL,
    description VARCHAR(500) NOT NULL
);

DROP TABLE IF EXISTS private_messages CASCADE;
CREATE TABLE private_messages (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at DATE NOT NULL,
    updated_at DATE, 
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    text VARCHAR(1000) NOT NULL,
    user_sender_id VARCHAR(36),
    user_reciver_id VARCHAR(36),
    FOREIGN KEY (user_sender_id) REFERENCES user (id),
    FOREIGN KEY (user_reciver_id) REFERENCES user (id),
);

DROP TABLE IF EXISTS group_messages CASCADE;
CREATE TABLE group_messages (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at DATE NOT NULL,
    updated_at DATE, 
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    text VARCHAR(1000) NOT NULL,
    user_sender_id VARCHAR(36),
    group_reciver_id VARCHAR(36),
    FOREIGN KEY (user_sender_id) REFERENCES user (id),
    FOREIGN KEY (group_reciver_id) REFERENCES group (id),
);
