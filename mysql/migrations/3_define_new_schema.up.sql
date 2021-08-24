CREATE TABLE IF NOT EXISTS `gacha` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `character_emissions`(
  `gacha_id` int NOT NULL,
  `character_id` int NOT NULL,
  `emission_weight` int NOT NULL,
  FOREIGN KEY (`gacha_id`) REFERENCES `gacha`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`character_id`) REFERENCES `character`(`id`) ON DELETE CASCADE
)