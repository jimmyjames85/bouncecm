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
    (2, 'Hadar', 'Ziv', 'hadarziv@sg.com', 'admin', '$2a$04$itSZdW4vcFm6q8OTBEJpO.5YIfAaoz7E.PnXN/WODM4Foq35Cc0Qq', NOW()),
    (3, 'Dana', 'Tran', 'danat@sg.com', 'admin', '$2a$04$PRK5zEW7ledmGkTZDfjkvO03BlMqLE1o0VPpsOJnJv44Qq4J6ruA6', NOW()),
    (4, 'Alfred', 'Lucero', 'alfredl@sg.com', 'admin', '$2a$04$qGL5V4TGTNDRbbXKHX9.W.OCD5pWyFuOUFMDW8OUB4CD2aH5kygQe', NOW()),
    (5, 'Dustin', 'Guerrero', 'dusting@sg.com', 'admin', '$2a$04$K3V5AIsCh46tq9kenOaixumeRHa10QR8Dpkb.r54ip6qFTovOnQ1K', NOW()),
    (6, 'Vinh', 'Lam', 'vinhl@sg.com', 'admin', '$2a$04$tiOmNTvo/5lGekgAO36f4e1V.jAxuRnl5jgsAAUt6qUKbdMyvXhF2', NOW());

    UNLOCK TABLES;