-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	5.7.27-0ubuntu0.18.04.1

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

CREATE DATABASE IF NOT EXISTS `test` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `test`;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `record_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `lastname` varchar(45) DEFAULT NULL,
  `gender` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`record_id`)
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Mario','Neri','M'),(2,'Mario','Neri','M'),(3,'Mario','Neri','M'),(4,'Mario','Neri','M'),(5,'Mario','Neri','M'),(6,'Mario','Neri','M'),(7,'Mario','Rossi','M'),(8,'Mario','Neri',NULL),(9,'Mario','Neri','M'),(10,'Mario','Neri','M'),(11,'Mario','Neri','M'),(12,'Mario','Neri','M'),(13,'Mario','Neri','M'),(14,'Mario','Neri','M'),(15,'Mario','Neri','M'),(16,'Mario','Neri','M'),(17,'Marco','Rossi','M'),(18,'Marco','Rossi','M'),(19,'Marco','Rossi','M'),(20,'Marco','Rossi','M'),(21,'Marco','Rossi','M'),(22,'Marco','Rossi','M'),(23,'Marco','Rossi','M'),(24,'Marco','Rossi','M'),(25,'Marco','Rossi','M'),(26,'Marco','Rossi','M'),(27,'Marco','Rossi','M'),(28,'Marco','Rossi','M'),(29,'Marco','Rossi','M'),(30,'Marco','Rossi','M'),(31,'Marco','Rossi','M'),(32,'Marco','Rossi','M'),(33,'Marco','Rossi','M'),(34,'Marco','Bossi','M'),(35,'Marco','Rossi','M'),(36,'Marco','Rossi','M'),(37,'Marco','Rossi','M'),(38,'Marco','Rossi','M'),(39,'Marco','Rossi','M'),(40,'Marco','Rossi','M'),(41,'Marco','Rossi','M'),(42,'Marco','Rossi','M'),(43,'Marco','Rossi','M'),(44,'Marco','Rossi','M'),(45,'Marco','Rossi','M'),(46,'Marco','Rossi','M'),(47,'Marco','Rossi','M'),(48,'Marco','Rossi','M'),(49,'Mario','Rossi',NULL),(50,'Mario','Rossi',NULL),(51,'Mario','Rossi',NULL),(52,'Mario','Rossi',NULL),(53,'Mario','Rossi',NULL),(54,'Mario','Rossi',NULL),(55,'Mario','Rossi',NULL),(56,'Mario','Rossi',NULL),(57,'Marco','Rossi','M'),(58,'Marco','Rossi','M'),(61,'Marco','Rossi','M'),(69,'Marco','Rossi','M'),(70,'Marco','Rossi','M'),(71,'Marco','Rossi','M'),(72,'Marco','Rossi','M'),(73,'Marco','Rossi','M'),(74,'Marco','Rossi','M'),(75,'Marco','Rossi','M'),(76,'Mario','Rossi',NULL),(77,'Marco','Rossi','M'),(78,'Marco','Rossi','M'),(79,'Mario','Rossi',NULL),(80,'Marco','Rossi','M'),(81,'Marco','Rossi','M');
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

-- Dump completed on 2019-08-28 18:10:31
