/*
 Navicat Premium Data Transfer

 Source Server         : 47.97.19.195
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : 47.97.19.195
 Source Database       : ngtesting

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : utf-8

 Date: 01/04/2018 17:02:41 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `sys_privilege`
-- ----------------------------
DROP TABLE IF EXISTS `sys_privilege`;
CREATE TABLE `sys_privilege` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_r_role_privilege`
-- ----------------------------
DROP TABLE IF EXISTS `sys_r_role_privilege`;
CREATE TABLE `sys_r_role_privilege` (
  `role_id` bigint(20) NOT NULL,
  `privilege_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`privilege_id`),
  KEY `FK_ky9ghoogn9iib4917xa0588ii` (`privilege_id`),
  CONSTRAINT `FK_ky9ghoogn9iib4917xa0588ii` FOREIGN KEY (`privilege_id`) REFERENCES `sys_privilege` (`id`),
  CONSTRAINT `FK_lafbrqm6tk3v0aj5wjan1afic` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_r_role_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_r_role_user`;
CREATE TABLE `sys_r_role_user` (
  `role_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`user_id`),
  KEY `FK_mp7eccpmrmommtiomo2hx94kq` (`user_id`),
  CONSTRAINT `FK_lnrx0pwvcwvfat4wno6ym36rk` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`),
  CONSTRAINT `FK_mp7eccpmrmommtiomo2hx94kq` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `descr` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_alert`
-- ----------------------------
DROP TABLE IF EXISTS `tst_alert`;
CREATE TABLE `tst_alert` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(10000) DEFAULT NULL,
  `due_time` datetime DEFAULT NULL,
  `is_read` bit(1) DEFAULT NULL,
  `opt_user_id` bigint(20) DEFAULT NULL,
  `sent` bit(1) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_bshl5kjhbto8rvbmolvdy2uiu` (`opt_user_id`),
  KEY `FK_b4fbqud01ub7bqahljyyux0ss` (`user_id`),
  CONSTRAINT `FK_b4fbqud01ub7bqahljyyux0ss` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_bshl5kjhbto8rvbmolvdy2uiu` FOREIGN KEY (`opt_user_id`) REFERENCES `tst_user` (`id`)
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
  `content` varchar(5000) DEFAULT NULL,
  `content_type` varchar(255) DEFAULT NULL,
  `create_by_id` bigint(20) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `estimate` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `objective` varchar(1000) DEFAULT NULL,
  `ordr` int(11) DEFAULT NULL,
  `p_id` bigint(20) DEFAULT NULL,
  `priority` varchar(255) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `prop01` varchar(255) DEFAULT NULL,
  `prop02` varchar(255) DEFAULT NULL,
  `prop03` varchar(255) DEFAULT NULL,
  `prop04` varchar(255) DEFAULT NULL,
  `prop05` varchar(255) DEFAULT NULL,
  `prop06` varchar(255) DEFAULT NULL,
  `prop07` varchar(255) DEFAULT NULL,
  `prop08` varchar(255) DEFAULT NULL,
  `prop09` varchar(255) DEFAULT NULL,
  `prop10` varchar(255) DEFAULT NULL,
  `prop11` varchar(255) DEFAULT NULL,
  `prop12` varchar(255) DEFAULT NULL,
  `prop13` varchar(255) DEFAULT NULL,
  `prop14` varchar(255) DEFAULT NULL,
  `prop15` varchar(255) DEFAULT NULL,
  `prop16` varchar(255) DEFAULT NULL,
  `prop17` varchar(255) DEFAULT NULL,
  `prop18` varchar(255) DEFAULT NULL,
  `prop19` varchar(255) DEFAULT NULL,
  `prop20` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `update_by_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_4paqpejxxg65icpu7asf9btow` (`create_by_id`),
  KEY `FK_le8suo2xxbcr036yaiivwkqn0` (`project_id`),
  KEY `FK_f3mtkmff26truvxmm897u8oeu` (`update_by_id`),
  CONSTRAINT `FK_4paqpejxxg65icpu7asf9btow` FOREIGN KEY (`create_by_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_f3mtkmff26truvxmm897u8oeu` FOREIGN KEY (`update_by_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_le8suo2xxbcr036yaiivwkqn0` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_case`
-- ----------------------------
BEGIN;
INSERT INTO `tst_case` VALUES ('1', '2018-01-04 14:10:54', b'0', b'0', null, '0', null, 'steps', '4', null, '10', '特性01', null, '0', null, 'middle', '2', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, 'functional', null), ('2', '2018-01-04 14:10:54', b'0', b'0', null, '0', null, 'steps', '4', null, '10', '用例01', null, '0', '1', 'middle', '2', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, 'functional', null), ('3', '2018-01-04 15:04:29', b'0', b'0', null, '0', null, 'steps', '5', null, '10', '特性01', null, '0', null, 'middle', '4', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, 'functional', null), ('4', '2018-01-04 15:04:29', b'0', b'0', '2018-01-04 15:05:20', '2', null, 'steps', '5', null, '10', '用例01', null, '0', '3', 'middle', '4', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, 'functional', '5');
COMMIT;

-- ----------------------------
--  Table structure for `tst_case_exe_status`
-- ----------------------------
DROP TABLE IF EXISTS `tst_case_exe_status`;
CREATE TABLE `tst_case_exe_status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `display_order` int(11) DEFAULT NULL,
  `is_build_in` bit(1) DEFAULT NULL,
  `is_final` bit(1) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_o4l4xg65y069b0ai5cgbfm175` (`org_id`),
  CONSTRAINT `FK_o4l4xg65y069b0ai5cgbfm175` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_case_exe_status`
-- ----------------------------
BEGIN;
INSERT INTO `tst_case_exe_status` VALUES ('5', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'untest', null, '10', b'0', b'0', '未执行', '3'), ('6', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'pass', null, '20', b'0', b'1', '成功', '3'), ('7', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'fail', null, '30', b'0', b'0', '失败', '3'), ('8', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'block', null, '40', b'0', b'0', '未执行', '3'), ('9', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'untest', null, '10', b'0', b'0', '未执行', '4'), ('10', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'pass', null, '20', b'0', b'1', '成功', '4'), ('11', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'fail', null, '30', b'0', b'0', '失败', '4'), ('12', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'block', null, '40', b'0', b'0', '未执行', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_case_in_run`
-- ----------------------------
DROP TABLE IF EXISTS `tst_case_in_run`;
CREATE TABLE `tst_case_in_run` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `case_id` bigint(20) DEFAULT NULL,
  `create_by_id` bigint(20) DEFAULT NULL,
  `ordr` int(11) DEFAULT NULL,
  `p_id` bigint(20) DEFAULT NULL,
  `result` varchar(255) DEFAULT NULL,
  `run_id` bigint(20) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_ahth2x2i7j9loamqyg3jcwfu6` (`create_by_id`),
  KEY `FK_8d38nl2cbd2ve2srlqrcur3qn` (`run_id`),
  KEY `FK_mwbiov88r7ppt8x9yunxr18pu` (`case_id`),
  CONSTRAINT `FK_8d38nl2cbd2ve2srlqrcur3qn` FOREIGN KEY (`run_id`) REFERENCES `tst_run` (`id`),
  CONSTRAINT `FK_ahth2x2i7j9loamqyg3jcwfu6` FOREIGN KEY (`create_by_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_mwbiov88r7ppt8x9yunxr18pu` FOREIGN KEY (`case_id`) REFERENCES `tst_case` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_case_priority`
-- ----------------------------
DROP TABLE IF EXISTS `tst_case_priority`;
CREATE TABLE `tst_case_priority` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `display_order` int(11) DEFAULT NULL,
  `is_build_in` bit(1) DEFAULT NULL,
  `is_default` bit(1) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_d8r4hkhobybms74u4vk43thj9` (`org_id`),
  CONSTRAINT `FK_d8r4hkhobybms74u4vk43thj9` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_case_priority`
-- ----------------------------
BEGIN;
INSERT INTO `tst_case_priority` VALUES ('4', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'high', null, '10', b'0', b'0', '高', '3'), ('5', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'medium', null, '20', b'0', b'1', '中', '3'), ('6', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'low', null, '30', b'0', b'0', '低', '3'), ('7', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'high', null, '10', b'0', b'0', '高', '4'), ('8', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'medium', null, '20', b'0', b'1', '中', '4'), ('9', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'low', null, '30', b'0', b'0', '低', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_case_step`
-- ----------------------------
DROP TABLE IF EXISTS `tst_case_step`;
CREATE TABLE `tst_case_step` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `expect` varchar(255) DEFAULT NULL,
  `opt` varchar(255) DEFAULT NULL,
  `ordr` int(11) DEFAULT NULL,
  `test_case_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_s8hj2viu2jtj1iwf4pgu789hi` (`test_case_id`),
  CONSTRAINT `FK_s8hj2viu2jtj1iwf4pgu789hi` FOREIGN KEY (`test_case_id`) REFERENCES `tst_case` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_case_step`
-- ----------------------------
BEGIN;
INSERT INTO `tst_case_step` VALUES ('1', null, b'0', b'0', null, '0', '', '是否', '1', '4'), ('2', null, b'0', b'0', null, '0', '', 'sdfds', '2', '4'), ('3', null, b'0', b'0', null, '0', '', 'sdfdsf', '4', '4'), ('4', null, b'0', b'0', null, '0', '', 'sdfdsf', '3', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_case_type`
-- ----------------------------
DROP TABLE IF EXISTS `tst_case_type`;
CREATE TABLE `tst_case_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `display_order` int(11) DEFAULT NULL,
  `is_build_in` bit(1) DEFAULT NULL,
  `is_default` bit(1) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_rqs9hiykm6kk5w8rewcy1uvy7` (`org_id`),
  CONSTRAINT `FK_rqs9hiykm6kk5w8rewcy1uvy7` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_case_type`
-- ----------------------------
BEGIN;
INSERT INTO `tst_case_type` VALUES ('7', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'functional', null, '10', b'0', b'0', '功能', '3'), ('8', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'performance', null, '20', b'0', b'1', '性能', '3'), ('9', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'ui', null, '30', b'0', b'1', '用户界面', '3'), ('10', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'compatibility', null, '40', b'0', b'0', '兼容性', '3'), ('11', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'security', null, '50', b'0', b'1', '安全', '3'), ('12', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'other', null, '60', b'0', b'0', '其它', '3'), ('13', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'functional', null, '10', b'0', b'0', '功能', '4'), ('14', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'performance', null, '20', b'0', b'1', '性能', '4'), ('15', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'ui', null, '30', b'0', b'1', '用户界面', '4'), ('16', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'compatibility', null, '40', b'0', b'0', '兼容性', '4'), ('17', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'security', null, '50', b'0', b'1', '安全', '4'), ('18', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'other', null, '60', b'0', b'0', '其它', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_custom_field`
-- ----------------------------
DROP TABLE IF EXISTS `tst_custom_field`;
CREATE TABLE `tst_custom_field` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `apply_to` varchar(255) DEFAULT NULL,
  `build_in` bit(1) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `format` varchar(255) DEFAULT NULL,
  `global` bit(1) DEFAULT NULL,
  `label` varchar(255) DEFAULT NULL,
  `my_column` varchar(255) DEFAULT NULL,
  `ordr` int(11) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  `required` bit(1) DEFAULT NULL,
  `rows` int(11) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_b1o40efa19tleean59bgg59jm` (`org_id`),
  CONSTRAINT `FK_b1o40efa19tleean59bgg59jm` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_document`
-- ----------------------------
DROP TABLE IF EXISTS `tst_document`;
CREATE TABLE `tst_document` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `msg` varchar(10000) DEFAULT NULL,
  `doc_type` varchar(255) DEFAULT NULL,
  `event_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_qqrnh8bqskjk1r56gflpm52yx` (`user_id`),
  CONSTRAINT `FK_qqrnh8bqskjk1r56gflpm52yx` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_history`
-- ----------------------------
DROP TABLE IF EXISTS `tst_history`;
CREATE TABLE `tst_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `msg` varchar(10000) DEFAULT NULL,
  `entity_id` bigint(20) DEFAULT NULL,
  `entity_type` varchar(255) DEFAULT NULL,
  `opt_user_id` bigint(20) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_sl3doy5ci4pnoo8958pi3vyvc` (`opt_user_id`),
  KEY `FK_j9m2m7ijlp9j2184nv0yiln9u` (`project_id`),
  KEY `FK_m4yjkr3nwc5y1fcjj1ke08xie` (`user_id`),
  CONSTRAINT `FK_j9m2m7ijlp9j2184nv0yiln9u` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_m4yjkr3nwc5y1fcjj1ke08xie` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_sl3doy5ci4pnoo8958pi3vyvc` FOREIGN KEY (`opt_user_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_org`
-- ----------------------------
DROP TABLE IF EXISTS `tst_org`;
CREATE TABLE `tst_org` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `admin_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_4oxyg49faexkfbphofgn8qnxr` (`admin_id`),
  CONSTRAINT `FK_4oxyg49faexkfbphofgn8qnxr` FOREIGN KEY (`admin_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_org`
-- ----------------------------
BEGIN;
INSERT INTO `tst_org` VALUES ('3', '2018-01-04 14:10:54', b'0', b'0', null, '0', '我的组织', null, null), ('4', '2018-01-04 15:04:29', b'0', b'0', null, '0', '我的组织', null, null);
COMMIT;

-- ----------------------------
--  Table structure for `tst_org_group`
-- ----------------------------
DROP TABLE IF EXISTS `tst_org_group`;
CREATE TABLE `tst_org_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_dlddwakgodocwt7n7abndkhtg` (`org_id`),
  CONSTRAINT `FK_dlddwakgodocwt7n7abndkhtg` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_org_group`
-- ----------------------------
BEGIN;
INSERT INTO `tst_org_group` VALUES ('2', '2018-01-04 14:10:54', b'0', b'0', null, '0', null, '所与人', '3'), ('3', '2018-01-04 15:04:29', b'0', b'0', null, '0', null, '所与人', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_org_privilege`
-- ----------------------------
DROP TABLE IF EXISTS `tst_org_privilege`;
CREATE TABLE `tst_org_privilege` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Records of `tst_org_privilege`
-- ----------------------------
BEGIN;
INSERT INTO `tst_org_privilege` VALUES ('1', '2017-04-05 09:39:15', b'0', b'0', '2017-04-05 09:39:20', '1', null, '管理公司', 'org_admin'), ('2', '2017-04-05 09:39:15', b'0', b'0', '2017-04-05 09:39:20', '1', null, '管理站点', 'site_admin'), ('3', '2017-04-05 09:39:15', b'0', b'0', '2017-04-05 09:39:20', '1', null, '管理项目', 'project_admin');
COMMIT;

-- ----------------------------
--  Table structure for `tst_org_role`
-- ----------------------------
DROP TABLE IF EXISTS `tst_org_role`;
CREATE TABLE `tst_org_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_q5g6x4w1pwr5ur4iwbg17nr9u` (`org_id`),
  CONSTRAINT `FK_q5g6x4w1pwr5ur4iwbg17nr9u` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_org_role`
-- ----------------------------
BEGIN;
INSERT INTO `tst_org_role` VALUES ('5', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'org_admin', null, '组织管理员', '3'), ('6', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'site_admin', null, '站点管理员', '3'), ('7', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'project_admin', null, '项目管理员', '3'), ('8', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'org_admin', null, '组织管理员', '4'), ('9', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'site_admin', null, '站点管理员', '4'), ('10', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'project_admin', null, '项目管理员', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_plan`
-- ----------------------------
DROP TABLE IF EXISTS `tst_plan`;
CREATE TABLE `tst_plan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `estimate` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_te991npw8lxmrtmt2gcjolimr` (`project_id`),
  CONSTRAINT `FK_te991npw8lxmrtmt2gcjolimr` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `descr` varchar(1000) DEFAULT NULL,
  `last_access_time` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `ord` int(11) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_avuusthsgk7g68bm0kiq6dix0` (`org_id`),
  KEY `FK_rm5uawwl53dtse1l5qhwci30v` (`parent_id`),
  CONSTRAINT `FK_avuusthsgk7g68bm0kiq6dix0` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`),
  CONSTRAINT `FK_rm5uawwl53dtse1l5qhwci30v` FOREIGN KEY (`parent_id`) REFERENCES `tst_project` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_project`
-- ----------------------------
BEGIN;
INSERT INTO `tst_project` VALUES ('1', '2018-01-04 14:10:54', b'0', b'0', null, '0', null, null, '默认项目组', null, '3', null, 'group'), ('2', '2018-01-04 14:10:54', b'0', b'0', null, '0', null, null, '默认项目', null, '3', '1', 'project'), ('3', '2018-01-04 15:04:29', b'0', b'0', null, '0', null, null, '默认项目组', null, '4', null, 'group'), ('4', '2018-01-04 15:04:29', b'0', b'0', null, '0', null, null, '默认项目', null, '4', '3', 'project');
COMMIT;

-- ----------------------------
--  Table structure for `tst_project_access_history`
-- ----------------------------
DROP TABLE IF EXISTS `tst_project_access_history`;
CREATE TABLE `tst_project_access_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `last_access_time` datetime DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_l0ifd62wftf6w81779j64rfmc` (`org_id`),
  KEY `FK_hv9vkb26yw1fluyh6thwh230h` (`project_id`),
  KEY `FK_dpcrx83ysgtel2eua0856xfk3` (`user_id`),
  CONSTRAINT `FK_dpcrx83ysgtel2eua0856xfk3` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_hv9vkb26yw1fluyh6thwh230h` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_l0ifd62wftf6w81779j64rfmc` FOREIGN KEY (`org_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_project_access_history`
-- ----------------------------
BEGIN;
INSERT INTO `tst_project_access_history` VALUES ('1', '2018-01-04 14:10:54', b'0', b'0', null, '0', '2018-01-04 14:10:54', '3', '2', '默认项目', '4'), ('2', '2018-01-04 15:04:29', b'0', b'0', null, '0', '2018-01-04 15:04:29', '4', '4', '默认项目', '5');
COMMIT;

-- ----------------------------
--  Table structure for `tst_project_privilege_define`
-- ----------------------------
DROP TABLE IF EXISTS `tst_project_privilege_define`;
CREATE TABLE `tst_project_privilege_define` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `action` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Records of `tst_project_privilege_define`
-- ----------------------------
BEGIN;
INSERT INTO `tst_project_privilege_define` VALUES ('100', '2017-12-26 10:13:08', b'0', b'0', '2017-12-26 10:13:11', '1', null, '测试计划', 'create', 'plan'), ('110', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试计划', 'update', 'plan'), ('120', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试计划', 'remove', 'plan'), ('130', '2017-12-26 10:11:16', b'0', b'0', '2017-12-26 10:11:18', '1', null, '测试用例', 'create', 'cases'), ('140', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试用例', 'update', 'cases'), ('150', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试用例', 'remove', 'cases'), ('160', '2017-12-26 10:18:29', b'0', b'0', '2017-12-26 10:18:38', '1', null, '测试集', 'create', 'round'), ('170', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试集', 'update', 'round'), ('180', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试集', 'remove', 'round'), ('190', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试集', 'close', 'round'), ('200', '2017-04-05 11:52:26', b'0', b'0', '2017-04-05 11:52:28', '1', null, '测试结果', 'update', 'result');
COMMIT;

-- ----------------------------
--  Table structure for `tst_project_role_for_org`
-- ----------------------------
DROP TABLE IF EXISTS `tst_project_role_for_org`;
CREATE TABLE `tst_project_role_for_org` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `descr` varchar(255) DEFAULT NULL,
  `is_build_in` bit(1) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_8eokjbtquljjgjahh7y0l0la6` (`org_id`),
  CONSTRAINT `FK_8eokjbtquljjgjahh7y0l0la6` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_project_role_for_org`
-- ----------------------------
BEGIN;
INSERT INTO `tst_project_role_for_org` VALUES ('2', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'test_leader', null, b'0', '测试主管', '3'), ('3', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'test_designer', null, b'0', '测试设计', '3'), ('4', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'tester', null, b'0', '测试执行', '3'), ('5', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'readonly', null, b'0', '只读用户', '3'), ('6', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'test_leader', null, b'0', '测试主管', '4'), ('7', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'test_designer', null, b'0', '测试设计', '4'), ('8', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'tester', null, b'0', '测试执行', '4'), ('9', '2018-01-04 15:04:29', b'0', b'0', null, '0', 'readonly', null, b'0', '只读用户', '4');
COMMIT;

-- ----------------------------
--  Table structure for `tst_project_role_priviledge_relation`
-- ----------------------------
DROP TABLE IF EXISTS `tst_project_role_priviledge_relation`;
CREATE TABLE `tst_project_role_priviledge_relation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `is_build_in` bit(1) DEFAULT NULL,
  `project_privilege_define_id` bigint(20) DEFAULT NULL,
  `project_role_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_6aiwgve7unve9rcj15v8woxyl` (`project_privilege_define_id`),
  KEY `FK_orqtwmqhjn4bih5y6pd5fla59` (`project_role_id`),
  CONSTRAINT `FK_6aiwgve7unve9rcj15v8woxyl` FOREIGN KEY (`project_privilege_define_id`) REFERENCES `tst_project_privilege_define` (`id`),
  CONSTRAINT `FK_orqtwmqhjn4bih5y6pd5fla59` FOREIGN KEY (`project_role_id`) REFERENCES `tst_project_role_for_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_project_role_priviledge_relation`
-- ----------------------------
BEGIN;
INSERT INTO `tst_project_role_priviledge_relation` VALUES ('2', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '100', '2'), ('3', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '110', '2'), ('4', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '120', '2'), ('5', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '130', '2'), ('6', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '140', '2'), ('7', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '150', '2'), ('8', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '160', '2'), ('9', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '170', '2'), ('10', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '180', '2'), ('11', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '190', '2'), ('12', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '200', '2'), ('13', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '130', '3'), ('14', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '140', '3'), ('15', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '150', '3'), ('16', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '160', '3'), ('17', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '170', '3'), ('18', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '180', '3'), ('19', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '190', '3'), ('20', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '200', '3'), ('21', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '160', '4'), ('22', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '170', '4'), ('23', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '180', '4'), ('24', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '190', '4'), ('25', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '200', '4'), ('26', '2018-01-04 14:10:54', b'0', b'0', null, null, b'0', '200', '5'), ('27', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '100', '6'), ('28', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '110', '6'), ('29', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '120', '6'), ('30', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '130', '6'), ('31', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '140', '6'), ('32', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '150', '6'), ('33', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '160', '6'), ('34', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '170', '6'), ('35', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '180', '6'), ('36', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '190', '6'), ('37', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '200', '6'), ('38', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '130', '7'), ('39', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '140', '7'), ('40', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '150', '7'), ('41', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '160', '7'), ('42', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '170', '7'), ('43', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '180', '7'), ('44', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '190', '7'), ('45', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '200', '7'), ('46', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '160', '8'), ('47', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '170', '8'), ('48', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '180', '8'), ('49', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '190', '8'), ('50', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '200', '8'), ('51', '2018-01-04 15:04:29', b'0', b'0', null, null, b'0', '200', '9');
COMMIT;

-- ----------------------------
--  Table structure for `tst_r_custom_field_project`
-- ----------------------------
DROP TABLE IF EXISTS `tst_r_custom_field_project`;
CREATE TABLE `tst_r_custom_field_project` (
  `custom_field_id` bigint(20) NOT NULL,
  `project_id` bigint(20) NOT NULL,
  PRIMARY KEY (`custom_field_id`,`project_id`),
  KEY `FK_5y5g3wjodtyxm3lpmmd04foy5` (`project_id`),
  CONSTRAINT `FK_5y5g3wjodtyxm3lpmmd04foy5` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_bo12oks940a30cyxlt39kiijc` FOREIGN KEY (`custom_field_id`) REFERENCES `tst_custom_field` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_r_org_group_user`
-- ----------------------------
DROP TABLE IF EXISTS `tst_r_org_group_user`;
CREATE TABLE `tst_r_org_group_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `org_group_id` bigint(20) DEFAULT NULL,
  `org_group_name` varchar(255) DEFAULT NULL,
  `org_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_oioog5ixo3vky1n5qhr55mjr6` (`org_group_id`),
  KEY `FK_96e8mkbgy9qly15goqecnson6` (`user_id`),
  CONSTRAINT `FK_96e8mkbgy9qly15goqecnson6` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_oioog5ixo3vky1n5qhr55mjr6` FOREIGN KEY (`org_group_id`) REFERENCES `tst_org_group` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_r_org_role_privilege`
-- ----------------------------
DROP TABLE IF EXISTS `tst_r_org_role_privilege`;
CREATE TABLE `tst_r_org_role_privilege` (
  `org_role_id` bigint(20) NOT NULL,
  `org_privilege_id` bigint(20) NOT NULL,
  PRIMARY KEY (`org_role_id`,`org_privilege_id`),
  KEY `FK_xrf0fqbnodxio07iqvttce72` (`org_privilege_id`),
  CONSTRAINT `FK_6kbys90ljdfp5dp7w5nb4d5ru` FOREIGN KEY (`org_role_id`) REFERENCES `tst_org_role` (`id`),
  CONSTRAINT `FK_xrf0fqbnodxio07iqvttce72` FOREIGN KEY (`org_privilege_id`) REFERENCES `tst_org_privilege` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_r_org_role_privilege`
-- ----------------------------
BEGIN;
INSERT INTO `tst_r_org_role_privilege` VALUES ('5', '1'), ('8', '1'), ('6', '2'), ('9', '2'), ('7', '3'), ('10', '3');
COMMIT;

-- ----------------------------
--  Table structure for `tst_r_org_role_user`
-- ----------------------------
DROP TABLE IF EXISTS `tst_r_org_role_user`;
CREATE TABLE `tst_r_org_role_user` (
  `org_role_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  PRIMARY KEY (`org_role_id`,`user_id`),
  KEY `FK_h6d5c2yfeaqitn4jb3fvkjtw6` (`user_id`),
  CONSTRAINT `FK_8cbhgbqt91ctmnw35ibtyofqg` FOREIGN KEY (`org_role_id`) REFERENCES `tst_org_role` (`id`),
  CONSTRAINT `FK_h6d5c2yfeaqitn4jb3fvkjtw6` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_r_org_role_user`
-- ----------------------------
BEGIN;
INSERT INTO `tst_r_org_role_user` VALUES ('5', '4'), ('8', '5');
COMMIT;

-- ----------------------------
--  Table structure for `tst_r_org_user`
-- ----------------------------
DROP TABLE IF EXISTS `tst_r_org_user`;
CREATE TABLE `tst_r_org_user` (
  `org_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  PRIMARY KEY (`org_id`,`user_id`),
  KEY `FK_dbrrq8bxgx5npl0wxialit7i2` (`user_id`),
  CONSTRAINT `FK_28gcxu8p61i0lao8unkaq5c6c` FOREIGN KEY (`org_id`) REFERENCES `tst_org` (`id`),
  CONSTRAINT `FK_dbrrq8bxgx5npl0wxialit7i2` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_r_org_user`
-- ----------------------------
BEGIN;
INSERT INTO `tst_r_org_user` VALUES ('3', '4'), ('4', '5');
COMMIT;

-- ----------------------------
--  Table structure for `tst_r_project_role_entity`
-- ----------------------------
DROP TABLE IF EXISTS `tst_r_project_role_entity`;
CREATE TABLE `tst_r_project_role_entity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `entity_id` bigint(20) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `project_role_id` bigint(20) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_e507ln5a5bxon2uyrs3b06bv8` (`project_role_id`),
  CONSTRAINT `FK_e507ln5a5bxon2uyrs3b06bv8` FOREIGN KEY (`project_role_id`) REFERENCES `tst_project_role_for_org` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_r_project_role_entity`
-- ----------------------------
BEGIN;
INSERT INTO `tst_r_project_role_entity` VALUES ('1', '2018-01-04 14:10:54', b'0', b'0', null, '0', '4', '2', '2', 'user'), ('2', '2018-01-04 15:04:29', b'0', b'0', null, '0', '5', '4', '6', 'user');
COMMIT;

-- ----------------------------
--  Table structure for `tst_run`
-- ----------------------------
DROP TABLE IF EXISTS `tst_run`;
CREATE TABLE `tst_run` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `estimate` int(11) DEFAULT NULL,
  `is_read` bit(1) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `ordr` int(11) DEFAULT NULL,
  `plan_id` bigint(20) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_3r1a8t5vxesj07c4kd5odc77y` (`plan_id`),
  KEY `FK_3yir1yvenq7mrnx44l4falpcq` (`project_id`),
  KEY `FK_iog5lfy5gnd0uccm0wgrlqcsd` (`user_id`),
  CONSTRAINT `FK_3r1a8t5vxesj07c4kd5odc77y` FOREIGN KEY (`plan_id`) REFERENCES `tst_plan` (`id`),
  CONSTRAINT `FK_3yir1yvenq7mrnx44l4falpcq` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_iog5lfy5gnd0uccm0wgrlqcsd` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_suite`
-- ----------------------------
DROP TABLE IF EXISTS `tst_suite`;
CREATE TABLE `tst_suite` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `descr` varchar(1000) DEFAULT NULL,
  `estimate` int(11) DEFAULT NULL,
  `objective` varchar(1000) DEFAULT NULL,
  `order_in_parent` int(11) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `priority` int(11) DEFAULT NULL,
  `project_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_k1x4ddnp0il0j0kd1qmardvl1` (`parent_id`),
  KEY `FK_bof1daqokqea3o5yfdlreg8jy` (`project_id`),
  KEY `FK_1r4cd0cr11rrevb0x5sj7w2pv` (`user_id`),
  CONSTRAINT `FK_1r4cd0cr11rrevb0x5sj7w2pv` FOREIGN KEY (`user_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_bof1daqokqea3o5yfdlreg8jy` FOREIGN KEY (`project_id`) REFERENCES `tst_project` (`id`),
  CONSTRAINT `FK_k1x4ddnp0il0j0kd1qmardvl1` FOREIGN KEY (`parent_id`) REFERENCES `tst_suite` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_thread`
-- ----------------------------
DROP TABLE IF EXISTS `tst_thread`;
CREATE TABLE `tst_thread` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `content` varchar(10000) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_hn7m54nygknarx9v1jn4phx81` (`author_id`),
  KEY `FK_mw7px95alyw1wrmwhlp96fbu5` (`parent_id`),
  CONSTRAINT `FK_hn7m54nygknarx9v1jn4phx81` FOREIGN KEY (`author_id`) REFERENCES `tst_user` (`id`),
  CONSTRAINT `FK_mw7px95alyw1wrmwhlp96fbu5` FOREIGN KEY (`parent_id`) REFERENCES `tst_thread` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tst_user`
-- ----------------------------
DROP TABLE IF EXISTS `tst_user`;
CREATE TABLE `tst_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `deleted` bit(1) DEFAULT NULL,
  `disabled` bit(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `case_board_left_size` int(11) DEFAULT NULL,
  `case_board_right_size` int(11) DEFAULT NULL,
  `default_org_id` bigint(20) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `last_login_time` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `verify_code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_user`
-- ----------------------------
BEGIN;
INSERT INTO `tst_user` VALUES ('1', '2018-01-04 14:06:38', b'0', b'0', null, '0', 'upload/sample/user/avatar.png', null, null, null, 'test@test.com', '2018-01-04 14:06:38', 'Aaron Chen', '111111', '11111111111', '87ec1ffa-92e2-4f2e-a55d-65fddc3b1f07', null), ('2', '2018-01-04 14:07:47', b'0', b'0', null, '0', 'upload/sample/user/avatar.png', null, null, null, 'test@test.com', '2018-01-04 14:07:47', 'Aaron Chen', '111111', '11111111111', '3882e9ed-560c-4a6d-b20c-da7b839e601b', null), ('3', '2018-01-04 14:09:42', b'0', b'0', null, '0', 'upload/sample/user/avatar.png', null, null, null, 'test@test.com', '2018-01-04 14:09:42', 'Aaron Chen', '111111', '11111111111', 'ea9928b6-96cd-42a0-ae06-75f658802dc7', null), ('4', '2018-01-04 14:10:54', b'0', b'0', null, '0', 'upload/sample/user/avatar.png', null, null, '3', 'test@test.com', '2018-01-04 14:10:54', 'Aaron Chen', '111111', '11111111111', 'd86331db-5418-4381-9988-89b12548bfe3', null), ('5', '2018-01-04 15:04:29', b'0', b'0', '2018-01-04 15:04:48', '1', 'upload/sample/user/avatar.png', null, null, '4', '462826@qq.com', '2018-01-04 15:04:48', 'Aaron Chen', '111111', '11111111111', '9460222c-b26a-466f-af03-440f6a2eeb57', null);
COMMIT;

-- ----------------------------
--  Table structure for `tst_verify_code`
-- ----------------------------
DROP TABLE IF EXISTS `tst_verify_code`;
CREATE TABLE `tst_verify_code` (
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `tst_verify_code`
-- ----------------------------
BEGIN;
INSERT INTO `tst_verify_code` VALUES ('1', '2018-01-04 14:10:54', b'0', b'0', null, '0', '624b40a9df5c4bc595d23dcc4370b453', '2018-01-04 14:20:54', '4'), ('2', '2018-01-04 15:04:29', b'1', b'0', '2018-01-04 15:04:48', '1', '75b181d6a6f54073b09e28d08d71cb7b', '2018-01-04 15:14:29', '5');
COMMIT;

-- ----------------------------
--  Procedure structure for `init_user`
-- ----------------------------
DROP PROCEDURE IF EXISTS `init_user`;
delimiter ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `init_user`(IN user_id  BIGINT)
BEGIN

DECLARE i BIGINT;

DECLARE sql_str varchar(5000);
DECLARE org_id BIGINT;
DECLARE org_role_id BIGINT;
DECLARE project_role_id BIGINT;
DECLARE project_role_leader_id BIGINT;
DECLARE project_id BIGINT;
DECLARE case_id BIGINT;

/*新增组织*/
insert into tst_org (name, disabled, deleted, create_time, version) values('我的组织', false, false, NOW(), 0);
select max(id) from tst_org into org_id;

/*更新用户默认的组织*/
update tst_user usr set usr.default_org_id = org_id where usr.id=user_id;

/*添加用户到组织*/
insert into tst_r_org_user (org_id, user_id) values(org_id, user_id);

/*初始化组织角色及其权限*/
insert into tst_org_role (code, name, org_id, disabled, deleted, create_time, version) values('org_admin', '组织管理员', org_id, false, false, NOW(), 0);
select max(id) from tst_org_role into org_role_id;
insert into tst_r_org_role_privilege (org_role_id, org_privilege_id) values(org_role_id, 1);
/*添加用户为组织管理员*/
insert into tst_r_org_role_user (org_role_id, user_id) values(org_role_id, user_id);

insert into tst_org_role (code, name, org_id, disabled, deleted, create_time, version) values('site_admin', '站点管理员', org_id, false, false, NOW(), 0);
select max(id) from tst_org_role into org_role_id;
insert into tst_r_org_role_privilege (org_role_id, org_privilege_id) values(org_role_id, 2);

insert into tst_org_role (code, name, org_id, disabled, deleted, create_time, version) values('project_admin', '项目管理员', org_id, false, false, NOW(), 0);
select max(id) from tst_org_role into org_role_id;
insert into tst_r_org_role_privilege (org_role_id, org_privilege_id) values(org_role_id, 3);

/*初始化组织群组*/
insert into tst_org_group (name, org_id, disabled, deleted, create_time, version) values('所与人', org_id, false, false, NOW(), 0);

/*初始化用例执行状态*/
insert into tst_case_exe_status (code, name, display_order, is_build_in, is_final, org_id, disabled, deleted, create_time, version) 
		   values('untest', '未执行', 10, false, false, org_id, false, false, NOW(), 0);
insert into tst_case_exe_status (code, name, display_order, is_build_in, is_final, org_id, disabled, deleted, create_time, version) 
		   values('pass', '成功', 20, false, true, org_id, false, false, NOW(), 0);
insert into tst_case_exe_status (code, name, display_order, is_build_in, is_final, org_id, disabled, deleted, create_time, version) 
		   values('fail', '失败', 30, false, false, org_id, false, false, NOW(), 0);
insert into tst_case_exe_status (code, name, display_order, is_build_in, is_final, org_id, disabled, deleted, create_time, version) 
		   values('block', '未执行', 40, false, false, org_id, false, false, NOW(), 0);

/*初始化用例优先级*/
insert into tst_case_priority (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('high', '高', 10, false, false, org_id, false, false, NOW(), 0);
insert into tst_case_priority (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('medium', '中', 20, false, true, org_id, false, false, NOW(), 0);
insert into tst_case_priority (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('low', '低', 30, false, false, org_id, false, false, NOW(), 0);

/*初始化用例类型*/
insert into tst_case_type (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('functional', '功能', 10, false, false, org_id, false, false, NOW(), 0);
insert into tst_case_type (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('performance', '性能', 20, false, true, org_id, false, false, NOW(), 0);
insert into tst_case_type (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('ui', '用户界面', 30, false, true, org_id, false, false, NOW(), 0);
insert into tst_case_type (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('compatibility', '兼容性', 40, false, false, org_id, false, false, NOW(), 0);
insert into tst_case_type (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('security', '安全', 50, false, true, org_id, false, false, NOW(), 0);
insert into tst_case_type (code, name, display_order, is_build_in, is_default, org_id, disabled, deleted, create_time, version) 
		   values('other', '其它', 60, false, false, org_id, false, false, NOW(), 0);

/*初始化项目角色*/
insert into tst_project_role_for_org (code, name, is_build_in, org_id, disabled, deleted, create_time, version) 
		   values('test_leader', '测试主管', false, org_id, false, false, NOW(), 0);
select max(id) from tst_project_role_for_org into project_role_id;
set project_role_leader_id=project_role_id;
/*初始化项目角色的权限*/
set i=100;
while i<=200 do
	insert into tst_project_role_priviledge_relation 
                ( project_privilege_define_id,   project_role_id,   create_time, deleted, disabled, is_build_in )
         VALUES ( i, project_role_id, now(),       false,   false,    false );
	set i=i+10;
end while;

insert into tst_project_role_for_org (code, name, is_build_in, org_id, disabled, deleted, create_time, version) 
		   values('test_designer', '测试设计', false, org_id, false, false, NOW(), 0);
select max(id) from tst_project_role_for_org into project_role_id;
/*初始化项目角色的权限*/
set i=130;
while i<=200 do
	insert into tst_project_role_priviledge_relation 
                ( project_privilege_define_id,   project_role_id,   create_time, deleted, disabled, is_build_in )
         VALUES ( i, project_role_id, now(),       false,   false,    false );
	set i=i+10;
end while;

insert into tst_project_role_for_org (code, name, is_build_in, org_id, disabled, deleted, create_time, version) 
		   values('tester', '测试执行', false, org_id, false, false, NOW(), 0);
select max(id) from tst_project_role_for_org into project_role_id;
/*初始化项目角色的权限*/
set i=160;
while i<=200 do
	insert into tst_project_role_priviledge_relation 
                ( project_privilege_define_id,   project_role_id,   create_time, deleted, disabled, is_build_in )
         VALUES ( i, project_role_id, now(),       false,   false,    false );
	set i=i+10;
end while;

insert into tst_project_role_for_org (code, name, is_build_in, org_id, disabled, deleted, create_time, version) 
		   values('readonly', '只读用户', false, org_id, false, false, NOW(), 0);
select max(id) from tst_project_role_for_org into project_role_id;
/*初始化项目角色的权限*/
insert into tst_project_role_priviledge_relation 
	( project_privilege_define_id,   project_role_id,   create_time, deleted, disabled, is_build_in )
	VALUES ( 200, project_role_id, now(),       false,   false,    false );

/*初始化项目组*/
insert into tst_project (name, type, parent_id, org_id, disabled, deleted, create_time, version) 
		   values('默认项目组', 'group', NULL, org_id, false, false, NOW(), 0);
select max(id) from tst_project into project_id;
/*初始化项目*/
insert into tst_project (name, type, parent_id, org_id, disabled, deleted, create_time, version) 
		   values('默认项目', 'project', project_id, org_id, false, false, NOW(), 0);
select max(id) from tst_project into project_id;

/*添加用户为项目主管*/
insert into tst_r_project_role_entity (project_id, project_role_id, entity_id, type, disabled, deleted, create_time, version) 
		   values(project_id, project_role_leader_id, user_id, 'user', false, false, NOW(), 0);

/*用户访问默认项目*/
insert into tst_project_access_history (org_id, project_id, user_id, project_name, last_access_time, disabled, deleted, create_time, version) 
		   values(org_id, project_id, user_id, '默认项目', NOW(), false, false, NOW(), 0);

/*添加测试用例*/
insert into tst_case (name, project_id, p_id, estimate, priority, type, ordr, create_by_id, content_type, disabled, deleted, create_time, version) 
		   values('特性01', project_id, null, 10, 'middle', 'functional', 0, user_id, 'steps', false, false, NOW(), 0);
select max(id) from tst_case into case_id;
insert into tst_case (name, project_id, p_id, estimate, priority, type, ordr, create_by_id, content_type, disabled, deleted, create_time, version) 
		   values('用例01', project_id, case_id, 10, 'middle', 'functional', 0, user_id, 'steps', false, false, NOW(), 0);

END
 ;;
delimiter ;

-- ----------------------------
--  Procedure structure for `move_node`
-- ----------------------------
DROP PROCEDURE IF EXISTS `move_node`;
delimiter ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `move_node`(IN node_table varchar(100), IN node_id BIGINT, IN parent_id BIGINT)
BEGIN
DECLARE sql_str varchar(5000);

/*获取老的原路径*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT node.path into @old_path FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

/*获取新的父路径*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT node.path into @node_path FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', parent_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

set @node_path = concat(@node_path, parent_id, '/');
set @child_path = concat(@node_path, node_id, '/');

/*更新自己*/
set sql_str = '';
set sql_str = concat(sql_str, '  UPDATE ', node_table , ' SET parent_id = ' , parent_id, ',');
set sql_str = concat(sql_str, '             path = ', '''' , @node_path, '''');
set sql_str = concat(sql_str, '  WHERE id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;

IF @old_path is null THEN 
    set @old_path = @old_path;
ELSE
    set @old_path = concat(@old_path, node_id, '/');
    /*更新后代*/
	set sql_str = '';
	set sql_str = concat(sql_str, '  UPDATE ', node_table);
	set sql_str = concat(sql_str, '   SET path = REPLACE(path, ', '''', @old_path , '''', ',', '''' , @child_path, '''', ')');
	set sql_str = concat(sql_str, '  WHERE path LIKE ', '''', @old_path, '%''');
	
	set @sql_str = sql_str;
	PREPARE stmt FROM @sql_str;
	EXECUTE stmt;
END IF;

/*查询*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT * FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

END
 ;;
delimiter ;

-- ----------------------------
--  Procedure structure for `update_node`
-- ----------------------------
DROP PROCEDURE IF EXISTS `update_node`;
delimiter ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `update_node`(IN node_table varchar(100), IN node_id BIGINT, IN status_name varchar(100), IN status_value varchar(100))
BEGIN
DECLARE sql_str varchar(5000);

/*获取路径*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT node.path into @node_path FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

set @node_path = concat(@node_path, node_id, '/');

/* 更新 */
set sql_str = '';
set sql_str = concat(sql_str, '  update ', node_table);
set sql_str = concat(sql_str, '    SET ', status_name, ' = ' , status_value);
set sql_str = concat(sql_str, '  WHERE id =', node_id, ' OR path LIKE ', '''', @node_path, '%''');

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

/*查询*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT * FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

END
 ;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
