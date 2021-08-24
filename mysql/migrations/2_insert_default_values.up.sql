INSERT INTO `users` (`name`,`token`)
VALUES ('Alice','2d4a6af7cd5aefea4edb8d867614702b2bf137f921a397cd904e91aa08954ea9');

INSERT INTO `characters` (`name`)
VALUES ('Jotaro'),('Joseph'),('Kakyoin'),('Polnareff'),('Avdol'),('Dio');

INSERT INTO `characters` (`name`)
VALUES ('Josuke'),('Koichi'),('Okuyasu'),('Rohan'),('Yukako'),('Yoshikage');

INSERT INTO `user_owned_characters` (`user_id`,`character_id`,`user_character_id`)
VALUES
(1,1,'1-1-202108240017121'),
(1,2,'1-2-202108240017122');