\c tele2_dbase
SET ROLE tele2;

CREATE TABLE user_group (
    userID VARCHAR(36),
    groupID VARCHAR(36),
    PRIMARY KEY (userID,groupID) 
)