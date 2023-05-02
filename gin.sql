/*
 Navicat MySQL Data Transfer

 Source Server         : 192.168.0.6
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : 192.168.0.6:3306
 Source Schema         : gin

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 01/09/2021 08:45:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cate_id` int(0) NOT NULL,
  `state` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cate_id`(`cate_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '8月份CPI同比上涨2.8% 猪肉价格上涨46.7%', 1, 1);
INSERT INTO `article` VALUES (2, '中国联通与中国电信共建共享5G网络 用户归属不变', 1, 1);
INSERT INTO `article` VALUES (3, '林郑月娥斥责暴徒破坏港铁:不能因为没生命就肆意破坏', 2, 1);
INSERT INTO `article` VALUES (4, '这些老师的口头禅，想起那些年“被支配的恐惧”了吗', 2, 1);
INSERT INTO `article` VALUES (5, '美国空军一号差点遭雷劈，特朗普惊呼：令人惊奇', 3, 1);
INSERT INTO `article` VALUES (6, '这些老师的口头禅，想起那些年“被支配的恐惧', 4, 1);

-- ----------------------------
-- Table structure for article_cate
-- ----------------------------
DROP TABLE IF EXISTS `article_cate`;
CREATE TABLE `article_cate`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `state` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_cate
-- ----------------------------
INSERT INTO `article_cate` VALUES (1, '国内', 1);
INSERT INTO `article_cate` VALUES (2, '国际', 1);
INSERT INTO `article_cate` VALUES (3, '娱乐', 1);
INSERT INTO `article_cate` VALUES (4, '互联网', 1);

-- ----------------------------
-- Table structure for bank
-- ----------------------------
DROP TABLE IF EXISTS `bank`;
CREATE TABLE `bank`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `balance` decimal(10, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bank
-- ----------------------------
INSERT INTO `bank` VALUES (1, '张三', 100.00);
INSERT INTO `bank` VALUES (2, '李四', 100.00);

-- ----------------------------
-- Table structure for lesson
-- ----------------------------
DROP TABLE IF EXISTS `lesson`;
CREATE TABLE `lesson`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lesson
-- ----------------------------
INSERT INTO `lesson` VALUES (1, '计算机网络');
INSERT INTO `lesson` VALUES (2, 'Java程序设计');
INSERT INTO `lesson` VALUES (3, '软件项目管理');
INSERT INTO `lesson` VALUES (4, '网络安全');

-- ----------------------------
-- Table structure for lesson_student
-- ----------------------------
DROP TABLE IF EXISTS `lesson_student`;
CREATE TABLE `lesson_student`  (
  `lesson_id` int(0) NOT NULL,
  `student_id` int(0) NOT NULL,
  PRIMARY KEY (`student_id`, `lesson_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lesson_student
-- ----------------------------
INSERT INTO `lesson_student` VALUES (1, 1);
INSERT INTO `lesson_student` VALUES (2, 1);
INSERT INTO `lesson_student` VALUES (4, 1);
INSERT INTO `lesson_student` VALUES (2, 2);
INSERT INTO `lesson_student` VALUES (4, 2);
INSERT INTO `lesson_student` VALUES (1, 3);
INSERT INTO `lesson_student` VALUES (3, 3);
INSERT INTO `lesson_student` VALUES (1, 4);
INSERT INTO `lesson_student` VALUES (2, 4);
INSERT INTO `lesson_student` VALUES (3, 4);
INSERT INTO `lesson_student` VALUES (4, 4);
INSERT INTO `lesson_student` VALUES (1, 5);
INSERT INTO `lesson_student` VALUES (2, 6);
INSERT INTO `lesson_student` VALUES (4, 6);

-- ----------------------------
-- Table structure for nav
-- ----------------------------
DROP TABLE IF EXISTS `nav`;
CREATE TABLE `nav`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(1) NULL DEFAULT 1,
  `sort` int(0) NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 22 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nav
-- ----------------------------
INSERT INTO `nav` VALUES (1, '商会组织', '', 1, 1);
INSERT INTO `nav` VALUES (2, '赣商人文', 'www.itying.com', 1, 1);
INSERT INTO `nav` VALUES (3, '商会动态', '', 1, 1);
INSERT INTO `nav` VALUES (4, '会员企业', '', 0, 3);
INSERT INTO `nav` VALUES (5, '商会服务', '', 0, 1);
INSERT INTO `nav` VALUES (6, '青海风情', '', 1, 5);
INSERT INTO `nav` VALUES (7, '商会简介', '', 1, 3);
INSERT INTO `nav` VALUES (21, '联系我们', 'www.itying.com', 1, 10);

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `number` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `class_id` int(0) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `number`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of student
-- ----------------------------
INSERT INTO `student` VALUES (1, '160101', '202cb962ac59075b964b07152d234b70', 1, '张三');
INSERT INTO `student` VALUES (2, '160201', '202cb962ac59075b964b07152d234b70', 2, '李四');
INSERT INTO `student` VALUES (3, '160102', '202cb962ac59075b964b07152d234b70', 1, '王五');
INSERT INTO `student` VALUES (4, '160103', '202cb962ac59075b964b07152d234b70', 1, '王麻子');
INSERT INTO `student` VALUES (5, '160104', '202cb962ac59075b964b07152d234b70', 1, '赵四');
INSERT INTO `student` VALUES (6, '160202', '202cb962ac59075b964b07152d234b70', 2, '刘能');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` tinyint(0) NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `add_time` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'gorm', 34, 'gorm@qq.com', 1590374917);
INSERT INTO `user` VALUES (2, '哈哈', 34, 'gorm@qq1.com', 1590374917);
INSERT INTO `user` VALUES (6, 'itying gin grom', 0, 'aaa@qqq.com', 1630372124);
INSERT INTO `user` VALUES (7, 'itying GORM', 22, '222@qq.con', 1630372165);

SET FOREIGN_KEY_CHECKS = 1;
