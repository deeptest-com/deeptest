/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50714
 Source Host           : localhost
 Source Database       : events

 Target Server Type    : MySQL
 Target Server Version : 50714
 File Encoding         : utf-8

 Date: 09/15/2016 12:13:00 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `clt_company`
-- ----------------------------
DROP TABLE IF EXISTS `clt_company`;
CREATE TABLE `clt_company` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `admin_id` bigint(20) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_rslfg5i8tu243j9jmkpjo6yv9` (`admin_id`),
  CONSTRAINT `FK_rslfg5i8tu243j9jmkpjo6yv9` FOREIGN KEY (`admin_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `clt_company`
-- ----------------------------
BEGIN;
INSERT INTO `clt_company` VALUES ('-1', '2016-08-27 09:58:03', b'0', b'0', '2016-08-27 09:58:12', '1', '工业园区 高和路79号', null, '苏州', 'mobiu.com', '江苏', null);
COMMIT;

-- ----------------------------
--  Table structure for `evt_around`
-- ----------------------------
DROP TABLE IF EXISTS `evt_around`;
CREATE TABLE `evt_around` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `creator_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_li16rcv3iofpeio5negfxc2ck` (`creator_id`),
  KEY `FK_9dsf853ti27237d2e1xvqj09c` (`event_id`),
  CONSTRAINT `FK_9dsf853ti27237d2e1xvqj09c` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_li16rcv3iofpeio5negfxc2ck` FOREIGN KEY (`creator_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_around`
-- ----------------------------
BEGIN;
INSERT INTO `evt_around` VALUES ('-6', null, b'0', b'0', null, null, null, '<p>\n    上海主题公园和展馆众多，形式多样。上海科技馆、杜莎夫人蜡像馆、上海野生动物园和锦江乐园等是小朋友们的最爱，寓教于乐都在其中；\n</p>\n<p>\n    上海博物馆、上海图书馆等则是增长学识的好去处；而去共青森林公园烧烤、去崇明岛东滩候鸟保护区观鸟则是年轻人的首选。\n</p>', '-1', null, 'entertainment'), ('-5', null, b'0', b'0', null, null, null, '<p>\n    上海的城隍庙始建于明代永乐年间，老城隍庙与豫园毗邻，位于方浜中路，东至安仁街，北通福佑路，西至旧校场路，经营各种上海风味<span style=\"white-space: normal;\">知名</span>小吃。\n</p>\n<p>\n    城隍庙美食街，可称得上是小吃的王国，绿波廊的特色点心、松月楼的素菜包、松云楼的八宝饭、南翔小笼、宁波汤团和酒酿圆子等，都是游上海不可错过的美味。\n</p>', '-1', null, 'food'), ('-4', null, b'0', b'0', null, null, null, '“香格里拉”是香港上市公司香格里拉（亚洲）有限公司的品牌，该酒店集团隶属于马来西亚著名华商--“糖王”郭鹤年的郭氏集团旗下。\n香格里拉的名称来自詹姆斯·希尔顿的小说《失落的地平线》里，在中国西藏群山中的世外桃源。', '-1', null, 'accommodation'), ('-3', null, b'0', b'0', null, null, null, '<p>\n    浦东香格里拉大酒店地理位置无与伦比，便捷通达虹桥及浦东国际机场。\n</p>\n<p>\n    毗邻着名地标性建筑、旅游景点、餐饮及购物商区，是商旅人士在上海下榻的至优之选。\n</p>\n<p>\n    与多条贯穿浦东及浦西的交通枢纽、隧道及地铁相接，轻松通达火车站，为远途客人提供无上便捷。\n</p>', '-1', null, 'transportation'), ('-2', null, b'0', b'0', null, null, null, '<p>\n    上海迪士尼乐园，是中国内地首座迪士尼主题乐园，位于上海市浦东新区川沙新镇，于2016年6月16日正式开园。\n</p>\n<p>\n    乐园拥有六大主题园区：米奇大街、奇想花园、探险岛、宝藏湾、明日世界、梦幻世界；两座主题酒店：上海迪士尼乐园酒店、玩具总动员酒店；一座地铁站：迪士尼站；并有许多全球首发游乐项目、精彩的现场演出和多种奇妙体验任何年龄段都能在这里收获快乐。\n</p>', '-1', null, 'tour'), ('-1', null, b'0', b'0', null, null, null, '\n<p>\n    国金中心：顶尖品牌打造顶级生活\n</p>\n<p>\n    静安嘉里中心：一大波肌肉男模来袭\n</p>\n<p>\n    K11：“文艺青年”的艺术天堂\n</p>\n<p>\n    港汇恒隆广场：给你惬意生活感受\n</p>\n<p>\n    IAPM：白天逛不够“夜场”来补\n</p>', '-1', null, 'shopping');
COMMIT;

-- ----------------------------
--  Table structure for `evt_business_card`
-- ----------------------------
DROP TABLE IF EXISTS `evt_business_card`;
CREATE TABLE `evt_business_card` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `dist_id` bigint(20) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `src_id` bigint(20) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_g2swt1jsuuc7nm39x3pfxlv02` (`dist_id`),
  KEY `FK_7wbhdul18e293dx90r3c9ki3y` (`src_id`),
  KEY `FK_9usao5qfjdy6beq6xf9wknp5i` (`event_id`),
  CONSTRAINT `FK_7wbhdul18e293dx90r3c9ki3y` FOREIGN KEY (`src_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_9usao5qfjdy6beq6xf9wknp5i` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_g2swt1jsuuc7nm39x3pfxlv02` FOREIGN KEY (`dist_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_business_card_exchange`
-- ----------------------------
DROP TABLE IF EXISTS `evt_business_card_exchange`;
CREATE TABLE `evt_business_card_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `dist_id` bigint(20) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `src_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_69h402gfqrdhdvdonmb5hpvjk` (`dist_id`),
  KEY `FK_qg2wl667724vnsxit93acem7o` (`event_id`),
  KEY `FK_tudbtfh22jl125q4uj41u57a` (`src_id`),
  CONSTRAINT `FK_69h402gfqrdhdvdonmb5hpvjk` FOREIGN KEY (`dist_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_qg2wl667724vnsxit93acem7o` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_tudbtfh22jl125q4uj41u57a` FOREIGN KEY (`src_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_client`
-- ----------------------------
DROP TABLE IF EXISTS `evt_client`;
CREATE TABLE `evt_client` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `company_id` bigint(20) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_oly9fts03llk1j4kmxjac1o9b` (`company_id`),
  CONSTRAINT `FK_oly9fts03llk1j4kmxjac1o9b` FOREIGN KEY (`company_id`) REFERENCES `clt_company` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_client`
-- ----------------------------
BEGIN;
INSERT INTO `evt_client` VALUES ('-1', '2016-08-27 09:57:46', b'0', b'0', '2016-08-27 09:57:50', '1', '-1', '462826@qq.com', 'Aaron Chen', '18626203266', 'CEO', '123456');
COMMIT;

-- ----------------------------
--  Table structure for `evt_document`
-- ----------------------------
DROP TABLE IF EXISTS `evt_document`;
CREATE TABLE `evt_document` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `msg` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `doc_type` varchar(255) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_gcvsidwip566v9022oxt0krqb` (`author_id`),
  KEY `FK_1mrhjjs9ajm36b295pd1g9vde` (`event_id`),
  CONSTRAINT `FK_1mrhjjs9ajm36b295pd1g9vde` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_gcvsidwip566v9022oxt0krqb` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_document`
-- ----------------------------
BEGIN;
INSERT INTO `evt_document` VALUES ('-6', '2016-08-29 09:09:58', b'0', b'0', '2016-08-29 09:10:01', '1', '-1', '', '-1', '引领时尚潮流', 'banner', 'upload/sample/document/ad01.jpg'), ('-5', '2016-08-29 09:09:58', b'0', b'0', '2016-08-29 09:10:01', '1', '-1', '', '-1', '引领时尚潮流', 'banner', 'upload/sample/document/ad02.jpg'), ('-4', '2016-08-29 09:09:58', b'0', b'0', '2016-08-29 09:10:01', '1', '-1', '', '-1', '引领时尚潮流', 'banner', 'upload/sample/document/ad03.jpg'), ('-3', '2016-08-29 09:09:58', b'0', b'0', '2016-08-29 09:10:01', '1', '-1', '', '-1', '探访时尚前沿', 'banner', 'upload/sample/document/ad04.jpg'), ('-2', '2016-08-29 09:09:58', b'0', b'0', '2016-08-29 09:10:01', '1', '-1', '', '-1', '把握流行趋势', 'banner', 'upload/sample/document/ad05.jpg'), ('-1', '2016-08-29 09:09:58', b'0', b'0', '2016-08-29 09:10:01', '1', '-1', '', '-1', '引领时尚潮流', 'banner', 'upload/sample/document/ad06.jpg');
COMMIT;

-- ----------------------------
--  Table structure for `evt_event`
-- ----------------------------
DROP TABLE IF EXISTS `evt_event`;
CREATE TABLE `evt_event` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `company_id` bigint(20) DEFAULT NULL,
  `contact` varchar(255) DEFAULT NULL,
  `creator_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `register_end_time` datetime DEFAULT NULL,
  `register_start_time` datetime DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `fax` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `qq` varchar(255) DEFAULT NULL,
  `wechat` varchar(255) DEFAULT NULL,
  `weibo` varchar(255) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `has_parallel_sessin` bit(1) DEFAULT NULL,
  `sponsor` varchar(255) DEFAULT NULL,
  `place` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_fgoyfc3vkqdi3towl6vwvv2l` (`company_id`),
  KEY `FK_o6uvxdk11le74m6wpw5i2wejb` (`creator_id`),
  CONSTRAINT `FK_fgoyfc3vkqdi3towl6vwvv2l` FOREIGN KEY (`company_id`) REFERENCES `clt_company` (`id`),
  CONSTRAINT `FK_o6uvxdk11le74m6wpw5i2wejb` FOREIGN KEY (`creator_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_event`
-- ----------------------------
BEGIN;
INSERT INTO `evt_event` VALUES ('-1', '2016-08-29 08:45:31', b'0', b'0', '2016-08-29 08:45:34', '1', '长寿路652号F幢1001室', '上海', '-1', 'Aaron Chen', '-1', '上海时装周聚焦全球对优秀原创设计师和时装品牌的关注，搭建促进产业商贸联动对接和展示交流的平台。', '2017-01-10 18:00:00', '孟非', '上海时装周 2017春夏', '', '2016-10-20 23:59:59', '2016-09-01 09:00:01', '2017-04-08 10:00:00', 'register', '', 'sifc@sifc.org.cn', '+86-021-32530455', '+86-021-32530463', null, 'shanghai_fashionweek', null, 'http://www.shanghaifashionweek.com', b'0', '上海时装周组委会办公室', '浦东香格里拉大酒店');
COMMIT;

-- ----------------------------
--  Table structure for `evt_guest`
-- ----------------------------
DROP TABLE IF EXISTS `evt_guest`;
CREATE TABLE `evt_guest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_4gh2ysmq5qswyc3lode6ks470` (`event_id`),
  CONSTRAINT `FK_4gh2ysmq5qswyc3lode6ks470` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_guest`
-- ----------------------------
BEGIN;
INSERT INTO `evt_guest` VALUES ('-5', null, b'0', b'0', null, null, 'upload/sample/guest/AlexisMabille.jpg', null, null, '-1', 'Alexis Mabille', null, '法国设计师', '无论是成衣设计还是高级定制，Alexis Mabille的时装都令人过目难忘。VOGUE对法国设计师Alexis Mabille进行了专访，听他讲一杯鸡尾酒是如何启发其2016春夏系列的。'), ('-4', null, b'0', b'0', null, null, 'upload/sample/guest/BhubawitKritpholnara.jpg', null, null, '-1', 'Bhubawit Kritpholnara', null, 'ISSUE创意总监', '无论是成衣设计还是高级定制，Alexis Mabille的时装都令人过目难忘。VOGUE对法国设计师Alexis Mabille进行了专访，听他讲一杯鸡尾酒是如何启发其2016春夏系列的。'), ('-3', null, b'0', b'0', null, null, 'upload/sample/guest/BoraAksu.jpg', null, null, '-1', 'Bora Aksu', null, '土耳其的设计师', '无论是成衣设计还是高级定制，Alexis Mabille的时装都令人过目难忘。VOGUE对法国设计师Alexis Mabille进行了专访，听他讲一杯鸡尾酒是如何启发其2016春夏系列的。'), ('-2', null, b'0', b'0', null, null, 'upload/sample/guest/CortoMoltedo.jpg', null, null, '-1', 'Corto Moltedo', null, '奢华包袋品牌Corto Moltedo创始人', '无论是成衣设计还是高级定制，Alexis Mabille的时装都令人过目难忘。VOGUE对法国设计师Alexis Mabille进行了专访，听他讲一杯鸡尾酒是如何启发其2016春夏系列的。'), ('-1', null, b'0', b'0', null, null, 'upload/sample/guest/TomasMaier.jpg', null, null, '-1', 'Tomas Maier', null, 'Bottega Veneta创意总监', '无论是成衣设计还是高级定制，Alexis Mabille的时装都令人过目难忘。VOGUE对法国设计师Alexis Mabille进行了专访，听他讲一杯鸡尾酒是如何启发其2016春夏系列的。');
COMMIT;

-- ----------------------------
--  Table structure for `evt_news`
-- ----------------------------
DROP TABLE IF EXISTS `evt_news`;
CREATE TABLE `evt_news` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `link` varchar(255) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_cjshg5ib1xy2419nlxkf48xur` (`author_id`),
  KEY `FK_6yykm1py80omkgcan8kif56yr` (`event_id`),
  CONSTRAINT `FK_6yykm1py80omkgcan8kif56yr` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_cjshg5ib1xy2419nlxkf48xur` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_notice`
-- ----------------------------
DROP TABLE IF EXISTS `evt_notice`;
CREATE TABLE `evt_notice` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `link` varchar(255) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_kd11alalg4xbbjprsoj6blqy2` (`author_id`),
  KEY `FK_9triihtti7whvjv2kgaof392g` (`event_id`),
  CONSTRAINT `FK_9triihtti7whvjv2kgaof392g` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_kd11alalg4xbbjprsoj6blqy2` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_organizer`
-- ----------------------------
DROP TABLE IF EXISTS `evt_organizer`;
CREATE TABLE `evt_organizer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_nad12bpe3psf8uxqnpv4jjutv` (`event_id`),
  CONSTRAINT `FK_nad12bpe3psf8uxqnpv4jjutv` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_organizer`
-- ----------------------------
BEGIN;
INSERT INTO `evt_organizer` VALUES ('-2', null, b'0', b'0', null, null, '-1', '上海服装设计协会', 'co_organizer', null), ('-1', null, b'0', b'0', null, null, '-1', '上海国际服装服饰中心', 'co_organizer', null);
COMMIT;

-- ----------------------------
--  Table structure for `evt_qa`
-- ----------------------------
DROP TABLE IF EXISTS `evt_qa`;
CREATE TABLE `evt_qa` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_ske0edj3rht1u5uqf4ueeruym` (`author_id`),
  KEY `FK_7775985cg1cdhddb07xtned4h` (`event_id`),
  KEY `FK_o6j8mpbaqpf8txvyptmo4wcu9` (`parent_id`),
  CONSTRAINT `FK_7775985cg1cdhddb07xtned4h` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_o6j8mpbaqpf8txvyptmo4wcu9` FOREIGN KEY (`parent_id`) REFERENCES `evt_qa` (`id`),
  CONSTRAINT `FK_ske0edj3rht1u5uqf4ueeruym` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_register_record`
-- ----------------------------
DROP TABLE IF EXISTS `evt_register_record`;
CREATE TABLE `evt_register_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `client_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `register_time` datetime DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_1vbgxu39moptqd7b51btgw9ae` (`client_id`),
  KEY `FK_ncokoe3h6mcgbuy6qews51isx` (`event_id`),
  CONSTRAINT `FK_1vbgxu39moptqd7b51btgw9ae` FOREIGN KEY (`client_id`) REFERENCES `evt_guest` (`id`),
  CONSTRAINT `FK_ncokoe3h6mcgbuy6qews51isx` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_schedule_item`
-- ----------------------------
DROP TABLE IF EXISTS `evt_schedule_item`;
CREATE TABLE `evt_schedule_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `guest_id` bigint(20) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `session_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_466i3y5kl5n5vr1kd8cwutrvu` (`event_id`),
  KEY `FK_b7jmmvw06fhk8v684l2hfqv0d` (`guest_id`),
  KEY `FK_enwjwmsfcm6f5o3khets5uwoa` (`session_id`),
  CONSTRAINT `FK_466i3y5kl5n5vr1kd8cwutrvu` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_b7jmmvw06fhk8v684l2hfqv0d` FOREIGN KEY (`guest_id`) REFERENCES `evt_guest` (`id`),
  CONSTRAINT `FK_enwjwmsfcm6f5o3khets5uwoa` FOREIGN KEY (`session_id`) REFERENCES `evt_session` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_schedule_item`
-- ----------------------------
BEGIN;
INSERT INTO `evt_schedule_item` VALUES ('-9', null, b'0', b'0', null, null, 'Catwork', '2017-04-08 14:00:00', '-1', null, '2017-04-08 10:00:00', 'LANYU', null, '-1'), ('-8', null, b'0', b'0', null, null, 'Trade Show', null, '-1', null, '2017-04-08 14:30:00', 'Burberry上海服装服饰展', null, '-2'), ('-7', null, b'0', b'0', null, null, 'Event', null, '-1', null, '2017-04-08 19:30:00', 'SIFS-DIESEL', null, '-1'), ('-6', null, b'0', b'0', null, null, 'Catwork', null, '-1', null, '2017-04-09 13:30:00', 'BAN XIAOXUE', null, '-1'), ('-5', null, b'0', b'0', null, null, 'Catwork', null, '-1', null, '2017-04-09 15:00:00', 'LANNERET', null, '-2'), ('-4', null, b'0', b'0', null, null, 'Catwork', null, '-1', null, '2017-04-09 17:00:00', 'LABORON', null, '-1'), ('-3', null, b'0', b'0', null, null, 'Catwork', null, '-1', null, '2017-04-09 18:00:00', 'C.J.YAO', null, '-1'), ('-2', null, b'0', b'0', null, null, 'Catwork', null, '-1', null, '2017-04-09 19:00:00', 'YUZZO LONDON', null, '-3'), ('-1', null, b'0', b'0', null, null, 'Catwork', null, '-1', null, '2017-04-09 20:00:00', 'Prolivon普洛利文', null, '-3');
COMMIT;

-- ----------------------------
--  Table structure for `evt_service`
-- ----------------------------
DROP TABLE IF EXISTS `evt_service`;
CREATE TABLE `evt_service` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `creator_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_iqnbnrxybhrg6jpo7py4yo9hb` (`creator_id`),
  KEY `FK_qlgkooutqxsphcv7xn0qsk0hx` (`event_id`),
  CONSTRAINT `FK_iqnbnrxybhrg6jpo7py4yo9hb` FOREIGN KEY (`creator_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_qlgkooutqxsphcv7xn0qsk0hx` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_service`
-- ----------------------------
BEGIN;
INSERT INTO `evt_service` VALUES ('-6', null, b'0', b'0', null, null, '-1', '<h3>\n    招车服务\n</h3>\n<p>\n    会议团体用车服务：021-51017070<br/>上海强生电调招车<span style=\"white-space: normal;\">：</span>021-62580000\n</p>', '-1', null, 'taxi'), ('-5', null, b'0', b'0', null, null, '-1', '<h3>\n    上海浦东香格里拉大酒店\n</h3>\n<p>\n    上海浦东香格里拉大酒店坐拥外滩及黄浦江美景，曾荣膺奖项，是上海顶级豪华酒店，提供热情周到的香格里拉特色服务及世界一流的就餐体验。\n</p>', '-1', null, 'accommodation'), ('-4', null, b'0', b'0', null, null, '-1', '<h3>\n    餐饮服务\n</h3>\n<p>\n    您既可光临翡翠36餐厅，品味现代创新菜肴，欣赏美不胜收的景观，又可停留滩万餐厅，尝尝做工精致的寿司的美味，或者光临最出色的中餐厅—桂花楼，我们诚挚邀请您享受香格里拉奉献的各种美食及殷勤周到的服务。\n</p>', '-1', null, 'food'), ('-3', null, b'0', b'0', null, null, '-1', '<h3>\n    酒店WIFI\n</h3>\n<p>\n    名称：shanghai&nbsp;\n</p>\n<p>\n    密码：12345678<br/>\n</p>\n<h3>\n    会议现场WIFI\n</h3>\n<p>\n    <span style=\"font-size: 16px;\">名称：</span><span style=\"font-size: 16px;\">meeting</span>\n</p>\n<p>\n    <span style=\"white-space: normal;\">密码：</span>abc123\n</p>', '-1', null, 'wifi'), ('-2', null, b'0', b'0', null, null, '-1', '<h3>\n    购物服务\n</h3>\n<p>\n    欢迎光临浦东香格里拉大酒店1楼商店，购买精美礼品。\n</p>\n<p>\n    酒店也提供礼品、地方特产代购服务，有意请拨打服务电话400-52012-34找陈小姐。\n</p>', '-1', null, 'shopping'), ('-1', null, b'0', b'0', null, null, '-1', '<h3>\n    打印服务\n</h3>\n<p>\n    如您需要会议打印等办公服务，请拨打我们的服务电话400-52012-34\n</p>', '-1', null, 'print');
COMMIT;

-- ----------------------------
--  Table structure for `evt_session`
-- ----------------------------
DROP TABLE IF EXISTS `evt_session`;
CREATE TABLE `evt_session` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `contact` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `fax` varchar(255) DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_sqkcig0k3tjqtdfcljdgjchcd` (`event_id`),
  CONSTRAINT `FK_sqkcig0k3tjqtdfcljdgjchcd` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `evt_session`
-- ----------------------------
BEGIN;
INSERT INTO `evt_session` VALUES ('-3', null, b'0', b'0', null, null, '上海展览中心友谊会堂', '上海', null, null, null, '-1', null, null, null, null, null), ('-2', null, b'0', b'0', null, null, '巨鹿路158号', '上海', null, null, null, '-1', null, null, null, null, null), ('-1', null, b'0', b'0', null, null, '新天地', '上海', null, null, null, '-1', null, null, null, null, null);
COMMIT;

-- ----------------------------
--  Table structure for `evt_thread`
-- ----------------------------
DROP TABLE IF EXISTS `evt_thread`;
CREATE TABLE `evt_thread` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_mnw0mv93g0t675ih01hbp14db` (`author_id`),
  KEY `FK_oibshoy3bj4mwtbm44nwt6l6n` (`event_id`),
  KEY `FK_e1eyy5a16vj80e57q2ee39dam` (`parent_id`),
  CONSTRAINT `FK_e1eyy5a16vj80e57q2ee39dam` FOREIGN KEY (`parent_id`) REFERENCES `evt_thread` (`id`),
  CONSTRAINT `FK_mnw0mv93g0t675ih01hbp14db` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_oibshoy3bj4mwtbm44nwt6l6n` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_consumer`
-- ----------------------------
DROP TABLE IF EXISTS `sys_consumer`;
CREATE TABLE `sys_consumer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `callback_url` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_device`
-- ----------------------------
DROP TABLE IF EXISTS `sys_device`;
CREATE TABLE `sys_device` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `brand` varchar(255) DEFAULT NULL,
  `chip_model` varchar(255) DEFAULT NULL,
  `chip_vender` varchar(255) DEFAULT NULL,
  `cpu` varchar(255) DEFAULT NULL,
  `dpi` int(11) DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  `host_ip` varchar(255) DEFAULT NULL,
  `keywords` varchar(255) DEFAULT NULL,
  `make` varchar(255) DEFAULT NULL,
  `model` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `operating_system` varchar(255) DEFAULT NULL,
  `pic_name` varchar(255) DEFAULT NULL,
  `popular` int(11) DEFAULT NULL,
  `port` int(11) DEFAULT NULL,
  `position` int(11) DEFAULT NULL,
  `ram` varchar(255) DEFAULT NULL,
  `screen_size` varchar(255) DEFAULT NULL,
  `sync_id` bigint(20) DEFAULT NULL,
  `sys_level` varchar(255) DEFAULT NULL,
  `sys_type` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `udid` varchar(255) DEFAULT NULL,
  `width` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_device_group`
-- ----------------------------
DROP TABLE IF EXISTS `sys_device_group`;
CREATE TABLE `sys_device_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_device_in_group`
-- ----------------------------
DROP TABLE IF EXISTS `sys_device_in_group`;
CREATE TABLE `sys_device_in_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `device_id` bigint(20) DEFAULT NULL,
  `group_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_806e597jr3jmbmt58ssmixw3b` (`device_id`),
  KEY `FK_jfefm6dxmm2for2brwc8qwvs5` (`group_id`),
  CONSTRAINT `FK_806e597jr3jmbmt58ssmixw3b` FOREIGN KEY (`device_id`) REFERENCES `sys_device` (`id`),
  CONSTRAINT `FK_jfefm6dxmm2for2brwc8qwvs5` FOREIGN KEY (`group_id`) REFERENCES `sys_device_group` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
