USE drop_rules;

DROP TABLE IF EXISTS `changelog`;

CREATE TABLE IF NOT EXISTS `changelog` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
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
  `operation` ENUM('Create', 'Delete', 'Update') NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


LOCK TABLES `changelog` WRITE;

INSERT INTO `changelog` (`id`, `rule_id`, `user_id`, `comment`, `created_at`, `response_code`, `enhanced_code`, `regex`, `priority`, `description`, `bounce_action`,`operation`) VALUES 
  (173, 173,1,'Inital Setup',1,421,'','',0,'RFC5321Servicenotavailable','retry','Create'),
  (174, 174,1,'Inital Setup',2,450,'','',0,'RFC5321Mailboxunavailable','retry','Create'),
  (175, 175,1,'Inital Setup',3,451,'','',0,'RFC5321Localerrorinprocessing','retry','Create'),
  (176, 176,1,'Inital Setup',4,452,'','',0,'RFC5321Insufficientsystemstorage','retry','Create'),
  (177, 177,1,'Inital Setup',5,454,'','',0,'RFC3207TLSnotavailable','retry','Create'),
  (178, 178,1,'Inital Setup',6,455,'','',0,'RFC5321Serverunabletoaccomodateparameters','retry','Create'),
  (179, 179,1,'Inital Setup',7,500,'','',0,'RFC5321Syntaxerrorcommandunrecognized','no_action','Create'),
  (180, 180,1,'Inital Setup',8,501,'','',0,'RFC5321Syntaxerrorinparametersorarguments','no_action','Create'),
  (181, 181,1,'Inital Setup',9,502,'','',0,'RFC5321Commandnotimplemented','no_action','Create'),
  (182, 182,1,'Inital Setup',10,503,'','',0,'RFC5321Badsequenceofcommands','no_action','Create'),
  (183, 183,1,'Inital Setup',11,521,'','',0,'RFC7504Serverdoesnotacceptmail','suppress','Create'),
  (184, 184,1,'Inital Setup',12,530,'','',0,'RFC4954MustissueaSTARTTLScommandfirst','no_action','Create'),
  (185, 185,1,'Inital Setup',13,550,'','',0,'RFC5321Mailboxunavailable','suppress','Create'),
  (186, 186,1,'Inital Setup',14,551,'','',0,'RFC5321Usernotlocal','suppress','Create'),
  (187, 187,1,'Inital Setup',15,552,'','',0,'RFC5321Exceededstorageallocation','no_action','Create'),
  (188, 188,1,'Inital Setup',16,553,'','',0,'RFC5321Mailboxnamenotallowed','suppress','Create'),
  (189, 189,1,'Inital Setup',17,554,'','',0,'RFC5321Transactionfailed','no_action','Create'),
  (190, 190,1,'Inital Setup',18,555,'','',0,'RFC5321Parametersnotrecognizedornotimplemented','no_action','Create'),
  (191, 191,1,'Inital Setup',19,556,'','',0,'RFC7504RecipientaddresshasnullMX','no_action','Create'),
  (192, 192,1,'Inital Setup',20,0,'5.1.1','',0,'RFC3463Baddestinationmailboxaddress','suppress','Create'),
  (193, 193,1,'Inital Setup',21,0,'5.1.2','',0,'RFC3463Baddestinationsystemaddress','suppress','Create'),
  (194, 194,1,'Inital Setup',22,0,'5.1.3','',0,'RFC3463Baddestinationmailboxaddresssyntax','suppress','Create'),
  (195, 195,1,'Inital Setup',23,0,'5.1.4','',0,'RFC3463Destinationmailboxaddressambiguous','no_action','Create'),
  (196, 196,1,'Inital Setup',24,0,'5.1.6','',0,'RFC3463Destinationmailboxhasmoved','suppress','Create'),
  (197, 197,1,'Inital Setup',25,0,'4.1.7','',0,'RFC3463Badsendersmailboxaddresssyntax','no_action','Create'),
  (198, 198,1,'Inital Setup',26,0,'5.1.7','',0,'RFC3463Badsendersmailboxaddresssyntax','no_action','Create'),
  (199, 199,1,'Inital Setup',27,0,'4.1.8','',0,'RFC3463Badsenderssystemaddress','no_action','Create'),
  (200, 200,1,'Inital Setup',28,0,'5.1.8','',0,'RFC3463Badsenderssystemaddress','no_action','Create'),
  (201, 201,1,'Inital Setup',29,0,'5.1.9','',0,'RFC3886Messagerelayedtonon-compliantmailer','no_action','Create'),
  (202, 202,1,'Inital Setup',30,0,'5.1.10','',0,'RFC7505RecipientaddresshasnullMX','no_action','Create'),
  (203, 203,1,'Inital Setup',31,0,'4.2.0','',0,'RFC3463Otherorundefinedmailboxstatus','retry','Create'),
  (204, 204,1,'Inital Setup',32,0,'5.2.0','',0,'RFC3463Otherorundefinedmailboxstatus','no_action','Create'),
  (205, 205,1,'Inital Setup',33,0,'4.2.1','',0,'RFC3463Mailboxdisablednotacceptingmessages','retry','Create'),
  (206, 206,1,'Inital Setup',34,0,'5.2.1','',0,'RFC3463Mailboxdisablednotacceptingmessages','suppress','Create'),
  (207, 207,1,'Inital Setup',35,0,'4.2.2','',0,'RFC3463Mailboxfull','retry','Create'),
  (208, 208,1,'Inital Setup',36,0,'5.2.2','',0,'RFC3463Mailboxfull','no_action','Create'),
  (209, 209,1,'Inital Setup',37,0,'5.2.3','',0,'RFC3463Messagelengthexceedsadministrativelimit','no_action','Create'),
  (210, 210,1,'Inital Setup',38,0,'5.2.4','',0,'RFC3463Mailinglistexpansionproblem','no_action','Create'),
  (211, 211,1,'Inital Setup',39,0,'4.3.0','',0,'RFC3463Otherorundefinedmailsystemstatus','retry','Create'),
  (212, 212,1,'Inital Setup',40,0,'5.3.0','',0,'RFC3463Otherorundefinedmailsystemstatus','no_action','Create'),
  (213, 213,1,'Inital Setup',41,0,'4.3.1','',0,'RFC3463Mailsystemfull','retry','Create'),
  (214, 214,1,'Inital Setup',42,0,'4.3.2','',0,'RFC3463Systemnotacceptingnetworkmessages','retry','Create'),
  (215, 215,1,'Inital Setup',43,0,'5.3.2','',0,'RFC3463Systemnotacceptingnetworkmessages','suppress','Create'),
  (216, 216,1,'Inital Setup',44,0,'4.3.3','',0,'RFC3463Systemnotcapableofselectedfeatures','no_action','Create'),
  (217, 217,1,'Inital Setup',45,0,'5.3.3','',0,'RFC3463Systemnotcapableofselectedfeatures','no_action','Create'),
  (218, 218,1,'Inital Setup',46,0,'5.3.4','',0,'RFC3463Messagetoobigforsystem','no_action','Create'),
  (219, 219,1,'Inital Setup',47,0,'4.3.5','',0,'RFC3463Systemincorrectlyconfigured','no_action','Create'),
  (220, 220,1,'Inital Setup',48,0,'5.3.5','',0,'RFC3463Systemincorrectlyconfigured','no_action','Create'),
  (221, 221,1,'Inital Setup',49,0,'4.4.0','',0,'RFC3463Otherorundefinednetworkorroutingstatus','retry','Create'),
  (222, 222,1,'Inital Setup',50,0,'5.4.0','',0,'RFC3463Otherorundefinednetworkorroutingstatus','no_action','Create'),
  (223, 223,1,'Inital Setup',51,0,'5.4.1','',0,'RFC3463Noanswerfromhost','no_action','Create'),
  (224, 224,1,'Inital Setup',52,0,'4.4.2','',0,'RFC3463Badconnection','retry','Create'),
  (225, 225,1,'Inital Setup',53,0,'5.4.2','',0,'RFC3463Badconnection','no_action','Create'),
  (226, 226,1,'Inital Setup',54,0,'4.4.3','',0,'RFC3463Directoryserverfailure','retry','Create'),
  (227, 227,1,'Inital Setup',55,0,'5.4.3','',0,'RFC3463Directoryserverfailure','no_action','Create'),
  (228, 228,1,'Inital Setup',56,0,'4.4.4','',0,'RFC3463Unabletoroute','retry','Create'),
  (229, 229,1,'Inital Setup',57,0,'5.4.4','',0,'RFC3463Unabletoroute','no_action','Create'),
  (230, 230,1,'Inital Setup',58,0,'4.4.5','',0,'RFC3463Mailsystemcongestion','retry','Create'),
  (231, 231,1,'Inital Setup',59,0,'4.4.6','',0,'RFC3463Routingloopdetected','retry','Create'),
  (232, 232,1,'Inital Setup',60,0,'5.4.6','',0,'RFC3463Routingloopdetected','no_action','Create'),
  (233, 233,1,'Inital Setup',61,0,'4.4.7','',0,'RFC3463Deliverytimeexpired','retry','Create'),
  (234, 234,1,'Inital Setup',62,0,'5.4.7','',0,'RFC3463Deliverytimeexpired','no_action','Create'),
  (235, 235,1,'Inital Setup',63,0,'4.5.0','',0,'RFC3463Otherorundefinedprotocolstatus','retry','Create'),
  (236, 236,1,'Inital Setup',64,0,'5.5.0','',0,'RFC3463Otherorundefinedprotocolstatus','suppress','Create'),
  (237, 237,1,'Inital Setup',65,0,'4.5.1','',0,'RFC3463Invalidcommand','no_action','Create'),
  (238, 238,1,'Inital Setup',66,0,'5.5.1','',0,'RFC3463Invalidcommand','no_action','Create'),
  (239, 239,1,'Inital Setup',67,0,'5.5.2','',0,'RFC3463Syntaxerror','no_action','Create'),
  (240, 240,1,'Inital Setup',68,0,'4.5.3','',0,'RFC3463Toomanyrecipients','retry','Create'),
  (241, 241,1,'Inital Setup',69,0,'5.5.3','',0,'RFC3463Toomanyrecipients','retry','Create'),
  (242, 242,1,'Inital Setup',70,0,'5.5.4','',0,'RFC3463Invalidcommandarguments','no_action','Create'),
  (243, 243,1,'Inital Setup',71,0,'4.5.5','',0,'RFC3463Wrongprotocolversion','retry','Create'),
  (244, 244,1,'Inital Setup',72,0,'5.5.5','',0,'RFC3463Wrongprotocolversion','no_action','Create'),
  (245, 245,1,'Inital Setup',73,0,'5.6.0','',0,'RFC3463Otherorundefinedmediaerror','no_action','Create'),
  (246, 246,1,'Inital Setup',74,0,'5.6.1','',0,'RFC3463Medianotsupported','no_action','Create'),
  (247, 247,1,'Inital Setup',75,0,'5.6.2','',0,'RFC3463Conversionrequiredandprohibited','no_action','Create'),
  (248, 248,1,'Inital Setup',76,0,'5.6.3','',0,'RFC3463Conversionrequiredbutnotsupported','no_action','Create'),
  (249, 249,1,'Inital Setup',77,0,'5.6.5','',0,'RFC3463Conversionfailed','no_action','Create'),
  (250, 250,1,'Inital Setup',78,0,'5.6.6','',0,'RFC4468Messagecontentnotavailable','no_action','Create'),
  (251, 251,1,'Inital Setup',79,0,'5.6.7','',0,'RFC6531Non-ASCIIaddressesnotpermittedforthatsender/recipient','no_action','Create'),
  (252, 252,1,'Inital Setup',80,0,'5.6.8','',0,'RFC6531UTF-8stringreplyisrequiredbutnotpermittedbytheSMTPclient','no_action','Create'),
  (253, 253,1,'Inital Setup',81,0,'5.6.9','',0,'RFC6531UTF-8headermessagecannotbetransferrredtooneormorerecipients','no_action','Create'),
  (254, 254,1,'Inital Setup',82,0,'4.7.0','',0,'RFC3463Otherorundefinedsecuritystatus','no_action','Create'),
  (255, 255,1,'Inital Setup',83,0,'5.7.0','',0,'RFC3463Otherorundefinedsecuritystatus','no_action','Create'),
  (256, 256,1,'Inital Setup',84,0,'4.7.1','',0,'RFC3463Deliverynotauthorized','retry','Create'),
  (257, 257,1,'Inital Setup',85,0,'5.7.1','',0,'RFC3463Deliverynotauthorized','no_action','Create'),
  (258, 258,1,'Inital Setup',86,0,'5.7.2','',0,'RFC3463Mailinglistexpansionprohibited','no_action','Create'),
  (259, 259,1,'Inital Setup',87,0,'5.7.3','',0,'RFC3463Securityconversionrequiredbutnotpossible','no_action','Create'),
  (260, 260,1,'Inital Setup',88,0,'5.7.4','',0,'RFC3463Securityfeaturesnotsupported','no_action','Create'),
  (261, 261,1,'Inital Setup',89,0,'5.7.5','',0,'RFC3463Cryptographicfailure','no_action','Create'),
  (262, 262,1,'Inital Setup',90,0,'5.7.6','',0,'RFC3463Cryptographicalgorithmnotsupported','no_action','Create'),
  (263, 263,1,'Inital Setup',91,0,'5.7.7','',0,'RFC3463Messageintegrityfailure','no_action','Create'),
  (264, 264,1,'Inital Setup',92,0,'5.7.8','',0,'RFC4954Authenticationcredentialsinvalid','no_action','Create'),
  (265, 265,1,'Inital Setup',93,0,'5.7.9','',0,'RFC4954Authenticationmechanismistooweak','no_action','Create'),
  (266, 266,1,'Inital Setup',94,0,'5.7.10','',0,'RFC5248EncryptionNeeded','no_action','Create'),
  (267, 267,1,'Inital Setup',95,0,'5.7.11','',0,'RFC4954Encryptionrequiredforrequestedauthenticationmechanism','no_action','Create'),
  (268, 268,1,'Inital Setup',96,0,'5.7.12','',0,'RFC4954Apasswordtransitionisneeded','no_action','Create'),
  (269, 269,1,'Inital Setup',97,0,'5.7.13','',0,'RFC5248UserAccountDisabled','no_action','Create'),
  (270, 270,1,'Inital Setup',98,0,'5.7.14','',0,'RFC5248Trustrelationshiprequired','no_action','Create'),
  (271, 271,1,'Inital Setup',99,0,'4.7.15','',0,'RFC6710PriorityLevelistoolow','retry','Create'),
  (272, 272,1,'Inital Setup',100,0,'5.7.15','',0,'RFC6710PriorityLevelistoolow','no_action','Create'),
  (273, 273,1,'Inital Setup',101,0,'4.7.16','',0,'RFC6710Messageistoobigforthespecifiedpriority','retry','Create'),
  (274, 274,1,'Inital Setup',102,0,'5.7.16','',0,'RFC6710Messageistoobigforthespecifiedpriority','no_action','Create'),
  (275, 275,1,'Inital Setup',103,0,'5.7.17','',0,'RFC7293Mailboxownerhaschanged','suppress','Create'),
  (276, 276,1,'Inital Setup',104,0,'5.7.18','',0,'RFC7293Domainownerhaschanged','suppress','Create'),
  (277, 277,1,'Inital Setup',105,0,'5.7.19','',0,'RFC7293RRVStestcannotbecompleted','no_action','Create'),
  (278, 278,1,'Inital Setup',106,0,'5.7.20','',0,'RFC7372NopassingDKIMsignaturefound','no_action','Create'),
  (279, 279,1,'Inital Setup',107,0,'5.7.21','',0,'RFC7372NoacceptableDKIMsignaturefound','no_action','Create'),
  (280, 280,1,'Inital Setup',108,0,'5.7.22','',0,'RFC7372Novalidauthor-matchedDKIMsignaturefound','no_action','Create'),
  (281, 281,1,'Inital Setup',109,0,'5.7.23','',0,'RFC7372SPFvalidationfailed','no_action','Create'),
  (282, 282,1,'Inital Setup',110,0,'4.7.24','',0,'RFC7372SPFvalidationerror','no_action','Create'),
  (283, 283,1,'Inital Setup',111,0,'5.7.24','',0,'RFC7372SPFvalidationerror','no_action','Create'),
  (284, 284,1,'Inital Setup',112,0,'5.7.25','',0,'RFC7372ReversDNSvalidationfalied','no_action','Create'),
  (285, 285,1,'Inital Setup',113,0,'5.7.26','',0,'RFC7372Multipleauthenticationchecksfailed','no_action','Create'),
  (286, 286,1,'Inital Setup',114,0,'5.7.27','',0,'RFC7505SenderaddresshasnullMX','no_action','Create'),
  (287, 287,1,'Inital Setup',115,0,'','policyreasons',0,'PolicyReasonsshouldneverbesuppressed','no_action','Create'),
  (288, 288,1,'Inital Setup',116,0,'','block|blacklist|bulkmail|virus|reputation|content|dmarc|spam',0,'Commonmessagestoavoidsuppression','no_action','Create'),
  (290, 290,1,'Inital Setup',117,0,'','Unauthenticatedemailisnotacceptedfromthisdomain',0,'','no_action','Create'),
  (291, 291,1,'Inital Setup',118,0,'','quota',0,'Quotaresponsesshouldnotresultinsuppression','no_action','Create'),
  (292, 292,1,'Inital Setup',119,0,'','DNSBL|SURBL|CURBL|symantec|brightmail',0,'Combinedregexforwell-knownblacklist/spamfilterresponses','no_action','Create'),
  (295, 295,1,'Inital Setup',120,554,'','Denied',0,'Deniedresponsetypicallyfrommxlogic.net','no_action','Create'),
  (296, 296,1,'Inital Setup',121,0,'','mailboxfull',0,'','no_action','Create'),
  (297, 297,1,'Inital Setup',122,521,'5.2.1','\\(CON:B1\\)',0,'AOLcodeindicatingthatthesendingIPhasbeenaddedtotheblacklist','no_action','Create'),
  (298, 298,1,'Inital Setup',123,521,'5.2.1','AOLwillnotacceptdeliveryofthismessage',0,'AOLcodethatlacksspecificityregardingrecipientvalidity','no_action','Create'),
  (301, 301,1,'Inital Setup',124,521,'5.2.1','HVU:B1',0,'AOLcodeindicatingcontentinemailheader/hostnameisgeneratingexcessivecomplaints','no_action','Create'),
  (309, 309,1,'Inital Setup',125,0,'','mailboxisfull',0,'','no_action','Create'),
  (310, 310,1,'Inital Setup',126,552,'5.2.2','',0,'Typicallyusedforfullmailbox/overquota','no_action','Create'),
  (311, 311,1,'Inital Setup',127,550,'','Connectionfrequencylimited',0,'Onlyusedbyqq.com','retry','Create'),
  (312, 312,1,'Inital Setup',128,550,'','AccessDenied\\.\\.\\.',0,'Netzero/Juno/UnitedOnlineaccessdeniedcode','no_action','Create'),
  (313, 313,1,'Inital Setup',129,550,'','550permanentfailureforoneormorerecipients',0,'','no_action','Create'),
  (314, 314,1,'Inital Setup',130,550,'5.0.0','Mailrejected',0,'','no_action','Create'),
  (319, 319,1,'Inital Setup',131,550,'5.2.1','receivingmailataratethatpreventsadditionalmessages',0,'gmailratelimitexceededforthisuser','no_action','Create'),
  (321, 321,1,'Inital Setup',132,554,'','Thisaccounthasbeentemporarilysuspended',0,'','no_action','Create'),
  (323, 323,1,'Inital Setup',133,554,'5.7.1','cannotfindyourhostname',0,'','no_action','Create'),
  (324, 324,1,'Inital Setup',134,554,'5.7.1','toomanydifferentIP\'sfordomain',0,'','no_action','Create'),
  (325, 325,1,'Inital Setup',135,550,'5.7.1','commandrejected',0,'','no_action','Create'),
  (326, 326,1,'Inital Setup',136,550,'','IPfrequencylimited',0,'','retry','Create'),
  (329, 329,1,'Inital Setup',137,0,'5.1.0','',0,'Usedprimarilybytier2ISPs.IndicatesaninvalidMAILFROMdomainormissingMXrecordforMAILFROMdomain.','no_action','Create'),
  (330, 330,1,'Inital Setup',138,452,'4.1.1','Toomuchmail',0,'UsedprimarilybyCharter.','retry','Create'),
  (331, 331,1,'Inital Setup',139,452,'4.1.1','temporaryfailure',0,'UsedprimarilybyHotmail.','retry','Create'),
  (332, 332,1,'Inital Setup',140,452,'4.1.1','Accounttemporarilyunavailable',0,'UsedprimarilybyExcite.com.','retry','Create'),
  (333, 333,1,'Inital Setup',141,452,'4.1.1','',0,'Anoverrideforwhatappearstobespammycontent','no_action','Create'),
  (334, 334,1,'Inital Setup',142,554,'','deliveryerror:.*',0,'HandlestwoknownUnknownUserbouncesfromyahoo.com','suppress','Create'),
  (336, 336,1,'Inital Setup',143,550,'5.2.1','Theemailaccountthatyoutriedtoreachisdisabled.*',10,'Gmail-Disabledemailaddressbounce','suppress','Create'),
  (337, 337,1,'Inital Setup',144,554,'5.4.14','Hopcountexceeded',0,'Asynchronousbounce.CheckBlockRegextable','no_action','Create'),
  (339, 339,1,'Inital Setup',145,554,'','Invalidrecipient',10,'StandardinvalidrecipientresponseforafewtiertwoISPs','suppress','Create'),
  (342, 342,1,'Inital Setup',146,451,'','greylist',0,'greylisting','retry','Create'),
  (343, 343,1,'Inital Setup',147,421,'','tryagain',0,'forcingdefaultretrybehaviorfor421','retry','Create'),
  (345, 345,1,'Inital Setup',148,553,'','Domainofsenderaddress',0,'appeartobelegitaddressesfailingduetotemporaryMXissue','retry','Create'),
  (346, 346,1,'Inital Setup',149,421,'4.7.0','[GL01].*',0,'yahoodeferralfortoomuchmailmaybereputationbased','retry','Create'),
  (348, 348,1,'Inital Setup',150,450,'4.3.0','Spamfiltermessageinitializationfailure',0,'verizonissuewithspamfilter.NOTSPAMRELATED','retry','Create'),
  (352, 352,1,'Inital Setup',151,550,'5.1.1','RESOLVER.ADR.RecipNotFound',0,'Thisresponseisreturnedwhenanaddressmigratestooffice365.Addressesalmostneveropenorclickemailswhichindicatesapermanentfailure.','suppress','Create'),
  (354, 354,1,'Inital Setup',152,451,'4.7.0','Unknowntemporaryconnectionproblem',0,NULL,'retry','Create'),
  (355, 355,1,'Inital Setup',153,550,'','550InvalidRDNSentryfor.*',0,'AfewsmallISPsreturnthisforrDNSissues','no_action','Create'),
  (356, 356,1,'Inital Setup',154,451,'4.7.0','',0,'TemporaryfailuresatalargenumberofsmallISPs','retry','Create'),
  (357, 357,1,'Inital Setup',155,451,'4.3.0','Tempfailedasanti-spammeasure',0,'Deferredasaresultofspamfilter.Messageshouldsucceedin20minutes.','retry','Create'),
  (359, 359,1,'Inital Setup',156,451,'4.3.0','',1,'Deferredasaresultofspamfilterorreputationissue.','retry','Create'),
  (360, 360,1,'Inital Setup',157,550,'','Administrativeprohibition',0,'Softbouncelikelyduetospamfiltersettings.','no_action','Create'),
  (362, 362,1,'Inital Setup',158,550,'','Deniedbypolicy',0,'Spamfilterrejectionfromsmallbusinessdomains.','no_action','Create'),
  (364, 364,1,'Inital Setup',159,550,'5.1.1','User[Uu]nknown',0,'unknownusersshouldbesuppressed','suppress','Create'),
  (367, 367,1,'Inital Setup',160,554,'','mailboxnotfound',0,'Invalidaddress','suppress','Create'),
  (368, 368,1,'Inital Setup',161,421,'','YourIPaddresshasbeentemporarilyblocked-SPAM.',1,'Messagethrottledduetospam.','retry','Create'),
  (369, 369,1,'Inital Setup',162,554,'5.7.1','Recipientaddressrejected.*doesnotexist',0,'Invalidaddress','suppress','Create'),
  (370, 370,1,'Inital Setup',163,550,'5.2.1','Mailboxnotavailable',0,'Invalidemailaddress','suppress','Create'),
  (371, 371,1,'Inital Setup',164,550,'5.2.0','Nosuchmailbox',0,'Invalidaddress','suppress','Create'),
  (372, 372,1,'Inital Setup',165,550,'5.2.1','Mailboxdisabledforthisrecipient',0,'Invalidemailaddress','suppress','Create'),
  (373, 373,1,'Inital Setup',166,550,'5.2.1','Mailboxunavailableforthisrecipient',0,'Invalidaddress','suppress','Create'),
  (374, 374,1,'Inital Setup',167,554,'','novalidrecipients',0,'Invalidemailaddress','suppress','Create'),
  (377, 377,1,'Inital Setup',168,450,'4.2.0','Pleasewait5minandthentryagain',0,'temporaryspam','retry','Create'),
  (378, 378,1,'Inital Setup',169,450,'4.7.1','temporarygreylistedbyCYRENIPreputation',0,'temporarilygreylistedbyCYREN','retry','Create'),
  (379, 379,1,'Inital Setup',170,551,'','[Ss]orry.\\(\\#5\\.1\\.1.',0,'nomailboxhereshouldbesuppressed','suppress','Create'),
  (381, 381,1,'Inital Setup',171,421,'4.5.1','',0,'Nomoremessagesonconnection','retry','Create'),
  (383, 383,1,'Inital Setup',172,532,'5.3.2','STOREDRV\\.Deliver.MissingorbadmailboxDatabaseproperty',0,'Missingorbadmailbox','no_action','Create'),
  (384, 384,1,'Inital Setup',173,501,'5.5.4','Invalidarguments',0,'Invalidarguments','no_action','Create'),
  (385, 385,1,'Inital Setup',174,550,'5.2.1','Addresseeunknown',0,'tier2inboxproviderinvalidaddress','suppress','Create'),
  (388, 388,1,'Inital Setup',175,553,'5.3.0','Nosuchuserhere',0,'InvalidEmailAddress','suppress','Create'),
  (389, 389,1,'Inital Setup',176,550,'','Recipientaddressrejected',2,'InvalidEmailAddress','suppress','Create'),
  (390, 390,1,'Inital Setup',177,554,'5.7.1','Recipientaddressrejected:Unknownuser',0,'InvalidEmailAddress','suppress','Create'),
  (391, 391,1,'Inital Setup',178,554,'5.7.1','Recipientaddressrejected:Unknownrecipient',0,'InvalidEmailAddress','suppress','Create'),
  (392, 392,1,'Inital Setup',179,550,'5.4.1','',0,'InvalidEmailAddress','suppress','Create'),
  (393, 393,1,'Inital Setup',180,521,'5.2.1','HVU:B2',0,'AOLBlockedforspam','no_action','Create'),
  (394, 394,1,'Inital Setup',181,550,'5.7.1','Nosuchuser',0,'Yandexinvalidemailaddress','suppress','Create'),
  (395, 395,1,'Inital Setup',182,452,'','Toomanyrecipientsreceivedthishour.Pleasesee',0,'RoadRunnerdeferralmessage','retry','Create'),
  (396, 396,1,'Inital Setup',183,450,'4.7.1','deferredusingTrendMicroEmailReputationdatabase',0,'MaildeferredduetoTrendMicrolisting.','retry','Create'),
  (397, 397,1,'Inital Setup',184,554,'5.7.1','nosuchfolderid',0,'UnknownBouncereasonfromAOL','no_action','Create'),
  (398, 398,1,'Inital Setup',185,550,'5.1.0','ThisIPhassenttoomanymessagesthishour',0,'SecureServerratelimitingdeferral.','retry','Create'),
  (399, 399,1,'Inital Setup',186,553,'','Senderisonuserdenylist',0,'smallinboxprovideruserfiltersetting','no_action','Create'),
  (400, 400,1,'Inital Setup',187,550,'','resolver.rst',0,'Office365recipientconfigurationissue','no_action','Create'),
  (401, 401,1,'Inital Setup',188,550,'5.4.300','',0,'Office365recipientconfigurationissue','no_action','Create'),
  (402, 402,1,'Inital Setup',189,550,'','themessagewasrejectedbyorganizationpolicy',0,'Office365recipientpolicyblock','no_action','Create'),
  (403, 403,1,'Inital Setup',190,550,'5.1.10','',0,'Office365configurationproblem.HardBounce','suppress','Create'),
  (404, 404,1,'Inital Setup',191,451,'4.3.5','serverconfigurationerror',0,'Unknownclientconfigurationerror','retry','Create'),
  (405, 405,1,'Inital Setup',192,552,'','NoSuchuser',0,'InvalidRecipientatseveraltypodomains','suppress','Create'),
  (406, 406,1,'Inital Setup',193,450,'4.2.0','4504.2.0<.+>:Recipientaddressrejected:.+greylistedfor\\d+secondsbyZEROSPAM',0,'Temporarygreylistblock','retry','Create'),
  (407, 407,1,'Inital Setup',194,550,'','spf',0,'SPFfailure','no_action','Create'),
  (408, 408,1,'Inital Setup',195,550,'','SenderIPaddressrejected',0,'IPaddressrejectionisatemporaryfail','no_action','Create'),
  (410, 410,1,'Inital Setup',196,554,'','Invalidmailbox',0,'Roadrunnerinvalidmailboxresponse','suppress','Create'),
  (411, 411,1,'Inital Setup',197,454,'','Thismessagecouldnotbescannedforviruses',0,'MXLogicspamfilterproblem.','retry','Create'),
  (412, 412,1,'Inital Setup',198,421,'','CloudmarkGatewayToomanyconnectionsfromsameIP',0,'temporaryblockcontainingthewordcloudmarkcausinganoaction','retry','Create'),
  (413, 413,1,'Inital Setup',199,550,'','PleaseturnonSMTPAuthenticationinyourmailclient',0,'Unknownauthenticationerror.','no_action','Create'),
  (414, 414,1,'Inital Setup',200,550,'','RCPTTO:<.*>Mail.*',0,'Libero.itMailboxfullblock','no_action','Create'),
  (415, 415,1,'Inital Setup',201,451,'','IPtemporarilyblacklisted',2,'Mimecasttemporaryreputationbaseddeferral','retry','Create'),
  (416, 416,1,'Inital Setup',202,550,'','RCPTTO:<.*>Relay',0,'UnknownAuthFailureatItalianDomains','no_action','Create'),
  (418, 418,1,'Inital Setup',203,550,'5.0.350','501Syntaxerror',0,'PoorlyformattedAsyncronousbouncenotindicatingabadaddress','no_action','Create'),
  (419, 419,1,'Inital Setup',204,0,'5.1.0','Unknownaddresserror',0,'WeirdinvalidaddressbouncefromafewsmallISPs','suppress','Create'),
  (420, 420,1,'Inital Setup',205,0,'','Baddestinationemailaddress',0,'InvlaidaddressbouncefromseveralsmallISPs','suppress','Create'),
  (421, 421,1,'Inital Setup',206,501,'5.5.4','InvalidAddress',0,'InvalidaddressresponsefromseveralsmallISPs','suppress','Create'),
  (422, 422,1,'Inital Setup',207,0,'','nomailboxherebythatname',0,'invalidaddressfromseveralsmallISPs','suppress','Create'),
  (423, 423,1,'Inital Setup',208,550,'5.7.1','Requestedactionnottaken',0,'InvalidaddressresponsefromseveralsmallISPs','suppress','Create'),
  (424, 424,1,'Inital Setup',209,550,'','Unabletodeliverto',0,'invalidaddressresponsefromseveralsmallISPs','suppress','Create'),
  (425, 425,1,'Inital Setup',210,553,'5.7.1','Nosuchuserhere',0,'InvalidaddressresponsefromseveralsmallISPs','suppress','Create'),
  (426, 426,1,'Inital Setup',211,451,'4.3.5','serverconfigurationproblem',0,'temporaryproblemmessagewilllikelydeliver','retry','Create'),
  (427, 427,1,'Inital Setup',212,451,'','Youhavebeengreylisted.',0,'Temporarygreylistinglikelymatchingsomeotherbouncerule.','retry','Create'),
  (429, 429,1,'Inital Setup',213,421,'','AntispamdaSoftSellimplementagreylist',0,'TemporaryGreylisting','retry','Create'),
  (430, 430,1,'Inital Setup',214,550,'','Domainfrequencylimited',0,'qq.comreputationbaseddeferral','retry','Create'),
  (431, 431,1,'Inital Setup',215,552,'','RCPTTO.*Mail.*',0,'','no_action','Create'),
  (432, 432,1,'Inital Setup',216,550,'','Maximumlinelengthexceeded',0,'Misused550response.Thisshouldbeablocknotabounce.','no_action','Create'),
  (433, 433,1,'Inital Setup',217,421,'4.7.1','willbepermanentlydeferred',0,'yahoopermanentdeferral','no_action','Create'),
  (436, 436,1,'Inital Setup',218,550,'','invalidDNSMXorA/AAAAresourcerecord',0,'DNSissue.LikelymissingMXrecord.','no_action','Create'),
  (438, 438,1,'Inital Setup',219,550,'','Theemailsentfrom',0,'MissingreqiredDNSrecords','no_action','Create'),
  (439, 439,1,'Inital Setup',220,550,'','Domainofsenderaddress',0,'MissingreqiredDNSrecords.','no_action','Create'),
  (440, 440,1,'Inital Setup',221,550,'','Verificationfailedfor',0,'MissingrequiredDNSrecords.','no_action','Create'),
  (441, 441,1,'Inital Setup',222,553,'','unabletoverifyaddress',0,'MissingrequiredDNSrecords','no_action','Create'),
  (442, 442,1,'Inital Setup',223,550,'','senderrejected',0,'Senderblacklistedormissingrequireddnsrecords','no_action','Create'),
  (443, 443,1,'Inital Setup',224,550,'','Senderemailaddressrejected',0,'Blockedforunknownreason','no_action','Create'),
  (444, 444,1,'Inital Setup',225,550,'','SenderIPreverselookuprejected',0,'MissingsomeDNSrecords','no_action','Create'),
  (445, 445,1,'Inital Setup',226,553,'','yourenvelopesenderdomain',0,'Missingrequireddnsrecords','no_action','Create'),
  (446, 446,1,'Inital Setup',227,550,'','invalidsender',0,'Senderlikelymissingrequireddnsrecords','no_action','Create'),
  (447, 447,1,'Inital Setup',228,550,'','Tryagainlater',0,'Temporarydelay.Possiblereputationissue.','retry','Create'),
  (448, 448,1,'Inital Setup',229,452,'','senderrejected',0,'Senderrejected.Shouldnotbeadeferral.Misused452','no_action','Create'),
  (449, 449,1,'Inital Setup',230,550,'','Toomanyconnect',0,'ToomanyconnectionstoSinaaddress','retry','Create'),
  (450, 450,1,'Inital Setup',231,550,'','Toomanyinvalid',0,'Throttlingforbouncesortoomuchmail','retry','Create'),
  (451, 451,1,'Inital Setup',232,550,'','Toomanyrecipients',0,'Messagegoingtotoomanyrecipientsunabletodeliver','no_action','Create'),
  (452, 452,1,'Inital Setup',233,554,'','usernofound',0,'InvalidrecipientfromSina','suppress','Create'),
  (455, 455,1,'Inital Setup',234,550,'5.7.1','unknownorillegaluser',0,'Hardbouncefromsomeuniversity','suppress','Create'),
  (456, 456,1,'Inital Setup',235,550,'','csi.cloudmark.com',0,'CloudmarkBlocking','no_action','Create'),
  (457, 457,1,'Inital Setup',236,550,'5.1.0','Recipientrejected',0,'InvalidEmail','suppress','Create'),
  (458, 458,1,'Inital Setup',237,451,'4.5.1','Messagegreylisted',0,'greylisting.shouldretry.','retry','Create'),
  (459, 459,1,'Inital Setup',238,550,'','doesnotpass',0,'SPFfailure.Shouldbeno_action','no_action','Create'),
  (460, 460,1,'Inital Setup',239,554,'','HL:ITC',0,'ChineseISPdeferralfortoomanyconnections','retry','Create'),
  (461, 461,1,'Inital Setup',240,550,'5.7.51','',0,'Office365configurationproblem.Notinvalidaddress','no_action','Create'),
  (462, 462,1,'Inital Setup',241,550,'5.7.511','',0,'Office365blocking.NotaninvalidAddress','no_action','Create'),
  (463, 463,1,'Inital Setup',242,0,'','connectionclosed',1,'Unknownconnectionclosure.primarilycomcast.net.alwaysinSTARTTLS','no_action','Create'),
  (464, 464,1,'Inital Setup',243,550,'','HELOhosthost_mismatch',0,'','no_action','Create'),
  (465, 465,1,'Inital Setup',244,550,'','rejectedasspam',0,'blockedforspamreasons','no_action','Create'),
  (467, 467,1,'Inital Setup',245,550,'5.7.1','messagerefused',1,'Aparentspamfilteringblockingfromanumberofsmalldomains','no_action','Create'),
  (468, 468,1,'Inital Setup',246,553,'5.1.2','Thesenderaddress',0,'Unknownsenderaddresserror','no_action','Create'),
  (469, 469,1,'Inital Setup',247,550,'','spf-checkfailed',0,'','no_action','Create'),
  (470, 470,1,'Inital Setup',248,550,'5.7.1','spfpolicy',0,'','no_action','Create'),
  (471, 471,1,'Inital Setup',249,554,'','transactionfailed',0,'','no_action','Create'),
  (472, 472,1,'Inital Setup',250,553,'5.3.0','Theemailaccountthatyoutriedtoreachdoesnotexist',0,'','suppress','Create'),
  (473, 473,1,'Inital Setup',251,550,'','nosuchuser',0,'','suppress','Create'),
  (474, 474,1,'Inital Setup',252,554,'','nosuchuser',0,'','suppress','Create'),
  (475, 475,1,'Inital Setup',253,553,'','nosuchuser',0,'','suppress','Create'),
  (476, 476,1,'Inital Setup',254,450,'4.2.0','greylist',1,'greylisting.shouldretry.','retry','Create'),
  (477, 477,1,'Inital Setup',255,550,'','Requestedactionnottaken:mailboxunavailableRejectduetopolicyrestrictions.Forexplanationvisit',0,'GMX.com','no_action','Create'),
  (478, 478,1,'Inital Setup',256,550,'','verifyfailedduetoroute',0,'Smallaudomainblocking','no_action','Create'),
  (479, 479,1,'Inital Setup',257,550,'','JunkMailrejected',0,'BlockedbysomeRBL.UsuallySpamCoporBarracuda','no_action','Create'),
  (480, 480,1,'Inital Setup',258,450,'4.7.0','Youremailisbeingverified',0,'Weirdemail/reputationverificationdeferral','retry','Create'),
  (481, 481,1,'Inital Setup',259,550,'','RejectingforSenderPolicyFramework',0,'MiscSPFpolicyrejection.Shouldbeblocknotbounce.','no_action','Create'),
  (482, 482,1,'Inital Setup',260,450,'4.1.1','userunknown',0,'malformedresponseusingthewrongcode.shouldbehardbounce.','suppress','Create'),
  (483, 483,1,'Inital Setup',261,550,'5.7.1','relayingdenied',2,'unknownblock.closematchtorule423.Hencethepriorityof2.','no_action','Create'),
  (484, 484,1,'Inital Setup',262,550,'','subjectcontainsinvalidcharacters',0,'Miscbounceforbadsubjectevenwhensubjectlooksgood.','no_action','Create'),
  (485, 485,1,'Inital Setup',263,550,'5.4.317','Messageexpired',0,'unknownasynchronousresponsefromsomeo365domains','no_action','Create'),
  (486, 486,1,'Inital Setup',264,550,'','RejectedbyheaderbasedAnti-Spoofingpolicy',0,'Mimecastanti-spoofingresponse','no_action','Create'),
  (487, 487,1,'Inital Setup',265,554,'','Undeliverablelocalhost',0,'SendGirdgeneratedresponseformessagessenttodomainswithlocalhostintheMXrecord.','no_action','Create'),
  (488, 488,1,'Inital Setup',266,450,'4.1.1','Recipientaddressrejected',5,'Malformedresponsefromabunchofsmalldomains.','suppress','Create'),
  (490, 490,1,'Inital Setup',267,450,'4.7.3','Organizationqueuequota',0,'MysteriousMicrosoftResponse','retry','Create'),
  (491, 491,1,'Inital Setup',268,550,'','Unroutablesenderaddress',0,'mail.rublockfornomxrecord(usuallywhitelabelrelated)','no_action','Create'),
  (492, 492,1,'Inital Setup',269,550,'','relaynotpermitted',0,'unknownblock.likelyreputationbased.','no_action','Create'),
  (493, 493,1,'Inital Setup',270,550,'','attachmentshere',0,'attachmentsnotallowed.','no_action','Create'),
  (494, 494,1,'Inital Setup',271,550,'','wedonotacceptmailfromthisaddress',0,'unknownblocking','no_action','Create'),
  (495, 495,1,'Inital Setup',272,451,'4.7.650','hasbeentemporarily',0,'Microsoftdeferralthatstartedonapril4th2018','retry','Create'),
  (496, 496,1,'Inital Setup',273,451,'4.7.651','hasbeentemporarily',0,'Microsoftdeferralthatstartedonapril4th2018','retry','Create'),
  (497, 497,1,'Inital Setup',274,554,'','unknownrecipient',0,'unknownrecipientshouldbeahardbounce','suppress','Create'),
  (498, 498,1,'Inital Setup',275,0,'','recipientaddressrejected:userunknown',0,'Alotofresponseswiththisstring.Unfortunatelyahugenumberofresponseandenhancedcodes.','suppress','Create'),
  (499, 499,1,'Inital Setup',276,550,'','fromcontainsinvalidcharacters',0,'fromaddresscontainsinvalidcharacters','no_action','Create'),
  (500, 500,1,'Inital Setup',277,450,'4.7.0','deferred',0,'proofpointdeferral','retry','Create'),
  (501, 501,1,'Inital Setup',278,450,'4.1.1','Addressverificationinprogress',1,'Failureisapparentlyrelatedtotherecipientaddressnotbeingverified.Lotsofengagementontheseaddresses.shouldnotsuppress','retry','Create'),
  (502, 502,1,'Inital Setup',279,550,'','erroroccurredfetchingthematchesdescription',0,'strangemimecastblock','no_action','Create'),
  (503, 503,1,'Inital Setup',280,450,'4.7.0','Youremailingisbeingverified',0,'finishdeferralsecmailfiltering','retry','Create'),
  (504, 504,1,'Inital Setup',281,550,'','noMXrecordfordomain',0,'mainlylibertydomainblockseeing~50%ofaddressesengagingSGwide','no_action','Create');

  UNLOCK TABLES;

