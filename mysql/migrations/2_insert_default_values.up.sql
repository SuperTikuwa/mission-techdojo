INSERT INTO `users` (`name`,`token`)
VALUES ('Alice','2d4a6af7cd5aefea4edb8d867614702b2bf137f921a397cd904e91aa08954ea9');

INSERT INTO `characters` (`name`,`weight`)
VALUES ('Josuke',1),('Koich',1),('Jotaro',2),('Joseph',2),('Okuyasu',3);

INSERT INTO `user_owned_characters` (`user_id`,`character_id`,`user_character_id`)
VALUES
(1,1,'1-1-1'),
(1,2,'1-2-1');