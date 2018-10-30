
-- USER AUTHENTICATION SCHEMA

CREATE TABLE IF NOT EXISTS `authentication` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` smallint(5) unsigned NOT NULL,
  `role` varchar(255) NOT NULL,
  'salt' varchar(255) NOT NULL
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `user` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- OR 


CREATE TABLE IF NOT EXISTS `user` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  'salt' varchar(255) NOT NULL,
  `created_at` datetime NOT NULL
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- CHANGELOG SCHEMA


-- use MyISAM table for logging?
-- potential column that references previous change made?

CREATE TABLE IF NOT EXISTS `log` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `rule_id` smallint(5) unsigned NOT NULL,
  `user_id` smallint(5) unsigned NOT NULL,
  `change_made` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL
  PRIMARY KEY (`id`),
  FOREIGN KEY (`rule_id`) REFERENCES bounce_rule(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;