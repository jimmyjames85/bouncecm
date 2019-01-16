# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: rp.dbmainendpointp1mdw1.sendgrid.net (MySQL 5.6.38-83.0-log)
# Database: mail
# Generation Time: 2018-10-05 19:26:26 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE IF NOT EXISTS drop_rules;

USE drop_rules;

# Dump of table bounce_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bounce_rule`;

CREATE TABLE IF NOT EXISTS `bounce_rule` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `response_code` smallint(5) unsigned NOT NULL DEFAULT '0',
  `enhanced_code` varchar(16) NOT NULL DEFAULT '',
  `regex` varchar(255) NOT NULL DEFAULT '',
  `priority` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `description` varchar(255) DEFAULT NULL,
  `bounce_action` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bounce_rule_components` (`response_code`,`enhanced_code`,`regex`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `bounce_rule` WRITE;
/*!40000 ALTER TABLE `bounce_rule` DISABLE KEYS */;

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(173,421,'','',0,'RFC5321 Service not available','retry'),
	(174,450,'','',0,'RFC5321 Mailbox unavailable','retry'),
	(175,451,'','',0,'RFC5321 Local error in processing','retry'),
	(176,452,'','',0,'RFC5321 Insufficient system storage','retry'),
	(177,454,'','',0,'RFC3207 TLS not available','retry'),
	(178,455,'','',0,'RFC5321 Server unable to accomodate parameters','retry'),
	(179,500,'','',0,'RFC5321 Syntax error command unrecognized','no_action'),
	(180,501,'','',0,'RFC5321 Syntax error in parameters or arguments','no_action'),
	(181,502,'','',0,'RFC5321 Command not implemented','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(182,503,'','',0,'RFC5321 Bad sequence of commands','no_action'),
	(183,521,'','',0,'RFC7504 Server does not accept mail','suppress'),
	(184,530,'','',0,'RFC4954 Must issue a STARTTLS command first','no_action'),
	(185,550,'','',0,'RFC5321 Mailbox unavailable','suppress'),
	(186,551,'','',0,'RFC5321 User not local','suppress'),
	(187,552,'','',0,'RFC5321 Exceeded storage allocation','no_action'),
	(188,553,'','',0,'RFC5321 Mailbox name not allowed','suppress'),
	(189,554,'','',0,'RFC5321 Transaction failed','no_action'),
	(190,555,'','',0,'RFC5321 Parameters not recognized or not implemented','no_action'),
	(191,556,'','',0,'RFC7504 Recipient address has null MX','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(192,0,'5.1.1','',0,'RFC3463 Bad destination mailbox address','suppress'),
	(193,0,'5.1.2','',0,'RFC3463 Bad destination system address','suppress'),
	(194,0,'5.1.3','',0,'RFC3463 Bad destination mailbox address syntax','suppress'),
	(195,0,'5.1.4','',0,'RFC3463 Destination mailbox address ambiguous','no_action'),
	(196,0,'5.1.6','',0,'RFC3463 Destination mailbox has moved','suppress'),
	(197,0,'4.1.7','',0,'RFC3463 Bad senders mailbox address syntax','no_action'),
	(198,0,'5.1.7','',0,'RFC3463 Bad senders mailbox address syntax','no_action'),
	(199,0,'4.1.8','',0,'RFC3463 Bad senders system address','no_action'),
	(200,0,'5.1.8','',0,'RFC3463 Bad senders system address','no_action'),
	(201,0,'5.1.9','',0,'RFC3886 Message relayed to non-compliant mailer','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(202,0,'5.1.10','',0,'RFC7505 Recipient address has null MX','no_action'),
	(203,0,'4.2.0','',0,'RFC3463 Other or undefined mailbox status','retry'),
	(204,0,'5.2.0','',0,'RFC3463 Other or undefined mailbox status','no_action'),
	(205,0,'4.2.1','',0,'RFC3463 Mailbox disabled not accepting messages','retry'),
	(206,0,'5.2.1','',0,'RFC3463 Mailbox disabled not accepting messages','suppress'),
	(207,0,'4.2.2','',0,'RFC3463 Mailbox full','retry'),
	(208,0,'5.2.2','',0,'RFC3463 Mailbox full','no_action'),
	(209,0,'5.2.3','',0,'RFC3463 Message length exceeds administrative limit','no_action'),
	(210,0,'5.2.4','',0,'RFC3463 Mailing list expansion problem','no_action'),
	(211,0,'4.3.0','',0,'RFC3463 Other or undefined mail system status','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(212,0,'5.3.0','',0,'RFC3463 Other or undefined mail system status','no_action'),
	(213,0,'4.3.1','',0,'RFC3463 Mail system full','retry'),
	(214,0,'4.3.2','',0,'RFC3463 System not accepting network messages','retry'),
	(215,0,'5.3.2','',0,'RFC3463 System not accepting network messages','suppress'),
	(216,0,'4.3.3','',0,'RFC3463 System not capable of selected features','no_action'),
	(217,0,'5.3.3','',0,'RFC3463 System not capable of selected features','no_action'),
	(218,0,'5.3.4','',0,'RFC3463 Message too big for system','no_action'),
	(219,0,'4.3.5','',0,'RFC3463 System incorrectly configured','no_action'),
	(220,0,'5.3.5','',0,'RFC3463 System incorrectly configured','no_action'),
	(221,0,'4.4.0','',0,'RFC3463 Other or undefined network or routing status','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(222,0,'5.4.0','',0,'RFC3463 Other or undefined network or routing status','no_action'),
	(223,0,'5.4.1','',0,'RFC3463 No answer from host','no_action'),
	(224,0,'4.4.2','',0,'RFC3463 Bad connection','retry'),
	(225,0,'5.4.2','',0,'RFC3463 Bad connection','no_action'),
	(226,0,'4.4.3','',0,'RFC3463 Directory server failure','retry'),
	(227,0,'5.4.3','',0,'RFC3463 Directory server failure','no_action'),
	(228,0,'4.4.4','',0,'RFC3463 Unable to route','retry'),
	(229,0,'5.4.4','',0,'RFC3463 Unable to route','no_action'),
	(230,0,'4.4.5','',0,'RFC3463 Mail system congestion','retry'),
	(231,0,'4.4.6','',0,'RFC3463 Routing loop detected','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(232,0,'5.4.6','',0,'RFC3463 Routing loop detected','no_action'),
	(233,0,'4.4.7','',0,'RFC3463 Delivery time expired','retry'),
	(234,0,'5.4.7','',0,'RFC3463 Delivery time expired','no_action'),
	(235,0,'4.5.0','',0,'RFC3463 Other or undefined protocol status','retry'),
	(236,0,'5.5.0','',0,'RFC3463 Other or undefined protocol status','suppress'),
	(237,0,'4.5.1','',0,'RFC3463 Invalid command','no_action'),
	(238,0,'5.5.1','',0,'RFC3463 Invalid command','no_action'),
	(239,0,'5.5.2','',0,'RFC3463 Syntax error','no_action'),
	(240,0,'4.5.3','',0,'RFC3463 Too many recipients','retry'),
	(241,0,'5.5.3','',0,'RFC3463 Too many recipients','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(242,0,'5.5.4','',0,'RFC3463 Invalid command arguments','no_action'),
	(243,0,'4.5.5','',0,'RFC3463 Wrong protocol version','retry'),
	(244,0,'5.5.5','',0,'RFC3463 Wrong protocol version','no_action'),
	(245,0,'5.6.0','',0,'RFC3463 Other or undefined media error','no_action'),
	(246,0,'5.6.1','',0,'RFC3463 Media not supported','no_action'),
	(247,0,'5.6.2','',0,'RFC3463 Conversion required and prohibited','no_action'),
	(248,0,'5.6.3','',0,'RFC3463 Conversion required but not supported','no_action'),
	(249,0,'5.6.5','',0,'RFC3463 Conversion failed','no_action'),
	(250,0,'5.6.6','',0,'RFC4468 Message content not available','no_action'),
	(251,0,'5.6.7','',0,'RFC6531 Non-ASCII addresses not permitted for that sender/recipient','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(252,0,'5.6.8','',0,'RFC6531 UTF-8 string reply is required but not permitted by the SMTP client','no_action'),
	(253,0,'5.6.9','',0,'RFC6531 UTF-8 header message cannot be transferrred to one or more recipients','no_action'),
	(254,0,'4.7.0','',0,'RFC3463 Other or undefined security status','no_action'),
	(255,0,'5.7.0','',0,'RFC3463 Other or undefined security status','no_action'),
	(256,0,'4.7.1','',0,'RFC3463 Delivery not authorized','retry'),
	(257,0,'5.7.1','',0,'RFC3463 Delivery not authorized','no_action'),
	(258,0,'5.7.2','',0,'RFC3463 Mailing list expansion prohibited','no_action'),
	(259,0,'5.7.3','',0,'RFC3463 Security conversion required but not possible','no_action'),
	(260,0,'5.7.4','',0,'RFC3463 Security features not supported','no_action'),
	(261,0,'5.7.5','',0,'RFC3463 Cryptographic failure','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(262,0,'5.7.6','',0,'RFC3463 Cryptographic algorithm not supported','no_action'),
	(263,0,'5.7.7','',0,'RFC3463 Message integrity failure','no_action'),
	(264,0,'5.7.8','',0,'RFC4954 Authentication credentials invalid','no_action'),
	(265,0,'5.7.9','',0,'RFC4954 Authentication mechanism is too weak','no_action'),
	(266,0,'5.7.10','',0,'RFC5248 Encryption Needed','no_action'),
	(267,0,'5.7.11','',0,'RFC4954 Encryption required for requested authentication mechanism','no_action'),
	(268,0,'5.7.12','',0,'RFC4954 A password transition is needed','no_action'),
	(269,0,'5.7.13','',0,'RFC5248 User Account Disabled','no_action'),
	(270,0,'5.7.14','',0,'RFC5248 Trust relationship required','no_action'),
	(271,0,'4.7.15','',0,'RFC6710 Priority Level is too low','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(272,0,'5.7.15','',0,'RFC6710 Priority Level is too low','no_action'),
	(273,0,'4.7.16','',0,'RFC6710 Message is too big for the specified priority','retry'),
	(274,0,'5.7.16','',0,'RFC6710 Message is too big for the specified priority','no_action'),
	(275,0,'5.7.17','',0,'RFC7293 Mailbox owner has changed','suppress'),
	(276,0,'5.7.18','',0,'RFC7293 Domain owner has changed','suppress'),
	(277,0,'5.7.19','',0,'RFC7293 RRVS test cannot be completed','no_action'),
	(278,0,'5.7.20','',0,'RFC7372 No passing DKIM signature found','no_action'),
	(279,0,'5.7.21','',0,'RFC7372 No acceptable DKIM signature found','no_action'),
	(280,0,'5.7.22','',0,'RFC7372 No valid author-matched DKIM signature found','no_action'),
	(281,0,'5.7.23','',0,'RFC7372 SPF validation failed','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(282,0,'4.7.24','',0,'RFC7372 SPF validation error','no_action'),
	(283,0,'5.7.24','',0,'RFC7372 SPF validation error','no_action'),
	(284,0,'5.7.25','',0,'RFC7372 Revers DNS validation falied','no_action'),
	(285,0,'5.7.26','',0,'RFC7372 Multiple authentication checks failed','no_action'),
	(286,0,'5.7.27','',0,'RFC7505 Sender address has null MX','no_action'),
	(287,0,'','policy reasons',0,'Policy Reasons should never be suppressed','no_action'),
	(288,0,'','block|blacklist|bulk mail|virus|reputation|content|dmarc|spam',0,'Common messages to avoid suppression','no_action'),
	(290,0,'','Unauthenticated email is not accepted from this domain',0,'','no_action'),
	(291,0,'','quota',0,'Quota responses should not result in suppression','no_action'),
	(292,0,'','DNSBL|SURBL|CURBL|symantec|brightmail',0,'Combined regex for well-known blacklist/spamfilter responses','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(295,554,'','Denied',0,'Denied response typically from mxlogic.net','no_action'),
	(296,0,'','mailbox full',0,'','no_action'),
	(297,521,'5.2.1','\\(CON:B1\\)',0,'AOL code indicating that the sending IP has been added to the blacklist','no_action'),
	(298,521,'5.2.1','AOL will not accept delivery of this message',0,'AOL code that lacks specificity regarding recipient validity','no_action'),
	(300,552,'4.2.2','User has full mailbox',0,'message will not succeed when retried, but address is likely valid','no_action'),
	(301,521,'5.2.1','HVU:B1',0,'AOL code indicating content in email header/hostname is generating excessive complaints','no_action'),
	(309,0,'','mail box is full',0,'','no_action'),
	(310,552,'5.2.2','',0,'Typically used for full mailbox / over quota','no_action'),
	(311,550,'','Connection frequency limited',0,'Only used by qq.com','retry'),
	(312,550,'','Access Denied\\.\\.\\.',0,'Netzero/Juno/UnitedOnline access denied code','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(313,550,'','550 permanent failure for one or more recipients',0,'','no_action'),
	(314,550,'5.0.0','Mail rejected',0,'','no_action'),
	(319,550,'5.2.1','receiving mail at a rate that prevents additional messages',0,'gmail rate limit exceeded for this user','no_action'),
	(321,554,'','This account has been temporarily suspended',0,'','no_action'),
	(323,554,'5.7.1','cannot find your hostname',0,'','no_action'),
	(324,554,'5.7.1','too many different IP\'s for domain',0,'','no_action'),
	(325,550,'5.7.1','command rejected',0,'','no_action'),
	(326,550,'','IP frequency limited',0,'','retry'),
	(329,0,'5.1.0','',0,'Used primarily by tier 2 ISPs. Indicates an invalid MAIL FROM domain or missing MX record for MAIL FROM domain.','no_action'),
	(330,452,'4.1.1','Too much mail',0,'Used primarily by Charter.','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(331,452,'4.1.1','temporary failure',0,'Used primarily by Hotmail.','retry'),
	(332,452,'4.1.1','Account temporarily unavailable',0,'Used primarily by Excite.com.','retry'),
	(333,452,'4.1.1','',0,'An override for what appears to be spammy content','no_action'),
	(334,554,'','delivery error: .*',0,'Handles two known Unknown User bounces from yahoo.com','suppress'),
	(336,550,'5.2.1','The email account that you tried to reach is disabled.*',10,'Gmail - Disabled email address bounce','suppress'),
	(337,554,'5.4.14','Hop count exceeded',0,'Asynchronous bounce. Check Block Regex table','no_action'),
	(339,554,'','Invalid recipient',10,'Standard invalid recipient response for a few tier two ISPs','suppress'),
	(342,451,'','greylist',0,'greylisting','retry'),
	(343,421,'','try again',0,'forcing default retry behavior for 421','retry'),
	(345,553,'','Domain of sender address',0,'appear to be legit addresses failing due to temporary MX issue','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(346,421,'4.7.0','[GL01].*',0,'yahoo deferral for too much mail may be reputation based','retry'),
	(348,450,'4.3.0','Spam filter message initialization failure',0,'verizon issue with spam filter. NOT SPAM RELATED','retry'),
	(352,550,'5.1.1','RESOLVER.ADR.RecipNotFound',0,'This response is returned when an address migrates to office 365. Addresses almost never open or click emails which indicates a permanent failure.','suppress'),
	(353,550,'','Service refuse. Veuillez essayer plus tard. service refused, please try later.*',0,'Laposte.net temporary failure.','retry'),
	(354,451,'4.7.0','Unknown temporary connection problem',0,NULL,'retry'),
	(355,550,'','550 Invalid RDNS entry for.*',0,'A few small ISPs return this for rDNS issues','no_action'),
	(356,451,'4.7.0','',0,'Temporary failures at a large number of small ISPs','retry'),
	(357,451,'4.3.0','Tempfailed as anti-spam measure',0,'Deferred as a result of spam filter. Message should succeed in 20 minutes.','retry'),
	(358,451,'','Message Tempfailed.*',0,'Temporarily deferred, sometimes due to spam reasons, sometimes not.','retry'),
	(359,451,'4.3.0','',1,'Deferred as a result of spam filter or reputation issue.','retry');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(360,550,'','Administrative prohibition',0,'Soft bounce likely due to spam filter settings.','no_action'),
	(362,550,'','Denied by policy',0,'Spam filter rejection from small business domains.','no_action'),
	(364,550,'5.1.1','User [Uu]nknown',0,'unknown users should be suppressed','suppress'),
	(367,554,'','mailbox not found',0,'Invalid address','suppress'),
	(368,421,'','Your IP address has been temporarily blocked - SPAM.',1,'Message throttled due to spam.','retry'),
	(369,554,'5.7.1','Recipient address rejected.*does not exist',0,'Invalid address','suppress'),
	(370,550,'5.2.1','Mailbox not available',0,'Invalid email address','suppress'),
	(371,550,'5.2.0','No such mailbox',0,'Invalid address','suppress'),
	(372,550,'5.2.1','Mailbox disabled for this recipient',0,'Invalid email address','suppress'),
	(373,550,'5.2.1','Mailbox unavailable for this recipient',0,'Invalid address','suppress');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(374,554,'','no valid recipients',0,'Invalid email address','suppress'),
	(375,554,'','Sorry, no mailbox here by that name.',0,'Invalid address','suppress'),
	(376,554,'5.7.1','<.*>: Client host rejected: User \\(.*\\) not known to us please verify your address book for any typos in this email address or inquire manually\\.',0,'Client host rejected, user not known','suppress'),
	(377,450,'4.2.0','Please wait 5min and then try again',0,'temporary spam','retry'),
	(378,450,'4.7.1','temporary greylisted by CYREN IP reputation',0,'temporarily greylisted by CYREN','retry'),
	(379,551,'','[Ss]orry. \\(\\#5\\.1\\.1. ',0,'no mailbox here should be suppressed','suppress'),
	(380,452,'4.1.1','Greylisting in action, please try again later',0,'Pandora greylisting','retry'),
	(381,421,'4.5.1','',0,'No more messages on connection','retry'),
	(382,421,'','mxcmd06.ad.aruba.it bizsmtp.*Too many connections, try later.',0,'Too many connections, mxcmd06.ad.aruba.it','retry'),
	(383,532,'5.3.2','STOREDRV\\.Deliver. Missing or bad mailbox Database property',0,'Missing or bad mailbox','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(384,501,'5.5.4','Invalid arguments',0,'Invalid arguments','no_action'),
	(385,550,'5.2.1','Addressee unknown',0,'tier 2 inbox provider invalid address','suppress'),
	(386,550,'5.1.0','\\<[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\> This IP has sent too many messages this hour\\. IB504 \\<http\\:\\/\\/x\\.co\\/rlbounce\\>',0,'legacy GoDaddy response should be retried','retry'),
	(387,550,'5.7.606','Access denied, banned sending IP .*.',0,'Microsoft Hard Bounce','no_action'),
	(388,553,'5.3.0','No such user here',0,'Invalid Email Address','suppress'),
	(389,550,'','Recipient address rejected',2,'Invalid Email Address','suppress'),
	(390,554,'5.7.1','Recipient address rejected: Unknown user',0,'Invalid Email Address','suppress'),
	(391,554,'5.7.1','Recipient address rejected: Unknown recipient',0,'Invalid Email Address','suppress'),
	(392,550,'5.4.1','',0,'Invalid Email Address','suppress'),
	(393,521,'5.2.1','HVU:B2',0,'AOL Blocked for spam','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(394,550,'5.7.1','No such user',0,'Yandex invalid email address','suppress'),
	(395,452,'','Too many recipients received this hour. Please see',0,'Road Runner deferral message','retry'),
	(396,450,'4.7.1','deferred using Trend Micro Email Reputation database',0,'Mail deferred due to TrendMicro listing.','retry'),
	(397,554,'5.7.1','no such folder id',0,'Unknown Bounce reason from AOL','no_action'),
	(398,550,'5.1.0','This IP has sent too many messages this hour',0,'Secure Server rate limiting deferral.','retry'),
	(399,553,'','Sender is on user denylist',0,'small inbox provider user filter setting','no_action'),
	(400,550,'','resolver.rst',0,'Office 365 recipient configuration issue','no_action'),
	(401,550,'5.4.300','',0,'Office 365 recipient configuration issue','no_action'),
	(402,550,'','the message was rejected by organization policy',0,'Office 365 recipient policy block','no_action'),
	(403,550,'5.1.10','',0,'Office365 configuration problem. Hard Bounce','suppress');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(404,451,'4.3.5','server configuration error',0,'Unknown client configuration error','retry'),
	(405,552,'','No Such user',0,'Invalid Recipient at several typo domains','suppress'),
	(406,450,'4.2.0','450 4.2.0 <.+>: Recipient address rejected: .+ greylisted for \\d+ seconds by ZEROSPAM',0,'Temporary greylist block','retry'),
	(407,550,'','spf',0,'SPF failure','no_action'),
	(408,550,'','Sender IP address rejected',0,'IP address rejection is a temporary fail','no_action'),
	(410,554,'','Invalid mailbox',0,'Roadrunner invalid mailbox response','suppress'),
	(411,454,'','This message could not be scanned for viruses',0,'MX Logic spam filter problem.','retry'),
	(412,421,'','Cloudmark Gateway Too many connections from same IP',0,'temporary block containing the word cloudmark causing a no action','retry'),
	(413,550,'','Please turn on SMTP Authentication in your mail client',0,'Unknown authentication error.','no_action'),
	(414,550,'','RCPT TO:<.*> Mail.*',0,'Libero.it Mailbox full block','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(415,451,'','IP temporarily blacklisted',2,'Mimecast temporary reputation based deferral','retry'),
	(416,550,'','RCPT TO:<.*> Relay',0,'Unknown Auth Failure at Italian Domains','no_action'),
	(418,550,'5.0.350','501 Syntax error',0,'Poorly formatted Asyncronous bounce not indicating a bad address','no_action'),
	(419,0,'5.1.0','Unknown address error',0,'Weird invalid address bounce from a few small ISPs','suppress'),
	(420,0,'','Bad destination email address',0,'Invlaid address bounce from several small ISPs','suppress'),
	(421,501,'5.5.4','Invalid Address',0,'Invalid address response from several small ISPs','suppress'),
	(422,0,'','no mailbox here by that name',0,'invalid address from several small ISPs','suppress'),
	(423,550,'5.7.1','Requested action not taken',0,'Invalid address response from several small ISPs','suppress'),
	(424,550,'','Unable to deliver to',0,'invalid address response from several small ISPs','suppress'),
	(425,553,'5.7.1','No such user here',0,'Invalid address response from several small ISPs','suppress');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(426,451,'4.3.5','server configuration problem',0,'temporary problem message will likely deliver','retry'),
	(427,451,'','You have been greylisted.',0,'Temporary greylisting likely matching some other bounce rule.','retry'),
	(428,550,'5.1.1','sorry, you are violating our security policies',0,'Response from domains hosted by Aurba.it domains. Unsure of what security policies means','no_action'),
	(429,421,'','Antispam da SoftSell implementa greylist',0,'Temporary Greylisting','retry'),
	(430,550,'','Domain frequency limited',0,'qq.com reputation based deferral','retry'),
	(431,552,'','RCPT TO.*Mail.*',0,'','no_action'),
	(432,550,'','Maximum line length exceeded',0,'Misused 550 response. This should be a block not a bounce.','no_action'),
	(433,421,'4.7.1','will be permanently deferred',0,'yahoo permanent deferral','no_action'),
	(435,550,'5.6.0','Invalid header found',0,'Usually a header that is too long, but not always','no_action'),
	(436,550,'','invalid DNS MX or A/AAAA resource record',0,'DNS issue. Likely missing MX record.','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(437,550,'','Sender has no A, AAAA, or MX DNS records.',0,'Sender is missing required DNS records.','no_action'),
	(438,550,'','The email sent from',0,'Missing reqired DNS records','no_action'),
	(439,550,'','Domain of sender address',0,'Missing reqired DNS records.','no_action'),
	(440,550,'','Verification failed for',0,'Missing required DNS records.','no_action'),
	(441,553,'','unable to verify address',0,'Missing required DNS records','no_action'),
	(442,550,'','sender rejected',0,'Sender blacklisted or missing required dns records','no_action'),
	(443,550,'','Sender email address rejected',0,'Blocked for unknown reason','no_action'),
	(444,550,'','Sender IP reverse lookup rejected',0,'Missing some DNS records','no_action'),
	(445,553,'','your envelope sender domain',0,'Missing required dns records','no_action'),
	(446,550,'','invalid sender',0,'Sender likely missing required dns records','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(447,550,'','Try again later',0,'Temporary delay. Possible reputation issue.','retry'),
	(448,452,'','sender rejected',0,'Sender rejected. Should not be a deferral. Misused 452','no_action'),
	(449,550,'','Too many connect',0,'Too many connections to Sina address','retry'),
	(450,550,'','Too many invalid',0,'Throttling for bounces or too much mail','retry'),
	(451,550,'','Too many recipients',0,'Message going to too many recipients unable to deliver','no_action'),
	(452,554,'','user no found',0,'Invalid recipient from Sina','suppress'),
	(453,550,'','recipient rejected. IB603a',0,'Bigpond.com bounce. Looks like a hard bounce, but can be temporary problem with mailbox.','no_action'),
	(454,0,'','Access denied, banned sending IP',0,'Blacklisted IP response','no_action'),
	(455,550,'5.7.1','unknown or illegal user',0,'Hard bounce from some university','suppress'),
	(456,550,'','csi.cloudmark.com',0,'Cloudmark Blocking','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(457,550,'5.1.0','Recipient rejected',0,'Invalid Email','suppress'),
	(458,451,'4.5.1','Message greylisted',0,'greylisting. should retry.','retry'),
	(459,550,'','does not pass',0,'SPF failure. Should be no_action','no_action'),
	(460,554,'','HL:ITC',0,'Chinese ISP deferral for too many connections','retry'),
	(461,550,'5.7.51','',0,'Office 365 configuration problem. Not invalid address','no_action'),
	(462,550,'5.7.511','',0,'Office 365 blocking. Not an invalid Address','no_action'),
	(463,0,'','connection closed',1,'Unknown connection closure. primarily comcast.net. always in STARTTLS','no_action'),
	(464,550,'','HELO host host_mismatch',0,'','no_action'),
	(465,550,'','rejected as spam',0,'blocked for spam reasons','no_action'),
	(467,550,'5.7.1','message refused',1,'Aparent spam filtering blocking from a number of small domains','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(468,553,'5.1.2','The sender address',0,'Unknown sender address error','no_action'),
	(469,550,'','spf-check failed',0,'','no_action'),
	(470,550,'5.7.1','spf policy',0,'','no_action'),
	(471,554,'','transaction failed',0,'','no_action'),
	(472,553,'5.3.0','The email account that you tried to reach does not exist',0,'','suppress'),
	(473,550,'','no such user',0,'','suppress'),
	(474,554,'','no such user',0,'','suppress'),
	(475,553,'','no such user',0,'','suppress'),
	(476,450,'4.2.0','greylist',1,'greylisting. should retry.','retry'),
	(477,550,'','Requested action not taken: mailbox unavailable Reject due to policy restrictions. For explanation visit',0,'GMX.com','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(478,550,'','verify failed due to route',0,'Small au domain blocking','no_action'),
	(479,550,'','JunkMail rejected',0,'Blocked by some RBL. Usually SpamCop or Barracuda','no_action'),
	(480,450,'4.7.0','Your email is being verified',0,'Weird email/reputation verification deferral','retry'),
	(481,550,'','Rejecting for Sender Policy Framework',0,'Misc SPF policy rejection. Should be block not bounce.','no_action'),
	(482,450,'4.1.1','user unknown',0,'malformed response using the wrong code. should be hard bounce.','suppress'),
	(483,550,'5.7.1','relaying denied',2,'unknown block. close match to rule 423. Hence the priority of 2.','no_action'),
	(484,550,'','subject contains invalid characters',0,'Misc bounce for bad subject even when subject looks good.','no_action'),
	(485,550,'5.4.317','Message expired',0,'unknown asynchronous response from some o365 domains','no_action'),
	(486,550,'','Rejected by header based Anti-Spoofing policy',0,'Mimecast anti-spoofing response','no_action'),
	(487,554,'','Undeliverable localhost',0,'SendGird generated response for messages sent to domains with localhost in the MX record.','no_action');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(488,450,'4.1.1','Recipient address rejected',5,'Malformed response from a bunch of small domains.','suppress'),
	(490,450,'4.7.3','Organization queue quota',0,'Mysterious Microsoft Response','retry'),
	(491,550,'','Unroutable sender address',0,'mail.ru block for no mx record (usually whitelabel related)','no_action'),
	(492,550,'','relay not permitted',0,'unknown block. likely reputation based.','no_action'),
	(493,550,'','attachments here',0,'attachments not allowed.','no_action'),
	(494,550,'','we do not accept mail from this address',0,'unknown blocking','no_action'),
	(495,451,'4.7.650','has been temporarily',0,'Microsoft deferral that started on april 4th 2018','retry'),
	(496,451,'4.7.651','has been temporarily',0,'Microsoft deferral that started on april 4th 2018','retry'),
	(497,554,'','unknown recipient',0,'unknown recipient should be a hard bounce','suppress'),
	(498,0,'','recipient address rejected: user unknown',0,'A lot of responses with this string. Unfortunately a huge number of response and enhanced codes.','suppress');

INSERT INTO `bounce_rule` (`id`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`)
VALUES
	(499,550,'','from contains invalid characters',0,'from address contains invalid characters','no_action'),
	(500,450,'4.7.0','deferred',0,'proofpoint deferral','retry'),
	(501,450,'4.1.1','Address verification in progress',1,'Failure is apparently related to the recipient address not being verified. Lots of engagement on these addresses. should not suppress','retry'),
	(502,550,'','error occurred fetching the matches description',0,'strange mimecast block','no_action'),
	(503,450,'4.7.0','Your emailing is being verified',0,'finish deferral secmail filtering','retry'),
	(504,550,'','no MX record for domain',0,'mainly liberty domain block seeing ~50% of addresses engaging SG wide','no_action');

/*!40000 ALTER TABLE `bounce_rule` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
