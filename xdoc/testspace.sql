/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50714
 Source Host           : localhost
 Source Database       : testspace

 Target Server Type    : MySQL
 Target Server Version : 50714
 File Encoding         : utf-8

 Date: 03/23/2017 11:11:17 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `evt_banner`
-- ----------------------------
DROP TABLE IF EXISTS `evt_banner`;
CREATE TABLE `evt_banner` (
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
  `uri` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_o3h1m2r595nkgbmairb8mf11s` (`author_id`),
  KEY `FK_sd83oieidto98jcqdh62aekdf` (`event_id`),
  CONSTRAINT `FK_o3h1m2r595nkgbmairb8mf11s` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_sd83oieidto98jcqdh62aekdf` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_bizcard`
-- ----------------------------
DROP TABLE IF EXISTS `evt_bizcard`;
CREATE TABLE `evt_bizcard` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `client_id` bigint(20) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_coj587kdy58u259hh95il6fqa` (`client_id`),
  CONSTRAINT `FK_coj587kdy58u259hh95il6fqa` FOREIGN KEY (`client_id`) REFERENCES `evt_client` (`id`)
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
  CONSTRAINT `FK_oly9fts03llk1j4kmxjac1o9b` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `client_id` bigint(20) DEFAULT NULL,
  `msg` varchar(10000) DEFAULT NULL,
  `doc_type` varchar(255) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_rtk3nqjyv2t9rbexla6eg6phy` (`client_id`),
  KEY `FK_1mrhjjs9ajm36b295pd1g9vde` (`event_id`),
  CONSTRAINT `FK_1mrhjjs9ajm36b295pd1g9vde` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_rtk3nqjyv2t9rbexla6eg6phy` FOREIGN KEY (`client_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `country` varchar(255) DEFAULT NULL,
  `creator_id` bigint(20) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `end_datetime` datetime DEFAULT NULL,
  `fax` varchar(255) DEFAULT NULL,
  `has_parallel_sessin` bit(1) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `place` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `qq` varchar(255) DEFAULT NULL,
  `register_end_datetime` datetime DEFAULT NULL,
  `register_start_datetime` datetime DEFAULT NULL,
  `sign_before` int(11) DEFAULT NULL,
  `sign_end_datetime` datetime DEFAULT NULL,
  `sign_start_datetime` datetime DEFAULT NULL,
  `sponsor` varchar(255) DEFAULT NULL,
  `start_datetime` datetime DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `wechat` varchar(255) DEFAULT NULL,
  `weibo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_fgoyfc3vkqdi3towl6vwvv2l` (`company_id`),
  KEY `FK_o6uvxdk11le74m6wpw5i2wejb` (`creator_id`),
  CONSTRAINT `FK_fgoyfc3vkqdi3towl6vwvv2l` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`),
  CONSTRAINT `FK_o6uvxdk11le74m6wpw5i2wejb` FOREIGN KEY (`creator_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `evt_feedback`
-- ----------------------------
DROP TABLE IF EXISTS `evt_feedback`;
CREATE TABLE `evt_feedback` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `content` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `feedback_type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_5jqeagvvgywp0aasocpftbqx3` (`author_id`),
  KEY `FK_4hws6qtt6l12b1d6i3y5ojyi` (`event_id`),
  CONSTRAINT `FK_4hws6qtt6l12b1d6i3y5ojyi` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_5jqeagvvgywp0aasocpftbqx3` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `descr` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_4gh2ysmq5qswyc3lode6ks470` (`event_id`),
  CONSTRAINT `FK_4gh2ysmq5qswyc3lode6ks470` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `answer` varchar(10000) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `question` varchar(10000) DEFAULT NULL,
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
  `session_id` bigint(20) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_1vbgxu39moptqd7b51btgw9ae` (`client_id`),
  KEY `FK_ncokoe3h6mcgbuy6qews51isx` (`event_id`),
  KEY `FK_m1dycsn9g246l1hwqj9rht7y0` (`session_id`),
  CONSTRAINT `FK_1vbgxu39moptqd7b51btgw9ae` FOREIGN KEY (`client_id`) REFERENCES `evt_guest` (`id`),
  CONSTRAINT `FK_m1dycsn9g246l1hwqj9rht7y0` FOREIGN KEY (`session_id`) REFERENCES `evt_session` (`id`),
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
  `end_datetime` datetime DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `session_id` bigint(20) DEFAULT NULL,
  `start_datetime` datetime DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_466i3y5kl5n5vr1kd8cwutrvu` (`event_id`),
  KEY `FK_enwjwmsfcm6f5o3khets5uwoa` (`session_id`),
  CONSTRAINT `FK_466i3y5kl5n5vr1kd8cwutrvu` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`),
  CONSTRAINT `FK_enwjwmsfcm6f5o3khets5uwoa` FOREIGN KEY (`session_id`) REFERENCES `evt_session` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `html` varchar(10000) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_iqnbnrxybhrg6jpo7py4yo9hb` (`creator_id`),
  KEY `FK_qlgkooutqxsphcv7xn0qsk0hx` (`event_id`),
  CONSTRAINT `FK_iqnbnrxybhrg6jpo7py4yo9hb` FOREIGN KEY (`creator_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_qlgkooutqxsphcv7xn0qsk0hx` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `default_session` bit(1) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `fax` varchar(255) DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_sqkcig0k3tjqtdfcljdgjchcd` (`event_id`),
  CONSTRAINT `FK_sqkcig0k3tjqtdfcljdgjchcd` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `content` varchar(10000) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_mnw0mv93g0t675ih01hbp14db` (`author_id`),
  KEY `FK_oibshoy3bj4mwtbm44nwt6l6n` (`event_id`),
  KEY `FK_e1eyy5a16vj80e57q2ee39dam` (`parent_id`),
  CONSTRAINT `FK_e1eyy5a16vj80e57q2ee39dam` FOREIGN KEY (`parent_id`) REFERENCES `evt_thread` (`id`),
  CONSTRAINT `FK_mnw0mv93g0t675ih01hbp14db` FOREIGN KEY (`author_id`) REFERENCES `evt_client` (`id`),
  CONSTRAINT `FK_oibshoy3bj4mwtbm44nwt6l6n` FOREIGN KEY (`event_id`) REFERENCES `evt_event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `r_user_role`
-- ----------------------------
DROP TABLE IF EXISTS `r_user_role`;
CREATE TABLE `r_user_role` (
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `FK_fulfqo67e102b3ued59m61h1j` (`role_id`),
  CONSTRAINT `FK_6klb3jmokm5ho9m2uwocipt8v` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`),
  CONSTRAINT `FK_fulfqo67e102b3ued59m61h1j` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_company`
-- ----------------------------
DROP TABLE IF EXISTS `sys_company`;
CREATE TABLE `sys_company` (
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
  KEY `FK_lfoeigyrb3g60xpqirbuc01cr` (`admin_id`),
  CONSTRAINT `FK_lfoeigyrb3g60xpqirbuc01cr` FOREIGN KEY (`admin_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_company`
-- ----------------------------
BEGIN;
INSERT INTO `sys_company` VALUES ('-1', '2017-03-17 09:43:07', b'0', b'0', '2017-03-17 09:40:41', '1', null, '-1', null, 'linkr.cn', null, 'linkr.cn');
COMMIT;

-- ----------------------------
--  Table structure for `sys_role`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
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
  `agent` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `company_id` bigint(20) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `last_login_time` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `platform` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `verify_code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_b9bngbew7788y87fndhgadegj` (`company_id`),
  CONSTRAINT `FK_b9bngbew7788y87fndhgadegj` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_user`
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES ('-1', '2017-03-17 09:42:47', b'0', b'0', '2017-03-22 15:57:15', '49', null, 'upload/event/20170212/129e383b-a4dc-4f49-80f7-42b58993a06e.jpg', '-1', '462826@qq.com', '2017-03-22 15:57:15', 'Aaron Chen', '111111', '18626203266', null, 'd15a2d42-2058-4505-8c25-8839f26dfa62', null);
COMMIT;

-- ----------------------------
--  Table structure for `sys_verify_code`
-- ----------------------------
DROP TABLE IF EXISTS `sys_verify_code`;
CREATE TABLE `sys_verify_code` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `expire_time` datetime DEFAULT NULL,
  `ref_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_case`
-- ----------------------------
DROP TABLE IF EXISTS `tst_case`;
CREATE TABLE `tst_case` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `estimate` int(11) DEFAULT NULL,
  `module_id` bigint(20) DEFAULT NULL,
  `priority` int(11) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_mqw7t5g6myypi1hob10xahm6g` (`module_id`),
  KEY `FK_d023qp0i62gmdoejro44i6pmh` (`project_id`),
  KEY `FK_78ndnyxj0eg3m260eiiyb2vv3` (`parent_id`),
  CONSTRAINT `FK_78ndnyxj0eg3m260eiiyb2vv3` FOREIGN KEY (`parent_id`) REFERENCES `tst_case` (`id`),
  CONSTRAINT `FK_d023qp0i62gmdoejro44i6pmh` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_mqw7t5g6myypi1hob10xahm6g` FOREIGN KEY (`module_id`) REFERENCES `tst_module` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_case`
-- ----------------------------
BEGIN;
INSERT INTO `tst_case` VALUES ('1', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '测试用例', '1', 'root', null, '/'), ('2', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '用户账号', '1', 'branch', '1', '/1/'), ('3', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '注册', '1', 'branch', '2', '/1/2/'), ('4', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '登录', '1', 'branch', '2', '/1/2/'), ('5', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '忘记密码', '1', 'branch', '2', '/1/2/'), ('6', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '测试管理', '1', 'branch', '1', '/1/'), ('7', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '测试设计', '1', 'branch', '6', '/1/6/'), ('8', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '测试用例管理', '1', 'branch', '7', '/1/6/7/'), ('9', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '测试规划', '1', 'branch', '6', '/1/6/'), ('10', '2017-02-28 14:37:47', b'0', b'0', '2017-02-28 14:37:50', '1', '描述', '10', '-1', '1', '测试执行', '1', 'branch', '6', '/1/6/'), ('11', '2017-03-01 17:34:05', b'0', b'0', '2017-03-01 17:34:08', '1', '描述', '10', '-1', '1', '用户登录成功', '1', 'leaf', '4', '/1/2/4/'), ('12', '2017-03-01 17:42:35', b'0', b'0', '2017-03-01 17:42:38', '1', '描述', '10', '-1', '1', '输入密码错误', '1', 'leaf', '4', '/1/2/4/'), ('13', '2017-03-01 17:44:08', b'0', b'0', '2017-03-01 17:44:11', '1', '描述', '10', '-1', '1', '新建用例', '1', 'leaf', '8', '/1/6/7/8/');
COMMIT;

-- ----------------------------
--  Table structure for `tst_company`
-- ----------------------------
DROP TABLE IF EXISTS `tst_company`;
CREATE TABLE `tst_company` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_company`
-- ----------------------------
BEGIN;
INSERT INTO `tst_company` VALUES ('-1', '2017-02-28 14:26:59', null, null, '2017-02-28 14:27:03', '1', null, 'linkr.cn');
COMMIT;

-- ----------------------------
--  Table structure for `tst_module`
-- ----------------------------
DROP TABLE IF EXISTS `tst_module`;
CREATE TABLE `tst_module` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `priority` int(11) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_my1pduski6bkr0slg7mkfjk54` (`project_id`),
  CONSTRAINT `FK_my1pduski6bkr0slg7mkfjk54` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_module`
-- ----------------------------
BEGIN;
INSERT INTO `tst_module` VALUES ('-1', '2017-02-28 14:28:17', null, null, '2017-02-28 14:28:19', '1', null, 'Web工作台', '1', '1');
COMMIT;

-- ----------------------------
--  Table structure for `tst_project`
-- ----------------------------
DROP TABLE IF EXISTS `tst_project`;
CREATE TABLE `tst_project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `company_id` bigint(20) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `is_active` bit(1) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_88h46xtiew6swo3puoi4wkhj5` (`company_id`),
  KEY `FK_rm5uawwl53dtse1l5qhwci30v` (`parent_id`),
  CONSTRAINT `FK_88h46xtiew6swo3puoi4wkhj5` FOREIGN KEY (`company_id`) REFERENCES `tst_company` (`id`),
  CONSTRAINT `FK_rm5uawwl53dtse1l5qhwci30v` FOREIGN KEY (`parent_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_tofebdwwpqg556vv34bvl15ay` FOREIGN KEY (`company_id`) REFERENCES `sys_company` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_project`
-- ----------------------------
BEGIN;
INSERT INTO `tst_project` VALUES ('0', '2017-03-18 15:22:39', b'0', b'0', '2017-03-18 15:22:42', '1', '-1', '描述', 'ROOT', b'1', null, '0', '/', 'root'), ('1', '2017-02-28 14:27:53', b'0', b'0', '2017-03-22 17:06:35', '5', '-1', '描述', 'TestSpace', b'1', '0', '1', '/0/', 'branch'), ('2', '2017-02-28 14:27:53', b'0', b'0', '2017-02-28 14:27:56', '1', '-1', '描述', 'TestSpace前端', b'1', '1', '2', '/0/1/', 'leaf'), ('3', '2017-02-28 14:27:53', b'0', b'0', '2017-02-28 14:27:56', '1', '-1', '描述', 'TestSpace后端', b'1', '1', '2', '/0/1/', 'leaf'), ('4', '2017-02-28 14:27:53', b'0', b'0', '2017-02-28 14:27:56', '1', '-1', '描述', 'Docker部署', b'1', '1', '2', '/0/1/', 'leaf'), ('5', '2017-02-28 14:27:53', b'0', b'0', '2017-03-22 19:33:56', '2', '-1', '描述', '项目组1', b'1', '0', '2', '/0/1/', 'branch'), ('6', '2017-02-28 14:27:53', b'0', b'0', '2017-02-28 14:27:56', '1', '-1', '描述', '项目A', b'1', '5', '2', '/0/5/', 'branch'), ('7', '2017-02-28 14:27:53', b'0', b'0', '2017-03-22 19:03:47', '3', '-1', '描述', '项目B', b'1', '5', '2', '/0/5/', 'leaf'), ('8', '2017-03-17 23:29:55', b'0', b'0', '2017-03-22 19:32:53', '17', '-1', '描述', '项目A1', b'1', '6', '3', '/0/5/6/', 'leaf'), ('9', '2017-03-18 02:31:28', b'0', b'0', '2017-03-18 02:31:32', '1', '-1', '描述', '项目A2', b'0', '6', '3', '/0/5/6/', 'leaf');
COMMIT;

-- ----------------------------
--  Function structure for `queryProjectChildren`
-- ----------------------------
DROP FUNCTION IF EXISTS `queryProjectChildren`;
delimiter ;;
CREATE DEFINER=`root`@`localhost` FUNCTION `queryProjectChildren`(companyId BIGINT, projectId BIGINT) RETURNS varchar(4000) CHARSET utf8
BEGIN
	DECLARE sTemp VARCHAR(4000);
	DECLARE sTempChd VARCHAR(4000);

	SET sTemp = '$';
	SET sTempChd = cast(projectId as char);

	WHILE sTempChd is not NULL DO
		SET sTemp = CONCAT(sTemp,',',sTempChd);
		SELECT group_concat(id) INTO sTempChd FROM tst_project where FIND_IN_SET(parent_id,sTempChd)>0 AND company_id = companyId;
	END WHILE;
	return sTemp;
END
 ;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
