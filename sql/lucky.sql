-- MySQL dump 10.13  Distrib 8.0.22, for macos10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: lucky
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `desire`
--

DROP TABLE IF EXISTS `desire`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `desire` (
  `id` int NOT NULL AUTO_INCREMENT,
  `desire` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `state` int NOT NULL,
  `type` int NOT NULL,
  `wishman_qq` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `wishman_wechat` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `wishman_tel` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `wishman_name` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `creat_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `light_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `desire`
--

LOCK TABLES `desire` WRITE;
/*!40000 ALTER TABLE `desire` DISABLE KEYS */;
INSERT INTO `desire` VALUES (6,'Lorem nostrud dolore',0,0,'mollit Ut incididunt in voluptate','commodo do velit veniam anim','13859835257','如战毛亲区','2021-04-16 18:50:09','2021-04-16 18:50:09'),(7,'veniam Lorem tempor dolor',0,0,'Ut commodo laborum','ipsum irure Duis eu','18183871354','价支厂根和千','2021-04-16 18:50:12','2021-04-16 18:50:12'),(8,'nulla exercitation aute occaecat culpa',0,0,'dolore esse dolor sunt','adipisicing','13244716412','入党构','2021-04-16 18:50:13','2021-04-16 18:50:13'),(9,'et quis aliquip esse',0,0,'nostrud reprehenderit in','culpa est','13281328748','四好向南领其','2021-04-16 18:50:16','2021-04-16 18:50:16');
/*!40000 ALTER TABLE `desire` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT,
  `desire_id` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `student_number` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `password` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `email` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `tel` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `wechat` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `school` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'12345','123456','1@1.com','123344','21213','whut');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_desire`
--

DROP TABLE IF EXISTS `user_desire`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_desire` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `desire_id` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_desire`
--

LOCK TABLES `user_desire` WRITE;
/*!40000 ALTER TABLE `user_desire` DISABLE KEYS */;
INSERT INTO `user_desire` VALUES (1,'1','6'),(2,'1','7'),(3,'1','8'),(4,'1','9');
/*!40000 ALTER TABLE `user_desire` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-04-18 10:52:48
