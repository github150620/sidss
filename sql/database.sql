# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 192.168.1.4 (MySQL 5.5.5-10.3.15-MariaDB-1)
# Database: sidss
# Generation Time: 2019-07-21 07:47:54 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table alarms
# ------------------------------------------------------------

CREATE TABLE `alarms` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `stock_id` varchar(8) NOT NULL DEFAULT '',
  `conditions` varchar(64) NOT NULL DEFAULT '',
  `sql` varchar(1024) NOT NULL DEFAULT '',
  `period` varchar(16) NOT NULL DEFAULT '',
  `status` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '0 - Not Happen; 1 - Happened',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table assets
# ------------------------------------------------------------

CREATE TABLE `assets` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `asset` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table favourites
# ------------------------------------------------------------

CREATE TABLE `favourites` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) DEFAULT NULL,
  `stock_id` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`stock_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table holdings
# ------------------------------------------------------------

CREATE TABLE `holdings` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) DEFAULT NULL,
  `stock_id` varchar(8) DEFAULT NULL,
  `shares` int(11) unsigned DEFAULT NULL,
  `cb` decimal(8,4) DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table klines
# ------------------------------------------------------------

CREATE TABLE `klines` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `stock_id` varchar(8) DEFAULT NULL,
  `date` varchar(6) DEFAULT NULL,
  `o` decimal(6,2) DEFAULT NULL,
  `c` decimal(6,2) DEFAULT NULL,
  `h` decimal(6,2) DEFAULT NULL,
  `l` decimal(6,2) DEFAULT NULL,
  `v` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stock_id` (`stock_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table ma
# ------------------------------------------------------------

CREATE TABLE `ma` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `stock_id` varchar(8) DEFAULT NULL,
  `date` varchar(6) DEFAULT NULL,
  `ma5` decimal(8,4) DEFAULT NULL,
  `ma10` decimal(8,4) DEFAULT NULL,
  `ma20` decimal(8,4) DEFAULT NULL,
  `ma30` decimal(8,4) DEFAULT NULL,
  `ma60` decimal(8,4) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stock_id` (`stock_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table prices
# ------------------------------------------------------------

CREATE TABLE `prices` (
  `stock_id` varchar(8) NOT NULL DEFAULT '',
  `price` decimal(6,2) DEFAULT 0.00,
  `chg` double DEFAULT NULL,
  `high` decimal(6,2) DEFAULT NULL,
  `low` decimal(6,2) DEFAULT NULL,
  `ttm_pe` double DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`stock_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table reports
# ------------------------------------------------------------

CREATE TABLE `reports` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `stock_id` varchar(8) DEFAULT NULL,
  `date` varchar(10) DEFAULT NULL,
  `eps` decimal(8,4) DEFAULT NULL COMMENT 'Earnings Per Share',
  `roe` decimal(5,2) DEFAULT NULL COMMENT 'Return On Equity',
  `gpr` decimal(4,2) DEFAULT NULL COMMENT 'Gross Profit Radio',
  `asset` decimal(15,2) DEFAULT NULL COMMENT 'Total Asset',
  `debt` decimal(15,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stock_id` (`stock_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table stocks
# ------------------------------------------------------------

CREATE TABLE `stocks` (
  `id` varchar(8) NOT NULL DEFAULT '',
  `name` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(32) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  `mail` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table workdays
# ------------------------------------------------------------

CREATE TABLE `workdays` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `date` varchar(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
