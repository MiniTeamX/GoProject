-- MySQL dump 10.13  Distrib 5.7.23, for Win64 (x86_64)
--
-- Host: localhost    Database: my_db
-- ------------------------------------------------------
-- Server version	5.7.23-log

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
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `comment_id` int(11) NOT NULL AUTO_INCREMENT,
  `essay_id` int(11) NOT NULL,
  `content` varchar(512) DEFAULT NULL,
  `from_id` int(11) DEFAULT NULL,
  `to_id` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`comment_id`,`essay_id`),
  UNIQUE KEY `content_UNIQUE` (`content`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,1,'写的真好',1,2,NULL),(2,2,'写的真烂',2,1,NULL);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `essay`
--

DROP TABLE IF EXISTS `essay`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `essay` (
  `essay_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) DEFAULT NULL,
  `content` text,
  `phrase_num` int(11) DEFAULT NULL,
  `comment_num` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`essay_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `essay`
--

LOCK TABLES `essay` WRITE;
/*!40000 ALTER TABLE `essay` DISABLE KEYS */;
INSERT INTO `essay` VALUES (1,'《冒险小虎队》的解密卡的原理是什么？','小时候书店故意把解密卡拿走，只有买书才能有解密卡看背后的答案。当时每天混迹于书店，没钱买书又想知道答案怎么办？一个方法是对着字硬瞅，基本上能认出几个来。但是很麻烦，而且要想象很久。后来找到一种类似丝袜那种丝网布。丝网布是相互垂直的丝线编制，线很细，中间孔比较大，不影响阅读。通过斜方向的拉扯，调节丝网的间距，直到达到一个特殊的间距（即频率和解密卡频率相同时），就可以比较轻松的阅读密码了。于是一块破布被我用了一年。。。。\n\n作者：孟诚博\n链接：https://www.zhihu.com/question/21650222/answer/484585732\n来源：知乎\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。',100,100,NULL),(2,'你看过/写过哪些有意思的代码？','作者：三级狗\n链接：https://www.zhihu.com/question/275611095/answer/407984155\n来源：知乎\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。\n\n关于外挂有很多中办法实现，我说几种主流的，大家评论里也已经多得五花八门了，1、其中很大一部分，也是绝大多是外挂的主流做法，就是直接在本机改内存。你的游戏运行在我这里，代码就得加载到我的内存里运行，那么内存里的数据再抽象，总有高手能给它鼓捣出来。就比如说这个连连看，我也可以通过读取内存的手段直接拿到它方块布局的数据，直接把这个数据全改成0，那立马就赢了。但这样做很麻烦，这里我一定要跟在座的提一下：无论是产品设计还是编码实现，一定要遵循一个核心原则：大道至简！什么叫大道至简？意思就是复杂的我真的不会。。。有一些游戏数据必须要在本地进行处理的，很容易遇到这种外挂，比如地下城与勇士无限刷图啦，更比如吃鸡，就说吃鸡，这样的第一人称射击游戏，打一枪子弹中没中，不可能放到服务器去判断，一是判断不过来，二是受网络的影响实时性根本达不到要求。所以你一枪子弹打出去中没中，一定是放在本地进行计算的，既然是在本地内存里，一旦防范不到位那就有人能给你改，我们所谓的“飞天遁地锁血金身”什么的。像LOL就不多存在这样类型的外挂，一方面肯定是反外挂投入的力度大，另一方面就是因为你的操作全部都是由服务器来进行计算并反馈的，不存在太多本地数据篡改的风险。',50,100,NULL),(3,'到目前为止，你做过的最骚的操作是什么？','作者：乔大雀\n链接：https://www.zhihu.com/question/293315890/answer/488157719\n来源：知乎\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。\n\n2017，北京地铁。某东北籍男子顶着胯蹭人一姑娘的屁股，那姑娘特好看，姑娘嚷嚷了我才发现。关键这东北男的还操着地方口音跟人继续耍横，我就走过去给这男的拎一边去。按理说北京唯一能打架的地就是地铁了，现在有带对讲机的乘务员了不好跑了，但是去年在地铁打架真是一点风险没有，打完就走没人管的，出警特慢找人特难。所以我老在地铁里揍这些摸人屁股的顶人屁股的，真不是地域歧视，让人发现还嘴硬的都是东北口音，真真的。我给这东北大哥拎一边去，按理说，应该他骂我，然后我骂他，然后他推我一把，然后我给他打躺下，我走。same story，old story。符合逻辑符合心理，揍得有理有据心悦诚服。但这大哥一张嘴我就跪了。“抄（四声）你妈咧凑外地的”我？？？？？丫在北京，在北京地铁，然后用辽宁话骂我臭外地？难就难在，我回他什么啊？这感觉回什么都像让人说中了恼羞成怒啊。我也回个操你妈？这简直就像人说“操你妈”我说“不让操”一样弱逼啊。关键这大哥说完也后悔了，一个劲的挠头，往窗外看。我也把手撒开了，我看看他他看看我，都挺尴尬的。我俩就这么不说话站了一站地，眼看下一站了，我看看他他看看我，我的尬点比较低，实在忍不住了，实在没辙没说法了，下车走了。刚刚下车，我忽然就来主意了。一拍脑门走回车厢“我抄（四声）你妈咧狗懒纸”我在北京，用东北话，骂骂我臭外地的东北人。太特么有才了！跟他对上了！这大哥一听，呜呜渣渣就走向我，然后如愿以偿的被打挺。',50,100,NULL);
/*!40000 ALTER TABLE `essay` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `user_id` bigint(15) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `nick_name` varchar(45) DEFAULT NULL,
  `gender` bigint(1) DEFAULT NULL,
  `photo_url` varchar(128) DEFAULT NULL,
  `introduction` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'11111111111','123','张三1',1,NULL,'哈哈，我是张三1'),(2,'222222222222','123','张三2',2,NULL,'哈哈，我是张三2'),(3,'33333333333','123','张三3',1,NULL,'哈哈，我是张三3'),(11,'7777777','sdafasf','张三6',1,'','哈哈，我是张三6'),(12,'7777777','sdafasf','张三6',1,'','哈哈，我是张三6'),(13,'7777777','sdafasf','张三6',1,'','哈哈，我是张三6'),(14,'88888888888','sdafasf','张三6',1,'','哈哈，我是张三6'),(15,'','','',0,'',''),(16,'','','',0,'','');
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

-- Dump completed on 2018-09-21 19:40:20
