/*
 Navicat Premium Data Transfer

 Source Server         : 宝塔
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : bt.start6.cn:3306
 Source Schema         : nest

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 02/01/2021 23:46:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `order_num` int(16) NOT NULL COMMENT '排序',
  `created_at` int(64) NULL DEFAULT NULL,
  `updated_at` int(64) NULL DEFAULT NULL,
  `deleted_at` int(64) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES ('2c88cf9803af40c2bc65b3a3ae439183', '0', '测试部门', 0, NULL, NULL, NULL);
INSERT INTO `sys_dept` VALUES ('627f1bcd757a49d9957da740748daa0f', '2c88cf9803af40c2bc65b3a3ae439183', '子级', 0, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `type` int(11) NOT NULL COMMENT '菜单类型： 1. 目录 2. 菜单  3. 按钮',
  `order_num` int(11) NOT NULL COMMENT '排序',
  `perms` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识，接口标识',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单标识，前端路由name',
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NULL DEFAULT NULL,
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路径',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件地址',
  `redirect` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `hidden` tinyint(1) NULL DEFAULT 0 COMMENT '左侧菜单是否隐藏',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('13f92c3e8c7342a9be1a6b52cbcc724b', '50a91160c62c4ca491cc1c73fcaa7d32', '菜单管理', 1, 1, '', 'menu', 1609504175, 1609506140, 0, 'menu', 'views/sys/menu.vue', '', '', 0);
INSERT INTO `sys_menu` VALUES ('173968f5f2214744b2ae74efe593d060', '61c78284a49c4c5698843988f84278fd', '删除', 3, 1, 'del', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('1ec486c473794f5194ed7dba52633858', '61c78284a49c4c5698843988f84278fd', '角色分配', 3, 1, 'power', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('3f0101fc4a2045639f66148b141c248f', 'ec4390b0af754ba4b88f325648af1d27', '编辑', 3, 1, 'edit', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('4a11bb1e364942b2b6a35932fd717c95', 'ec4390b0af754ba4b88f325648af1d27', '添加', 3, 1, 'add', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('50a91160c62c4ca491cc1c73fcaa7d32', 'cf723a77df7441f7b62aa6d78e4909ba', '系统管理', 1, 2, '', 'sys', 1609504175, 1609506781, 0, '/sys', 'views/layout-children.vue', '', '', 0);
INSERT INTO `sys_menu` VALUES ('568d06c571e2462982a5b19329546362', '5eac78683b5d4c52b8f34e3eb85ad4ff', '删除', 3, 1, 'del', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('5eac78683b5d4c52b8f34e3eb85ad4ff', '50a91160c62c4ca491cc1c73fcaa7d32', '角色管理', 2, 1, '', 'role', 1609504175, 1609506140, 0, 'role', 'views/sys/role.vue', '', '', 0);
INSERT INTO `sys_menu` VALUES ('61c78284a49c4c5698843988f84278fd', '50a91160c62c4ca491cc1c73fcaa7d32', '用户管理', 2, 1, '', 'user', 1609504175, 1609506140, 0, 'user', 'views/sys/user.vue', '', '', 0);
INSERT INTO `sys_menu` VALUES ('7f266e6a435d40219817400a63ad868b', '13f92c3e8c7342a9be1a6b52cbcc724b', '添加', 3, 1, 'add', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('7f575e798c0a4b388a9edd6fa999e9c6', '5eac78683b5d4c52b8f34e3eb85ad4ff', '修改', 3, 1, 'edit', 'edit', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('8ab88ff40e184241a3bbf335c324e39b', 'cf723a77df7441f7b62aa6d78e4909ba', '首页', 2, 1, '', 'home', 1609504175, 1609506234, 0, '/home', 'views/home/home.vue', '', '', 0);
INSERT INTO `sys_menu` VALUES ('8fd813159b65483ebe187327ab6c4b50', '61c78284a49c4c5698843988f84278fd', '添加', 3, 1, 'add', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('a726fe51eda949338b6ca9773eb6a8d3', '13f92c3e8c7342a9be1a6b52cbcc724b', '编辑', 3, 1, 'edit', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('cd29666b9ba04cb08c8c153e10e2f563', '5eac78683b5d4c52b8f34e3eb85ad4ff', '权限分配', 3, 1, 'power', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('cf723a77df7441f7b62aa6d78e4909ba', '0', 'layout', 1, 1, '', 'layout', 1609504126, 1609504126, 0, '/', 'layout/layout/layout.vue', '/home', '', 1);
INSERT INTO `sys_menu` VALUES ('ddc0da9446d348ab8c023a1621f414ca', '61c78284a49c4c5698843988f84278fd', '修改', 3, 1, 'edit', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('ec4390b0af754ba4b88f325648af1d27', '50a91160c62c4ca491cc1c73fcaa7d32', '部门管理', 2, 1, '', 'depart', 1609504175, 1609506140, 0, 'depart', 'views/sys/depart.vue', '', '', 0);
INSERT INTO `sys_menu` VALUES ('f1e104d205af43f8ad158dda108a2ebf', '13f92c3e8c7342a9be1a6b52cbcc724b', '删除', 3, 1, 'del', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('f2610cb627974fab8c962d2b30b61e7f', '5eac78683b5d4c52b8f34e3eb85ad4ff', '添加', 3, 1, 'add', '', 1609591250, 1609591250, 0, '', '', '', '', 0);
INSERT INTO `sys_menu` VALUES ('fa1a90a20d5e45af80c0dc3e811f7f05', 'ec4390b0af754ba4b88f325648af1d27', '删除', 3, 1, 'del', '', 1609591250, 1609591250, 0, '', '', '', '', 0);

-- ----------------------------
-- Table structure for sys_oss
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss`;
CREATE TABLE `sys_oss`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '图片等url链接',
  `location` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件存放位置，暂存本地服务器，不采用云 oss',
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件类型',
  `size` int(11) NOT NULL COMMENT '文件大小 size',
  `old_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原文件名称',
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_oss
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色备注',
  `role_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('77092119fb7e4aec8959f7aa3706fd84', '超级管理员', 'admin', 1609506919, 1609506919, 0);
INSERT INTO `sys_role` VALUES ('93e30b154d2d4e3f847c866556301ea3', '测试', 'test', 1609597744, 1609597744, 0);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色ID',
  `dept_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门ID',
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色与部门对应关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `menu_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_b65fa84413c357d7282153b4a88`(`role_id`) USING BTREE,
  INDEX `FK_543ffcaa38d767909d9022f2522`(`menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色与菜单对应关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('09555b1d49cf4f1baa046ef64fef20aa', '77092119fb7e4aec8959f7aa3706fd84', '568d06c571e2462982a5b19329546362', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('0dc9d0c7e31649fa94f199d65877dd2c', '77092119fb7e4aec8959f7aa3706fd84', 'a726fe51eda949338b6ca9773eb6a8d3', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('23e6351bc3ae4b50bdb03548d147ce4d', '77092119fb7e4aec8959f7aa3706fd84', '1ec486c473794f5194ed7dba52633858', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('2d4f865384904c029fdae8c47b9fb81f', '77092119fb7e4aec8959f7aa3706fd84', '8fd813159b65483ebe187327ab6c4b50', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('3282bd770d7d4c84b4635567a5f6e45c', '0b1459e1cc3f4602a1f931795fab468e', '510802ba6ac74efeb20e2131a1d74aa6', 1608882063, 1608882063, 0);
INSERT INTO `sys_role_menu` VALUES ('3687a7d9086a4fc28ebfc19422f10714', '93e30b154d2d4e3f847c866556301ea3', '50a91160c62c4ca491cc1c73fcaa7d32', 1609597755, 1609597755, 0);
INSERT INTO `sys_role_menu` VALUES ('3ba558991d5241fa913c9d0bf00404ab', '77092119fb7e4aec8959f7aa3706fd84', '5eac78683b5d4c52b8f34e3eb85ad4ff', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('497b178e8fdc42d9b852e665a619124b', '77092119fb7e4aec8959f7aa3706fd84', '13f92c3e8c7342a9be1a6b52cbcc724b', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('6c89874a759a4501ae14379a77f7af5b', '0b1459e1cc3f4602a1f931795fab468e', 'c583848d23924d12be14af1a6029e747', 1608882063, 1608882063, 0);
INSERT INTO `sys_role_menu` VALUES ('6d77162c916a414bb6f7953c091779bb', '77092119fb7e4aec8959f7aa3706fd84', 'cf723a77df7441f7b62aa6d78e4909ba', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('701c556bf29b43c4825a5952be47f8e7', '77092119fb7e4aec8959f7aa3706fd84', '61c78284a49c4c5698843988f84278fd', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('74937fe88df74b039a89f675d5e75ca5', '77092119fb7e4aec8959f7aa3706fd84', '3f0101fc4a2045639f66148b141c248f', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('78447c70c62c4a2ba6f0adfd406d5b81', '93e30b154d2d4e3f847c866556301ea3', '13f92c3e8c7342a9be1a6b52cbcc724b', 1609597755, 1609597755, 0);
INSERT INTO `sys_role_menu` VALUES ('7bd38cc07ba64aca9efe52f6c9bf4902', 'ffd9e663ec604360a4d5b97e10aac986', '510802ba6ac74efeb20e2131a1d74aa6', 1609253044, 1609253044, 0);
INSERT INTO `sys_role_menu` VALUES ('9c2796a49701454b94333c3ee28a3cef', '93e30b154d2d4e3f847c866556301ea3', 'ec4390b0af754ba4b88f325648af1d27', 1609597755, 1609597755, 0);
INSERT INTO `sys_role_menu` VALUES ('a1074873972140259d511f3e4c7bae7d', '77092119fb7e4aec8959f7aa3706fd84', 'f2610cb627974fab8c962d2b30b61e7f', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('a675bdbc6afc40d7a4f61ef95e7db24b', '77092119fb7e4aec8959f7aa3706fd84', '8ab88ff40e184241a3bbf335c324e39b', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('abe6f0ea104a4696a3b3c449bc943119', '77092119fb7e4aec8959f7aa3706fd84', '7f575e798c0a4b388a9edd6fa999e9c6', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('bd5bfd3396d8449f9fbc9f13d2d40f08', '93e30b154d2d4e3f847c866556301ea3', '61c78284a49c4c5698843988f84278fd', 1609597755, 1609597755, 0);
INSERT INTO `sys_role_menu` VALUES ('c46699464778417cb86bae83fe71401e', '77092119fb7e4aec8959f7aa3706fd84', '50a91160c62c4ca491cc1c73fcaa7d32', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('c81e0b33b29b41baa32fc66797ba3b21', '77092119fb7e4aec8959f7aa3706fd84', 'f1e104d205af43f8ad158dda108a2ebf', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('d009e6eb88a5475080d13c247290003f', '93e30b154d2d4e3f847c866556301ea3', 'cf723a77df7441f7b62aa6d78e4909ba', 1609597755, 1609597755, 0);
INSERT INTO `sys_role_menu` VALUES ('d1a390d8849448f295f20a4300ce683d', '77092119fb7e4aec8959f7aa3706fd84', '7f266e6a435d40219817400a63ad868b', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('d1fc4774ded8400eb8ac1481e3a9fd67', '77092119fb7e4aec8959f7aa3706fd84', 'ec4390b0af754ba4b88f325648af1d27', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('db724651e0bf414497602eacf29f5136', '77092119fb7e4aec8959f7aa3706fd84', 'ddc0da9446d348ab8c023a1621f414ca', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('dc6c6670068e4cd487b1c6c6502e97c5', '77092119fb7e4aec8959f7aa3706fd84', '173968f5f2214744b2ae74efe593d060', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('e5f87d92385d49f18b889f6dbb290e43', '77092119fb7e4aec8959f7aa3706fd84', '4a11bb1e364942b2b6a35932fd717c95', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('e698368f0a2e480d8a63eea225fbb87a', '93e30b154d2d4e3f847c866556301ea3', '5eac78683b5d4c52b8f34e3eb85ad4ff', 1609597755, 1609597755, 0);
INSERT INTO `sys_role_menu` VALUES ('ed2b84bd92fd4dc4b2d1082874966044', '77092119fb7e4aec8959f7aa3706fd84', 'cd29666b9ba04cb08c8c153e10e2f563', 1609597080, 1609597080, 0);
INSERT INTO `sys_role_menu` VALUES ('ff1f2c11bc1b4ca5b97fd70c0a8ffe6c', '93e30b154d2d4e3f847c866556301ea3', '8ab88ff40e184241a3bbf335c324e39b', 1609597755, 1609597755, 0);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pass_word` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户登录密码',
  `account` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户登录账号',
  `nick_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户显示昵称',
  `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮箱地址',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '所属状态是否有效  1是有效 0是失效',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '头像地址',
  `dept_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `updated_at` int(64) NULL DEFAULT NULL COMMENT '更新时间',
  `phone_num` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户手机号码',
  `deleted_at` int(64) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sys_user_FK`(`dept_id`) USING BTREE,
  CONSTRAINT `sys_user_FK` FOREIGN KEY (`dept_id`) REFERENCES `sys_dept` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('22723c6941c1469aa41eb21cb4e8e724', '098f6bcd4621d373cade4e832627b4f6', 'test', '测试', '', 1, '', '627f1bcd757a49d9957da740748daa0f', 1609597358, 1609601636, '', 0);
INSERT INTO `sys_user` VALUES ('ff3f16866746402b8b257257d16048d0', '5211bae8172c37f40cfa6789a77a7ae7', 'admin', '超级管理员', '', 1, '', '2c88cf9803af40c2bc65b3a3ae439183', 1609506902, 1609602218, '', 0);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` int(64) NOT NULL,
  `updated_at` int(64) NOT NULL,
  `deleted_at` int(64) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_71b4edf9aedbd3e5707156e80a2`(`user_id`) USING BTREE,
  INDEX `FK_e8300bfcf561ed417f5f02c6776`(`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与角色对应关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES ('3a286ad09fb54c67a54d2fb9de5ba94e', 'ff3f16866746402b8b257257d16048d0', '77092119fb7e4aec8959f7aa3706fd84', 1609507009, 1609507009, 0);
INSERT INTO `sys_user_role` VALUES ('9b9fe9be101b4083a029ab57b6ae5092', '22723c6941c1469aa41eb21cb4e8e724', '93e30b154d2d4e3f847c866556301ea3', 1609600197, 1609600197, 0);
INSERT INTO `sys_user_role` VALUES ('a14eb50be3d8468db64db48f6bf6a1eb', '88b06e001e794f4fa9c83bf089a58366', '0b1459e1cc3f4602a1f931795fab468e', 1608891730, 1608891730, 0);

SET FOREIGN_KEY_CHECKS = 1;
