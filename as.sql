-- MySQL dump 10.13  Distrib 5.6.23, for osx10.8 (x86_64)
--
-- Host: localhost    Database: mu77
-- ------------------------------------------------------
-- Server version	5.6.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `mu77`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `platform` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `platform`;

--
-- Table structure for table `users`
--

-- DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
/*
CREATE TABLE IF NOT EXISTS `users` (
  `uid` varchar(64) NOT NULL,
  `platform` varchar(64) NOT NULL default '',
  `channel` varchar(64) NOT NULL default '',
  `user` varchar(64) NOT NULL default '',
  `passwd` varchar(64) NOT NULL default '',
  `admin_level` tinyint(1) DEFAULT '0',
  `email` varchar(64),
  `phone` varchar(64),
  `device_id` varchar(255) NOT NULL,
  `date` datetime NOT NULL DEFAULT '1900-1-1 00:00:00',
  PRIMARY KEY (`uid`),
  KEY `user` (`user`),
  KEY `email` (`email`),
  KEY `phone` (`phone`),
  KEY `device_id` (`device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
/*!40101 SET character_set_client = @saved_cs_client */;

CREATE TABLE IF NOT EXISTS `users` (
    `uid` varchar(64) not null,
    `portrait_url` varchar(256) default '',
    `name` varchar(64) default '',
    `gender` tinyint(1) default 1,
    `email` varchar(128) default '',
    `phone` varchar(64) default '',
    `admin_level` tinyint(1) default 1,
    `create_time` datetime not null default '1900-1-1 00:00:00',
    primary key (`uid`),
    KEY `name` (`name`),
    KEY `email` (`email`),
    KEY `phone` (`phone`)
    ) default charset=utf8mb4;

create table if not exists `third_platform_user_map` (
    id int not null auto_increment,
    `platform` varchar(32) not null,
    `channel` varchar(64) default '',
    `third_openid` varchar(64) default '',
    `third_unionid` varchar(64) default '',
    `third_token` varchar(256) default '',
    `device_id` varchar(256) not null,
    `uid` varchar(64) not null,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_openid` (`platform`, `third_openid`),
    KEY `third_openid` (`third_openid`),
    KEY `third_unionid` (`third_unionid`),
    KEY `channel` (`channel`),
    KEY `device_id` (`device_id`)
    ) default charset=utf8mb4;
--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-10-09 16:07:24
