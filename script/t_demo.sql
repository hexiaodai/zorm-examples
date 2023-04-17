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

 Date: 03/08/2022 16:34:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_demo
-- ----------------------------
DROP TABLE IF EXISTS `t_demo`;
CREATE TABLE `t_demo` (
  `id` varchar(50) NOT NULL COMMENT '主键',
  `userName` varchar(30) NOT NULL COMMENT '姓名',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `active` int DEFAULT NULL COMMENT '是否有效(0否,1是)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_demo
-- ----------------------------
BEGIN;
INSERT INTO `t_demo` (`id`, `userName`, `password`, `createTime`, `active`) VALUES ('20210630163227149563000042432429', '陈翔', '123456', '2022-07-25 21:42:14', 0);
INSERT INTO `t_demo` (`id`, `userName`, `password`, `createTime`, `active`) VALUES ('20210630163227149563000042432430', '妹总', '123456', '2022-07-25 22:31:30', 1);
INSERT INTO `t_demo` (`id`, `userName`, `password`, `createTime`, `active`) VALUES ('20210630163227149563000042432431', '蘑菇头', '123456', '2022-07-25 22:32:11', 0);
INSERT INTO `t_demo` (`id`, `userName`, `password`, `createTime`, `active`) VALUES ('20210630163227149563000042432432', '毛台', '123456', '2022-07-25 22:32:53', 1);
INSERT INTO `t_demo` (`id`, `userName`, `password`, `createTime`, `active`) VALUES ('20210630163227149563000042432433', '球球', '123456', '2022-07-25 22:36:04', NULL);
INSERT INTO `t_demo` (`id`, `userName`, `password`, `createTime`, `active`) VALUES ('20210630163227149563000042432434', '闰土', '123456', '2022-07-25 22:36:43', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;


-- 自定义函数
set global log_bin_trust_function_creators=1;
DROP FUNCTION IF EXISTS testfunc;

DELIMITER $
CREATE FUNCTION testfunc(id_in VARCHAR(100))
RETURNS VARCHAR(100)
BEGIN
	DECLARE userName_out VARCHAR(100) DEFAULT '20220726125301346422000491406956';
	SELECT userName INTO userName_out FROM t_demo WHERE id=id_in;
	RETURN userName_out;
END$
DELIMITER;

-- 调用函数
-- SELECT testfunc("20220726125301346422000491406956")

-- 定义存储过程
DELIMITER $$
DROP PROCEDURE IF EXISTS testproc;
CREATE PROCEDURE testproc(IN id VARCHAR(255))
BEGIN
	SELECT * FROM t_demo WHERE id = id;
	COMMIT;
END $$

-- 调用存储过程
-- DELIMITER ;
-- CALL testproc("20210630163227149563000042432429");