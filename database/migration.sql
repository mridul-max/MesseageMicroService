CREATE TABLE `message`
(
    messageID   bigint auto_increment,
    content varchar(255) NOT NULL,
    PRIMARY KEY (`messageID`)
);

INSERT INTO `message` (`content`)
VALUES ('I will be absent tomorrow'),
       ('I will handout the assignment in upcoming days');