USE drop_rules;

DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `hash` char(60) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `user` WRITE;

INSERT INTO `user` (`id`, `first_name`, `last_name`, `email`, `role`, `hash`, `created_at`) 
VALUES 
    (1, 'Jim', 'John', 'jimjohn@jimjohn.com', 'admin', '$2a$04$AoMothWIRl.kQnGQz4jaBeN0y5PEGjRhkyaZgvPx/FJPXOkcrAKny', NOW()),
    (2, 'Hadar', 'Ziv', 'hadarziv@sg.com', 'admin', '$2a$04$itSZdW4vcFm6q8OTBEJpO.5YIfAaoz7E.PnXN/WODM4Foq35Cc0Qq', NOW());

    UNLOCK TABLES;