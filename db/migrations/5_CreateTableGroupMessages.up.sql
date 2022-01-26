\c tele2_dbase
SET ROLE tele2;

CREATE TABLE group_messages (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,  
    deleted_at DATETIME,  
    text VARCHAR(1000) NOT NULL,
    user_sender_id VARCHAR(36),
    group_reciver_id VARCHAR(36),
    FOREIGN KEY (user_sender_id) REFERENCES user (id),
    FOREIGN KEY (group_reciver_id) REFERENCES group (id),
);