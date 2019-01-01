USE drop_rules;

DROP TABLE IF EXISTS `changelog`;

CREATE TABLE IF NOT EXISTS `changelog` (
  `rule_id` smallint(5) unsigned NOT NULL,
  `user_id` smallint(5) unsigned NOT NULL,
  `comment` varchar(255) NOT NULL,
  `created_at` int(11) NOT NULL,
  `response_code` smallint(5) unsigned NOT NULL DEFAULT '0',
  `enhanced_code` varchar(16) NOT NULL DEFAULT '',
  `regex` varchar(255) NOT NULL DEFAULT '',
  `priority` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `description` varchar(255) DEFAULT NULL,
  `bounce_action` varchar(255) NOT NULL,
  PRIMARY KEY (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


LOCK TABLES `changelog` WRITE;

INSERT INTO `changelog` (`rule_id`, `user_id`, `comment`, `created_at`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`) 
VALUES 
	(300,1,'Testing changelog 1','1542594709',552 ,'4.2.2','User has full mailbox',0,'message will not succeed when retried, but address is likely valid','no_action'),
	(301,2,'Testing changelog 2','1542594710',552 ,'5.2.1','HVU:B1',0,'AOL code indicating content in email header/hostname is generating excessive complaints','no_action');

UNLOCK TABLES;