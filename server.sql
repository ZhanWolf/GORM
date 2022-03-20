-- MySQL dump 10.13  Distrib 5.7.35, for Linux (x86_64)
--
-- Host: localhost    Database: login
-- ------------------------------------------------------
-- Server version	5.7.35

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
-- Table structure for table `childcomment`
--

DROP TABLE IF EXISTS `childcomment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `childcomment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `from_id` int(11) DEFAULT NULL,
  `from_username` varchar(20) DEFAULT NULL,
  `content` longtext,
  `theday` date DEFAULT NULL,
  `Useful` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `childcomment_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `childcomment`
--

LOCK TABLES `childcomment` WRITE;
/*!40000 ALTER TABLE `childcomment` DISABLE KEYS */;
INSERT INTO `childcomment` VALUES (1,1,1,'asd123','确实好看','2022-01-22',1),(2,1,1,'asd123','一般般吧','2022-01-22',0),(3,1,1,'asd123','一般般吧','2022-01-22',0),(4,1,1,'asd123','一般般吧','2022-02-09',0);
/*!40000 ALTER TABLE `childcomment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `childcomment2`
--

DROP TABLE IF EXISTS `childcomment2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `childcomment2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `userid` int(11) DEFAULT NULL,
  `username` varchar(20) DEFAULT NULL,
  `content` longtext,
  `theday` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `childcomment2`
--

LOCK TABLES `childcomment2` WRITE;
/*!40000 ALTER TABLE `childcomment2` DISABLE KEYS */;
INSERT INTO `childcomment2` VALUES (1,2,1,'asd123','e','2022-02-13'),(2,3,3,'bbb11111','ok','2022-02-14');
/*!40000 ALTER TABLE `childcomment2` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `from_username` varchar(20) DEFAULT NULL,
  `from_id` int(11) DEFAULT NULL,
  `Content` longtext,
  `theday` date DEFAULT NULL,
  `usenum` int(11) DEFAULT NULL,
  `unusenum` int(11) DEFAULT NULL,
  `score` double DEFAULT NULL,
  `movie_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Comment_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,'asd123',1,'真好看','2022-01-22',1,3,4,1),(2,'asd123',1,'真不错','2022-01-22',0,0,4,1);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment2`
--

DROP TABLE IF EXISTS `comment2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `talking_id` int(11) DEFAULT NULL,
  `content` longtext,
  `username` varchar(20) DEFAULT NULL,
  `userid` int(11) DEFAULT NULL,
  `theday` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment2`
--

LOCK TABLES `comment2` WRITE;
/*!40000 ALTER TABLE `comment2` DISABLE KEYS */;
INSERT INTO `comment2` VALUES (1,0,'e','asd123',1,'2022-02-13'),(2,1,'e','asd123',1,'2022-02-13'),(3,2,'d','bbb11111',3,'2022-02-14');
/*!40000 ALTER TABLE `comment2` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movie`
--

DROP TABLE IF EXISTS `movie`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `movie` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `moviename` varchar(20) DEFAULT NULL,
  `yyear` int(11) DEFAULT NULL,
  `introduction` longtext,
  `ddate` date DEFAULT NULL,
  `posterurl` longtext,
  `length` varchar(20) DEFAULT NULL,
  `area` varchar(20) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `feature` varchar(20) DEFAULT NULL,
  `releasing` int(11) DEFAULT NULL,
  `score` double DEFAULT NULL,
  `language` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `movie_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie`
--

LOCK TABLES `movie` WRITE;
/*!40000 ALTER TABLE `movie` DISABLE KEYS */;
INSERT INTO `movie` VALUES (1,'杰伊·比姆 Jai Bhim',2021,'当一名部落男子因涉嫌盗窃而被捕时，他的妻子求助于一名人权律师以帮助伸张正义。','2021-11-02','https://sm.ms/image/a3dkRFSrHfPx8Lt','164分钟','印度','剧情','犯罪',0,2,'印度语'),(2,'误杀2',2021,'林日朗（肖央 饰）与妻子阿玲（文咏珊 饰）、儿子小虫（王昊泽 饰）一直过着清贫但幸福的生活，直到儿子小虫突发意外急需救治，几经周折，走投无路的林日朗为了救儿子决定放手一搏。他制定了一个惊天计划.....\n　　本片改编自电影《迫在眉梢》。','2021-12-17','https://sm.ms/image/qcWUvTCBJVt843r','118分钟','中国大陆','剧情','犯罪',1,4,'中文'),(3,'好想去你的世界爱你',2022,'这世上是否有个人知你冷暖，懂你悲喜？身处北京的助理建筑师安易（周依然 饰）与远在德国的调音师高昂（施柏宇 饰），因一场意外脑电波相连，从此他们听觉、味觉、触觉神奇共享。“被迫绑定”的日子，使他们方寸大乱，笑料百出，却也因为无时无刻的陪伴成为了互相最懂彼此的人...... 但一次误会使二人产生了隔阂，面对遥远的距离，未知的将来，他们是否会坚定地去往对方的世界？\n　　感谢那些看透了我们，却一直陪在我们身边的人。','2022-02-14','https://img9.doubanio.com/view/photo/l/public/p2867146164.webp','95分钟','中国大陆','喜剧','奇幻',1,0,'中文'),(4,'特送 특송',2022,'该片是一部动作犯罪电影，被称为女版《大叔》，讲述了一个女人和一个少年不择手段地将所有给钱的物品全部送出的故事，朴素谈在片中饰演负责送货的张银河。','2022-01-02','https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2799120314.webp','108分钟','韩国','动作','犯罪',0,0,'韩语');
/*!40000 ALTER TABLE `movie` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moviepic`
--

DROP TABLE IF EXISTS `moviepic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `moviepic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `url` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `moviepic_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moviepic`
--

LOCK TABLES `moviepic` WRITE;
/*!40000 ALTER TABLE `moviepic` DISABLE KEYS */;
INSERT INTO `moviepic` VALUES (1,3,'https://img2.doubanio.com/view/photo/l/public/p2694443712.webp'),(2,3,'https://img1.doubanio.com/view/photo/l/public/p2637460127.webp');
/*!40000 ALTER TABLE `moviepic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `person`
--

DROP TABLE IF EXISTS `person`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `introduction` longtext,
  `birthday` date DEFAULT NULL,
  `Constellations` varchar(20) DEFAULT NULL,
  `chinesename` varchar(20) DEFAULT NULL,
  `englishname` varchar(20) DEFAULT NULL,
  `birthplace` varchar(20) DEFAULT NULL,
  `jobs` varchar(20) DEFAULT NULL,
  `posterurl` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `person`
--

LOCK TABLES `person` WRITE;
/*!40000 ALTER TABLE `person` DISABLE KEYS */;
INSERT INTO `person` VALUES (1,'戴墨','1982-11-26','射手座','戴墨','Mo Dai','中国,辽宁,沈阳','演员 / 导演','https://img9.doubanio.com/view/personage/raw/public/74af5876a496edaa16ffe5e5024dece5.jpg'),(2,'none','1986-02-05','none','朴大民','Dae-min Park','韩国','导演','https://img1.doubanio.com/view/celebrity/raw/public/pHuk4wrQuSTwcel_avatar_uploaded1481126465.7.jpg'),(3,'So-dam Park was born on September 8, 1991 in South Korea. She is an actress, known for The Priests (2015), The Silenced (2015) and Ode to the Goose (2018).','1991-09-08','处女座','朴素丹','So-dam Park','韩国','演员/主持人/配音','https://img9.doubanio.com/view/celebrity/raw/public/p1610886913.74.jpg'),(4,'none','1970-12-26','摩羯座','宋清晨','Sae-byeok Song','韩国','演员','https://img3.doubanio.com/view/celebrity/raw/public/p13690.jpg'),(5,'内地编剧/导演,北京电影学院电影学硕士毕业,攻读导演创作及理论研究方向。曾任耀莱影视电影项目开发总监,现任恐龙影业首席内容官。先后参与《功夫之王》、《大兵小将》导演组工作以及《天将雄师》、《大约在冬季》、《我的日记》编剧、策划工作,同时一手打造 JC 电影编剧团队,开发多部成龙电影。\n','1978-03-08','双鱼座','孙琳','Lin Sun','中国，山东','制片人 / 编剧 / 副导演 / 导演','https://img2.doubanio.com/view/celebrity/raw/public/p1615529451.1.jpg'),(6,'印度电影《杰伊·比姆》的导演。','0000-01-01','none','T.J. Gnanavel','T.J. Gnanavel','none','导演','https://img3.doubanio.com/f/movie/8dd0c794499fe925ae2ae89ee30cd225750457b4/pics/movie/celebrity-default-medium.png'),(7,'苏利耶是泰米尔电影演员Sivakumar的长子。2006年和女演员Jyothika结婚。从影多年后，开创了一间“工作室绿色”负责做他自己和他的兄弟Karthi的一些电影，成为了一名电影制片人。2008年,他开始加入Agaram基金会，做各种慈善活动。他现已成为许多国际品牌到达南印度市场的形象大使。2012年开始与 STAR Vijay主持“gaming show”游戏节目，后参与主持“Neengalum Vellalam Oru Kodi”和泰米尔语引进节目“谁想成为百万富翁?”等电视节目。\n','1975-07-23','none','苏利耶·西瓦库马','Suriya Sivakumar','印度,钦奈','演员 / 编剧','https://img2.doubanio.com/view/celebrity/raw/public/p1580357897.41.jpg'),(8,'印度电影演员，导演，制片人，电视节目主持人，主要活跃在南印度电影业，也参与少数宝莱坞电影。早期从事电视行业和坎纳达语电影，直到出演泰米尔语电影才开始走红。 \n\n普拉卡什·拉贾除了母语坎纳达语，还熟练掌握泰米尔语，泰卢固语，马拉雅拉姆语，印地语和英语，是印度电影界炙手可热的黄金配角。其出演角色性格各异，造型百变，无论什么人物经过他的演绎都变的真实可信，令人印象深刻。 \n\n凭借其出色的演技，普拉卡什·拉贾获得了印度奥斯卡（Filmfare Awards South）一次最佳男演员，2次最佳男配角，2次最佳反派。国家电影奖（National Film Awards）1次最佳男主角，一次最佳男配角，2次特别奖（Special Jury Award）\n','1965-03-26','白羊座','普拉卡什·拉贾','Prakash Raj','印度,卡纳塔克邦,班加罗尔','演员 / 导演 / 制片人','https://img1.doubanio.com/view/celebrity/raw/public/p1481167938.87.jpg'),(9,'胖三井，男，中国内地编剧。本科生物技术专业，硕士在北京电影学院读导演。早年撰写推理故事和做影评节目，后从事编剧工作。编剧代表作品有武侠推理剧《侠探简不知》、古装喜剧《欢乐英雄》、电影《我们的日记》等。\n','1981-01-01','none','胖三井','Pang San Jing','none','编剧','https://img1.doubanio.com/view/celebrity/raw/public/p1588122889.59.jpg'),(10,'中国内地女演员，毕业于四川音乐学院古典舞专业，大隐光时的签约艺人，因与经纪人的解约风波引起广泛关注。2018年，周依然凭借出演喜剧电影《我是你妈》进入演艺圈。','1995-10-30','天蝎座','周依然','Yiran Zhou','中国,重庆','演员','https://img2.doubanio.com/view/personage/raw/public/79c2a0e368dce27bb441f72c5d51bd91.jpg');
/*!40000 ALTER TABLE `person` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `personpic`
--

DROP TABLE IF EXISTS `personpic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `personpic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `url` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personpic_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `personpic`
--

LOCK TABLES `personpic` WRITE;
/*!40000 ALTER TABLE `personpic` DISABLE KEYS */;
INSERT INTO `personpic` VALUES (1,10,'https://bkimg.cdn.bcebos.com/pic/8326cffc1e178a82b901c12ba848648da977381242ae?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyNzI=,g_7,xp_5,yp_5/format,f_auto'),(2,10,'https://bkimg.cdn.bcebos.com/pic/9358d109b3de9c82d158fc0632ca970a19d8bd3ee5ad?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyNzI=,g_7,xp_5,yp_5/format,f_auto');
/*!40000 ALTER TABLE `personpic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `record_act`
--

DROP TABLE IF EXISTS `record_act`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `record_act` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `personid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `record_act_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `record_act`
--

LOCK TABLES `record_act` WRITE;
/*!40000 ALTER TABLE `record_act` DISABLE KEYS */;
INSERT INTO `record_act` VALUES (2,4,3),(3,4,4),(4,1,7),(5,1,8),(6,3,10);
/*!40000 ALTER TABLE `record_act` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `record_all`
--

DROP TABLE IF EXISTS `record_all`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `record_all` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `personid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `record_all`
--

LOCK TABLES `record_all` WRITE;
/*!40000 ALTER TABLE `record_all` DISABLE KEYS */;
INSERT INTO `record_all` VALUES (1,2,1),(2,4,3),(3,4,4),(4,1,6),(5,1,7),(6,1,8),(7,3,9),(8,3,10);
/*!40000 ALTER TABLE `record_all` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `record_direct`
--

DROP TABLE IF EXISTS `record_direct`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `record_direct` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `personid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `record_direct_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `record_direct`
--

LOCK TABLES `record_direct` WRITE;
/*!40000 ALTER TABLE `record_direct` DISABLE KEYS */;
INSERT INTO `record_direct` VALUES (1,2,1),(2,4,2),(3,1,6);
/*!40000 ALTER TABLE `record_direct` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `record_script`
--

DROP TABLE IF EXISTS `record_script`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `record_script` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `personid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `record_script_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `record_script`
--

LOCK TABLES `record_script` WRITE;
/*!40000 ALTER TABLE `record_script` DISABLE KEYS */;
INSERT INTO `record_script` VALUES (1,1,6),(2,3,9);
/*!40000 ALTER TABLE `record_script` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shortcomment`
--

DROP TABLE IF EXISTS `shortcomment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shortcomment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `from_username` varchar(20) DEFAULT NULL,
  `from_id` int(11) DEFAULT NULL,
  `content` longtext,
  `theday` date DEFAULT NULL,
  `lorw` int(11) DEFAULT NULL,
  `score` double DEFAULT NULL,
  `usenum` int(11) DEFAULT NULL,
  `nouse` int(11) DEFAULT NULL,
  `movie_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shortcomment`
--

LOCK TABLES `shortcomment` WRITE;
/*!40000 ALTER TABLE `shortcomment` DISABLE KEYS */;
INSERT INTO `shortcomment` VALUES (1,'asd123',1,'想看','2022-02-07',1,8,3,2,2);
/*!40000 ALTER TABLE `shortcomment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `talking`
--

DROP TABLE IF EXISTS `talking`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `talking` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) DEFAULT NULL,
  `title` varchar(20) DEFAULT NULL,
  `content` longtext,
  `theday` date DEFAULT NULL,
  `username` varchar(20) DEFAULT NULL,
  `userid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `talking`
--

LOCK TABLES `talking` WRITE;
/*!40000 ALTER TABLE `talking` DISABLE KEYS */;
INSERT INTO `talking` VALUES (1,1,'e','e','2022-02-13','asd123',1),(2,2,'e','e','2022-02-14','bbb11111',3),(3,3,'e','e','2022-02-19','bbb11111',3);
/*!40000 ALTER TABLE `talking` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `protectionQ` varchar(20) DEFAULT NULL,
  `protectionA` varchar(20) DEFAULT NULL,
  `introduction` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'asd123','1234','我的大学','CQUPT','真开心'),(2,'asd111','1234','123','123','0'),(3,'bbb11111','1223','我的大学','CQUPT','0');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-19 17:46:19
