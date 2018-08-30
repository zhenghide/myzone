/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : zone

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 28/08/2018 15:22:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for diary
-- ----------------------------
DROP TABLE IF EXISTS `diary`;
CREATE TABLE `diary`  (
  `id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of diary
-- ----------------------------
INSERT INTO `diary` VALUES ('5203e2d2-a650-443a-bf7b-a56d8099c76f', 'admin', '测试第一篇', '哈哈哈哈哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈，哈哈哈哈哈哈！！！', '2018-08-27 13:57:41');
INSERT INTO `diary` VALUES ('7e6715e2-a3ef-4b25-9574-a4bb3ea5e948', 'admin', '测试格式', '       哈哈哈哈哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈\n       哈哈哈哈哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈哈哈哈哈，哈哈哈哈哈哈', '2018-08-28 10:54:37');

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo`  (
  `id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `photo_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `creater` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of photo
-- ----------------------------
INSERT INTO `photo` VALUES ('5f067fcf-b25e-45db-92db-34e400e22870', 'LOGO', '2018-08-27', 'admin');
INSERT INTO `photo` VALUES ('af5e0665-32ff-43a8-b6c1-c00b20975c3c', '头像', '2018-08-27', 'admin');

-- ----------------------------
-- Table structure for picture
-- ----------------------------
DROP TABLE IF EXISTS `picture`;
CREATE TABLE `picture`  (
  `id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pic_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pic_size` double NOT NULL,
  `pic_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `photo_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `photo_id`(`photo_id`) USING BTREE,
  CONSTRAINT `picture_ibfk_1` FOREIGN KEY (`photo_id`) REFERENCES `photo` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of picture
-- ----------------------------
INSERT INTO `picture` VALUES ('31d122fd-3053-4491-8960-423c112f1a3a', 'photo.png', 368829, 'static/upload/admin/20180827114047photo.png', 'af5e0665-32ff-43a8-b6c1-c00b20975c3c');
INSERT INTO `picture` VALUES ('50b4d07c-27d3-4b0f-84e4-1c3cae097ff0', 'logo_x.png', 6917, 'static/upload/admin/20180827120134logo_x.png', '5f067fcf-b25e-45db-92db-34e400e22870');
INSERT INTO `picture` VALUES ('c9333423-92fe-4a57-b4e0-edf73364ed36', 'logo.png', 6118, 'static/upload/admin/20180827114111logo.png', '5f067fcf-b25e-45db-92db-34e400e22870');

-- ----------------------------
-- Table structure for say
-- ----------------------------
DROP TABLE IF EXISTS `say`;
CREATE TABLE `say`  (
  `id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of say
-- ----------------------------
INSERT INTO `say` VALUES ('200a64a3-e8b8-4f56-acdb-9dabd20f749e', 'admin', '第一条说说！！！', '2018-08-27 11:37:17');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nick_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `phone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('06b4f8c8-5839-41ca-8191-a0a77524cfc6', 'admin', '000', '管理员', '110', '110@qq.com');

SET FOREIGN_KEY_CHECKS = 1;
