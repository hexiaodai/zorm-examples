/*
 Navicat Premium Data Transfer

 Source Server         : dbms
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : rm-bp1d4145wfeb6z4stso.mysql.rds.aliyuncs.com:3306
 Source Schema         : zorm_t

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 03/08/2022 16:33:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_actor
-- ----------------------------
DROP TABLE IF EXISTS `t_actor`;
CREATE TABLE `t_actor` (
  `id` int NOT NULL,
  `stageName` varchar(255)  NOT NULL COMMENT '艺名',
  `realName` varchar(255)  NOT NULL COMMENT '实名',
  `company` varchar(255)  NOT NULL COMMENT '经济公司',
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_actor
-- ----------------------------
BEGIN;
INSERT INTO `t_actor` (`id`, `stageName`, `realName`, `company`, `createTime`) VALUES (1, '陈翔', '陈翔', '陈翔六点半', '2022-07-26 07:41:48');
INSERT INTO `t_actor` (`id`, `stageName`, `realName`, `company`, `createTime`) VALUES (2, '妹总', '应宝林', '陈翔六点半', '2022-07-26 07:44:20');
INSERT INTO `t_actor` (`id`, `stageName`, `realName`, `company`, `createTime`) VALUES (3, '球球', '纪文君', '陈翔六点半', '2022-07-26 07:45:22');
INSERT INTO `t_actor` (`id`, `stageName`, `realName`, `company`, `createTime`) VALUES (4, '毛台', '邰光远', '陈翔六点半', '2022-07-26 07:48:51');
INSERT INTO `t_actor` (`id`, `stageName`, `realName`, `company`, `createTime`) VALUES (5, '闰土', '李闰刚', '陈翔六点半', '2022-07-26 07:50:08');
INSERT INTO `t_actor` (`id`, `stageName`, `realName`, `company`, `createTime`) VALUES (6, '蘑菇头', '黄晓飞', '陈翔六点半', '2022-07-26 07:51:02');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
