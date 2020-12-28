/*
 Navicat Premium Data Transfer

 Source Server         : nest-go
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : bt.start6.cn:3306
 Source Schema         : nest

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 28/12/2020 15:53:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` varchar(32) NOT NULL,
  `parent_id` varchar(32) NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '部门名称',
  `order_num` int(16) NOT NULL COMMENT '排序',
  `created_at` int(64) DEFAULT NULL,
  `updated_at` int(64) DEFAULT NULL,
  `deleted_at` int(64) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='部门管理';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` VALUES ('2c88cf9803af40c2bc65b3a3ae439183', '0', '测试部门', 0, NULL, NULL, NULL);
INSERT INTO `sys_dept` VALUES ('627f1bcd757a49d9957da740748daa0f', '2c88cf9803af40c2bc65b3a3ae439183', '子级', 0, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` varchar(32) NOT NULL,
  `parent_id` varchar(32) NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '菜单名称',
  `type` int(11) NOT NULL COMMENT '菜单类型： 1. 目录 2. 菜单  3. 按钮',
  `order_num` int(11) NOT NULL COMMENT '排序',
  `perms` varchar(255) DEFAULT NULL COMMENT '权限标识，接口标识',
  `code` varchar(30) NOT NULL COMMENT '菜单标识，前端路由name',
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单管理';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES ('510802ba6ac74efeb20e2131a1d74aa6', '0', '54', 1, 1, '5', '5', 1608862787, 1608862787, 0);
INSERT INTO `sys_menu` VALUES ('6eac043a5f3a44c280a360934d492da4', 'c583848d23924d12be14af1a6029e747', '1', 2, 1, '123', 'aaa', 1608862780, 1608862780, 0);
INSERT INTO `sys_menu` VALUES ('c583848d23924d12be14af1a6029e747', '0', '测试', 1, 1, '', '123', 1608718710, 1608718710, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_oss
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss`;
CREATE TABLE `sys_oss` (
  `id` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '图片等url链接',
  `location` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件存放位置，暂存本地服务器，不采用云 oss',
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件类型',
  `size` int(11) NOT NULL COMMENT '文件大小 size',
  `old_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原文件名称',
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_oss
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` varchar(32) NOT NULL,
  `remark` varchar(100) NOT NULL COMMENT '角色备注',
  `role_name` varchar(100) NOT NULL COMMENT '角色名称',
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES ('0b1459e1cc3f4602a1f931795fab468e', '55', '8', 1608648565, 1608648571, 0);
INSERT INTO `sys_role` VALUES ('ffd9e663ec604360a4d5b97e10aac986', '测试角色', 'test', 1608890039, 1608890039, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `id` varchar(32) NOT NULL,
  `role_id` varchar(32) DEFAULT NULL COMMENT '角色ID',
  `dept_id` varchar(32) DEFAULT NULL COMMENT '部门ID',
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与部门对应关系';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` varchar(32) NOT NULL,
  `role_id` varchar(32) NOT NULL,
  `menu_id` varchar(32) NOT NULL,
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FK_b65fa84413c357d7282153b4a88` (`role_id`) USING BTREE,
  KEY `FK_543ffcaa38d767909d9022f2522` (`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与菜单对应关系';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` VALUES ('3282bd770d7d4c84b4635567a5f6e45c', '0b1459e1cc3f4602a1f931795fab468e', '510802ba6ac74efeb20e2131a1d74aa6', 1608882063, 1608882063, 0);
INSERT INTO `sys_role_menu` VALUES ('6c89874a759a4501ae14379a77f7af5b', '0b1459e1cc3f4602a1f931795fab468e', 'c583848d23924d12be14af1a6029e747', 1608882063, 1608882063, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pass_word` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户登录密码',
  `account` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户登录账号',
  `nick_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户显示昵称',
  `email` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱地址',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '所属状态是否有效  1是有效 0是失效',
  `avatar` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像地址',
  `dept_id` varchar(32) CHARACTER SET utf8mb4 NOT NULL,
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `updated_at` int(64) DEFAULT NULL COMMENT '更新时间',
  `phone_num` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户手机号码',
  `deleted_at` int(64) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sys_user_FK` (`dept_id`),
  CONSTRAINT `sys_user_FK` FOREIGN KEY (`dept_id`) REFERENCES `sys_dept` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES ('4b6bfb615f294e9f91b366e7b24e6e14', '6537e99af2c2223642df9f70a0b5afca', 'test', '555', '', 0, '', '2c88cf9803af40c2bc65b3a3ae439183', 1608562179, 1608562179, '', 0);
INSERT INTO `sys_user` VALUES ('85573297efe44472945df8c98b5ba9b9', 'e10adc3949ba59abbe56e057f20f883e', 'test1', '123', '12312', 1, '', '2c88cf9803af40c2bc65b3a3ae439183', 1608993584, 1608993584, '12312', 0);
INSERT INTO `sys_user` VALUES ('88b06e001e794f4fa9c83bf089a58366', '220466675e31b9d20c051d5e57974150', 'admin', '超级管理员', '', 1, '', '627f1bcd757a49d9957da740748daa0f', 1608545956, 1608562121, '', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `id` varchar(32) NOT NULL,
  `user_id` varchar(32) NOT NULL,
  `role_id` varchar(32) NOT NULL,
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FK_71b4edf9aedbd3e5707156e80a2` (`user_id`) USING BTREE,
  KEY `FK_e8300bfcf561ed417f5f02c6776` (`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户与角色对应关系';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` VALUES ('3aaec12a1d6e4a53bd8d3e2c566af93d', '85573297efe44472945df8c98b5ba9b9', 'ffd9e663ec604360a4d5b97e10aac986', 1608891712, 1608891712, 0);
INSERT INTO `sys_user_role` VALUES ('a14eb50be3d8468db64db48f6bf6a1eb', '88b06e001e794f4fa9c83bf089a58366', '0b1459e1cc3f4602a1f931795fab468e', 1608891730, 1608891730, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
