-- MySQL dump 10.14  Distrib 5.5.37-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: whoIsInTheLab
-- ------------------------------------------------------
-- Server version	5.5.37-MariaDB-log

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
-- Table structure for table `who_blacklist`
--

DROP TABLE IF EXISTS `who_blacklist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `who_blacklist` (
  `blacklist_id` int(6) NOT NULL AUTO_INCREMENT,
  `blacklist_MAC` varchar(17) COLLATE cp1251_bulgarian_ci NOT NULL,
  `blacklist_comment` varchar(100) COLLATE cp1251_bulgarian_ci NOT NULL,
  PRIMARY KEY (`blacklist_id`),
  UNIQUE KEY `blacklist_MAC` (`blacklist_MAC`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=cp1251 COLLATE=cp1251_bulgarian_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `who_devices`
--

DROP TABLE IF EXISTS `who_devices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `who_devices` (
  `device_id` int(10) NOT NULL AUTO_INCREMENT,
  `device_MAC` varchar(17) COLLATE cp1251_bulgarian_ci NOT NULL,
  `device_uid` int(6) NOT NULL,
  `device_comment` varchar(100) COLLATE cp1251_bulgarian_ci NOT NULL,
  PRIMARY KEY (`device_id`),
  UNIQUE KEY `device_MAC` (`device_MAC`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=cp1251 COLLATE=cp1251_bulgarian_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `who_history`
--

DROP TABLE IF EXISTS `who_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `who_history` (
  `history_id` int(11) NOT NULL AUTO_INCREMENT,
  `history_MAC` varchar(17) COLLATE cp1251_bulgarian_ci NOT NULL,
  `history_to` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `history_from` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`history_id`)
) ENGINE=InnoDB AUTO_INCREMENT=36603 DEFAULT CHARSET=cp1251 COLLATE=cp1251_bulgarian_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `who_online`
--

DROP TABLE IF EXISTS `who_online`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `who_online` (
  `online_id` int(6) NOT NULL AUTO_INCREMENT,
  `online_MAC` varchar(17) COLLATE cp1251_bulgarian_ci NOT NULL,
  `online_IP` varchar(30) COLLATE cp1251_bulgarian_ci NOT NULL,
  `online_since` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `online_last` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`online_id`),
  UNIQUE KEY `online_MAC` (`online_MAC`,`online_IP`),
  KEY `online_last` (`online_last`)
) ENGINE=InnoDB AUTO_INCREMENT=552 DEFAULT CHARSET=cp1251 COLLATE=cp1251_bulgarian_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `who_users`
--

DROP TABLE IF EXISTS `who_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE IF NOT EXISTS `who_users` (
  `user_id` int(6) NOT NULL AUTO_INCREMENT,
  `user_name1` varchar(40) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_name2` varchar(40) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_twitter` varchar(100) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_facebook` varchar(100) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_tel` varchar(16) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_email` varchar(100) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_google_plus` varchar(100) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_website` varchar(200) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_fstoken` varchar(255) COLLATE cp1251_bulgarian_ci NOT NULL,
  `user_fscheckin` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=cp1251 COLLATE=cp1251_bulgarian_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-06-05 14:58:45
